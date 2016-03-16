// Code generated by go-bindata.
// sources:
// api/companies.raml
// api/contracts.raml
// api/organizations.raml
// api/securitySchemes/oauth_2_0.raml
// api/userorganizations.raml
// api/users.raml
// DO NOT EDIT!

package specifications

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

var _companiesRaml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xa4\x55\xdf\x4f\xe3\x38\x10\x7e\xcf\x5f\x31\xa7\xd3\xbd\xd1\x34\x29\x3f\x54\xe5\xe9\xb8\x1e\x07\xbd\x03\x71\x82\xe5\xa1\x5a\xad\x54\x37\x9e\x06\x2f\x8e\x6d\x6c\xa7\x6d\x58\xed\xff\xbe\x4e\xd2\x90\xa4\x04\xe8\xb2\x96\x2a\x55\x1e\xcf\xcc\xf7\x7d\xf3\x23\xbf\xff\x71\x73\x7a\x75\x09\xa1\x1f\x78\x96\x59\x8e\x11\x4c\xad\x99\xc9\xec\x5a\x70\x26\xd0\x5b\xa1\x36\x4c\x8a\x08\x02\x3f\xf4\x16\xc4\xe0\x9d\x66\x11\x0c\x3d\x9b\x2b\x34\x91\x07\x30\x91\xa9\x22\x22\x2f\xfe\x02\x28\x2d\x15\x6a\xcb\x2a\x53\x75\x12\x2e\x17\x84\x33\xda\xdc\x14\xa7\xf0\x8f\xc0\x58\xcd\x44\xd2\x31\xa4\x4c\x5c\xa2\x48\xec\x7d\x04\x87\x5d\x03\xd9\xd4\x86\xf0\x38\x78\x36\xa9\x6c\xc1\x59\xfc\x1f\xe6\xe6\xf5\x04\x9f\xbf\xec\x46\x9a\x5a\x4c\x4d\x04\xa3\x26\x0e\x6e\x14\xd3\xd8\x17\x83\x12\x8b\x9d\x6b\x8d\x8f\x99\x7b\x4b\x23\x58\x12\x6e\x1a\x9b\xd4\x09\x11\xec\x89\x58\x27\xd8\x47\xd0\x84\x41\xb0\x57\x1e\x26\x96\xf2\xd7\xc8\xbe\x15\xdd\x92\x8d\xd0\x7b\x16\xab\x2f\x06\x6e\x48\xaa\x38\xf6\x34\x00\x4c\x2f\xd9\x03\x4e\x88\xa0\xf9\xbb\xe5\x1b\x40\x38\xbe\x3d\xbf\x98\xdd\xb2\xf1\xbf\x93\xd5\xe9\xea\xd3\xdd\xc9\x2c\x5d\x85\x17\x37\x69\x9a\x8d\x4f\xae\x1e\x47\x33\xf5\xb4\x5b\x3b\xc7\x30\x1c\x0f\xc2\x60\xd0\x62\xba\xab\xd5\xa0\xbc\xf9\x93\x71\x87\x24\x2e\x90\xf8\xb1\x4c\x77\xb8\xc3\x5f\x67\xe1\xe8\xf0\xa8\xfd\xf3\xdc\x93\xb8\xea\xf4\x15\xc3\x75\x13\xb1\x92\x66\x3b\x04\x9e\x37\xac\x1e\x6d\x07\x40\x49\x63\xab\xa7\x14\x4d\xac\x99\xb2\xe5\x28\xdd\x60\xc2\x8c\x45\x0d\x04\x04\xae\xeb\xb8\xe5\xbb\x85\xa4\x79\x13\x9c\x28\xe5\xc4\x29\xfb\x69\xf8\xd5\x38\xd7\x9e\xaa\x4c\x9e\xbd\x87\xdf\x2a\xa5\xa7\xf4\xfb\x76\x1a\x33\x5b\x7b\x74\xf2\xdf\xa9\xa2\xa5\x9d\x6a\x0e\x85\xab\x68\x0d\xc0\xaf\x0c\xc5\xcd\x7c\x5e\x87\x9a\xcf\x81\x19\x10\xd2\x02\xe1\x5c\xae\x91\xfa\x5e\x5d\x7a\xa3\x5c\x9b\xb7\x27\x7d\x14\x04\x6d\x84\x5d\x2e\xef\xf3\x79\xc9\xa8\x3a\x47\xc1\x61\xfb\xe1\x4b\x2a\x05\xe2\x1a\x6f\x81\xf6\x1f\xa9\x17\x8c\x52\x14\xbf\x95\x6e\xc3\x76\x0f\x24\x68\x9b\x60\x3d\x1c\x7a\x78\xbc\xc6\x66\x5f\x56\xbb\xfc\x5a\x7d\x54\xe1\x5b\x15\xd3\xe1\x2a\xd2\x8b\xf1\x31\x43\x9d\xff\x4f\x34\x49\xd1\xb5\xcc\xee\x5a\x91\x0f\xd8\x93\xaf\x33\xaf\x55\x8e\x58\x0a\xab\x49\x6c\x4d\x6f\x92\x8e\xa6\xe7\x68\xc1\xde\x23\x3c\xbb\xc0\xfa\x1e\x35\x96\x77\xed\x15\x57\x48\x1d\x82\x5c\x96\x06\x45\xca\xb5\xef\xc3\xb5\xa6\xae\xb3\x8b\x80\x28\x68\x51\x99\x45\x5e\x6e\x50\x7f\x1f\x46\x4c\xc4\x3c\xa3\x78\x56\xce\x33\xed\x6f\x8f\x85\x94\x1c\x89\xd8\xb1\x75\x18\x4c\xab\x30\x25\xb0\x6a\x37\xd0\x86\xcd\x41\x89\x08\x97\x24\xe3\x16\xa4\xe0\x79\xf9\xcc\x59\xd8\xaa\xcd\x99\x38\xc6\x1a\x6d\xa6\x45\xd3\xf3\xf5\x79\x6d\x75\x82\xd3\xdc\xe9\xd0\x8f\x9b\x09\x8b\x09\xea\xb7\x70\xdf\x16\xde\x4e\xd1\xa5\x41\x7b\x00\x99\xc1\x65\xc6\x61\x29\xb5\x53\x37\x71\x52\xfa\xf0\xf7\x16\xb6\x53\x7e\x1e\xcc\xf7\x87\xe5\x3e\x02\x1f\x06\x75\x45\x36\x45\x7e\x04\xc3\x9e\xf0\x5d\x54\xc7\x3f\x03\x6b\x0b\x8d\xa5\x59\xea\x96\xb7\xfb\xa8\xff\x08\x00\x00\xff\xff\x69\x83\xe2\xb8\x8d\x08\x00\x00")

