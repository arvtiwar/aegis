// Code generated by go-bindata.
// sources:
// bindata.go
// index_http_proxy.js
// index_stdio.js
// shim.go
// DO NOT EDIT!

package shim

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _bindataGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func bindataGoBytes() ([]byte, error) {
	return bindataRead(
		_bindataGo,
		"bindata.go",
	)
}

func bindataGo() (*asset, error) {
	bytes, err := bindataGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "bindata.go", size: 0, mode: os.FileMode(436), modTime: time.Unix(1512878110, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _index_http_proxyJs = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x56\xdf\x6f\xdb\x36\x10\x7e\xa6\xfe\x8a\x9b\x30\x54\x12\xe2\x49\xd9\xb0\x3d\x4c\x86\x07\xb4\x9d\xb7\x75\x68\xd3\x6e\xf1\x8a\x3d\x0c\x48\x68\xf1\x64\x31\x95\x49\x8d\xa4\xea\x18\xa9\xff\xf7\xe1\x48\xc9\x91\x93\x0c\x1b\xf6\xb4\x3e\x34\x21\xf9\xdd\x7d\xdf\xfd\x54\xa2\x4a\x2b\xeb\xe0\xcd\xf3\xdf\xaf\x7e\x78\xfe\xea\xf5\x25\x2c\xe0\xeb\x79\x14\x7d\xe4\x06\xaa\x46\xb6\xe2\xaa\x33\xba\x42\x6b\x61\x01\x06\xff\xec\xa5\xc1\x34\x39\x79\x48\xb2\x59\xc4\x36\xda\x1f\x61\x01\xaa\x6f\xdb\x59\xc4\x84\x56\x08\x0b\x20\xe7\xba\xc5\xbc\xd5\x9b\x7c\x2d\x95\x48\x87\x0b\xb2\xa9\xb9\x6c\xc9\xed\x79\xc4\x1a\xe7\xba\x29\x01\x9d\x93\x6c\x1e\x45\x69\xdd\xab\xca\x49\xad\x40\xe1\xee\x6a\x60\x49\x33\xb8\x8b\x22\x56\x14\xd0\xc9\x0e\xc1\x3a\x21\x55\xa1\x7b\x37\x83\x75\x2b\x95\x80\x8e\x5b\xeb\x1a\xa3\xfb\x4d\x43\x8f\x68\xcc\x54\xe0\x89\xf8\xdc\x76\x7c\xa7\xd2\x24\x2f\x38\x6e\xa4\xbd\xe2\x5d\xe0\x1d\xf1\xb9\x56\x69\x82\xc6\x68\x93\xcc\x60\xd4\x92\xa2\x31\x24\x81\xb1\xa3\x17\xcf\x92\xef\x8c\x74\x98\xc6\x23\x97\xb7\x43\x51\x42\x7c\xf6\xf3\xe5\xdb\x8b\xdc\x3a\x23\xd5\x46\xd6\x7b\xef\xe0\x2c\xfe\x43\xc5\xd9\x3c\x62\x4c\xd6\x90\x9e\x9d\x85\x74\x7c\x77\x5f\x89\x40\x71\xe4\xc0\x5b\xe9\xd2\x2f\xb3\x39\x14\x05\xd4\xda\x54\x48\xc9\x75\x5c\x2a\x34\x60\xd0\x3a\x6e\x1c\xf0\xda\xa1\x01\xa7\x35\x6c\xb9\xda\x83\x77\x19\x31\x76\x88\x18\x3b\xc9\x1f\xb1\x52\x85\xbc\x90\x79\xc4\x0e\x8f\x63\xbe\x95\x6e\x1a\x72\xa5\x05\xfe\xab\x98\x6f\xa5\x43\x01\x9d\xc1\x2d\x77\xbd\xc1\x76\x0f\x3b\xe9\x1a\x20\x7b\x4a\x04\xfd\xfc\x5f\x44\xae\x70\x07\x4b\x2a\x50\x1a\x2f\x83\xe6\xa3\xce\x41\x66\x36\x66\xe6\x90\x91\x5d\x84\xb7\x9d\x36\xce\xe6\x0d\x57\xa2\x45\x03\x8b\x49\x43\x7c\x44\xe5\x66\x5e\x16\xde\xba\x63\x7b\xf2\x76\xc7\xf7\x16\x74\xef\xba\xde\x81\xd3\x50\xf5\xc6\xa0\x72\x23\x30\xb1\x40\x5a\xa6\xd3\x42\xd7\x39\x1d\x8f\xe3\xe2\x3d\xce\x83\xc3\xb1\x42\xbe\xe5\x87\xd4\x3f\x68\xad\x3b\x0f\x64\xb1\x97\x14\x97\x10\xa4\x0d\x97\x83\xbb\xb8\x1c\xa9\xfc\xfd\xe1\xd8\x8b\x11\xa3\xb9\xa7\x30\x61\x01\xdf\x7e\x73\x7e\x3e\x8f\x98\xac\x43\x78\xb9\x75\x7c\x83\xef\xb9\x91\x7c\xdd\xa2\x85\xcf\x16\x0b\xe8\x95\xc0\x5a\x2a\x14\xf0\xec\x19\x3c\x85\xca\x5b\xbe\x5d\x0b\xfe\x8e\x3c\xfe\x27\x03\x5a\x27\x43\xdf\x05\x55\xff\x60\x44\x25\x0b\xb9\xf2\x3b\x65\xb2\x0a\x42\x6c\xba\xa3\x82\xd1\xda\xb9\x8b\x18\x40\xa3\xad\x53\x7c\x8b\x25\x24\xad\xae\x78\x4b\xe7\x64\x46\x2f\x44\x57\xfa\xff\xc3\x91\xbb\x66\x48\x66\x4e\xbf\xc3\xa7\x4f\x90\x14\x01\xba\x45\xd7\x68\x31\xbe\x12\xef\x1b\x7f\xe3\x31\x3f\x2e\x57\x01\xd5\x20\x17\x68\xec\x11\x16\x8e\x84\xf1\x4a\x00\x92\x97\x54\x13\xe5\xbe\x58\xed\x3b\x4c\x4a\x48\x78\xd7\xb5\xb2\xe2\x24\xb8\xb8\xb1\x5a\x05\x3f\x13\xe0\x6b\x54\x1b\xd7\x24\x25\xbc\xe8\xeb\x1a\x4d\xbe\xde\x3b\x0c\x77\x43\xc9\xd6\x5a\xec\x33\x32\x3a\x44\xec\x30\xf4\xd0\x64\x25\xa7\xf1\xf2\xfd\xf2\x62\x55\xfa\x61\x7c\xf0\xe4\x1d\x64\x4f\xd9\xbc\x7d\xb7\x7a\xf5\xf6\xe2\xf2\x49\xab\x21\xbd\xde\xae\xe2\x6d\xbb\xe6\xd5\x87\xe9\xa0\x18\xb4\x9d\x56\x36\xac\x12\x00\xaa\x88\x41\xfb\x42\x8b\x3d\x2c\x20\x49\xc8\x0c\xa0\x28\xb8\xd2\xae\x41\xfa\x02\xf5\xea\x03\xe8\x1a\x04\x77\x1c\x1a\x6e\x61\x8d\xa8\xc0\x60\x25\xf1\x23\x8a\x19\x58\x0d\xbc\xeb\x50\x09\x90\x7e\xc0\xae\xad\x33\xd7\xe4\x63\x24\xf2\xcb\x8c\xac\x27\xcb\x0c\x52\xef\x37\x1b\x13\x3f\x2a\x38\x5b\x04\xc2\xb9\xcf\x58\x16\xc4\x9c\x38\x7a\xfc\x25\x08\xbd\x59\x14\xd3\x24\x5c\x77\x46\xaf\x5b\xdc\x86\x8d\x42\x1f\x35\xb4\xae\x84\xcf\xef\x30\xdf\xa2\xb5\x7c\x83\x87\xeb\xc9\x0a\xce\x4e\xf8\x8a\xc2\x35\x08\xbb\x46\xb7\x78\xe4\xfe\x9b\xc8\x77\x08\x37\xbd\x75\xd0\x19\xa9\x1c\x25\x40\xf7\x0e\x1a\x34\xf8\x58\xb7\x12\x27\xf1\x0f\xa1\xb3\x93\xba\xfe\xb4\x5a\xbd\x83\x5f\x97\xbf\xfc\xb6\xbc\x5c\xc1\xf2\xe2\xfb\x38\x08\x3b\x01\x0d\xa9\xf2\x3d\xc5\x8a\x82\xea\x77\x03\x0b\xf0\x3b\xa8\xe3\xc6\xe2\x11\x31\x9f\x40\xec\x08\xb9\x5f\x53\x37\x01\x00\x10\x96\x31\xfd\xd1\x00\x77\x11\x0c\xff\x58\x6c\x1d\x77\xbd\x7d\xa9\x05\xc6\x25\xc4\x5f\x9d\x9f\xc7\xb3\xc9\xeb\x30\x3c\x71\x39\xb5\x61\xf1\x74\x80\xc8\xec\xe1\x00\xc5\xf7\xe0\xc3\xc4\x5d\x51\xc4\x34\x28\x71\xf9\x54\x1c\x13\xd6\x29\xe8\x3e\x92\x47\xc0\xc3\x7d\x41\xe9\xfb\x43\xfb\x20\x1f\x9a\x60\x9c\x8f\x19\x8c\xc3\x91\xe5\xa8\x84\xff\xbc\x44\x87\xbf\x02\x00\x00\xff\xff\x90\xbd\x5d\x5e\x8b\x09\x00\x00")

func index_http_proxyJsBytes() ([]byte, error) {
	return bindataRead(
		_index_http_proxyJs,
		"index_http_proxy.js",
	)
}

func index_http_proxyJs() (*asset, error) {
	bytes, err := index_http_proxyJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "index_http_proxy.js", size: 2443, mode: os.FileMode(436), modTime: time.Unix(1510187626, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _index_stdioJs = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x57\xd1\x6f\xe3\x36\xf2\x7e\xd7\x5f\xf1\xc1\x40\x2b\x19\xeb\xc8\xd9\x87\x1f\xf0\x83\x7d\x7e\x48\x77\xdb\xdd\xb4\x69\xb3\x68\x52\xec\x43\x2e\x08\x68\x71\x2c\x71\x43\x93\x3a\x92\xb2\xe2\x4b\xf3\xbf\x1f\x48\x4a\x8a\x9c\xb8\x8b\x6b\x71\x38\xdc\xcb\x6e\x4c\x0d\xbf\xf9\x38\xf3\xcd\x70\x58\x1b\x5d\x90\xb5\x39\xa9\xdd\x4d\xfa\xe9\xec\xfa\x63\x7a\x8b\x15\x8e\xad\xbe\x41\xba\x48\xf1\xe6\xf0\xdb\xc5\xd9\xcf\xdf\xbd\x3f\xbb\xbb\x3e\xbb\xfa\xe9\xee\xd7\xcb\xcb\xeb\x91\xdd\xdd\x1d\x17\x46\xb1\x2d\x2d\x93\xf9\x1c\x85\x56\x56\x4b\xca\xa5\x2e\xb3\x23\xe8\xd3\x65\x92\xec\x98\x41\x51\x09\xc9\xef\x3a\x03\xac\x60\xe8\x1f\x8d\x30\x94\xa5\x07\x1f\xd2\xe9\x32\x11\x1b\x64\xa5\x0e\x2b\x58\xad\x56\x68\x14\xa7\x8d\x50\xc4\xa7\x78\x4c\x70\xe0\x6f\x62\x1d\x33\x4e\xa8\x12\x8a\xda\xe8\xa2\x3f\x05\xd6\x54\xb0\xc6\x12\x7a\xa8\x96\xd9\x67\xa8\xc9\x74\x99\x00\x9e\xd7\xe0\xe9\x90\x61\x6e\x6b\xd6\xaa\x2c\xcd\xe7\x8c\x4a\x61\xef\x58\x5d\xa7\x33\x3c\xc2\x3a\x2e\xf4\x02\x37\x69\x2d\x6a\x4a\x67\xe8\xff\x1f\xb6\x39\x4e\xc6\xdc\xe2\x69\xba\x4c\x9e\x12\x1f\x9f\xf3\x0d\x5c\x45\x08\x78\xc4\x5f\x90\xac\x98\x05\x53\x20\x63\xb4\x99\x41\xea\x12\x4c\x71\x18\x72\x8d\x51\x70\x15\x73\x79\xd2\x11\xcc\xb5\xca\xd2\x60\x97\xce\xb0\x69\x54\xe1\x84\x56\x19\x19\x13\xa3\x32\x9f\xbf\xa0\x90\xb7\x46\x38\xca\x26\xfd\xf9\xc2\x56\xe2\x0b\x4c\xde\xfc\x78\x75\xf9\x4b\x6e\x9d\x11\xaa\x14\x9b\x7d\xc0\x78\x33\xf9\xbb\x8a\x31\x99\xcf\xc1\xb5\xa2\xb0\xba\x4c\xe2\xc2\x85\x2e\x21\x1c\x9c\xc6\x3b\xa9\x1b\xfe\x99\xb9\xa2\x42\x56\x1b\xda\x09\xdd\x58\xb9\x9f\x81\x39\x47\xdb\xda\x11\xf7\x46\x03\x7d\xa3\x9b\xb2\xc2\xd9\xa7\x73\x7c\x60\x8e\x5a\xb6\x9f\x8e\xf2\x17\x08\x65\xe9\x2b\x7e\xdf\xd8\x74\x86\xb1\xf7\x5f\x34\xa7\xfc\x8b\x1d\x62\x46\x0f\xc2\xa1\x15\x52\xa2\x30\xc4\x1c\x81\x29\xed\x2a\x32\x28\xb4\xe4\x08\x8a\xc8\x13\x3c\xab\xf9\x41\xb8\xec\xad\x4f\x87\x47\x8c\x09\x11\x2e\xa0\xd8\x19\x0a\x26\x65\x3c\xf0\x14\xad\x70\xd5\x90\x8d\x90\x08\x9f\xb8\xe0\xae\xf0\x1c\x0e\x53\xf1\x20\xdc\x38\x13\xde\x62\x06\x2b\x4a\xc5\xe4\x90\x92\xb3\x92\x09\x35\xc3\x97\xc6\xba\x90\xdc\x57\x41\xf4\x4e\x3c\x52\xfe\x95\xc0\x3c\x08\xb7\x08\x04\x56\xdf\xd8\xce\xc1\x2a\x04\xe9\xc0\xe5\xf2\x45\x69\xa4\xfe\x64\x6b\x56\xdc\x5b\xb0\x78\xda\x45\x1a\x8f\x1b\x16\x83\xfd\xd7\xa2\xf4\x91\x29\x2e\x09\x9c\x39\x86\x42\x6f\x7d\x8d\xf9\x8d\xd0\x8d\x0b\xf5\x1c\x3e\xac\xa0\x1a\x29\x97\xfe\xa8\x7f\x3b\x39\x81\x50\xc2\x09\x26\xb1\x63\xb2\xa1\x19\x0c\x59\x0a\x07\xf6\x46\xd0\x0a\xc4\x8a\x0a\x55\x00\x36\xd9\x74\x08\xa7\x75\x5c\x37\x2e\x44\xd5\xa3\x1e\x44\xb5\x6a\xd4\x7d\x08\x67\x02\x6c\x98\x90\xbe\x75\x9c\x06\x87\x11\x3d\xac\x25\x80\x6f\x1a\x91\xd2\x2a\x92\x8a\x39\xc0\xc0\x93\x5a\x7c\xd7\x6c\x36\x64\x3a\x4c\x7f\xfe\x27\x90\xb4\x34\x32\xec\x8a\x66\x64\x11\xf3\x58\x54\x54\xdc\x63\xa3\x8d\xc7\x91\x42\x11\x98\x2d\x84\x40\x51\x31\x83\xb7\xa7\xd1\xa8\xa7\x90\x4b\x52\xa5\xab\xf0\xed\xb7\x01\xf3\x66\xb4\x76\xf2\xf6\x16\xab\x15\xde\x9e\x0e\x0a\x89\x4d\x48\x37\xae\x6e\x1c\x56\x08\xa5\x59\x33\x63\x29\x42\x39\x7d\x15\xea\x34\x4b\x7f\xbb\xfe\xe1\xe4\xff\xd3\xe9\x50\x15\xfe\x9f\x83\x56\xd8\x61\x6c\x8c\xde\xa2\xd4\x8b\xc9\xac\x43\xed\xab\xfa\xd5\x86\x41\x0b\xde\x76\x2c\x8c\x67\x7b\x7c\x20\x07\x06\x43\x1b\x32\xa4\x0a\xf2\xc9\xf4\x35\xd1\x5b\x77\xdd\x6a\xab\x77\xe4\xb5\x1d\x7c\xfb\xef\x35\x33\xa4\x1c\x6c\xa1\x6b\x82\xd5\x10\x2e\xb5\xd0\x4a\xee\x61\xfd\x7a\x14\x92\x2a\x28\x1f\x07\x21\xf4\xe0\x9e\xc6\x4d\x24\x9f\x57\xc4\x38\x19\x7b\x33\xf1\xf7\x05\x59\x77\x22\xf8\xe4\xf6\x16\xbf\xff\x7e\xd0\x03\x67\x21\xd2\x53\x3c\xe2\xe8\x09\xa1\xb4\xc3\x46\x37\xca\xf7\xfd\x43\x93\x9a\x59\x4b\x3c\x6c\xf7\x61\x08\x30\x4b\x3c\x0d\x31\xe3\x24\xc9\xd1\xbf\xcb\x6b\x14\x6a\xfc\x1a\xf4\x19\xe4\x57\x91\xf1\x8a\x41\x4b\x52\x0e\xc0\xa3\xfa\x19\x22\x7e\x9c\xbd\xd7\x5d\xe7\x06\xc2\x77\xf0\x3e\xb5\xc7\x59\x8c\x13\x1e\x5b\x9b\x77\xf2\x2c\x87\x50\xaa\x5a\xf2\xc1\xe9\x8b\xef\x71\x7d\x54\x1b\xaf\xa9\x29\x1d\x0a\xca\x57\xc2\xe4\xd9\x5d\x0f\xf4\x68\x1d\x73\x8d\x7d\xa7\x39\x2d\xf0\x7f\xa7\xa7\x33\xac\x35\xdf\x2f\x30\x99\x3c\x0d\xe8\x49\xf2\x57\x95\xff\x67\x44\xff\x67\xf4\xfe\x5f\x90\xfa\xff\xa2\xca\xff\x92\xc0\xff\x58\xdb\x2f\x65\xfd\x1f\x55\xf4\xd7\xc5\xfc\x4a\xc7\xfd\x55\xf6\x9b\xa5\x90\xaa\x30\x23\xf6\x14\x84\xb2\x8e\x18\x87\xde\xf4\xd7\xbf\x6d\x8a\x82\x88\x67\x53\x90\x2b\x72\xbf\xf1\x27\xa2\x1a\x0c\x5b\x56\x7b\xb3\x6d\x23\x9d\xa8\xe5\x28\x56\x10\x2a\x8e\x08\x3b\x9f\x6a\xbd\xf1\xa7\x2d\x1a\x13\x04\xd1\xf1\xb7\x79\x9c\x7e\x87\x2d\x2b\x3c\x3e\x2d\x93\x84\x1e\x6a\x6d\x9c\xcd\xbb\xdb\x10\xab\x51\xa6\x3d\x9a\xbf\xdd\x95\xa3\x07\xff\xc7\x7a\xb8\x2d\x2a\xe7\xea\xc5\x7c\xce\x75\x61\x73\xd6\xda\x9c\x6d\xd9\x3f\xb5\xca\x0b\xbd\x9d\x4b\xb6\x5d\x73\x36\x97\xcc\x91\x75\x73\x5e\xce\x95\xe6\xf4\xc5\x9e\xd4\x46\x97\x27\x5b\xcd\x49\x9e\x74\x88\x79\xe5\xb6\xb2\x1f\xad\x1c\x6c\x63\x08\xe7\x68\x99\x72\x70\x95\xb0\x79\xde\x0d\x23\xc1\xb6\x27\xfe\x99\x09\x67\x7f\xd0\xe6\xfb\x6d\xed\xf6\xdf\x7b\x86\x17\x5a\xd7\x9e\x36\x93\x96\xba\x61\x62\x3e\x07\xe3\x61\x04\x8c\x11\xf1\xb1\x11\x6a\xa7\xef\x09\x4e\x6c\x09\x99\xe6\x5c\xee\x83\x4e\x6b\x7f\x7d\x2b\xe7\x23\xd8\xb9\x82\x36\x71\xdb\x34\x22\x9d\xa3\x15\xb6\xf2\xb5\xe6\x27\xf7\x50\x6f\x6d\x45\x6a\x3c\x4e\xc2\x50\x41\x62\x47\x71\x50\xeb\x22\x9e\xe7\xb9\xea\x8f\x25\xc2\xec\x6d\x28\xb5\x60\xf0\x1b\x9c\x86\xaf\x85\x30\x5a\x43\xa8\x8d\xee\xcb\xb2\x32\xd7\x9e\xe1\xf3\x03\xa9\x32\x9e\x72\x16\xa4\x17\x68\xf5\xa9\xba\xf2\x03\xe6\xc7\xde\x3c\xee\xeb\xf4\xe9\x5b\x84\xd2\x66\xcb\x24\x7e\x64\x3b\x76\x55\x18\x51\x3b\x6c\x85\x94\xc2\x52\xa1\x15\xb7\x47\xc1\x3c\xc2\xcf\xb6\x1b\x52\xde\x33\x47\xd9\x34\x2f\x29\x2c\x67\x43\x7b\x62\xb2\x65\x7b\xdb\xf7\x4b\xa7\xd1\x0b\xad\x0b\x5f\x6a\x83\x8e\x9f\x0b\xc5\xf7\x99\x2e\x89\xfe\x67\xbe\x16\x8a\x67\xdd\x4a\x57\x37\x97\x17\xef\x8f\xf4\xd4\x40\xd0\xb7\x89\x98\x8d\x63\x9d\x34\xa2\x84\x3e\xda\x03\x46\xc7\x7d\xa9\x9d\xbf\x47\x29\x76\xa4\x3c\xd1\xc6\x62\xbd\xc7\x45\xd4\xe7\x28\x7b\x41\x66\x43\xdb\x89\x61\xe9\x72\xf8\xae\x23\xde\xfd\x3c\xe7\xfe\xe9\x5a\xac\x8f\xdd\x00\x51\x69\xa3\x66\xd2\x13\xff\x43\xac\x57\xa3\xf2\xf3\xcd\x80\x35\x6d\xb4\xa1\xf0\xca\x0b\xdc\x59\x5d\x1f\xbf\x2d\x3e\x89\xda\x6b\xbb\xbb\x24\x3e\x04\xcb\x51\x2f\x08\x6d\x90\xd9\x70\xab\x21\x3e\xb5\x12\x60\x34\xf0\x0a\xd5\xcd\x9a\x2f\x1e\x63\x71\x16\x8d\xa7\x9a\x2c\x22\xda\x2c\xae\x75\xa1\x9e\x2c\xfa\xa0\xfb\xf9\x74\x78\xb9\x45\x5a\x57\xce\xbf\x8b\x14\xc5\x87\xd8\xd0\xa4\x85\x3a\xbc\xa7\x46\x33\xb9\x2f\xf9\x20\xbc\x4e\x91\xdd\x48\x76\xe9\xeb\xa6\x15\x96\xd0\x92\x97\x53\xea\x70\xaf\x74\x8b\xb6\x12\x45\xd5\xab\xb0\xd4\x64\xe3\xab\x29\x2e\x1b\xb2\xb5\x56\x96\x72\x7c\xf6\x95\x2a\x5c\x7c\x8d\xd9\x39\x29\x6e\xf3\x97\xb7\x03\x90\x3c\xfd\x2b\x00\x00\xff\xff\x75\xdb\x65\x80\xa5\x10\x00\x00")

func index_stdioJsBytes() ([]byte, error) {
	return bindataRead(
		_index_stdioJs,
		"index_stdio.js",
	)
}

func index_stdioJs() (*asset, error) {
	bytes, err := index_stdioJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "index_stdio.js", size: 4261, mode: os.FileMode(436), modTime: time.Unix(1512878099, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _shimGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4c\xcc\xb1\x11\xc2\x30\x0c\x85\xe1\xde\x53\xbc\x05\x12\xf7\xcc\x40\xc1\x0a\xcf\x58\x08\x5f\x88\xe4\x93\x1d\xee\xd8\x9e\x22\x4d\xca\xff\x2f\xbe\x9c\xd5\x6f\x2a\x26\xc1\x29\x50\x5f\x4a\xb3\xca\x49\x2c\x7d\x53\x8c\x77\xdb\xb1\xa6\x94\x33\x1e\x7c\x6e\x54\x39\x57\x0f\xff\xb6\x2a\x03\x3c\xfb\xe5\x81\x38\xcc\x9a\x29\x18\xa5\xcd\x60\xfc\xf0\xa1\xe9\x41\x95\x01\x37\xdc\xb9\x97\xca\x35\xf5\x8b\x93\xfe\x01\x00\x00\xff\xff\xb4\x05\x8d\x5f\x7e\x00\x00\x00")

func shimGoBytes() ([]byte, error) {
	return bindataRead(
		_shimGo,
		"shim.go",
	)
}

func shimGo() (*asset, error) {
	bytes, err := shimGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "shim.go", size: 126, mode: os.FileMode(436), modTime: time.Unix(1510187626, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"bindata.go":          bindataGo,
	"index_http_proxy.js": index_http_proxyJs,
	"index_stdio.js":      index_stdioJs,
	"shim.go":             shimGo,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"bindata.go":          &bintree{bindataGo, map[string]*bintree{}},
	"index_http_proxy.js": &bintree{index_http_proxyJs, map[string]*bintree{}},
	"index_stdio.js":      &bintree{index_stdioJs, map[string]*bintree{}},
	"shim.go":             &bintree{shimGo, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
