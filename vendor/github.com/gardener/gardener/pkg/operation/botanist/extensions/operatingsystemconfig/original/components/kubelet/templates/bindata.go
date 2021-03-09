// Code generated by go-bindata. (@generated) DO NOT EDIT.

// Package templates generated by go-bindata.// sources:
// scripts/health-monitor.tpl.sh
package templates

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
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// ModTime return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _scriptsHealthMonitorTplSh = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x58\x7b\x4f\xe3\x48\x12\xff\xdf\x9f\xa2\xe8\xf5\x41\xa2\x59\xdb\x49\x76\x6f\x6e\xc5\x4c\x56\x87\x66\x91\x76\x34\x73\xec\x88\x70\x0f\x89\x20\xab\x63\x97\xe3\x1e\x9c\x6e\x5f\x77\x19\xc8\x42\xbe\xfb\xaa\xfd\x88\x6d\x12\x98\x87\xe0\x1f\xfa\x59\xfd\xab\xd7\xaf\x2a\xfe\xe1\x20\x58\x08\x19\x2c\xb8\x49\x1d\x83\x04\x9e\x02\xa9\x0a\x69\x90\x9a\x69\x2e\x72\x4c\xb8\xc8\x1c\x27\x29\x64\x44\x42\x49\xb8\x2e\x16\x98\x21\x85\x2b\x25\x05\x29\x2d\xe4\x12\xee\x1d\x00\x8c\x52\x05\xec\xbf\x5c\x10\x24\x4a\xc3\x04\x56\x42\x16\x84\xa6\x9c\xd5\x77\x80\x14\x2c\x10\x1a\x51\x3c\x63\x0e\x80\xc9\x10\x73\x18\x4f\x46\x0e\x40\xa6\x22\x9e\x81\xa7\x61\xc5\xef\x42\x83\x91\x92\xb1\x99\x8e\xdb\x1d\x55\x50\x5e\xd0\x94\x31\xc7\x01\xe8\x21\x8a\x28\x2b\x61\x00\x04\x2a\xa7\x52\xad\x66\xd9\xf3\xca\x91\x92\x89\x58\xc2\xfd\x3d\xf8\x39\xa7\xf4\x43\x85\xe8\xc3\x76\xe7\x1c\x79\x06\x9b\x0d\x30\xf7\x9f\x16\xd5\xa6\xf7\x82\x46\x43\x5c\x53\xd8\xe8\x51\xbd\x94\x5f\x8b\x2c\x03\x2f\x01\x56\xaf\xef\x5e\xcc\x39\x45\x69\x28\x24\xa1\x96\x3c\x0b\x45\x5e\x5f\xad\xac\xf5\xef\x3c\xe6\x64\x0d\x78\xa6\x62\x04\xb5\xf8\x8c\x11\x81\x3b\x81\x5b\x41\x29\xbc\xaf\x2f\xbd\xff\x04\xee\x4f\x3e\x2b\xaf\x45\x85\xce\x60\x5e\x0e\x01\xbc\xff\x7d\x3a\xb9\x78\xf7\x7b\x3b\xff\x1d\xd8\x3b\x25\x09\x25\x79\x17\xeb\x1c\x8f\x81\xe7\x79\x26\x22\x6e\x91\x04\x86\x34\x27\x5c\x8a\xc8\x5b\xa1\x5e\xa2\x57\x22\x7b\xf5\xd9\x28\xc9\x7a\x12\x4e\xa2\x08\x73\xea\xdf\xed\x9f\x62\xee\x38\xe0\xb9\x08\x6e\xc6\x81\x54\x31\x9a\xc0\x9d\x04\x86\x38\x15\xa6\x23\xc9\x8b\x39\x71\x60\xf7\x73\x56\x6d\xcd\xd9\xf1\xfd\x9c\xf1\x38\xd6\x68\x0c\xda\xe9\x65\x3b\x9f\xb3\x63\x98\x33\xf7\xa7\x39\xfb\x11\xe6\x8c\xd6\x39\xce\xd9\xf1\x9c\xb5\x16\x98\xb3\xcd\xd5\x66\xd3\x15\x1f\xf1\x08\x35\xc1\xdb\xc1\x82\x1b\x7c\xfd\x33\x78\x31\xbc\x7d\xfb\x16\xdc\x41\xe3\xf5\xda\xdf\x37\x02\x6f\x6d\x20\x5b\x15\xac\xdf\xa7\xf7\x7e\x94\x15\x86\x50\x9b\xcb\xd1\x55\x33\xf6\xad\x30\x91\x58\x7d\xd1\xe3\x05\xa5\x4a\x0b\x5a\x97\x3a\x6c\xc0\xf3\x34\xbf\x1d\x0e\x3b\x8f\x5f\xe3\x1a\x82\x1b\xae\x83\x4c\x2c\x82\xda\xf7\x41\x7e\x2d\x9a\xb1\x17\x65\xc2\x7a\x21\x2a\xb4\x46\x49\x7e\x8e\xab\x2e\x74\x0b\xfc\xfb\xae\xff\x0a\x41\x8c\x37\x81\x2c\xb2\x0c\x26\xbf\x1e\x8e\x9b\x78\x23\xb1\xc2\x44\xf3\x15\x4e\x5f\x8f\x6c\xb6\x90\x5a\x2e\x33\x0c\x29\xd5\x68\x52\x95\xc5\xd3\xbf\x3b\x00\x91\x2a\xe4\x36\x84\x43\x9e\x95\xc6\xb5\xe1\x17\x2e\x90\x6e\x11\x65\xa8\x91\xc7\xeb\x90\xcb\x38\x94\x8a\xea\x99\x8d\x45\x21\xc3\xf6\x85\x51\xfd\xde\x56\x52\x7b\x36\x11\xda\x50\xa8\xa2\x0a\x78\x54\x9d\xcd\xb8\x69\x5f\xad\xce\xd9\x90\xc0\x29\xbb\xd0\x05\x96\x99\x7c\x9b\x8a\x0c\xe1\x12\xc6\x70\xf5\x06\x62\x55\xaa\xfa\x03\xbc\x4b\x31\xba\x86\xdb\x14\x29\x45\x0d\x94\x62\x43\x23\x47\x06\x82\x14\x79\x46\xe9\x9f\x80\x32\xce\x95\x90\x04\x1a\x73\xa5\xc9\x40\x21\xab\x2d\x21\xd1\x98\x52\x92\x48\xe0\xa0\x61\x0e\x77\x50\xa6\x90\xb7\x02\xb7\x43\x31\x36\x87\x3d\x03\xde\x0c\x52\xa2\xfc\x38\x08\xc6\x93\x7f\xf8\x23\x7f\xe4\x8f\x8f\xc7\xa3\xc9\xcf\xbf\x6c\x5f\xb3\x36\x1f\xbe\xb1\x58\x64\xed\xce\x32\x91\xdd\x4a\x7a\x77\x89\xd5\x04\x03\xa2\x85\xb4\x3e\x60\xf5\x91\x47\x7c\x52\xaf\x56\x4c\xf8\x7a\x54\x4f\x23\x25\x49\xc8\x02\xcb\x69\x22\x9c\xf2\xbf\xcd\xb8\xb0\x62\x8a\x29\x6b\xa3\x7d\x89\x54\x6e\x19\xf0\xb2\xd2\x4c\x5a\x22\xa1\xf1\x85\x0a\x52\x65\x48\x5a\xc7\xb9\x83\x66\x38\x6c\x32\x62\xc8\x5a\xa1\x55\x9e\x5a\xa1\x95\x4e\x9d\x97\xe0\x01\x3e\xff\xdf\xd2\xf2\x91\x2f\x08\x57\x65\xe6\x54\xc7\x8f\x6a\x09\x22\x81\xcb\x4b\xf0\xfe\x04\xe6\x76\x84\x31\xb8\xba\x82\x87\x07\xbb\xf5\x68\x7d\x3a\x05\x66\xa3\xd8\x9e\xd8\xb5\x27\xeb\xf2\xa1\xad\x1f\x94\x0a\x03\x0d\x7a\x90\xca\xae\x16\x32\x06\x21\xcb\xb8\x30\x6b\x43\xb8\xfa\x11\x6e\xb9\xb0\xf1\xec\xb3\x9e\x41\x27\xad\x41\x5f\x2c\x01\xec\xdf\xb7\x24\xc1\x93\xfe\x7c\x26\xcc\x21\x56\x68\xea\xb8\x06\x2e\xbb\x05\xc1\xda\x12\x6a\xf2\x6c\x3d\x28\xf2\x6d\xa5\x79\xe4\xc6\xca\xec\x1d\x37\x6e\x89\xf8\xf2\x0a\x1e\xc0\x60\x86\x11\x0d\x7c\xcb\xbc\xd3\x69\x87\x77\xd9\x10\x1e\xa0\x39\x7c\xd4\x8d\x16\x91\x87\x78\xf7\x02\x6f\x9d\xde\x3d\xfb\xd6\xe3\xb8\xea\xa8\x58\x06\xd7\xe1\xe1\xce\x7e\x03\xeb\x89\xd0\x6a\xf2\x32\xe5\xa6\x0c\xa4\xca\xbe\x18\xef\x58\x58\xdb\x95\x16\x5e\xcf\xe6\xb0\x46\xf2\xd9\x9b\x5a\x70\x49\x30\x15\x0c\xf7\xfe\xc3\x2f\xb3\xf0\xec\x8f\xdf\x4e\xc3\xf7\x9f\xc2\xf7\x67\x17\xa7\xe7\x67\x27\x1f\xc3\x8f\x27\xb3\x8b\x70\x76\x7a\x7a\xf6\xea\x6e\xf3\x18\x56\x03\xac\x8a\x03\x91\x94\x84\x09\x06\xb1\x87\x87\xb9\xcf\x0b\x66\x10\x71\x69\x9b\xab\xc2\x60\xbc\x45\x56\x62\x13\xf9\x16\xb6\x49\xd5\x2d\x3c\xc0\x52\x63\x0e\x5f\x10\xd8\xad\x35\x8f\xf0\x36\x88\x3f\x7e\x1f\x50\x61\xc0\x90\xed\x9c\x8a\xdc\x23\x65\x8b\x2c\x76\x00\x03\x18\xd4\x37\xa8\xbb\xec\xf6\x6d\xb5\xbc\xba\xbf\x19\xb2\x8e\xcc\x32\x38\x4a\x16\xfc\x5a\x7e\x5b\x21\x71\x5b\xff\x7d\x7b\xeb\xa8\x27\x4c\x24\x7b\xfa\x3a\xb7\x7a\xb6\x96\x5b\xf2\xd4\x17\x2c\xb1\x63\xd5\xc6\xae\xb3\x22\x8a\xd0\x98\xa4\xc8\xb2\x35\x14\xb6\x47\xc4\xb8\xdb\x22\xfa\xac\x77\xa9\xc7\x2b\xb5\x9c\xcc\xe0\x1e\xc1\x27\x12\x50\x6b\xa5\xa1\x66\xa7\xb8\xae\xbc\x45\xd3\x87\x5a\xf6\x79\xea\xa1\x44\x38\x3b\xc3\xed\xe0\x99\x7e\x56\x54\x79\x96\x2b\x63\xc4\x22\x43\x1f\xce\xab\xe2\x67\xcf\xd5\xc9\xd8\xa6\xd2\x73\x85\x71\xb2\x97\x48\x31\xeb\xa4\xdf\x7e\x96\xe8\xb3\xc0\x5d\xc9\xa7\xcf\xbb\x66\xba\x47\xd0\xb3\xb4\xdd\x50\x76\x59\x02\xa0\xa6\xc0\xaa\x0b\x33\x50\x57\x17\x20\x5d\x20\x70\x19\x43\xc2\x33\x53\x8d\x34\x2e\x94\x22\xf8\xcf\xbf\x6c\x50\xa5\x3c\xcf\x51\x62\x0c\xa4\x14\xa8\x84\x50\xfa\x0d\x09\xee\x2d\xce\x3b\x4c\x6b\x9b\x19\x61\x3b\xf6\xbd\x54\x7b\x6e\xb1\x55\x2c\xbb\x2d\xde\x3d\xd3\x54\x64\xcb\xdc\xa6\x48\x1f\x4c\xa1\xea\xd1\x76\x29\xab\x3a\xea\x7e\x65\x09\xb4\xe5\x7e\xb4\x2b\xe4\xeb\x4b\xa8\x3b\xb0\x49\x00\xaf\xfe\x66\x86\x3b\x2c\x34\xb3\x01\x03\xa4\x79\x74\x6d\x43\xea\x39\x4f\x74\xe2\xb9\x8d\xdc\x6e\xaa\x7c\xa3\x5a\x07\x4f\xa8\xd5\x18\xf2\xa9\xd6\xb7\x32\xad\xdb\x36\x49\x7b\xa8\xe0\x85\x9a\x15\x77\x30\x78\x19\x49\xaf\xc6\xc3\xe1\x1e\x4a\x79\x29\x98\x2f\x23\xa7\x4f\x8c\xb5\x37\x5f\x46\x34\x78\x4b\x04\xf7\xf1\x0f\xab\xfd\xbe\x03\x30\x45\xac\xea\xe4\xee\xed\x74\x48\xf4\x49\x46\x75\xbe\x3f\x16\x0f\x0f\xc1\x1d\x0c\x3a\xd9\xe2\x7d\xed\xfd\xe1\xb0\x56\x70\xab\xef\x8e\x66\x2f\xdc\x3e\x7f\x7b\x03\xdd\xc4\xdc\x39\x1a\x24\xfa\x42\xb2\x6f\x19\x61\x9b\xf4\xad\x69\x9f\xfe\x49\xda\xe4\x64\x8f\xec\xab\xf2\xe3\xce\x3e\x9e\x9e\x7e\x0a\x67\xa7\xef\xfe\x38\xfb\x6d\xe6\x00\xc4\x4a\xa2\xb3\x71\x9c\xde\xfa\x74\x3c\x72\xba\xbc\x54\xfd\xf2\x83\xce\xc7\xb2\xce\xf7\x30\xe6\xec\x7e\x4c\x73\xfe\x0a\x00\x00\xff\xff\xf5\x35\xb2\x29\x95\x13\x00\x00")

func scriptsHealthMonitorTplShBytes() ([]byte, error) {
	return bindataRead(
		_scriptsHealthMonitorTplSh,
		"scripts/health-monitor.tpl.sh",
	)
}

func scriptsHealthMonitorTplSh() (*asset, error) {
	bytes, err := scriptsHealthMonitorTplShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "scripts/health-monitor.tpl.sh", size: 5013, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
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
	"scripts/health-monitor.tpl.sh": scriptsHealthMonitorTplSh,
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
// AssetDir("foo.txt") and AssetDir("nonexistent") would return an error
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
	"scripts": &bintree{nil, map[string]*bintree{
		"health-monitor.tpl.sh": &bintree{scriptsHealthMonitorTplSh, map[string]*bintree{}},
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
	canonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(canonicalName, "/")...)...)
}