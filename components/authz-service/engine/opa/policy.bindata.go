// Code generated by go-bindata. DO NOT EDIT.
// sources:
// policy/authz.rego (1.074kB)
// policy/authz_v2.rego (2.229kB)
// policy/common.rego (2.009kB)
// policy/introspection.rego (866B)
// policy/introspection_v2.rego (1.87kB)

package opa

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
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
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes  []byte
	info   os.FileInfo
	digest [sha256.Size]byte
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

var _policyAuthzRego = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x53\xc1\x6e\xdb\x3a\x10\x3c\x9b\x5f\x31\xb0\x2f\x2f\x0f\xb2\x82\xa6\x2d\x02\x04\xf0\xa1\x40\xfe\xa0\xbd\x05\x81\xb1\x11\x57\x16\x6b\x9a\x2b\x90\x94\x5b\xa7\xc8\xbf\x17\xa4\x44\xc3\x76\x0c\xe4\xd2\x9b\x34\xb3\x9c\x19\x0e\xc9\x9e\x9a\x2d\x6d\x18\x34\xc4\xee\x55\x29\xb3\xeb\xc5\x47\x68\x8a\x54\x37\xb2\xdb\x89\x3b\x83\x7a\xb1\xa6\x31\x1c\xb0\x80\x0c\x1e\xe5\x57\x29\xcd\x2d\x0d\x36\x66\x19\xf1\xe6\x95\x35\x56\x68\xc9\x06\x56\x6a\x81\xef\xc3\xcb\x4f\x6e\x62\xa8\x30\x38\x6b\xb6\x0c\xcf\x41\x06\xdf\x70\x00\x39\x0d\x6a\xa2\x11\x17\x2a\x90\x67\xf4\x5e\xf6\x46\xb3\x06\x25\x12\xe4\x3d\x1d\x6a\xb5\xc0\x8f\xce\x04\x34\x1d\x37\xdb\x80\xd8\x51\x04\x45\x58\xa6\x10\x21\x8e\x21\x2d\x62\xc7\x30\xae\x1f\x22\xc2\x64\x87\x1d\xc5\xa6\x4b\x26\x05\x82\x71\x79\x2e\x07\x3f\xd4\xaa\xa3\xb0\x9e\xa8\xa7\x5e\xec\xda\xe8\x67\xfc\x51\xb3\xb2\xaf\x82\xd5\x45\xf1\x69\xfd\x8c\x55\x5a\x9d\x56\xa9\x59\xb6\xbb\x20\x33\x36\xd2\x63\x81\x85\x5f\x4f\x69\xfe\x3b\x4e\x54\x45\xe9\x46\xbd\xa9\x9c\xa5\xf4\x72\x1a\x66\x52\x29\xd4\xb9\xcc\x11\xae\xf0\x2e\x74\xa1\x8e\xea\x63\xcf\xa7\xda\x23\x72\x21\x39\x82\x57\x04\x47\x22\xcb\xbd\x5b\x58\x21\x44\xf1\xac\x6f\x92\xac\x71\x58\x4d\xff\x1f\x0f\x8f\x9f\x58\x61\xfe\xff\x3c\x4f\x5b\x2b\xbf\xae\x9e\x02\xb7\x6d\x3a\xc3\x15\xe6\x79\x66\xae\x66\x57\xce\x6f\x04\x2f\x8b\x1c\xd1\xf3\x02\x92\x99\x66\x77\xf8\xc0\x2b\x8d\xfc\x0b\xab\xf3\x27\x93\xef\xfa\xa3\x09\xe2\x35\x7b\x7b\xa8\xb0\x63\x72\xc6\x6d\xf0\x98\x12\xc9\x9e\xbd\x37\x9a\x03\xbe\xa5\x9d\xaa\x05\x5a\xf1\x68\x07\x1f\x3b\xf6\xe0\xdf\xbd\x25\x47\xc9\xe0\x01\x5d\x8c\x7d\x78\xb8\xbd\x7d\xb1\xb2\xa9\xa5\x67\x37\xde\x6d\xda\xb0\x8b\xb5\xf8\xcd\xed\xe4\xb0\xdc\xb3\x0f\x43\x58\xea\xa3\xe7\xb2\x64\x59\xde\x7f\xba\xff\x72\xff\xb5\xb9\xfb\xac\xef\x5a\x75\xf2\x84\xd3\x0d\xc9\xfe\x33\x27\x11\xa9\x08\xf5\xa6\xfe\x06\x00\x00\xff\xff\xe1\x34\x20\xc8\x32\x04\x00\x00")

