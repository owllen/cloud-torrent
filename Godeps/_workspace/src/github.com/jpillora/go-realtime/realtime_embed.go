// Code generated by go-bindata.
// sources:
// realtime.js
// DO NOT EDIT!

package realtime

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
	"os"
	"time"
	"io/ioutil"
	"path"
	"path/filepath"
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
	name string
	size int64
	mode os.FileMode
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

var _realtimeJs = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xd4\x3c\xfb\x5b\xdb\x48\x92\xbf\xf3\x57\x34\xfe\x72\x83\x4c\x8c\x0c\x33\x3b\xf7\xc0\x43\x6e\x19\xe2\xdc\x91\x63\x70\x8e\x90\xec\xdd\xc7\xb0\x46\x96\xda\xb6\x82\x2c\x69\xd5\x32\xe0\x4b\xd8\xbf\xfd\xaa\xfa\xa5\x6e\xa9\x65\x4c\x26\xb7\x37\xeb\x6f\x26\xb6\xd5\xd5\xd5\xd5\xf5\xae\xea\x36\xfd\x5d\x32\x4b\xb2\x49\x90\x90\xfb\x38\x8d\xb2\xfb\xde\x9f\xe8\xe4\x7d\x16\xde\xd2\x92\xec\xf6\xb7\xfa\x7d\x72\x41\x83\xa4\x8c\x17\x94\xec\x91\xbb\x7d\xff\xc0\xdf\x87\x0f\xf3\xb2\xcc\xd9\x61\xbf\x3f\x8b\xcb\xf9\x72\xe2\x87\xd9\xa2\xff\x29\x8f\x93\x24\x2b\x82\xfe\x2c\xdb\x2b\xe4\x14\x9c\xfe\x36\xc0\xb9\xef\xc4\x20\xf9\x29\xa2\x77\x7f\x54\xa0\x38\xef\x15\x60\xfb\xe5\xf4\x92\x9c\x64\xf9\xaa\x88\x67\xf3\x92\x7c\xbf\x7f\xf0\xe3\x96\x37\x5d\xa6\x61\x19\x67\xa9\x27\xa9\x22\x51\x16\x2e\x17\x34\x2d\xbb\xe4\xf3\xd6\x16\x21\xf1\xd4\xdb\x16\x43\xbe\x26\xb8\x0b\xcf\x09\x29\x68\xb9\x2c\x52\x12\x24\xb4\x28\xbd\xce\xe5\x3c\x66\x64\x52\x64\xf7\x8c\x16\x80\x83\x32\x92\x66\x25\x61\xcb\x3c\xcf\x8a\x92\xe8\xa9\xac\xd3\x1d\x20\xda\x7e\x5f\xd1\x4e\xf2\x22\x2b\xb3\x30\x4b\xc8\x1d\x2d\x18\x50\x02\xa3\x77\x41\x21\x1e\x93\x23\xd2\xb9\x3b\xe8\xc8\x29\xf9\x72\x92\xc4\x21\x59\xd0\x72\x9e\x45\x12\x4c\xa3\x39\x22\x7a\x2b\xcb\x22\x41\xea\x91\x48\x0e\x52\xc2\x60\x4a\xef\x35\x87\x39\xc0\x40\x6c\xa2\x64\x7e\xbe\x64\x73\xaf\x28\xd5\x13\xb1\xad\xa2\xc4\xaf\x8f\xf8\x8f\x5a\xc2\x57\x24\xf1\x77\x6b\x24\x4b\x93\x38\x45\x1a\xca\x62\x49\x25\xb5\x2c\xa7\x61\x0c\xe2\x5e\xd0\x62\x86\x42\x8d\x67\x69\x56\x50\xf2\x02\x66\xe7\xc0\xb2\x98\x32\x0e\x46\x1e\xc8\x4f\x7b\x64\x25\xb7\x23\x80\x8d\xbd\x3c\xf4\xc8\x4a\x6d\x26\x9e\x12\x6f\xfb\x81\x7c\xf9\x42\xca\x55\x4e\xb3\x29\x4c\xdd\x3e\x02\x0e\x65\x93\x4f\x34\x2c\x3b\x30\xc0\xc1\x08\xd9\x5e\x19\x40\x2b\x0b\xa8\x2b\x41\xe4\x36\x57\x03\xcd\xa6\xdb\x81\x5e\xe4\x81\xc4\x29\x2b\x83\x34\xc4\xf9\xc7\x45\x11\xac\xc8\x77\xdf\x01\xa2\xfa\x53\x85\xec\x7e\x1e\x27\x14\xa6\xf9\x09\x4d\x67\xe5\x9c\xbc\x22\x2b\xf9\x51\x41\x10\xf2\xe0\xe7\x59\xee\x49\x26\xd3\x84\x51\x39\x32\xcd\x0a\xe2\xdd\x02\x6a\xf2\x50\x01\x23\x11\xb7\x57\xfb\xd7\x82\xf4\x17\x1d\x5c\x7e\x5b\x40\xad\xba\x15\x18\x21\x11\x4d\x68\x49\xc9\xc3\xd5\xed\xb5\xc0\x5c\xa1\xd3\xd4\xe1\x20\xb0\x94\xb3\xd6\xc3\x2f\xc0\x52\xf8\xd7\x96\xf7\x83\x14\xb7\x52\xab\x92\xc1\x94\x2b\x8e\x14\xcc\x4f\x18\x2e\x6c\xbe\x5c\x32\x12\xce\x83\x14\x84\x04\xff\x46\xa0\xfb\x00\xa0\x84\x45\xb2\x54\x80\x78\xf4\x4e\x5a\x90\x58\xa1\xae\x25\x69\x70\x17\xcf\x82\x32\x2b\xe0\xd1\x19\x3c\xd2\xa4\x7b\xb8\x76\x0c\x10\xfb\x03\x78\xfb\x89\xab\xa7\x60\x24\x7c\x7f\xf9\x52\x61\xe4\x66\x59\x47\x0b\x2c\x02\xf0\xab\xf8\xda\x0f\x96\x80\x9a\x96\xc5\xaa\xe2\x94\x1c\xe1\x4f\x95\x10\x1e\xb7\xc4\xff\xd2\xbc\x83\x28\x1a\x22\xd9\x67\x31\x2b\x69\x4a\x0b\x6f\x47\x20\xde\xe9\x11\xbd\x31\x3e\xb3\x1d\x7e\x3a\x95\x13\x0c\x78\xce\xc0\x39\x4d\x40\xe7\x99\x64\x2e\xe7\x0e\xe7\x6f\x67\x41\x19\x0b\x66\xb4\xd3\xeb\xd0\xa2\xc8\x0a\x78\x07\xe3\x48\xe1\x2d\x4c\x32\x46\x3b\x9c\xff\x38\x25\xc9\x42\x80\x97\x2b\xc3\x97\x00\xf9\x2d\x84\xa3\xfd\x66\x98\x04\x8c\x81\xa1\x15\x34\x2f\x28\xe3\x4b\x04\x84\xc5\xe9\x0c\x74\xf3\x9e\x4e\x98\x70\xb6\xde\x07\xf4\x50\x20\xac\x72\x4e\x09\x7c\x04\xaf\xb3\xc7\xe2\x88\x76\x4d\x41\x5a\x9e\x42\x5b\x9f\xb7\x8d\x5f\x25\x4f\xe1\x23\x7a\x27\xed\xc7\x3a\x03\x0d\xe5\xf5\xff\xcc\x1d\xf7\xbf\x1e\xf6\xfd\x92\xb2\x92\x63\xe9\xda\x13\x61\x0f\xbe\xf6\x7c\x2f\x01\x4f\xbf\x03\x6f\xf8\x74\x9e\xb1\x12\x3e\x02\x58\x03\xa3\x07\x28\x7f\xed\xff\xda\xf7\x5f\x76\x5f\xb8\x50\x97\x73\x70\xc0\xa4\x73\x9a\xde\x05\x49\x1c\x91\x0f\x17\x67\x87\xa4\x63\xe2\x2a\xc1\x4b\xfb\x92\xf2\x7b\x86\x43\x17\x74\x36\x7c\xc8\xfd\x17\x07\x06\x40\x98\xa5\x29\xf8\x0a\xa5\x27\xfc\x19\xb8\x0f\x94\xd8\xe7\x47\xe3\x19\x5b\x4e\x1a\xcf\xb2\x74\x99\x47\x01\x90\x56\x1f\x90\x48\x69\x84\xde\x2d\x00\xfb\x1f\x48\xf5\xbb\xb0\xdc\x2b\x3a\x2d\x9c\xca\x27\x82\x8a\x1d\x56\xae\xf0\x96\xae\x40\xb3\x26\x9f\x50\xbd\xc4\x22\x96\x3d\x48\x77\x07\x50\xc2\x6b\xb0\xb2\x00\xd9\x77\x2a\x1b\xa8\x71\x07\x01\xf7\xc8\x62\x09\xdc\x9e\x80\x26\x08\xe8\x41\x85\x6f\x1b\x96\x32\xdc\x28\x7e\x73\x39\xd2\x06\x5e\x31\x6e\xa0\x0e\x52\xf9\xcc\xc4\xae\x99\x7a\x05\x74\x5c\x37\x90\xbd\x5e\xe6\x10\xea\x60\x8b\x92\xcc\x20\x01\x3d\x8b\x56\xc8\x11\x1a\x69\x3c\x36\x12\x60\x1b\x7c\xb6\xc6\x50\x42\x6a\x6c\xdf\x9e\xa5\xe4\xa4\xa7\xca\x07\x8d\xf9\x61\x11\x4f\xa8\xf6\x19\x3d\xfe\x26\x85\x69\x08\xa7\x12\x05\x9f\xa8\x5d\x90\x8e\x89\xc6\xa0\xed\x85\x7a\xca\x0d\x17\x2b\x27\xbe\x30\xa1\x41\x71\x09\x0a\x92\x2d\x4b\xaf\x9a\xef\xab\x80\x5d\xb1\xf3\x9e\x99\x7c\x44\x9d\x83\xa9\xb0\x2d\xcf\x84\xdc\xe6\x23\x10\x37\x82\x55\x0d\x9a\x3f\x03\x7a\x0f\xf6\x6d\x56\xdd\x33\x99\x3f\xe8\x2c\xc6\x53\x76\xa4\x11\xa3\x8f\x1a\xe3\x53\xdc\x2f\xbc\xa9\xe7\xc2\xd5\xf9\xe0\xda\x87\x41\x38\xaf\xb2\x2d\x43\x75\x01\x08\xcd\x31\x4b\x3b\x2f\x35\x9b\x88\x40\x06\x4b\x5f\x51\x14\x0e\xff\x06\x1f\xfd\x09\x78\x40\x8f\x7f\xd3\x4b\x3f\x76\x2d\x6a\x73\x50\x63\x1f\x33\x1e\x46\xcb\xd3\xb4\x04\x07\x17\x24\x5e\x35\xc2\x11\xf0\xf9\x3d\xf2\xc3\x3e\xd9\xc5\xdd\xee\xdb\xa2\x88\x62\xf6\x2c\xf9\x6a\x6b\x6e\xe3\xbb\x52\x1a\xf1\xd4\x44\xfa\xb9\x2e\x17\x53\x84\x22\x32\x7f\x13\x0e\x6b\x6e\x0a\x36\x23\x4b\xd3\x65\x92\x34\x59\x58\xa9\x92\xcf\x0d\xee\x7d\x89\x36\x88\x76\xaf\x85\xef\x9f\x9c\x8d\xde\x0f\x5f\xd7\x94\xe7\x1e\x77\x0d\x21\xcb\xeb\x3a\x74\xc7\x58\x8a\x6b\xb3\x43\x2e\xa5\xcd\x2b\x08\x5f\xa6\xeb\x03\xc3\x0c\x6c\x67\xd7\x24\xf2\xc8\x22\x72\xf4\x6e\x78\x5e\x67\xa5\xa6\x14\xb1\x0b\x9c\xd6\xa2\x48\x48\xbb\xc8\xf9\xa4\x4e\xce\x7d\xaa\x4d\xab\x72\x12\x4f\xcc\x7d\xfb\x7e\x74\xee\x0b\x3f\x1b\x4f\x57\x5e\x25\x1d\x15\x07\x0f\xc5\xa7\x9e\x1e\x10\x6e\xf3\xa3\xa8\x0b\xd8\x61\xe5\x93\xb4\xdc\x6c\x4a\xb2\x54\xe6\x13\x06\x25\x56\x2a\x26\xd4\x08\x68\x00\xa9\xf0\x01\x1f\xb9\x50\x09\x9f\x78\x7c\x0c\xdd\xbc\xd8\xa8\x56\x42\x63\xba\x74\x9c\x5a\xce\x60\x04\xd5\x5e\xaa\xe8\xc7\xb7\x9b\x07\x05\xe8\x04\x20\xad\xcc\x95\x80\x5b\x07\x45\x85\x7c\xc7\xd4\x50\x29\x21\x30\x3c\x96\x25\xd4\xbf\x0f\x8a\x14\x41\x7a\xc4\x9a\xbb\x55\x25\xcd\xf5\x54\x51\x2e\xec\x4e\x17\x25\xe1\x00\xad\xbc\x7e\x7c\x3d\xb0\xc6\x30\xcc\xc0\xa8\xff\x1f\x74\x65\x0f\x44\xac\x94\x16\x57\x45\x1a\x1b\x82\x15\x21\x9f\xfa\xda\x60\xa5\x30\x69\x1c\x81\x08\xba\x0d\x38\xcc\xa4\x1d\x36\x59\xc6\xa9\x2c\x94\x34\x34\x20\xa0\x09\xa8\xa4\x01\xf8\x89\x65\x69\x8e\xdc\xf2\x83\x3c\x4f\x56\x1e\xe0\xe9\xe1\x72\xdd\x6a\x19\xa3\x92\xc0\x97\x48\xf2\x0d\x38\x73\x05\x19\xc9\x61\xd4\x7f\xc1\x11\x0a\x49\x2b\x55\xe9\x58\x85\x85\x06\xf2\x4c\x2c\xb8\x5f\x15\x29\x35\x5b\xac\x50\x3a\x70\xac\x58\xcd\x68\x5d\x4f\x81\x58\xab\x35\x62\xf8\xd2\x97\xc6\x50\x69\x84\x78\x87\x72\x73\x19\x86\xa0\xfc\xd3\x25\x54\x9c\x6c\x46\x30\x03\x86\x04\x58\xf8\xe8\x30\x5b\xa2\xc7\x31\x2d\xb2\x11\xec\xb4\x09\x61\x02\xde\x6e\xc9\x66\x1a\x67\x06\x76\x9d\xd2\xa8\xcc\x9f\x58\x5f\x3d\x84\xed\x6e\x94\x5b\x64\x29\x77\xa4\x9b\x91\x60\xc5\x9e\xa7\x68\xe0\xc0\x36\x11\x82\x0d\xbb\x47\xe4\xfb\x3a\x92\xaa\x88\x32\xac\xc8\x4c\x3f\x44\x88\xb5\x12\x13\x49\x99\x15\x64\x8d\x74\xc3\x16\x9a\xde\x2e\xaf\x7c\x4c\x8f\x65\x7a\x86\x7e\x5f\xfb\x04\x0e\xe7\x75\xaa\x52\x46\x4e\xfc\x07\xd6\xe9\xe1\x67\xb3\xac\x13\xa5\x91\x68\x98\xc4\xdc\x42\x64\xed\x64\xf4\x4b\xd4\xc7\xc1\xd6\x63\xa3\xfb\xd3\x23\xcb\x34\xa2\x53\xa8\xe6\x22\xf4\xb1\x5b\xfd\xdd\xed\xad\x5d\x57\x3b\x0a\x62\x4f\x21\xd5\x6b\xef\x6d\x10\xde\xf6\xd1\xed\xed\xbd\x43\xa3\x85\x09\x68\xc0\x7b\xdc\x82\xf7\x22\xc8\x68\xe9\x83\x0f\x55\x84\xec\xf3\x1c\x92\x7d\xff\x47\xff\x0f\x00\xe5\x85\x5d\xec\x46\xfd\x40\xde\x66\x10\xc2\xe3\x05\x84\x31\x86\xea\xba\xcb\x3b\x56\xb0\x03\x9a\xc2\x16\x76\xfb\x5b\x3c\x07\x18\xd3\x07\xa8\x37\x23\x95\x07\xf8\xd5\x03\x70\x35\xba\x82\xf3\xa2\x1e\x99\x28\x36\xf2\xa6\x00\x6f\x2a\x61\x63\x00\x1e\xa3\xa7\x9f\xf8\xf3\x80\x8d\xee\xd3\x77\xa2\x1f\xb3\xf2\xf2\x6e\x97\x44\x57\x39\xda\xd9\x04\xde\x64\x51\xae\x10\x8e\xc7\xa8\x87\x5a\x03\xc1\x2d\x2f\x43\xd0\x10\x00\x8e\x06\x52\xa2\xe3\xb1\x55\xc3\x4c\xaa\x6f\x02\x57\x64\x0d\x63\x46\x89\x48\x81\xfd\x03\xbe\xb3\x51\x11\xcf\xe2\x34\x48\x86\x28\x56\x00\xe0\xef\xc0\x7b\x1c\xd3\x9e\x70\x50\xb5\xec\x88\xa7\x9f\xaa\x8d\xf6\x77\xc9\xeb\x0c\x3b\x6f\x73\x88\x5e\xb8\xcb\x45\x16\x2d\xa1\xfc\x85\xac\x49\x55\x10\x52\xaa\x3e\x87\x7f\x9d\x51\x96\xee\x94\x50\x77\x66\xb7\x24\x05\x46\xf7\x48\xc0\xa0\x58\x86\x18\x95\xf2\xfe\x5d\xbc\x40\x37\x99\x2f\x4b\x0e\x7e\xb3\xad\x57\x24\xdf\x7d\x77\x03\xc5\xcd\x14\x5b\x5a\x3c\x2b\x8b\x17\x0b\x1a\xc5\xe8\xe4\x34\x7d\x61\x90\x24\xa2\x84\x4d\xc9\x25\x6c\xfa\x3d\x98\x7b\x5e\x8a\x95\x41\x9a\xc2\xda\x8c\x5d\x08\x2f\xdf\x8c\x8a\x4a\xaf\xb7\x94\x07\x1e\x8b\xc4\x00\x62\x15\x6a\x81\xc1\x10\x73\x2e\xa2\x1e\x71\x38\x1f\xbc\x27\x33\x9d\xad\x0e\xb7\xc6\xb8\xe1\x76\xe5\x60\x85\x36\x33\xf1\x2a\x22\x6e\xc5\xea\x57\x86\xc7\xb7\x74\x2d\xc6\x6d\x37\x66\x2a\xca\xb2\xba\xf6\xc5\x5d\x17\x28\xbe\x70\x21\xd1\xa4\x8c\xbb\x83\x06\xc4\xe3\x56\xfb\x37\xb9\x11\xb1\x3d\x0d\x22\xd9\xd9\xd5\xa1\xa6\xd2\x71\xfa\x97\x25\xf8\x48\x2f\x30\x6c\x07\x5f\xec\x3e\x46\x91\xab\x70\x16\xd4\x29\x0d\x03\x46\xc9\x8e\xf6\x18\x3b\x87\x8e\xd1\x49\x96\x61\xfe\xef\x1c\x13\x49\xa1\x73\x28\x5d\x2e\x26\xb4\xa8\x0d\x19\x7b\x0b\x78\x4c\x9d\x0c\x1c\x53\x85\x92\x38\xa6\x22\xff\xc5\x3c\xcc\xcc\xbb\x4e\xa6\x4b\xf4\x13\x0d\xd6\xe4\x3c\xa2\x19\xc7\x8c\x77\x42\xbd\xa0\x55\x7c\xbc\x6b\xab\xe1\x80\xaf\xe0\xa6\x02\xd5\x29\xc5\xca\x62\xd2\xe8\x95\xb6\x10\x23\x83\x9d\x13\xcc\x50\x3b\x48\x09\x7b\x04\x5b\x3c\x41\x95\x0c\x42\x82\x98\x88\x9c\xb0\x75\x11\x41\xa7\xd2\x00\x48\x10\x41\x09\xe0\xdf\x6e\xfb\x8c\x0d\x49\x53\xc5\x87\x91\x2d\x54\xaf\xc7\xe6\x1c\xdc\xc5\x44\xda\xb6\x61\xe9\xc0\xba\xe6\x74\x0e\x7b\x26\x78\x79\x24\x66\xa9\x4d\xbb\xe5\x65\xe0\x0b\xba\x96\x14\xce\xd6\x08\xe1\x89\x5d\xda\xcc\x17\xec\x96\xf8\xd6\x30\xfd\xf9\x0c\x7f\x82\x0c\x8b\xd1\xd6\x28\x18\x66\xb0\x4c\xca\x56\x2b\xb2\xd2\x28\x9d\xa0\x6c\xa9\x98\xf2\x27\x4a\x96\x60\x51\x01\x79\x1b\xdc\x05\x8c\x7b\x71\x02\xee\x6b\x4e\xca\x0c\xea\x13\x0c\x01\x14\x02\xb7\xe5\x4b\x7c\x82\xe5\xb8\x80\x82\x7c\x02\x92\x50\x4f\x1e\x7c\x80\x83\x07\x64\xa2\x84\xc3\x1e\x2c\x3e\xe4\xed\x5c\x12\x47\x00\x19\x4f\x63\xc8\x0f\x08\x3f\x3d\x81\x8f\x11\x7a\xd1\x62\x1a\xfe\xe3\xbf\xec\x7f\x2f\x22\xc7\xa9\xa8\x63\xc9\x7d\xb0\xea\xc9\x50\x45\x16\x41\xce\x49\x20\x22\x36\x69\x94\x62\x91\x8c\xc4\x90\x0a\x47\x10\x9e\xb0\xa3\x16\x55\xfe\x0e\x50\xd3\xe9\x14\x32\x24\x58\x17\xf1\x59\xa1\x09\xf6\x7d\x69\x92\x07\x21\x34\xe7\x4d\xb9\x09\x04\x54\xc0\xa9\xdb\x7a\x6a\x02\xaf\x0d\x26\x9f\x46\x39\xd3\x0d\x4c\x7c\x59\x4d\x4c\x70\xfc\xd8\xc0\x04\x8f\x5c\x77\x17\xf0\x5c\x65\xf9\x3c\xc7\xb8\x0b\x92\xba\xbd\x38\x0d\xe9\xb1\x67\xc4\xad\x45\x76\x47\x37\x59\x4c\x1e\x9b\xa8\x35\x9f\xbb\x4c\x9e\x04\xe1\x46\xeb\x7c\x83\x4d\xb5\x6d\x09\x32\xeb\x82\x52\x57\x60\x2e\xe9\x22\x47\x09\x80\xe8\x0e\x49\x67\x3c\xa3\x25\xe4\xc7\xa0\x18\x73\xd9\x40\x98\x16\xd9\x42\x45\x41\x2d\x24\x5e\xe9\x21\xc6\x1e\xb9\x42\x04\xd7\xdd\x35\x00\x0d\x33\x92\x6b\x09\x01\xb8\x56\xb3\x66\x7c\x15\x6e\x50\x23\x1b\x31\x7e\xec\x11\xce\xd1\x43\xbe\x67\xc1\xdd\x27\xd6\x7a\x8a\xdd\x61\x96\xaf\xfe\x5e\xd8\xfd\x37\x62\x09\x9e\xb3\x6c\xa2\xec\x12\x8f\xf2\xe8\x4a\xf7\x7b\x86\xea\x77\x9d\x0b\x20\xcf\x36\x59\xa0\xc2\x23\x5a\xfd\x35\xdb\x95\x5e\x7b\xb0\xb5\xa9\xfb\x0a\x30\x23\xf1\xc9\x2f\x41\xba\x82\xcf\x54\x9c\x88\x05\x50\x1b\x42\x0d\x80\x21\x8d\x7b\xe7\x86\x8b\x83\x59\x4f\xb8\xb8\x00\xfb\x56\x71\x9d\x7a\x78\xea\x33\xa4\x80\x7a\x71\x0f\x53\x14\x27\x5b\x36\x11\x48\xd3\xcf\x6d\xb2\xe2\xc1\xb3\x57\x69\xb8\xb9\xd6\x65\x20\x6c\xff\x66\x17\x27\x42\x87\x8f\x5f\xea\xe6\x28\x87\xf0\x4b\x5d\x2d\xe5\x10\x7e\xa9\x2b\x94\x1c\xc2\x2f\xcf\xd2\x0c\x29\xf2\x22\xcb\xca\x67\x2b\x07\x4e\x7a\x3a\x00\x36\x2c\x47\xcc\xf2\x85\x60\x7d\xac\x1b\x79\xfb\x84\x1f\xf7\x75\x5b\xea\x2b\xec\x58\xc6\xa9\xa9\x45\x2d\xa5\x56\x05\x51\xaf\xb9\xd0\xc6\xda\xd2\x76\x67\xe8\x72\x44\x4b\xce\xd9\xa7\x6b\xb0\xe7\x05\xed\x3a\x51\xf5\x4d\x3b\x40\xd4\x6e\x61\xe8\x99\xdb\x6c\xe5\xbc\x70\x42\xdf\x7a\xbb\xae\xe4\xa1\x61\x52\xdc\xfd\x73\x12\x9e\x15\x6c\xf1\xe3\x46\xc1\x76\x3d\xea\x27\x02\x4b\xe5\x86\x7f\x5b\x60\xf9\x2d\x76\xbf\xde\x9e\xc4\xc2\xf5\x13\x17\x0e\x89\xf5\x6c\xed\xb9\x61\x42\x1b\x07\xa8\xa7\x62\x53\x5b\x58\xe2\xed\x1b\x59\x0c\x0b\x18\x54\x5a\xfe\xd5\x97\x8f\xad\x83\x3b\xf9\x0c\x90\x5a\x30\xb2\x87\xc1\xdb\xff\x6e\xf0\x4d\x18\x84\xb6\x82\x8d\x15\xbc\x3b\x53\x1d\xf3\xab\xaa\x10\x39\xa5\xfa\x0f\xcd\xf6\x89\xf4\xa5\xfd\x1f\x1e\xa0\x7a\xc2\x76\x25\xb0\x20\xc0\x9e\x57\x38\x87\x12\xa3\xff\xe7\x5f\xa3\x97\xea\x8e\x06\x1e\xe1\x58\xc5\x11\x89\x19\x1e\x01\xce\x68\xc1\xc7\x6a\xc7\x34\xb1\x79\x58\x2f\xaf\xbf\xd0\x14\xdb\xcd\x65\xd1\x28\x6e\x71\x34\x9c\x07\xc5\x49\x16\x19\x4a\x26\x2f\x64\xf1\xb2\x9f\xa6\x8d\x76\x8d\x84\x97\x28\xd5\xd7\xe3\xb2\xd1\x60\x42\xd9\x68\xe8\x57\x47\xe4\x0f\xff\x8c\xac\xd2\x4f\x7e\x3a\x22\x3f\xfe\x93\xd3\x11\xbd\x7c\xd9\x74\x1b\xd5\xb1\x8f\xf9\xd4\xe9\x3e\x9c\xf5\xa8\x31\x5e\x99\x94\x96\x43\x9f\x1c\xf3\x23\x9d\xc0\x68\x3a\x1b\xb5\x25\xfc\x57\x15\x6c\x98\x57\xda\x02\x31\xf3\x4d\x3e\x95\x32\x6e\xef\xb1\x7d\xcd\x44\xdd\x39\x64\x50\x4c\xab\xd3\x07\x98\x20\x5a\x2e\xb9\x90\x92\x9c\x2e\x25\x25\xd1\x71\x57\xda\x90\x4f\x0e\xf2\xc9\x1d\x02\x12\xc4\x6b\x54\xba\x19\xad\xc7\x91\xbd\xd6\x13\xd8\xfe\x9b\x38\x8d\x8c\xa8\xdc\xc8\xd3\xb9\x6f\x94\x38\x85\xa3\xfc\xf2\x85\x74\x3a\x83\x06\xa0\x6c\x6c\x22\x08\x4f\x9f\x4a\x6f\xa7\xbf\xd3\x6d\xc2\xe1\x8d\x98\x23\xce\xcb\x1a\x31\xbc\x2a\xc0\xb3\xa5\xe6\x1c\xc1\xa2\xdb\xb6\x36\x0d\xbf\x1d\xf6\x10\x33\xd0\x93\xd9\x3b\x58\xff\x4d\x11\xcc\xf0\x38\x02\x8f\xbe\x54\x77\xb1\xb6\x98\x64\x25\x3f\x5f\x72\x68\xa2\x38\xd7\xc4\x05\xaf\xca\x6b\x47\xcf\x04\x35\xdc\x25\xe7\x3a\x8c\x9b\xaa\x23\x83\xae\xb6\xd9\x0a\x43\x95\x50\x6c\x3a\x0b\x5f\x2d\xdc\xe0\x1c\x64\x3c\xb5\xc5\x4c\xba\xeb\x7f\xca\xe2\xd4\x21\x27\xf3\x25\xbd\x25\x4f\x89\x80\x08\x2e\x8b\x3d\x48\x8b\xbf\x8e\x82\x4a\x8f\xd6\xac\xb8\x96\x21\x4e\xc4\xdb\xcf\x61\x8e\x0a\x3a\x28\xbd\xac\xf0\xa4\xa5\xe5\xb8\x2b\x51\xaf\xf6\x9c\xd4\xaf\xe3\x91\x73\xa4\xf9\xd4\xd1\xb6\x2c\x5d\x3e\x8f\x5f\x64\xa5\xab\xcd\x64\x2e\x24\xf3\xea\xc8\xe5\xb3\xcd\x97\xf6\x3f\x32\x6f\xbe\x12\xb2\xc8\xf2\x6b\x91\xc0\x49\x3e\xd4\x4a\xf7\x01\xd1\x3e\x92\x03\xb4\xa2\x9f\x14\x34\xb8\x75\xb3\xc8\xc5\x08\xd7\x2e\x74\xbb\x1b\x63\xef\xba\xed\x2a\xde\xec\xec\xed\xac\xdb\xb1\xb0\xe2\x2a\x36\xb7\x50\x57\x4f\x07\x5c\x2b\x2a\x63\xe7\x57\x8a\xab\x38\xbc\x2e\x47\x56\x2f\x71\x37\x0f\xcf\xed\xde\x42\x88\xe1\x67\x9c\x43\x71\x26\x3b\x7c\xc8\xc5\x89\x34\xc4\x98\x65\xca\xe2\x19\x48\x99\x4c\x02\x46\xf7\x0e\xf6\x21\x5f\xe7\x4b\x88\xf4\xb1\x47\x16\xc1\x2d\x1e\xce\xa1\xa7\x46\x54\x05\x9d\xd2\x82\xa6\x21\x4c\x10\xe9\x13\x0e\xf0\xca\x1c\x76\x43\xb9\x4d\xdc\xc7\xe0\xac\xf1\xf1\xff\xd0\x22\xdb\x43\xb4\xd8\x15\x8d\xe8\x03\x24\xa9\x9d\xd1\xbb\xe1\xc5\xf1\xe5\xe9\xe8\x7c\xfc\xee\xf8\xf2\xdf\xc7\xa7\x67\x67\xc3\x7f\x3b\x3e\x1b\x1f\x5f\x5c\x1c\xff\xf7\xf8\xf4\xfc\xf5\xf0\xbf\x3a\xda\x24\x2a\x9b\x95\x9f\x9f\x6d\x09\x95\x30\xf8\xcd\x16\xe0\x9e\xb8\x96\x79\xb0\xdf\x82\xca\x8d\x66\x53\x4d\xaf\x0b\x4c\x29\xba\xb8\x48\x81\x69\x3a\x3e\x45\x82\x5e\x19\xda\xf1\xf5\x72\xc4\xba\xd8\xec\x3c\x03\x8f\xc9\x2f\x1f\xde\x5f\x92\xf3\xd1\x25\x5e\xe5\x9c\x81\x69\xe8\x54\x8f\x4b\x90\xa7\x87\x04\x12\x47\x29\x2c\x26\x6a\x52\x29\x42\x5b\x40\x1f\x8f\xcf\x3e\x0c\xc7\xa3\x0f\x97\xe3\xd1\x9b\xf1\xcf\xa3\x0f\xe7\xaf\xdf\x7f\x6b\xd9\x68\xd7\x20\xba\x34\xff\xaf\x9e\x61\x9d\x35\x2a\xdb\x17\xd2\xf3\x39\xa7\x47\x53\x6f\xe7\xaf\xe0\x06\xb6\x8f\xc8\xde\x41\xfb\x21\x88\x8e\xe8\xbe\xac\x25\xbd\xfe\x5f\x0f\xfa\xb3\x1e\xc1\xe8\x67\x3c\xdb\xe7\xcf\x00\x21\xdf\x1e\x65\x61\x90\x53\x9e\xba\xb2\x6f\xe3\x7c\x45\x71\xf6\x7b\xf3\xbd\xfc\xda\xb2\x9b\x7e\xf3\x6e\x19\x78\x20\xaf\xc3\x92\x80\x41\x7a\x29\x1b\x0a\x98\x18\xf3\xd4\xcf\xf8\xe1\x0e\x8d\x3a\x4e\x35\x6c\x21\xb3\x49\x90\xc8\x11\xdd\xe7\x0c\x8f\xed\xf9\xbd\x60\xb2\xca\xf0\xf1\xdf\xda\x25\x01\xd4\x70\x7c\x37\xaa\x4a\xdb\x96\xed\x7b\x01\x63\xd8\x0f\xb5\x0a\x2d\x7d\x6b\xc4\xb3\xe7\xf5\x88\x84\xad\x68\xd5\x68\x6a\xde\x42\xde\x35\xec\x91\x34\x58\xc0\xbf\x5c\x87\x7b\x55\xc9\xe1\xee\x98\x0b\xec\x66\xa3\x45\xa2\xa9\xf1\x99\xa7\x36\x72\x88\xff\x78\x85\x7f\x72\xc0\xe0\xe2\xfc\x27\x25\x0b\xd7\xa8\x70\x61\x47\x82\x38\xc7\x78\x55\x20\x1d\x55\x94\x3b\xe0\x70\x2b\x3a\xd9\x6f\x15\x9b\xcd\x20\x29\xbe\xae\x67\x5d\x6c\x91\xfb\xac\xe4\xd9\x90\x5b\x1d\x4b\x0d\x7e\x2d\x58\x7f\x77\x97\xbf\xef\x92\x0b\x1a\x2e\x0b\x16\xdf\x51\x50\x16\x28\xa2\xc2\x5b\x06\xa5\x02\x05\xcf\x5c\x18\x85\xe0\x3c\x60\x04\x1b\x9c\x3a\x35\x13\x31\x18\x4d\x02\x7f\x13\x62\x9d\x07\x6a\x35\x80\x49\x1f\x14\x7c\xbd\xb9\x20\x13\xfd\xf5\xf9\x9e\xbb\x2f\xb4\x65\x21\x31\x7e\x75\xc0\x43\x1d\x8f\x26\xb5\x9f\x23\x98\xbf\x46\x68\x6b\x18\xc6\x4f\xb4\x0b\xeb\x9b\xe1\x07\xd0\x2d\xee\xef\x89\x83\x7c\xf7\xb7\xc7\xe6\xbd\x9a\xaa\xb8\x7f\xac\x4b\xed\xa3\x0c\xf6\xc6\x4f\x77\xb4\x5a\xfa\xe4\x04\xac\x06\xcf\x6f\xf1\x44\xe9\xa6\xd2\x08\x95\x21\xdc\xf8\xe4\x12\x43\x3b\x23\x37\xb6\x72\xdc\x20\x13\xf8\xa5\x10\xbc\xbd\x92\x8a\x8b\x73\x52\xb8\xe4\x8f\x90\xc9\x04\x0b\xf2\x59\x70\xf2\xd1\xe8\x19\xec\x99\xfd\x03\xa1\x30\xc2\xbf\x77\x6b\x53\x45\x16\xf0\x28\x13\x86\x3d\xf9\x8e\x52\xaa\xce\xb6\xd5\x0f\x8e\xfe\xb2\xc4\x64\xaf\x65\xed\x2b\x34\xb0\x6b\x5c\x58\x2c\x07\x1a\x2b\x5b\xef\x06\x26\x26\xbc\x33\x66\x81\x65\xc6\x7f\x60\x82\xed\x7b\x1a\xd5\x70\x8a\xfe\x1e\xe0\x74\x95\x40\xb8\x46\x98\x2d\x90\xcf\x49\x06\xd9\x28\x4f\x31\x6f\x70\xf5\x1b\xa7\xce\x57\x45\x96\xe1\xdf\xa4\xc7\x5b\x53\x68\xd5\x2c\x43\xe9\xae\xde\x0b\x56\x7b\xea\x9e\x0e\x2a\xb7\xe1\x8e\xe4\xb5\x1b\x7c\x5a\x55\x13\x6a\xb8\xa1\xa2\xad\x29\xdd\xce\xc8\x64\x1c\x86\x36\x6d\xfd\x3b\x90\x1e\x54\xc9\x19\x64\x78\xe3\xe3\xf3\xf1\xe8\xe7\xb7\xc3\x93\xcb\x9d\x56\x77\x6e\x18\x6c\x55\x4e\x6f\xcb\x5c\xa0\x52\x55\xc8\x07\xbe\x86\xc4\x9b\x2c\xbf\x21\xea\x9e\x84\x22\x38\x4b\xa9\xa5\x4b\x4c\xdd\xe1\x43\xa5\xba\x78\x73\xb2\x87\x77\x23\xec\xcd\x8c\xde\x41\xea\x0f\xe9\xe6\xe9\xeb\x67\xee\xa4\x2e\x21\xd1\x28\xe2\x62\x92\x97\xb4\xbe\x6a\x5b\x88\xa5\xb9\xb1\x40\xfe\xea\xca\xa6\x5d\x94\x31\x5f\x45\xbd\x67\x0a\x40\xd4\x95\xd8\x56\xb7\x55\x4b\x0f\x61\x5b\x1d\xb6\x63\xf4\x7e\x35\x08\xf7\x30\xbf\x79\xd7\x88\xa5\xb9\x6b\xf9\xe3\x44\xe2\x19\xa7\x6e\x20\xc8\x1b\xa4\xf4\x06\xb4\x33\x22\x37\x48\xd9\x8d\x21\xf0\xae\xcd\xa1\x37\x17\xa3\x5f\xc6\x17\xc3\xff\xfc\x70\x7a\x31\xfc\x06\x2c\x82\x32\xaa\x85\x43\x32\x99\x6e\x19\xc5\xde\xb6\xe0\x5f\x35\x26\x4f\x00\xd6\x45\xbf\x8d\x58\xc7\xf1\x6c\xce\x3b\xd8\xc2\x4d\x8f\xdc\x48\x7a\x25\x17\x91\xbe\x35\x5c\x14\xd5\xd8\xef\x89\x8d\x76\x2c\xb6\x79\xfa\x55\x1e\xef\x6f\xc7\xc6\x93\xe3\x73\x74\xa0\x27\xa3\xf3\xcb\xe3\xd3\xf3\x31\x94\xb6\xc3\x37\xa7\xe7\xcf\x66\xab\x2b\x63\xe6\x29\x95\xcd\x35\x51\xfa\xbb\xd2\x14\xd5\xe8\x3e\xa3\x56\x46\xeb\x1b\xcd\xec\x4e\xbf\xd3\x6d\x6d\x23\xd5\x1b\xd0\x02\x8f\x2b\xb6\x6d\x82\x0d\x49\x57\xd4\xa0\x37\xa9\x63\x7e\x49\x0e\x64\x57\xa3\x0d\xa4\x2d\x15\x6b\x97\xfe\x89\xb8\xb9\x0d\x3b\x87\x1c\x70\x81\xd1\x8e\x8b\xd5\x08\xad\x41\xc9\x33\x8a\x88\xb2\xb8\x80\x20\x82\xab\x3b\x3c\xb0\x94\xe8\xf1\xeb\x8d\x64\xa8\x65\x69\x67\x81\x95\x64\x9f\x6b\x15\xe2\x74\xb7\x65\x10\xcf\x23\x9d\xfd\x42\x7b\x25\x1d\xbb\x9e\xca\x4d\xbe\x8a\xb3\x76\x56\x16\x60\x38\xe3\x0b\x96\x73\xf8\xac\xff\x02\x06\x5f\xda\xc1\xde\x0f\xe7\x17\xc3\xf7\xa3\xb3\x8f\xc7\x3f\x9f\x0d\xff\x8f\x18\xfc\x74\xe8\x73\xdb\x8f\xe2\xd6\x47\x79\x9a\xeb\xba\xd9\x65\x87\x4a\x7d\x34\x5e\xd5\x51\x8f\x2d\xb6\x25\x8b\xb7\x66\xf2\xee\x5d\x59\xeb\x5e\xb7\x72\x80\x1f\x1d\x70\x34\x60\x3a\x22\x93\x17\x35\x30\x6e\x6c\x1d\x97\xbf\xb1\xc4\x79\x8e\xb0\xa9\xcc\x79\xc8\xfe\x46\x32\xaf\xd5\xdf\xf5\xfe\x88\x4e\xd5\x81\xcd\xfa\x73\xa3\x3c\xb6\x0a\x2d\x59\x94\xd8\x59\xa6\x4f\x4e\xa7\xb2\x24\x20\xbc\xa4\xa0\xd8\xf9\x84\x08\x02\xd1\xe4\x0e\x8a\xe3\xa8\x67\x55\x34\xfc\x37\x26\x51\x14\xe3\x5c\x28\xd3\x56\xfa\xf4\x34\x22\xc1\x2c\xc0\xbf\x1c\x62\x5e\x19\x42\xb4\xaa\x00\x83\x65\x84\x3c\x01\x03\x60\x12\x3f\x26\x42\xf4\xa2\x5e\x44\x0a\x6b\x3d\x02\xe3\x84\x53\xd7\x3b\xee\xca\x4a\x9f\xf4\xc2\x13\x85\xee\xb3\x8d\xed\x8b\x56\xda\xc7\x75\xa5\x0f\xf5\xd4\x0a\xcd\xc6\x0e\x5e\x6f\x76\x44\xad\xea\xd6\xbf\x9a\xea\x2c\xad\xdb\xf5\x8f\x7f\xae\x38\x6c\xfc\x41\x01\xde\x0f\x40\x15\x7b\x0f\xf9\xcb\xf0\xfc\x64\xa8\xea\x17\xde\xf7\xaf\x9f\x02\xd6\x4e\xab\xda\x22\xad\xd8\x0a\x6f\xf4\x18\x3f\x5e\xad\x5f\x16\xc1\x89\xbc\xa1\x19\x26\x58\x99\xf0\x09\x2c\x13\x26\x20\xaf\x63\xb3\x60\x8a\x6d\x17\x64\x0b\xef\xd2\xe1\x79\x47\xa5\x58\x8d\x35\x39\x8c\xd9\x12\x13\x15\xa5\xc9\xef\xc6\x1d\xc1\xd6\xee\xb2\xeb\x2e\xbe\xc2\xd4\xf6\xe3\x58\x5b\x1a\xd6\x99\xa2\x9a\xca\xaf\xe7\x3f\xe7\xf7\x37\xf2\x77\xbe\xc4\x73\x26\x34\xd4\xfc\x5b\x3a\xb6\xd8\x5d\x84\xc9\xce\x49\xfd\xea\x44\x1b\x0f\x84\x46\xb5\x5e\xb4\x58\xeb\x37\x68\xe5\x36\xf0\x97\x81\xdd\xea\xe7\x59\x18\x4a\x8c\x6f\xf8\xf7\x3d\xf8\x6f\x03\x8d\xba\x91\x3e\x60\xd7\x98\x89\x3f\x94\xa1\xed\x4a\x67\x6a\x72\x58\xb7\x70\x6b\x4d\xdd\x81\x05\x64\xd0\xd3\x24\xd2\x09\xea\x8e\x2c\xaa\x11\xa9\x80\x1b\xfd\xc6\xb6\x56\xa4\x3d\xad\x09\x2d\x81\x1e\xb7\xfe\x37\x00\x00\xff\xff\x59\xbc\x34\x20\xdb\x4b\x00\x00")

func realtimeJsBytes() ([]byte, error) {
	return bindataRead(
		_realtimeJs,
		"realtime.js",
	)
}

func realtimeJs() (*asset, error) {
	bytes, err := realtimeJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "realtime.js", size: 19419, mode: os.FileMode(420), modTime: time.Unix(1441426880, 0)}
	a := &asset{bytes: bytes, info:  info}
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
	if (err != nil) {
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
	"realtime.js": realtimeJs,
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
	Func func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"realtime.js": &bintree{realtimeJs, map[string]*bintree{
	}},
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
        err = os.MkdirAll(_filePath(dir, path.Dir(name)), os.FileMode(0755))
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
                err = RestoreAssets(dir, path.Join(name, child))
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
