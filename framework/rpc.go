package framework

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	lambdaSDK "github.com/aws/aws-sdk-go/service/lambda"
)

// RPCRouter struct provides an interface to handle remote procedures (other Lambdas invoking the one listening via AWS SDK)
type RPCRouter struct {
	handlers map[string]RPCHandler
	Tracer   TraceStrategy
}

// RPCHandler is similar to and other router/handler but it returns a map[string]interface{} in addition to an error
type RPCHandler func(context.Context, map[string]interface{}) (map[string]interface{}, error)

// LambdaHandler is a native AWS Lambda Go handler function. Handles a remote procedure call (invocation via SDK with a special event format).
func (r *RPCRouter) LambdaHandler(ctx context.Context, evt map[string]interface{}) (map[string]interface{}, error) {
	var err error
	var response map[string]interface{}

	handled := false
	procedureName := ""
	if name, ok := evt["_rpcName"]; ok {
		procedureName = name.(string)
	}

	if r.handlers != nil {
		// If there's a _rpcName, use the registered handler if it exists.
		if handler, ok := r.handlers[procedureName]; ok {
			handled = true
			// Trace (default is to use XRay)
			// Annotations can be searched in XRay.
			// For example: annotation.RPCName = "myProcedure"
			r.Tracer.Annotations = map[string]interface{}{
				"RPCName": procedureName,
			}
			err = r.Tracer.Capture(ctx, "RPCHandler", func(ctx1 context.Context) error {
				r.Tracer.AddAnnotations(ctx1)
				r.Tracer.AddMetadata(ctx1)
				response, err = handler(ctx, evt)
				return err
			})
		}
		// Otherwise, use the catch all (router "fallthrough" equivalent) handler.
		// The application can inspect the map and make a decision on what to do, if anything.
		// This is optional.
		if !handled {
			// It's possible that the RPCRouter wasn't created with NewRPCRouter, so check for this still.
			if handler, ok := r.handlers["*"]; ok {
				// Capture the handler (in XRay by default) automatically
				r.Tracer.Annotations = map[string]interface{}{
					"RPCName":            procedureName,
					"FallthroughHandler": true,
				}
				err = r.Tracer.Capture(ctx, "RPCHandler", func(ctx1 context.Context) error {
					r.Tracer.AddAnnotations(ctx1)
					r.Tracer.AddMetadata(ctx1)
					response, err = handler(ctx, evt)
					return err
				})
			}
		}
	}

	return response, err
}

// Listen will start an RPC listener which acts much like a task handler except that it handles special RPC events instead
func (r *RPCRouter) Listen() {
	lambda.Start(r.LambdaHandler)
}

// NewRPCRouter simply returns a new RPCRouter struct and behaves a bit like Router, it even takes an optional rootHandler or "fall through" catch all
func NewRPCRouter(rootHandler ...RPCHandler) *RPCRouter {
	// The catch all is optional, if not provided, an empty handler is still called and it returns nothing.
	handler := func(context.Context, map[string]interface{}) (map[string]interface{}, error) {
		return map[string]interface{}{}, nil
	}
	if len(rootHandler) > 0 {
		handler = rootHandler[0]
	}
	return &RPCRouter{
		handlers: map[string]RPCHandler{
			"*": handler,
		},
	}
}

// Handle will register a handler for a given remote procedure name
func (r *RPCRouter) Handle(name string, handler RPCHandler) {
	if r.handlers == nil {
		r.handlers = make(map[string]RPCHandler)
	}
	r.handlers[name] = handler
}

// RPC will make the remote procedure call (invoke another lambda)
func RPC(ctx context.Context, functionName string, message map[string]interface{}) (map[string]interface{}, error) {
	sess, err := session.NewSession()
	if err != nil {
		log.Println("could not make remote procedure call, session could not be created")
		return nil, err
	}

	// Payload will need JSON bytes
	jsonBytes, err := json.Marshal(message)
	if err != nil {
		log.Println("could not marshal remote procedure call message")
		return nil, err
	}

	// region? cross account?
	svc := lambdaSDK.New(sess)
	// Wrap in XRay so it gets logged and appears in service map
	// TODO: We don't hav r *RPCRouter here...So we don't have a configurable interface to use...
	// TraceStrategy.AWS()
	// xray.AWS(svc.Client)
	AWSClientTracer(svc.Client)

	// TODO: Look into this more. So many interesting options here. InvocationType and LogType could be interesting outside of defaults
	output, err := svc.InvokeWithContext(ctx, &lambdaSDK.InvokeInput{
		// ClientContext // TODO: think about this...
		FunctionName: aws.String(functionName),
		// JSON bytes, sadly it does not pass just any old byte array. It's going to come in as a map to the handler.
		// That's a map[string]interface{} from JSON. I saw byte array at first and got excited.
		Payload: jsonBytes,
		// Qualifier ... this is an interesting one. We use latest by default...But we also want to work in a circuit breaker
		// here, so we'll need to set the qualifier at some point.
	})

	// Unmarshal response.
	var resp map[string]interface{}
	if err == nil {
		err = json.Unmarshal(output.Payload, &resp)
	}

	return resp, err
}