func policyAuthzRegoBytes() ([]byte, error) {
	return bindataRead(
		_policyAuthzRego,
		"policy/authz.rego",
	)
}

func policyAuthzRego() (*asset, error) {
	bytes, err := policyAuthzRegoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "policy/authz.rego", size: 1074, mode: os.FileMode(0644), modTime: time.Unix(1570679371, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x83, 0x72, 0x37, 0x57, 0x0, 0x46, 0xe3, 0x48, 0xef, 0x17, 0xae, 0xe9, 0xa3, 0x1b, 0xdd, 0x62, 0xfa, 0xab, 0xe9, 0x3a, 0x8e, 0xa0, 0xc0, 0x89, 0xa8, 0x69, 0x18, 0xf5, 0xf, 0xf, 0xfb, 0x73}}
	return a, nil
}

var _policyAuthz_v2Rego = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x55\xc1\x6e\xa3\x30\x10\x3d\xc7\x5f\x31\xcb\x29\x54\x28\x87\x3d\x56\xe2\x4b\x10\xb2\x1c\x70\x36\xee\x1a\x8c\x6c\x93\xaa\xad\xf2\xef\x2b\x7b\x6c\x07\x1a\x52\x82\xb4\x27\x98\xe7\x37\xf3\x9e\x87\xb1\x19\x58\xf3\x97\xfd\xe1\xc0\x46\x7b\xfe\xa4\x97\xdf\x84\x88\x6e\x50\xda\x42\xcb\x2c\x3b\x34\xaa\xeb\x54\x3f\x83\x06\x25\x45\x23\xb8\x99\x81\x5a\x49\x6e\x08\x69\xf9\x89\x8d\xd2\xfa\x62\x4a\x8b\x4f\xde\x42\x09\x27\x26\x0d\x27\xe4\xcc\x0c\xed\x78\x77\xe4\xba\x1a\x94\xa4\xa2\xad\xe1\x8b\xec\xdc\xab\x19\x8f\xf0\x5a\x42\x2c\x1c\x97\x0f\xc8\x36\x15\xad\xc9\x4e\xf4\xc3\x68\x23\xd3\x07\x07\x33\x1e\xdf\x78\x63\x71\x1d\x8d\x46\x8c\x76\xcc\x36\x67\x6e\xf6\x29\xad\x80\xa0\x94\x93\x2b\x7a\xd1\xdc\xa8\x51\x37\xbc\x0a\x7a\x05\x18\xcb\x2c\xef\x78\x6f\x9d\xba\x77\x77\x43\x22\x7b\xd1\x68\xa2\x99\x6a\x56\xe3\x10\xb3\x66\x1e\x23\x38\x37\x99\xe0\xa9\x8f\x88\x79\xd3\xbd\xa2\xef\x42\xb6\x0d\xd3\xed\x9e\xe5\xce\x5e\xa3\x7a\xcb\x44\x6f\xf6\xac\x80\xec\x25\xcb\xa1\x8c\xdd\xbe\x12\xc2\x1a\x2b\x54\x3f\x11\x71\x85\x95\xe6\xad\x4f\x9d\x16\x0b\xb0\x6b\xb2\xab\x80\xe1\x4a\x89\x12\xa6\x8b\x7b\x33\x48\x61\x43\xa1\x02\xb2\xd7\x2c\x2f\x00\x31\x97\xe4\xe2\x7c\x5e\x6e\x5f\x19\xae\x2f\xc2\x6d\x37\x7b\xc9\xea\x02\x6e\x31\x2d\x80\xd6\x4e\xc1\xea\x91\x3f\xcc\xb2\x1f\xc3\x42\x2e\xa2\x4f\xa4\x67\x2f\x59\x01\x17\xae\x8f\xdf\xa5\x3d\xf6\x28\x7d\x96\x45\x9f\xe2\xd7\x05\xd0\xdb\xb2\x1b\x3c\xa4\x3c\x37\x76\xc8\xdd\x36\x74\x98\x83\x23\x77\xf7\x05\xdd\xa4\x21\x38\x15\x46\x24\x1d\x8d\x75\x87\x1b\x8e\x80\x92\x1c\x4a\x70\x0f\x2a\x5a\xb2\xf3\x37\x45\x15\xc2\xa9\xd9\x48\x42\x64\xc5\xfa\x84\x99\x5c\x0f\x5a\xb9\xb3\x5f\x55\xe1\x05\x8f\xfc\x03\xff\x5a\xbd\x6d\xeb\x6a\x28\x8a\x6d\x8d\x81\x2f\x81\xef\xc9\xa8\x8b\xf3\x44\x99\x30\x7c\xea\x95\x90\xc5\x04\xb8\xd1\x92\xbf\xb2\x84\x70\x65\x34\xaa\x37\x96\x32\x29\xe3\x26\xcd\xdc\x03\x36\x27\xad\x6d\x11\xf9\xb5\x2a\x02\xdf\xeb\xdf\xb5\xc0\x77\xdc\xef\xcd\xcb\x55\x15\x3f\x9d\x56\x3e\x00\x32\xb6\x7d\x02\xcc\x21\xbb\xfb\x3f\x09\x62\x6b\x37\x3a\xb2\x7e\x1e\x6e\x7f\xed\x49\xa9\xde\x9d\xc9\xb0\x9b\xcc\x03\x59\xb8\x99\x3c\xa5\xe5\xfd\xc7\x94\xe1\xe2\x19\x61\xf2\x07\xfc\x22\x3b\x5f\xc0\x5d\xba\x16\x1c\x33\x89\xf0\x36\x4d\x6d\x78\xd6\x4b\xba\x3f\x6d\xe7\xd9\xa9\x0f\xae\xc5\x8a\x64\xd8\xc8\xff\x52\xbc\xb5\x61\x51\xf5\x51\x13\x52\xab\x96\xec\x92\x2b\xf9\x17\x00\x00\xff\xff\xb4\xf9\x2e\x84\xb5\x08\x00\x00")