func companiesRamlBytes() ([]byte, error) {
	return bindataRead(
		_companiesRaml,
		"companies.raml",
	)
}

func companiesRaml() (*asset, error) {
	bytes, err := companiesRamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "companies.raml", size: 2189, mode: os.FileMode(420), modTime: time.Unix(1457705563, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _contractsRaml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xbc\x55\x4f\x6f\x13\x3f\x10\xbd\xef\xa7\x18\xe9\xa7\x9f\x00\xa9\x4a\x36\x55\x5b\x55\xb9\xa0\x12\x5a\x88\x68\x01\xf5\xcf\x01\x21\x0e\xce\xee\x24\x31\xdd\xd8\xc6\x1e\xb7\x09\x7f\xbe\x3b\x63\xef\xae\xb3\x49\x28\x15\x1c\xe8\xa1\xaa\x67\x3d\xf3\xde\x9b\x79\x9e\xfe\xf7\xff\xe5\xc9\xc5\x39\x0c\x7a\x79\x46\x92\x2a\x1c\xc2\x98\xdc\x07\xed\xdf\xa9\x4a\x2a\xcc\xee\xd0\x3a\xa9\xd5\x10\xf2\xde\x20\x9b\x08\x87\x37\x56\x0e\xa1\x9f\x65\xb4\x32\xe8\x86\x19\xc0\x95\x9c\x29\x41\xde\x62\x38\x00\x18\xab\x0d\x5a\x92\xf5\xc7\xf0\xe3\xf8\x02\x96\x2f\x56\x43\x70\x64\xa5\x9a\x35\xe1\x52\x10\x83\x85\xdf\x4d\xc0\xf8\x49\x25\x8b\x37\xb8\x6a\x13\x01\x02\xc8\x56\x1a\x27\xa2\x2b\xac\x34\x14\x69\x5d\xcf\xb1\x49\x84\x5b\x5c\x81\x77\x58\x76\x60\x3b\xbc\x76\xca\x71\x74\xa4\x15\x59\x51\xd0\x43\xcc\x8d\xd8\x38\x6e\x16\xf8\xf8\x29\x85\x17\x52\x8d\x09\x17\x6e\x08\xfb\xeb\x98\x58\xb6\xb1\x3c\x05\xbd\x92\x5f\x3c\x36\x71\xb2\xbe\x95\x5e\x30\x11\x54\xb4\x25\xb5\x68\xe8\x5d\x07\xd4\xdf\xf7\x84\xd1\xce\x51\xcd\x68\x3e\x84\x83\xfc\xd7\x9d\x3a\xb3\x88\x30\x95\x58\x95\x40\x1a\x26\x18\x7b\x05\x53\x6d\x81\xc4\xcc\xc1\x64\x05\x14\x7a\xc9\x92\x57\x50\x58\x14\xc4\xf5\x41\x24\x12\x4d\x51\xa9\xee\x44\x25\xc3\xd4\xdc\xf3\xc7\xfb\x92\x7a\x30\x78\x80\x54\x25\x1d\x81\x9e\x26\x94\x71\xf9\xc4\x31\x0f\xe9\x52\xa4\x8b\xd8\x6b\x8a\xe0\x92\xbb\x55\xfe\x33\x02\x0e\x84\xaa\x31\xc3\x53\x00\x6f\xb4\x4a\x44\x8c\xb4\x6c\x90\xae\x8d\xd7\x95\xfe\xc4\xc7\xdf\x53\x18\xa2\xa7\xd7\x55\x02\x7e\x98\xcc\x5c\xb8\x79\x60\xda\x4e\x89\x8d\xb9\x17\x0f\x8d\x79\xd6\x87\xd6\x33\x7b\xc0\xc4\xcb\x18\x8e\x44\x45\x80\x8a\x54\x7b\x5d\x34\x0d\x85\xa8\x0a\x5f\x71\x3c\x01\x71\x31\x71\x5b\x1f\x3f\x3b\x4e\xb2\x68\x58\x27\xa3\xd4\x35\x1a\x1a\xa9\x45\xf7\x92\xe6\xda\x53\x0c\xca\x32\xa2\xa6\xe7\x17\x58\x5a\xb9\x88\xdf\x14\xde\x87\xad\xe2\x80\x5d\xe7\x8c\x28\xf8\xaf\xa7\xb8\x2c\xd0\x70\x89\x39\xaa\x70\x67\x05\xc2\x18\x14\x96\xe7\x1e\x53\x78\xf8\x1e\xdd\xb3\x58\x13\x55\xa1\x4b\x0c\x5f\xc0\xd3\xf4\xb8\xab\xe2\xcc\xea\x45\x3d\x37\x46\xf4\x15\x35\x02\x04\x5c\xbd\x3e\xd9\x3f\x3c\x62\x9f\x57\x95\xbe\x67\xc3\xb3\xcf\x05\x5c\x8e\xdf\x9f\x5e\xbc\x1c\x1c\xe5\x7b\x4c\x69\x49\x10\x16\xdb\xe1\xf1\x68\x8e\xc5\xed\x69\x8d\x11\xa0\xeb\x4a\x51\x1c\xe7\xe4\xcb\x3c\xe7\x1d\x81\x53\xb9\xec\x65\xdb\x3b\x86\x2d\x90\xf6\x20\xfb\x2f\xeb\xb7\xad\x89\xcb\xc3\x68\xd7\xec\x98\x8d\x99\x8f\xc2\x2b\x0b\x1c\xb9\x2d\xa9\x97\xb5\x28\xae\xc8\x2e\x73\xeb\xdd\xb3\x9f\x0f\xd6\x76\x9a\xe8\xb2\xb3\x24\x21\x74\x8c\x17\x60\x1c\x4d\x3f\x8c\xab\xfb\xad\x35\xdf\x68\xf3\x1d\x1f\x74\xeb\x6d\xb0\xba\x51\xc2\xf3\x34\xad\xfc\xca\xab\x94\xaf\xf4\xbf\xad\xbd\xf8\xa3\xce\x99\x21\xb5\xc9\x1b\xa9\xaf\x90\x76\x37\xc6\x8e\x94\x20\x26\xef\x52\xdc\x96\xf3\x98\xa0\x07\x24\x05\x51\x07\xdd\x9b\x1b\xdc\xde\x6a\x62\x13\x78\x55\x6e\x4b\xea\x77\x86\x58\xff\x1f\x48\xd3\xda\x29\x12\x66\xbc\xab\xf0\xef\x06\x92\xfc\x92\xfd\x0c\x00\x00\xff\xff\x50\xa8\x00\x0d\x83\x07\x00\x00")

func contractsRamlBytes() ([]byte, error) {
	return bindataRead(
		_contractsRaml,
		"contracts.raml",
	)
}

func contractsRaml() (*asset, error) {
	bytes, err := contractsRamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "contracts.raml", size: 1923, mode: os.FileMode(420), modTime: time.Unix(1457614661, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _organizationsRaml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xdc\x58\x4d\x6f\xdb\x38\x10\xbd\xeb\x57\x10\x58\xec\x2d\x91\xe5\xd4\x6d\x37\xbe\x05\x6d\x11\x04\x9b\x8f\x22\x69\x1a\x6c\x8b\x05\x42\x49\x63\x9b\xad\x44\xaa\x24\x95\xd8\xdd\xed\x7f\xdf\x21\x29\x59\xdf\xb2\x93\xed\xa9\xb9\x04\x96\x38\x9c\x99\xf7\xde\xcc\x90\xfa\xed\xf7\xeb\x93\x8b\x73\x32\xf5\x03\x4f\x33\x9d\xc0\x9c\x9c\x69\xf5\x97\xc8\xaf\x78\xc2\x38\x78\x0f\x20\x15\x13\x7c\x4e\x02\x7f\xea\x85\x54\xc1\xad\x64\x73\x32\xf1\x3c\xbd\xc9\x40\xcd\x3d\x42\xae\xe4\x92\x72\xf6\x9d\x6a\xb3\x0c\x7f\x13\x92\x49\x91\x81\xd4\xcc\xbd\x37\x7f\xcb\x44\x84\x34\x61\x71\xf9\x9b\x10\x63\x3e\x27\x4a\x4b\xc6\x97\xdb\x87\x29\xe3\xe7\xc0\x97\x7a\x35\x27\x2f\xaa\x87\x74\x5d\x3e\x9c\xbe\x0c\x8a\xc7\x59\x1e\x26\x2c\xfa\x13\x36\xaa\x7f\xcb\xcf\x7f\xd7\xed\xcf\x34\xa4\x6a\x4e\x8e\x4a\xeb\x98\x3f\xc5\x6c\x1a\x04\xdb\xa7\x31\xa8\x48\xb2\xcc\xe6\x4a\xce\x99\xd2\x44\x2c\x88\xa8\x21\x40\xde\x5e\xde\xf8\xc5\x72\xf1\xc8\x11\xbd\xe7\x04\x38\xe0\x28\x57\xb8\x1f\xb9\x37\xff\x38\x4d\x41\xdd\x97\x9e\x52\x48\xc3\xa7\xba\xda\x95\x95\x73\xf6\xb8\x12\x84\x4a\x28\x3c\x10\xc6\x89\x5e\x31\xd5\x48\xb9\x0c\x82\xf1\x28\xc9\x63\xf8\x29\xd0\x3a\xc1\x9c\xc5\x26\x10\x95\x87\x0d\x7f\xca\xb3\x36\xb0\xa6\x69\x86\x7a\x6d\x4b\x8c\x2c\x25\x00\x67\xda\x3c\x80\x11\xb9\x1c\x92\xe9\xc9\xf1\xdd\xdd\xea\x15\x3b\x79\x37\xbb\xfe\x74\x7a\xf9\x9a\xae\x37\xaf\xd6\x9f\xc2\xfc\x2e\x3f\xff\xca\xbf\xdd\x9d\xcb\x8f\x7d\x7a\x39\x6c\x38\xf0\x23\x91\x0e\xd0\x7d\x68\x21\x9c\xb6\x7e\x1f\xb5\x7e\xbf\x18\xa2\xd0\xbd\x9e\x0d\x82\x7b\x48\x68\x8c\x05\xe3\x77\xa2\x69\xd5\x64\x61\xb2\x53\xb9\xca\xab\x28\xab\xdb\x23\x71\x5e\x19\xde\x50\x7d\x97\x8a\xdc\x51\xdf\x8d\x18\x6e\x15\xc4\x28\x2f\xe0\x84\x2a\xc5\x96\x1c\x17\x12\x5a\xea\x4c\xa3\xea\x78\x4b\x65\x7d\xac\x6f\x1d\x93\x50\x84\x66\xc5\x19\x7f\x60\x7a\xb4\x17\x39\xa3\x4e\x6c\x52\x24\xb5\xf0\x07\x12\x40\xff\x3c\x4f\xe7\xe4\xb3\xe5\xfa\xa0\x88\xb6\x52\x76\x24\x81\x6a\x40\x0d\xc6\xf8\x6f\x30\x60\x17\x6c\xcd\xad\x93\x8e\xd7\xda\xe3\x26\xe7\x07\x24\x78\x45\x2e\xc5\x03\x99\x1e\x1f\xcf\x48\xf0\xc7\x7c\x76\x3c\x7f\xf1\x9a\x9c\x5e\x7c\xf0\xbc\x89\x68\x93\x9c\x09\xa5\x7b\xc8\x7e\x63\x37\x44\x6c\x39\x3c\x36\x21\x25\x53\x1b\x0f\x51\x2b\x91\x27\x31\x09\xc1\xd5\x37\x14\x52\x26\x09\xca\xc4\x27\x1f\x4d\x59\xb9\xe6\x86\x95\x8f\x68\x2e\x84\x4c\x91\x3b\xe4\x28\x5a\x41\xf4\x95\xb0\x85\x35\x52\x10\xe5\x92\xe9\xcd\x0d\x3e\x4d\xd1\x5f\x92\x88\x47\x85\xc5\xce\xe9\x12\x7f\x73\x14\x5c\xd1\x3c\x8c\x4f\xd7\x34\x42\x11\x6f\x4a\x64\x68\x96\x61\x8d\x5a\x3f\x93\x2f\xaa\xe4\xaf\x5f\x93\xf6\x95\x04\x95\x61\xe6\x15\xb3\x47\xc1\xb4\x32\xaa\x6f\x3d\xbe\x7d\xbf\x0b\x55\xbc\x9f\xd5\x37\x6d\x0a\x98\xd3\x5c\xaf\x84\x64\xdf\x21\xc6\x15\x4b\xe8\x03\xff\xdf\xc2\xf6\x14\x74\xb3\xda\x7c\x72\x52\x58\x3b\x68\x13\x96\x32\xad\x6c\xaf\xb5\xa1\x3a\x80\x25\x7c\xcb\x41\x69\x53\x1b\x15\x6a\x3d\x89\x07\x3f\x33\x71\x5c\x30\xf9\xa7\xec\xa7\x3f\xdc\xfa\x6d\x76\xad\xfc\xda\x79\xa1\x82\x16\xa2\xd4\x76\x3b\xce\x56\xa4\xdd\x58\x77\x45\x3b\xa8\x85\x0e\x53\xe3\x5c\x95\x06\xb3\x41\x83\x4b\xa1\xc9\x42\xe4\xdc\xad\xce\xf2\xfe\xec\x6f\x33\x53\xe8\x83\x00\x34\xd3\x1b\x4b\x6d\x30\xad\x5f\x03\x43\xbb\x62\xd2\x1a\x71\x55\xbb\xea\x58\x9e\xd8\x89\xd0\x18\x07\x7d\x27\x8e\xe7\x89\xdd\xed\x59\x35\xfe\x2e\xc2\xad\x56\xd2\x09\xef\xc2\x45\xe5\xe6\x16\x16\xaa\xca\xa3\x08\x94\x5a\xe4\x49\xb2\x69\x58\x75\xb9\xd9\xcd\xce\x40\x98\x1d\x6e\xf6\x61\xa7\xc3\xcf\x38\x43\xa6\xee\xcb\x89\xfa\xa3\x4e\x4d\x02\x1a\x06\x69\xbe\x86\x54\x3c\x40\x45\xd6\x42\x8a\xb4\x41\x57\xcd\xb0\x17\x6c\x03\xf7\xac\x0d\x43\x1f\xe0\x2e\x90\x11\xbc\x3b\x10\xed\x07\x52\x0f\x4c\x2d\xc3\x0f\x38\xdf\xec\xa8\x14\xd2\x0d\xc8\x7a\xc5\xc7\x02\x14\xe1\x08\x25\xac\xcd\xc0\x6c\x6d\x7c\xbc\xdf\xc6\x34\xc1\x31\x1d\x6f\xcc\x88\x2d\x91\x2c\x8e\x3c\x93\xe6\xc1\xb2\x59\x36\xad\xed\xec\xf1\xc7\x70\x61\xf7\xc4\xb2\x09\x01\x8f\x84\xc5\x40\x37\x27\xbe\xce\xa9\xaa\xda\xe8\xb9\x8d\xa4\x23\xd4\x41\x92\xc7\xc9\x29\x62\x2f\xce\x3f\xc3\x24\xf7\x97\xd5\x3e\x85\x35\x18\xf1\x2f\xa3\x1c\xee\x98\xf6\x47\xeb\xb9\xaf\xa2\x07\x6a\xba\xd8\x6f\xb4\xa6\x07\x09\xef\xad\xeb\x96\xa3\x2b\xbb\xbd\xb4\xee\x46\x39\xef\x21\x68\x5f\x8a\x7a\x49\x1a\xee\x84\xae\xe8\x22\xc1\xb5\xa4\x91\xae\xbe\x61\xc0\xd0\xb4\x32\x67\x20\x43\xed\xd6\xc4\x5c\x6b\xf0\x1c\xd7\xa1\x1b\x19\x9a\x9a\x22\x34\x2f\x32\x6a\x6f\x25\x3e\xce\xe1\xd8\xb6\x36\x15\x01\x8f\xcd\x41\x2f\xdc\xd8\xeb\x43\xa5\x06\x3c\x02\xca\xcd\x7b\x2a\x91\x45\xdd\xb8\x22\x6e\xaf\x85\xef\xd6\x19\x93\x10\xf7\xd7\x67\x28\xf0\x8e\x41\xf9\x78\xe9\xd9\x6d\x6c\x60\xe0\xf6\xaa\xb2\x39\xb0\x11\xc1\x82\xe6\x89\x39\xc3\x27\x1b\xbb\x0c\xdf\xb0\x87\x7a\xce\xe6\xe4\x2a\x41\xe7\x12\x27\xa2\xdf\x72\x66\x8e\xb1\x36\x40\xb2\xa0\x89\x82\xda\x5b\xa5\x11\x87\xfe\xb8\x19\xd7\xb0\x6c\x95\x69\x2b\xee\x1b\x63\x8d\x88\x2e\x14\xe8\x03\x53\x0d\xa8\x1c\xa4\x51\x22\xba\x4b\x84\xd2\x27\x6f\x8b\xb0\x11\xf9\xfb\xe0\x7e\xff\xb0\x52\xba\x7e\x76\x50\x17\x74\x6d\xfc\xe3\x85\x08\x95\xb8\x33\xaa\x97\x4f\x09\xab\x08\x8d\xa5\xe6\x12\x7a\xf4\x32\x28\xb4\xca\xb6\x97\xde\x01\xb5\x32\x95\x25\x74\x73\x69\x2f\xca\xa8\xd6\xf7\x4e\x69\xd5\x5d\x59\x8d\x0b\x3b\x29\x3e\x17\x64\x85\x42\x6b\xfe\x6c\x5e\xee\x53\x11\x4e\x9b\x2f\x62\xec\x03\xd1\xc8\x64\x08\xda\x68\xff\xdf\x16\x5f\xe5\xe6\xbe\x5f\x58\x9c\xf6\x3b\xdc\xb4\xf0\x72\xad\xb0\x03\xd9\x70\xef\x7c\x43\x79\x04\x09\xce\xe0\x2e\x5c\xfe\x7e\x6d\x73\xa0\x71\xf6\xcd\x4b\xd7\x58\x22\xeb\x32\xc1\xbe\xf7\x5f\x00\x00\x00\xff\xff\xaf\x14\xd4\x11\xcf\x15\x00\x00")

func organizationsRamlBytes() ([]byte, error) {
	return bindataRead(
		_organizationsRaml,
		"organizations.raml",
	)
}

func organizationsRaml() (*asset, error) {
	bytes, err := organizationsRamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "organizations.raml", size: 5583, mode: os.FileMode(420), modTime: time.Unix(1458148761, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _securityschemesOauth_2_0Raml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xbc\x92\x3f\x6f\xdb\x40\x0c\xc5\x77\x7d\x0a\x42\x4b\xa7\x4a\x86\xe1\x49\x9b\x8b\x02\x1d\xdb\xa1\x9e\x8a\xa2\xb8\xea\x08\xeb\x10\xe5\x78\x21\x29\x3b\x0a\xf2\xe1\xc3\xd3\x1f\xc3\x36\x32\x64\x8a\xa6\x03\xa9\xf7\xf8\xe3\xbb\xf3\x28\x2d\x87\xa4\x81\x62\x03\xaf\x05\xc0\xcf\xfd\xa0\xdd\x16\x82\x80\x83\xc4\xa4\xd4\x52\x0f\xda\x39\x85\x1e\x55\x00\x9f\x15\x39\xba\x1e\x5c\x4a\x02\x8c\x4f\x03\x8a\x82\x33\x0d\x71\x78\x71\xd9\x07\x94\x4c\x19\x4e\x4e\xd1\xfc\x3c\xaa\x0b\xbd\x40\x88\x66\x38\x08\xf2\x17\x73\x6e\x5b\x1a\xa2\xc2\x39\x98\x6c\x50\x38\xa2\x6a\x88\x47\x1b\x83\x81\x21\x39\x91\x33\xb1\xaf\x0a\x1d\x13\x36\x33\x11\x6c\xab\x4d\x31\xc3\xfe\x47\xff\x6d\x6c\xcc\xba\x43\xe7\x91\x25\x1f\x01\xf6\xd7\x08\x73\x29\x0f\xbf\xdb\x6e\xfe\x0e\x82\x3e\x53\x0a\x46\x6f\x54\x27\xd7\x07\xbf\x8e\xc9\x70\x28\x62\xed\x07\x8c\x15\x7c\x27\x88\xa4\x19\xdc\x2a\xc6\xd9\x21\x4f\xd8\x17\x2f\xab\x40\x39\x6b\xfe\x4d\x9a\x12\x2c\x13\x1e\x41\x94\xf3\x4e\xc9\xb1\x7b\x44\x0b\xad\x32\xc9\xd4\xf9\xb5\x56\x16\xf2\x6b\xf1\x67\x83\xdf\xa4\x56\x2e\x89\xda\x2f\x8c\x92\x28\x0a\x2e\x88\xbb\xcd\xee\x5d\xb2\x43\x5c\x6f\x1e\x7d\x21\xf3\x2d\x4e\x92\x9b\x07\x71\xe0\xd0\x40\xa7\x9a\xa4\xa9\xeb\xa0\x32\xd2\x50\x51\xec\x43\xc4\xba\xa7\x63\x88\x35\xe5\xdf\xeb\x8b\x55\xb1\x66\xf2\x3b\xaf\xf2\x51\xf5\x55\x8a\xf7\x00\x3f\xd8\x45\x95\x06\xfe\x40\x4b\x1e\xe1\xaf\xf5\xa5\xa5\xb4\x6e\xf7\x75\x7a\x97\xcb\xb1\xcc\xe7\x86\x2d\x88\xb2\x78\x0b\x00\x00\xff\xff\x88\xbd\xe6\x45\x1d\x03\x00\x00")

func securityschemesOauth_2_0RamlBytes() ([]byte, error) {
	return bindataRead(
		_securityschemesOauth_2_0Raml,
		"securitySchemes/oauth_2_0.raml",
	)
}

func securityschemesOauth_2_0Raml() (*asset, error) {
	bytes, err := securityschemesOauth_2_0RamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "securitySchemes/oauth_2_0.raml", size: 797, mode: os.FileMode(420), modTime: time.Unix(1454406267, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _userorganizationsRaml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\x92\x41\x8f\xd3\x30\x10\x85\xef\xfe\x15\x23\x21\x6e\x28\x49\x11\xa7\xdc\xf6\x84\x16\x81\x56\x5a\xc4\x01\x55\x3d\x38\xc9\x34\x9d\xca\xb1\x2d\x8f\x53\x54\x4a\xff\x3b\x76\x92\x52\xa7\xa4\x85\xb9\xb8\xea\x8c\xdf\x7b\xf3\x39\x6f\xde\xbe\x3e\x7d\xf9\x0c\xab\xac\x10\x9e\xbc\xc2\x12\x9e\x3d\x7f\x37\xfd\x8b\x56\xa4\x51\x1c\xd0\x31\x19\x5d\x42\x91\xad\x44\x25\x19\xbf\x39\x2a\x21\x17\xfe\x68\x91\x4b\x21\x60\xa8\x4f\x86\xf4\x8b\x6b\xa5\xa6\x9f\xd2\x87\xf1\x67\x7d\x20\x3f\xfc\x2a\xa7\x09\x00\xeb\x8c\x45\xe7\x29\x5e\x83\xa4\x4c\x72\xaf\x04\xf6\x8e\x74\x3b\x1b\xe8\x19\xdd\x62\xc3\x99\x10\x77\xf6\x4f\xac\x98\x6c\x71\x3c\x16\xea\xbe\x2b\x61\x6d\x7e\x68\x74\xef\xa0\xc3\xae\x42\xb7\x11\xa1\xf2\xe8\xc2\xf9\x29\x1e\x5a\x76\x78\xce\xd3\x5c\x43\xe4\x16\xfd\xe8\xd6\x20\xd7\x8e\xec\x18\xf8\x23\x7a\xf0\x3b\x04\x45\xec\x67\xbb\x30\xc8\x21\x3a\x10\xc3\xe0\x17\xba\x93\x23\x98\xed\x20\xe4\x90\x6d\x18\xbc\x12\x79\x5f\x14\xd7\x85\x2a\xd3\x1c\xd3\xf5\xa4\xb5\x8a\xea\x41\x3b\xdf\x73\x8a\xf6\x11\x5e\x18\xcd\x2f\x44\xd6\x9b\x9b\xee\x98\x28\x69\x87\x7e\x7e\x6a\x95\xa9\xa4\xa2\xe6\x9c\x47\xc8\x81\x4b\x3c\xce\xa3\xb2\x35\xec\x2f\x1e\x33\x14\x4f\x75\x8d\xd6\x4f\x8a\xbc\x23\x0b\xa4\x67\x48\xc4\xd2\x62\x8f\xd6\x1a\xdf\xf2\xfe\xd7\x35\x8d\xfe\xc5\x31\x92\x5c\xa5\x42\xb7\x2c\xff\x45\xf3\x3f\xad\x1b\x54\xe8\x71\x11\xc6\xaf\x3f\x7a\xaf\xb8\xc7\xfa\x06\xcb\x45\x25\x12\x92\x73\x48\xd9\xa3\xa5\x3e\xa4\x31\xef\xf8\xc5\xfa\xda\x87\xb7\xe0\x6d\xaf\xd4\x31\x08\x45\x7f\x6c\x12\xdb\x4c\xfc\x0e\x00\x00\xff\xff\x2b\x39\x6a\x40\xf7\x03\x00\x00")

func userorganizationsRamlBytes() ([]byte, error) {
	return bindataRead(
		_userorganizationsRaml,
		"userorganizations.raml",
	)
}

func userorganizationsRaml() (*asset, error) {
	bytes, err := userorganizationsRamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "userorganizations.raml", size: 1015, mode: os.FileMode(420), modTime: time.Unix(1457621431, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _usersRaml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x5a\x7b\x73\xdb\x36\x12\xff\x5f\x9f\x02\xe7\xbb\x9b\xbb\x4e\x6c\xbd\x2c\xa7\x89\xfe\x49\x1c\xe7\x51\xa7\x76\x92\x89\xed\x76\x3a\x6e\x3a\x06\x49\x48\x42\x0d\x02\x2c\x00\x4a\x56\x32\xfd\xee\x5d\x80\xa4\x04\x92\xa0\x2d\xf9\x51\x3b\x99\x6a\x26\xb1\x44\x2c\x16\xbf\x7d\x62\x17\xe0\xbf\xff\xfb\x71\xf7\xf0\x00\xf5\xda\xdd\x96\xa6\x9a\x91\x21\xda\xd7\xea\x17\x91\xbe\xe7\x8c\x72\xd2\x9a\x12\xa9\xa8\xe0\x43\xd4\x6d\xf7\x5a\x01\x56\xe4\x44\xd2\x21\xea\xb4\xf4\x3c\x21\x6a\xd8\x42\xe8\x28\x14\x09\x31\x5f\x10\x8a\x88\x0a\x25\x4d\xb4\xa5\x7f\x2d\x24\xc2\x1c\x91\x8b\x84\x61\x8e\xcd\x33\x84\x03\x91\x6a\xa4\xcc\x04\x05\x63\x51\xf6\x35\xc6\x49\x42\xf9\x78\x13\x29\x42\xd0\x44\xeb\x44\x0d\x3b\x9d\x31\xd5\x93\x34\x68\x87\x22\xee\x50\xad\xe6\x22\x15\x16\x4e\x87\x46\x84\x03\xcc\xb9\x22\x12\x90\x75\x02\x26\x82\x4e\x8c\x95\x86\xef\x91\x08\x55\x47\xe0\x54\x4f\xfa\x9d\x6c\x8d\x76\x1c\x59\x5c\x89\x84\x5f\x52\xd3\x0c\x70\xf6\x49\x81\x03\xc7\x31\x88\xab\xb4\x84\xe5\x17\x03\x63\x89\xb9\x26\xd1\xb1\x58\xd2\x22\x64\xa4\xad\x51\x56\x04\x3e\x9e\x10\x24\xe4\x18\x73\xfa\x39\x13\x57\x4f\xb0\x46\x54\x15\x1c\x91\x24\x38\x42\x38\x0c\x89\x52\x48\x0b\xa4\x44\x0c\x13\x46\x40\x46\x2c\x9a\xff\x29\x44\xf9\x48\xc8\xd8\xce\x5e\xac\x82\xa3\x48\xc2\x8c\x67\xc3\xa6\x75\x0f\x33\xfd\x15\xac\x18\x0e\x08\x53\xc5\xaf\x62\x6d\xc1\x89\x5d\xd4\xa1\xd0\x55\xbc\x92\xfc\x91\x12\x50\x65\xd4\x76\x96\xf2\xe9\xce\x7c\x36\x4e\x3f\x6d\x94\x9f\x34\x68\x29\x49\x03\x46\xc3\x1f\xc9\x1c\x44\xc8\xc7\x4e\x3f\x2d\x46\xc1\xfa\x94\x7d\xbd\xb2\x4d\x60\xed\xaf\x16\x7d\x80\xf9\xf9\x57\x0b\x7e\x84\x43\x12\x08\x01\x02\x20\xf8\xc3\x08\x5e\x46\x4c\x96\x3b\x3c\x03\x2e\x2c\xd5\x14\xde\x8e\x73\x56\x34\x72\x40\x95\x36\x0a\x28\xb1\xb1\xf2\xe6\x22\x5a\x75\xb9\xa2\x87\x90\x00\x4d\x56\x83\xfc\x25\x51\x4c\xe2\x00\x92\xe9\x84\x26\xc0\xa4\xdd\x82\x55\x5e\x80\x01\x76\xc3\x50\xa4\x5c\x0f\x1b\x13\x15\x05\x33\xad\x90\x8a\x62\x7c\x71\x40\xf8\x58\x4f\x86\x68\xbb\xbb\x34\x31\x0d\xd7\x9b\xdb\x5b\xce\xb5\xb8\xe4\x7c\xbd\xf9\x83\xae\x11\xec\x83\x89\x0b\x9e\x1a\x79\x73\xc1\xb0\x86\x14\x0d\x3a\xfc\xed\xd7\x47\xa7\xdd\xad\xa7\x9f\x1e\xfd\xc7\xd0\xed\x66\xb9\xad\x59\xf8\x10\x52\x7d\xd9\x1d\x1a\x20\x34\x29\x00\x08\x09\xd1\xeb\xb3\xd8\x59\xb2\xe0\x72\xfd\xe9\x8e\x1a\x05\x38\x88\x7c\x76\x23\x21\x3c\x96\x58\x8d\xc7\x60\xc9\x23\x11\x4a\x63\x16\x8a\x88\xac\xcf\xa6\x6f\xd8\x9c\xa8\x85\x31\x2f\xdb\x4e\x6f\x22\xa7\x1d\xa1\x7c\xb1\xac\x67\x0f\xf1\x6d\x21\x17\x09\x95\x90\x85\x51\x84\x35\xb9\x64\x63\xb9\x83\x6c\xbf\x3e\x4b\x27\x32\x2e\xdd\xe1\xd7\xe7\x9c\xc7\x52\x73\x76\x5f\x9f\xa5\x93\x9f\x3c\x79\xf7\x26\xde\x58\xa4\xe8\xeb\xc4\xa5\x1d\x22\x17\x38\x4e\x18\xf1\xd5\x73\x81\x08\x7c\x8e\x53\xe2\xb8\x85\x7a\x6f\x7e\x3a\xfa\x7e\x67\x30\xd0\xd3\x9f\x0f\xfb\xc7\x87\xbd\xed\xe9\xbb\x1f\x9e\xa8\x40\xe0\xf8\xed\xe7\x83\xc7\xf1\xdb\xed\xc7\x15\x07\x33\x61\xd0\x7b\xb2\xd5\xeb\x6e\xf5\xbb\x65\x2f\x2b\xb3\x9e\x09\x79\x6e\x41\x3c\x87\xe2\x35\xc1\x7c\x6e\x8a\xd8\x12\xc5\x44\xe4\x30\x9f\xe7\x52\xd4\x28\x36\x22\xcc\x43\x82\x42\x96\x06\x1b\x19\xa5\x29\x7e\xed\x43\xf3\xac\x5d\xcc\x83\xcd\xa6\xec\x9d\x2e\x94\x00\x16\x79\xb4\xdd\xef\xf5\xb7\x07\xf9\x3f\x67\x90\x8c\x61\x10\x04\xf2\x0c\x62\x37\x29\x97\x30\x57\x9c\x24\x4f\xcf\xe8\x28\x31\x16\x1b\x51\xc2\xa2\x1a\x45\x9e\x7f\xa1\x9a\xa0\x3c\xff\x51\xa3\x81\x04\x8b\x06\x3b\x2f\xea\xdc\xf3\xb4\x87\x8e\x45\x2c\xa4\x14\x33\x68\x24\xea\x2b\x38\x89\x0d\x1d\x6f\x3d\x1d\xec\x74\xeb\xe6\x68\x00\xbe\x87\x13\x0a\x73\xd1\x1e\xfc\x6a\x44\x5e\x10\x5d\x02\xbe\x77\x2b\xd0\x7b\xdd\x6e\xb7\x14\xbd\x65\xd8\xe7\x41\x58\x97\xc3\x16\x07\xe8\xf8\xc0\xd8\x70\x67\xf9\x5f\x8d\xce\xd4\x01\x68\xf7\xc5\xde\xcb\x57\xaf\x57\xc4\xda\xca\x42\x6a\x4a\xc9\xec\x56\xb2\xfe\xba\x6d\xc0\x3f\xd9\xfa\xdb\xc8\xd6\xc5\xd3\x72\xed\xbd\xf4\x00\x20\x78\x2b\x28\x7f\xef\x0c\xef\xf3\x29\x44\x9c\xad\xb9\x1b\x3d\xcf\x65\x57\x43\x61\xdc\xb2\xf6\x50\x0a\xb6\x6a\x75\x62\x4c\x3e\x44\xa7\x62\xc6\x89\xdc\xcc\xeb\x76\x8b\x74\x4f\x40\xa4\xe0\x50\x1f\xd1\x31\x87\x49\x1f\xb3\xb2\xff\x92\xf2\x35\xa7\xdf\x8f\xea\xee\x89\xa5\x49\x41\xf9\xd3\x56\xc7\x60\xb6\x33\x4d\x52\xf0\x1c\xa7\xec\x49\x02\xe5\x0d\xc2\x88\x93\x99\x15\xd0\x92\x04\x22\x72\xca\x42\xe8\xd8\x20\xc6\xac\x4e\x3a\xbf\x2b\xc1\x7d\xe2\x9e\x64\x53\x3b\x5f\x8a\xd8\xfd\x33\xa3\x1a\xbb\x55\x32\xb8\x66\x02\x66\xaa\xfa\x5e\xbf\xdb\xad\xa7\xa0\x32\x04\xf7\x73\x39\x9c\x46\x68\x26\x57\x2c\xb0\x94\x74\x70\x92\x44\x56\x07\xe6\x54\x89\x66\xed\x96\x91\xa2\x9d\x0d\x98\x9f\x67\x67\x85\x58\x67\x67\xe6\xe8\x85\x0b\x8d\x30\x63\x62\x06\xdd\xa7\x3d\xa0\xb1\x4d\x2a\x83\xce\x95\x01\x03\x65\x1b\x34\x43\x13\x10\x94\x5a\xe6\x11\xf4\x73\x30\x6f\x86\xe7\x9b\xb6\xb3\xe3\xe6\x3c\x86\x39\xc6\xb5\x53\xfe\xcf\xe8\x39\xc9\xc3\xc2\x1e\x64\x15\x71\x66\x8e\x77\x4c\x00\x7e\x57\xf4\xba\x1e\x4d\x56\xb4\x58\xd7\xdf\x55\x7a\xab\x68\xcb\x7c\x06\xdd\xed\xc6\x46\x7e\xa1\x9c\x42\x35\x46\x31\xaf\x85\x0c\x68\x14\x11\xfe\xaf\xb2\x33\x74\xcc\x11\x94\xc7\x23\x22\xaa\x12\x86\xe7\xef\x6c\x85\xf5\x86\x68\xb3\xfa\xbe\xe7\xb4\xea\x7e\x3c\xa7\xd8\xa9\x2a\xb2\x4c\x31\xa3\xc6\xa8\x1e\x79\x20\x70\xe5\xfc\x03\x96\x40\xa7\xf3\xc0\x73\x98\x8a\x73\xe2\x59\xb3\x92\x2f\xdc\x95\xf2\x4d\xa1\x10\xfb\x5e\x83\xa9\x69\xe7\x28\x7f\x7c\xfb\x88\x4f\x5c\x77\xa3\x5a\xa6\x26\x7f\x72\x32\xa9\x09\x97\x36\xb6\x75\xf2\x53\x33\x6e\x3f\x56\xdf\x46\xea\x0d\xb7\xde\xcd\xc2\xed\x2a\x8d\x36\x6b\xb2\x8a\xd0\xeb\x31\x9d\x2f\x36\x27\xdd\x7f\x1a\xf6\xd8\x3b\x75\x33\x80\x27\x13\x0b\x89\xc2\x7c\x5f\x72\x72\x72\x2e\xda\xf2\xb8\xef\x3a\x3a\xaf\xa3\x89\x08\x23\x9a\x34\x00\x7a\x69\x07\x0d\x0c\xec\x55\x77\xb2\xac\xe4\xea\x31\x7a\x17\x49\xfa\xa6\x5e\x53\x2d\x3d\x57\x89\xbd\xa4\x56\xae\xde\x7d\xfc\xd5\x4b\xe4\xbb\x88\x41\xdf\x5a\x4d\xe6\x7d\x30\x01\xd5\x68\xca\xeb\x85\x95\x23\xe3\x6d\x84\x56\x15\xd5\x4a\xe1\x55\x71\x31\xd7\x04\xa6\x83\x29\x4e\x72\x57\xf0\x55\x43\x5e\x54\x4b\x7f\x9b\xb3\xd6\x5b\xa4\x87\xb8\x61\x54\x51\x7e\x43\xd5\x44\x59\xb4\x9a\xfb\x3c\xb8\xd0\xad\x9a\xe2\x9a\xa1\x6b\x84\xcb\x7d\xfd\x36\x42\xb7\x8a\x6a\xb5\xd0\x6d\xd6\x3d\x74\x40\x74\x94\x2f\xec\xab\x60\x4b\xfc\xa0\xf8\xcf\xae\xfb\xf2\xdb\xb0\xd2\x64\xdb\x34\x29\x58\x4d\x12\x94\x10\x1e\x19\xe1\xe9\xa2\x9d\x57\x46\x3d\x20\xa4\x14\x50\x99\xab\x07\xee\xd2\xc8\x05\x3e\xbc\xe4\x90\xa2\x74\x61\xe8\xc5\x95\xc9\x7b\x03\x1e\xc5\x39\x42\x7e\xe0\x00\xac\xfc\x27\x11\xd9\x79\x8a\x6b\xda\x62\x66\xdd\xac\x5e\xa3\x2e\xc8\xd1\x6c\x42\xc0\x86\xc5\xbb\x09\xa6\x65\xec\x15\x97\xbf\xe6\xe8\x02\x94\xd7\x46\xef\x65\x04\x43\x86\x51\x6e\xea\x60\x6e\x2f\x61\x0a\x1f\x6f\x6c\xb2\x28\x0f\x59\x1a\x91\x57\xf6\x60\x3d\xf2\x1d\x53\x54\xaf\x6e\x6b\x88\xf7\x33\x16\x16\x50\x76\x40\x1f\x2d\xd1\x6f\x5a\x24\x64\x84\x53\x06\x3e\xca\xd9\xdc\x92\xc1\x08\x9d\xba\x32\x1a\x2f\x95\x44\xa7\x92\x97\x2f\xa7\x91\xbd\xd0\xb5\xd0\xa0\xb3\x67\x6a\x79\xab\xa4\x34\xc8\xee\xc3\x4b\xb9\x26\x63\xa7\x21\xaf\xe1\x3d\x32\x33\x41\x83\x23\x45\xf4\xa6\xd1\xe9\x28\x65\x08\xfa\x67\xd0\xe6\x18\x54\xd7\x86\x38\xcd\xe0\x82\xa6\xcf\xba\x67\xab\xc1\x89\xf1\xc5\xb5\xc0\x1c\xe2\x0b\xb3\x2e\x41\x8a\x7e\x26\x57\xa2\xd9\x59\x15\x4e\x0e\x89\xc6\xe6\xf8\xac\x9f\x5d\xd5\xb8\xce\x98\xbd\x95\xb3\x6e\x82\x31\x6f\xf4\x08\x59\x5c\xa9\x67\x3c\xee\x79\x5b\xb0\xef\x3c\xd9\x98\xad\x8b\xd7\xf9\xb2\x78\x83\xc8\xb7\x8f\xd5\x44\xc5\x48\x25\x24\x34\x29\xb4\x2c\xe9\x43\x10\xf1\xea\xdd\xe5\x23\x89\xc5\xd4\xec\x2e\x96\x7e\xb3\xfc\x72\x88\xfb\x2a\xc4\x8c\x32\x06\x7b\x05\x62\x82\x83\x6b\xa2\x09\x36\xb3\xf2\x77\xa1\xbc\x2f\x40\xb5\x57\xdb\x70\xf3\x95\x5b\x7f\x05\x00\x00\xff\xff\x72\x44\xd0\x0d\xc7\x26\x00\x00")

func usersRamlBytes() ([]byte, error) {
	return bindataRead(
		_usersRaml,
		"users.raml",
	)
}

func usersRaml() (*asset, error) {
	bytes, err := usersRamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "users.raml", size: 9927, mode: os.FileMode(420), modTime: time.Unix(1457621611, 0)}
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
	"companies.raml": companiesRaml,
	"contracts.raml": contractsRaml,
	"organizations.raml": organizationsRaml,
	"securitySchemes/oauth_2_0.raml": securityschemesOauth_2_0Raml,
	"userorganizations.raml": userorganizationsRaml,
	"users.raml": usersRaml,
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
	"companies.raml": &bintree{companiesRaml, map[string]*bintree{}},
	"contracts.raml": &bintree{contractsRaml, map[string]*bintree{}},
	"organizations.raml": &bintree{organizationsRaml, map[string]*bintree{}},
	"securitySchemes": &bintree{nil, map[string]*bintree{
		"oauth_2_0.raml": &bintree{securityschemesOauth_2_0Raml, map[string]*bintree{}},
	}},
	"userorganizations.raml": &bintree{userorganizationsRaml, map[string]*bintree{}},
	"users.raml": &bintree{usersRaml, map[string]*bintree{}},
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