func policyAuthz_v2RegoBytes() ([]byte, error) {
	return bindataRead(
		_policyAuthz_v2Rego,
		"policy/authz_v2.rego",
	)
}

func policyAuthz_v2Rego() (*asset, error) {
	bytes, err := policyAuthz_v2RegoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "policy/authz_v2.rego", size: 2229, mode: os.FileMode(0644), modTime: time.Unix(1560797954, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xf6, 0x0, 0x7f, 0xa6, 0x28, 0xae, 0xee, 0x29, 0x43, 0x1f, 0x1d, 0xc2, 0x1, 0x3a, 0x7a, 0x58, 0x6c, 0xfc, 0xd, 0x56, 0x82, 0x38, 0x6d, 0xd4, 0xdf, 0x1a, 0x55, 0xb2, 0x93, 0x93, 0x63, 0xe7}}
	return a, nil
}

var _policyCommonRego = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x54\x4f\x6f\x1b\xb7\x13\x3d\x9b\x9f\x62\x7e\xbb\x01\x24\x19\x6b\xe9\xd7\x5c\x8a\x2e\xaa\x83\x91\xf4\x52\x04\x69\x90\x18\xbd\x04\x81\x30\xcb\x1d\x69\x59\xaf\x48\x82\x33\x94\xac\x1a\xf6\x67\x2f\xc8\xd5\xca\x7f\xd2\x22\x4e\xab\x93\xc8\x9d\xf7\x1e\xe7\xcd\x23\x4b\xf8\x80\xfa\x1a\x37\x04\xda\x6d\xb7\xce\x82\x76\x56\xd0\x58\x86\x75\xb4\x5a\x8c\xb3\x0c\x68\x5b\x08\xb1\x27\x06\xe9\x50\x00\x03\x01\x77\x18\xa8\x85\x86\x64\x4f\x64\x61\xf7\x43\x2e\xda\xbd\x56\x25\xb8\x35\xb8\x18\x20\xd0\xc6\x41\xef\x36\x46\x2b\xff\x44\x41\x29\xed\x2c\xcb\x0a\xfb\x7e\xe5\x83\xfb\x83\xb4\x30\x2c\xa1\xb8\xbf\xbf\x7c\xf7\xee\xe2\xc3\xc7\xdf\x7e\xfd\xe5\xcd\xd5\xa7\xfb\xfb\x42\xa9\x52\x95\xf0\x3b\x06\x83\x4d\x4f\x40\x37\x1e\x2d\x1b\x67\x55\xa9\x94\x75\xab\xdd\xf1\x03\x4f\x71\x06\xb7\xea\x6c\x3c\xf8\x14\x2b\x28\x5e\xdd\x16\x33\x58\x2e\x61\x8d\x3d\x93\xba\x53\xea\x79\xb5\xb1\x2d\xdd\xb8\xf5\x43\xf1\xcf\xf0\x78\xeb\xae\x98\x25\x54\x09\x6f\x69\x6d\x6c\xee\x9c\x4e\x8e\xc0\x24\x9f\xa5\x9d\xc0\xbe\x33\xba\x83\x40\x12\x83\x65\x30\xc2\xb0\xc3\x3e\x12\xec\x0c\x66\x84\x8b\xe2\xa3\xc0\x28\xae\xca\x11\x4a\xed\x64\xae\x4a\x78\xef\x84\x6a\xd0\x31\x04\xb2\xd2\x1f\x2a\x70\xb6\x3f\x0c\x9d\xb6\x83\xa6\xb3\x74\x82\xc3\x9e\xe0\xda\xba\x7d\x0d\xaf\x6e\xf1\x75\x1d\x99\x82\xc5\x2d\xdd\xcd\xd5\x80\x98\xba\x60\x36\x33\x58\xc2\xa8\x91\x1a\x65\xdf\x1b\x99\x1a\xeb\xa3\xcc\x39\x36\xd9\xee\xcf\xab\x2f\x15\x14\x75\x51\xc1\xe7\x22\xb1\x14\x15\xac\x2a\x18\xf9\xbe\xcc\xd4\xd9\x89\xa1\x5e\x42\x20\xdf\xa3\xa6\xcc\x9e\xdd\x7a\xac\x5d\x3c\xe0\xb2\x63\x7b\xd3\xb7\x1a\x43\x7b\xb4\x99\x6c\xcb\x7b\x23\x5d\x36\xb5\x3e\x1f\x5d\x7d\xd3\x91\xbe\x1e\xd2\x64\x04\x5a\x47\x0c\xd6\x09\x90\x6d\x21\x55\xe7\x52\xb8\x7c\xff\xf6\x54\x62\x86\x02\x04\x76\xbd\x11\x0c\x07\x28\xce\x8b\x07\x07\xaf\x3a\x82\x1e\x45\x28\xa4\xca\x36\x99\xc6\x6e\x00\xef\x29\xad\x27\x03\x79\xf4\x03\xff\x64\xf0\x23\x10\xbb\x18\x34\xc1\x12\xce\x27\xaa\x3c\x86\xdc\xd8\x1c\x60\x8f\x41\x0c\xf6\x10\x88\x63\x2f\x3c\x8a\x9d\x58\x71\xe7\x4c\x0b\x85\x75\x52\x54\xc7\x20\x74\x29\x42\x81\x9f\x41\xc1\x79\x31\x5b\xf3\x27\xe6\xdb\x54\x01\x53\x0a\x42\x27\xe2\xb9\x5e\x2c\x36\x46\xba\xd8\xcc\xb5\xdb\x2e\x9c\x27\x7b\xe1\x5d\x6f\xf4\xe1\x02\x37\x64\x65\xe1\x3c\x2e\x0c\x73\x24\x5e\xfc\xf8\xff\x9f\xe6\xca\x3a\x59\x7d\xcb\xe1\x87\xd8\x9f\x21\xfc\x6f\x99\x8c\x1a\x4c\xbf\xea\x0c\x03\x47\xef\x5d\x90\x1c\x2e\x26\x68\x22\xa7\x78\xf3\xd0\x7a\xad\x4a\x48\xb4\x97\x30\x8a\xc0\x16\x0f\x43\x28\x9d\xd6\x31\x24\x6f\x24\x7b\xcd\x02\x4c\xf9\x36\x24\x63\xa6\xcd\x57\xa0\x34\xae\x26\xdf\xf9\xc6\x58\x3a\xce\x15\xc1\x07\x5a\x9b\x1b\x98\xd2\x7c\x33\x07\x8d\x36\x95\x31\x1e\xa0\xb8\xa9\x0f\xf5\xda\xb9\xf3\x62\x96\x09\xf5\x13\x42\xf4\xbe\x37\xe9\x16\xba\x2c\x7f\xbc\x30\xe3\x09\xf2\xdb\x83\xf6\x00\x2d\x91\xa7\x30\x6e\xb3\x2a\x21\xfd\x06\xad\x02\x53\xa6\xb6\x28\xba\x23\x4e\xab\xa6\xc8\xb8\xf4\xaf\xd6\x45\x05\x24\x7a\x3e\x9b\x9f\x02\xbc\xca\xa5\xc9\xd5\x26\xfb\xcc\x82\x41\x4e\x4e\x4b\x30\xdb\x69\x53\x25\x73\x67\x43\xa6\x55\x09\x1f\xc7\x38\x65\xa8\xb1\x1b\x55\xaa\x31\x62\xab\xa3\xf2\xd4\xd8\x0a\x58\x5c\xa0\x36\xd3\x3e\x79\xc8\x8e\xfb\x69\xf7\xd1\x9c\x4f\xbb\xc6\xa6\xd9\x0e\xcb\xa4\xf9\xef\xb8\xbf\xe6\x7d\xd6\xf2\x23\x96\x97\xa8\x7c\xe7\xf1\x8f\x4f\xd5\x7f\x13\x78\x51\x0f\x4f\x95\xfe\x5e\x6a\x35\x8c\x10\x96\x20\x21\xd2\x30\xc6\x4f\xc3\x2b\xf9\x78\x8a\xc7\x87\xf3\x9f\x8d\x7e\xc9\xb8\xbe\x41\xf2\xbd\x73\x79\x4e\xf7\xac\x95\xbf\x02\x00\x00\xff\xff\xdf\x3d\xa3\xb4\xd9\x07\x00\x00")

func policyCommonRegoBytes() ([]byte, error) {
	return bindataRead(
		_policyCommonRego,
		"policy/common.rego",
	)
}

func policyCommonRego() (*asset, error) {
	bytes, err := policyCommonRegoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "policy/common.rego", size: 2009, mode: os.FileMode(0644), modTime: time.Unix(1557960818, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x3c, 0x27, 0xd8, 0x68, 0x58, 0xbf, 0xf1, 0xd5, 0xaf, 0x60, 0x6e, 0xe1, 0x2b, 0x1f, 0x72, 0xb6, 0xc5, 0xf5, 0x47, 0x42, 0xc, 0x86, 0x57, 0xcb, 0x6b, 0x8b, 0x85, 0x39, 0x5e, 0x28, 0xd7, 0xe5}}
	return a, nil
}

var _policyIntrospectionRego = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x51\x4d\x6a\xf3\x30\x10\x5d\x7b\x4e\x31\x38\x8b\x7c\x1f\x04\x1f\xa0\x90\x2b\xf4\x02\x26\x18\x45\x9a\xd4\x6a\x6d\x8d\x90\x46\x2d\x4d\xc9\xdd\x8b\xe4\x38\xc5\x49\x0a\xed\xa6\x4b\xbd\x79\xa3\xf7\x33\x5e\xe9\x17\xf5\x44\xa8\x92\xf4\xc7\xc6\x3a\x09\x1c\x3d\x69\xb1\xec\x00\xec\xe8\x39\x08\x1a\x25\xaa\x29\x84\x05\xa2\x79\x1c\xd9\x2d\x20\xcf\x83\xd5\x96\x22\x80\x57\x36\x74\xa3\x12\xdd\x53\xec\x02\x45\x4e\x41\x53\xdb\x7a\x1e\x3a\x6b\x36\x98\xc7\xbb\x1d\x7e\x40\x35\xaf\x9c\x47\x3b\xdc\xa2\xe7\x01\x2a\xeb\x7c\x92\x26\xf3\x62\xdb\x15\x54\xd9\x00\xd5\x24\xda\xcc\x3f\xce\x12\xff\xf2\xf4\x82\x6e\xf2\x17\x97\xd7\x7f\x38\x5d\xf9\x51\x25\xdf\x3d\x37\xdf\xa8\x4e\xed\x4c\x6b\x4b\xc9\x09\x2b\x82\x8b\x18\xe7\x41\xd1\x5e\xe1\x23\x0b\x3d\xa0\x30\x06\x92\x14\x1c\x4a\x4f\x18\xd3\x3e\x92\x20\x1f\xca\x2b\x2b\x70\xb0\x47\x32\x45\x32\xce\xb8\x0f\xfc\x6a\x0d\x19\x2c\xce\x36\xb0\x42\x4e\x01\x43\x1a\x28\xe2\x98\xa2\x60\x3d\x7d\x59\x17\xf6\x3a\xef\xae\xa7\x5b\x80\x1a\x06\x7e\x23\xd3\x65\xac\x2d\x11\xef\xf5\xdd\xd0\xe1\x40\x5a\x70\x8b\x75\x59\xa8\xe7\xb4\xbd\x8a\x5d\x4c\xfb\x67\xd2\x32\x73\xa1\xfa\xd1\x59\xaf\x68\xf7\xdb\xce\xc5\x18\x72\xf6\x37\x06\x0d\xb9\xf7\xbf\xf4\xf7\x75\x94\x2b\x8f\xb7\xd5\x42\xe5\x58\xf0\x26\x10\x9c\xe0\x33\x00\x00\xff\xff\x8b\x33\xc3\xfe\x62\x03\x00\x00")

func policyIntrospectionRegoBytes() ([]byte, error) {
	return bindataRead(
		_policyIntrospectionRego,
		"policy/introspection.rego",
	)
}

func policyIntrospectionRego() (*asset, error) {
	bytes, err := policyIntrospectionRegoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "policy/introspection.rego", size: 866, mode: os.FileMode(0644), modTime: time.Unix(1554905064, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x4e, 0xbc, 0xd9, 0x1d, 0xd5, 0x12, 0xf1, 0x18, 0x79, 0xf3, 0x71, 0x4c, 0x7d, 0xb4, 0xb9, 0xd7, 0x1a, 0xaf, 0x10, 0x9b, 0x59, 0xdd, 0xac, 0xf, 0xbc, 0xf7, 0xa, 0xc7, 0xf, 0x81, 0x3a, 0xe2}}
	return a, nil
}

var _policyIntrospection_v2Rego = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x55\x4f\x6f\xdb\x3e\x0c\x3d\x5b\x9f\x82\x48\x0e\xfd\xfd\x80\xcc\x87\x1d\x0b\xf8\x23\x6c\xa7\xdd\x8c\xc0\x60\x24\xa6\xd6\x6a\x8b\x86\x44\xa7\x48\x87\x7e\xf7\x41\x92\xdd\x38\x7f\x3a\x74\xdd\xb0\x93\x2d\xea\xf1\x89\xef\x91\xb2\x07\xd4\x8f\xf8\x40\x80\xa3\xb4\xcf\xcd\xe1\x73\x69\x9d\x78\x0e\x03\x69\xb1\xec\x94\xb2\xfd\xc0\x5e\xc0\xa0\x60\x39\x63\x00\x43\xc6\x9f\xed\x6a\xee\x7b\x76\x67\xa1\x81\x3b\xab\x2d\x85\xb3\xa0\xe7\x8e\x82\x52\x9a\x5d\x90\x26\x1c\x83\x50\xdf\xc8\x71\x20\xa8\x60\x95\x97\x2b\xa5\x06\xb4\xbe\xe9\x51\x74\x4b\xa1\xf1\x14\x78\xf4\x9a\xea\x7a\xe0\xae\xb1\x66\x03\x41\x50\xa8\x27\x27\x69\x15\xc1\xdb\x2d\xfc\x50\xc5\x7c\xe0\x04\xdc\x96\xaf\xc0\x50\x2f\x73\xb6\xe5\xcc\x19\xea\x66\x0b\xd5\x82\x70\xde\x50\x85\x75\xc3\x28\x65\x24\x9f\x40\xf1\x55\x15\x59\xe7\x2b\xc1\x5c\xe5\x7f\x71\xf7\x35\xba\xb9\xc1\xf8\xbf\x7a\xb9\x10\x86\xc9\xe4\xbf\x29\x2b\x33\x5e\x89\xca\xe1\x37\x25\xa5\x66\x4e\xb9\xe7\x7a\x72\x6c\x73\x45\xf5\x2f\xb4\xc4\x41\x81\x0a\xe2\xa3\xb1\x46\x15\x69\x70\xea\x69\x79\x21\x35\x45\xff\x58\xe5\x82\x25\x09\x5c\xc3\xb7\xd6\x06\xd0\x38\x06\x0a\x20\x2d\x41\x8b\xa1\xe9\xa9\xdf\x91\x87\x40\x02\xc2\xb0\x23\x78\x20\x47\x1e\x85\x0c\xa0\x33\xd0\x53\xcf\xf6\x99\x0c\xb0\xd3\x04\xd6\x05\x21\x34\xc0\x7b\xb5\x86\x3d\x7b\x20\xd4\x2d\x24\x13\x8e\x25\x7c\x61\x4f\x7c\x20\x0f\x56\x60\xf0\x74\x88\x2e\xa4\x73\x76\xa8\x1f\xc5\xa3\x7e\xb4\xee\x01\xf6\x9e\xfb\xf9\x90\xb8\x36\xe3\xd0\x59\x8d\x42\x6a\x0d\xe8\xc2\x13\xf9\x98\x84\x02\x2d\x1e\x68\xaa\x89\x0e\xd8\x8d\xa9\xa6\xdd\x31\x31\xce\x53\x98\x6a\xcc\x1a\x21\x9b\xe0\x43\xa9\x16\xba\xaa\x7c\xb7\xcb\x53\x48\xa9\x04\x6c\xa2\x57\x75\x4d\xfb\x3d\x69\xc9\x2d\xdd\xc0\xad\x76\xe7\x46\x67\x1c\xdc\x57\xf0\xfe\x96\xe7\x1c\x55\x9c\xce\x9e\x73\x54\xf1\x81\x4f\xc2\x45\xd2\x7b\x46\x34\xb7\xfd\x2b\x0b\xdd\x47\x27\x3d\xc9\xe8\x5d\x32\x30\x8c\xbb\xd8\x72\xde\xa7\x55\xf4\x88\x7d\xea\x73\x1a\xb4\x39\x3e\x78\x3e\x58\x43\x06\xd2\x10\x6e\xd4\x1a\x78\xf4\xe0\xc7\x8e\x02\xf4\x63\x10\x58\x65\xca\x55\x42\xdf\xc5\xdc\xbb\xfc\x5d\x54\xd8\x75\xfc\x44\x26\xdb\x9c\xaa\x89\x36\x2e\xad\x5f\x25\xc8\x6a\xf6\xbe\xd9\x40\x93\x2b\x36\xe4\xec\xaf\x33\x0d\xb9\xe3\x8d\xc4\x93\x8c\x04\x8c\xcd\x5f\x54\x01\x9f\x60\xc1\xac\x4e\x15\x7a\xfe\x4e\x5a\xea\xe9\x99\xef\x75\x7e\xff\xbd\x7e\x4f\x49\xf1\x9a\xaa\x62\x52\x07\xd5\x47\x26\xe6\x72\x66\x4f\x73\xe3\x58\xae\xf9\xf2\x0f\xa7\x82\xab\xbf\xd0\xa5\x29\x37\x94\xbe\xe5\x82\x7a\x51\x3f\x03\x00\x00\xff\xff\x5f\xfa\x7e\x42\x4e\x07\x00\x00")

func policyIntrospection_v2RegoBytes() ([]byte, error) {
	return bindataRead(
		_policyIntrospection_v2Rego,
		"policy/introspection_v2.rego",
	)
}

func policyIntrospection_v2Rego() (*asset, error) {
	bytes, err := policyIntrospection_v2RegoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "policy/introspection_v2.rego", size: 1870, mode: os.FileMode(0644), modTime: time.Unix(1572470115, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xc, 0xc4, 0x1b, 0x3e, 0xb9, 0x9f, 0x3e, 0x38, 0x1, 0x7d, 0x22, 0x3e, 0xd5, 0x47, 0xac, 0xb6, 0x5e, 0x2d, 0xc4, 0xcf, 0x60, 0x39, 0x9f, 0x73, 0x19, 0x44, 0xcc, 0x68, 0xdb, 0x88, 0xfd, 0x52}}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetString returns the asset contents as a string (instead of a []byte).
func AssetString(name string) (string, error) {
	data, err := Asset(name)
	return string(data), err
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

// MustAssetString is like AssetString but panics when Asset would return an
// error. It simplifies safe initialization of global variables.
func MustAssetString(name string) string {
	return string(MustAsset(name))
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetDigest returns the digest of the file with the given name. It returns an
// error if the asset could not be found or the digest could not be loaded.
func AssetDigest(name string) ([sha256.Size]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s can't read by error: %v", name, err)
		}
		return a.digest, nil
	}
	return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s not found", name)
}

// Digests returns a map of all known files and their checksums.
func Digests() (map[string][sha256.Size]byte, error) {
	mp := make(map[string][sha256.Size]byte, len(_bindata))
	for name := range _bindata {
		a, err := _bindata[name]()
		if err != nil {
			return nil, err
		}
		mp[name] = a.digest
	}
	return mp, nil
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
	"policy/authz.rego": policyAuthzRego,

	"policy/authz_v2.rego": policyAuthz_v2Rego,

	"policy/common.rego": policyCommonRego,

	"policy/introspection.rego": policyIntrospectionRego,

	"policy/introspection_v2.rego": policyIntrospection_v2Rego,
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
// then AssetDir("data") would return []string{"foo.txt", "img"},
// AssetDir("data/img") would return []string{"a.png", "b.png"},
// AssetDir("foo.txt") and AssetDir("notexist") would return an error, and
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		canonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(canonicalName, "/")
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
	"policy": &bintree{nil, map[string]*bintree{
		"authz.rego":            &bintree{policyAuthzRego, map[string]*bintree{}},
		"authz_v2.rego":         &bintree{policyAuthz_v2Rego, map[string]*bintree{}},
		"common.rego":           &bintree{policyCommonRego, map[string]*bintree{}},
		"introspection.rego":    &bintree{policyIntrospectionRego, map[string]*bintree{}},
		"introspection_v2.rego": &bintree{policyIntrospection_v2Rego, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory.
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
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
}

// RestoreAssets restores an asset under the given directory recursively.
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
	canonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(canonicalName, "/")...)...)
}
