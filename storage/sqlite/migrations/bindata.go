// Code generated by go-bindata. DO NOT EDIT.
// sources:
// 01_init_database.down.sql
// 01_init_database.up.sql
// 02_add_directory_table.down.sql
// 02_add_directory_table.up.sql
// 03_add_curse_release_game_version_table.down.sql
// 03_add_curse_release_game_version_table.up.sql
package migrations

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

var __01_init_databaseDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\xf0\x74\x53\x70\x8d\xf0\x0c\x0e\x09\x56\x48\x2e\x2d\x2a\x4e\x8d\x4f\x4c\x49\xc9\xcf\xb3\xe6\xc2\xa3\xa2\x28\x35\x27\x35\xb1\x38\xd5\x9a\x0b\x10\x00\x00\xff\xff\xd7\xd0\xf8\xf1\x46\x00\x00\x00")

func _01_init_databaseDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__01_init_databaseDownSql,
		"01_init_database.down.sql",
	)
}

func _01_init_databaseDownSql() (*asset, error) {
	bytes, err := _01_init_databaseDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "01_init_database.down.sql", size: 70, mode: os.FileMode(420), modTime: time.Unix(1529317648, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __01_init_databaseUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x90\x41\x6b\xc4\x20\x10\x85\xef\xfe\x8a\x39\x26\xb0\xff\xa0\xa7\x74\x99\x2c\xd2\xd4\x2d\xae\x85\xdd\x53\x18\xe2\x14\x04\x63\xc0\x28\xa5\xff\xbe\x58\x58\x69\x4a\xca\xde\x1c\xdf\xf3\xf9\xcd\x13\x47\x8d\x9d\x41\x30\xdd\xf3\x80\x20\x7b\x50\x67\x03\x78\x95\x17\x73\x81\x29\xc7\x95\x47\xb2\x76\x09\xd0\x08\x00\x67\x41\x2a\x83\x27\xd4\xf0\xa6\xe5\x6b\xa7\x6f\xf0\x82\xb7\x83\x00\x08\x34\x33\x18\xbc\x9a\x32\xe4\xe8\xeb\x79\xcd\xf3\x4c\xf1\xab\xce\x76\xf9\x0c\x7e\x21\x3b\x2d\x39\xa4\x7b\x9a\x68\x9f\x84\x78\x0c\x12\xd9\x33\xad\xfc\x17\xa5\xf8\xd4\xfb\x30\x94\xf8\x0f\xe7\xb9\xb2\x6c\x94\x29\x32\x25\xb6\x23\xa5\xdd\x87\x77\xe6\xcd\xe5\xcf\xe6\xe3\x3f\x5f\xb9\x75\x24\x9f\x38\x06\x4a\xbc\x6b\xf8\x55\x11\x34\xce\x1e\x6a\x5c\x5b\xd4\xfe\xac\x51\x9e\x54\x51\x9b\x2a\x80\xc6\x1e\x35\xaa\x23\x6e\xca\x6f\x9c\x6d\x4b\x47\xdf\x01\x00\x00\xff\xff\x78\xb0\x06\x80\xac\x01\x00\x00")

func _01_init_databaseUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__01_init_databaseUpSql,
		"01_init_database.up.sql",
	)
}

func _01_init_databaseUpSql() (*asset, error) {
	bytes, err := _01_init_databaseUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "01_init_database.up.sql", size: 428, mode: os.FileMode(420), modTime: time.Unix(1529783136, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __02_add_directory_tableDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\xf0\x74\x53\x70\x8d\xf0\x0c\x0e\x09\x56\x48\x2e\x2d\x2a\x4e\x8d\x2f\x4a\xcd\x49\x4d\x2c\x4e\x8d\x4f\xc9\x2c\x4a\x4d\x2e\xc9\x2f\xaa\xb4\xe6\x02\x04\x00\x00\xff\xff\xdf\xa2\xb7\x73\x2e\x00\x00\x00")

func _02_add_directory_tableDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__02_add_directory_tableDownSql,
		"02_add_directory_table.down.sql",
	)
}

func _02_add_directory_tableDownSql() (*asset, error) {
	bytes, err := _02_add_directory_tableDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "02_add_directory_table.down.sql", size: 46, mode: os.FileMode(420), modTime: time.Unix(1529575953, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __02_add_directory_tableUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\x8e\xc1\x0a\x82\x40\x18\x84\xef\xfb\x14\x73\x74\xc1\x37\xe8\x64\x32\xca\x92\xad\xf1\xfb\x07\x7a\x92\xd0\x3d\x2c\x04\xc1\x5a\x87\xde\x3e\x0a\xca\xea\x3a\xf3\x7d\xc3\x94\xc2\x42\x09\x2d\xb6\x0d\xe1\x2a\xf8\x56\xc1\xde\x75\xda\x61\xba\xa5\x25\x8c\x29\x9c\xc3\x69\x09\xe3\x1c\x53\x98\xae\x97\x74\xcf\x0c\xf0\x0e\xe3\x0c\xe7\x95\x35\xe5\x25\xfa\x63\xd3\xe4\x06\xf8\xb0\x50\xf6\xfa\x53\x1d\xc4\xed\x0b\x19\xb0\xe3\x80\x6c\x9d\xc9\x57\xc7\x1a\xa0\x6a\x85\xae\xf6\x4f\xea\x0b\xb2\x10\x56\x14\xfa\x92\x7f\xef\xb2\x38\x5b\x63\x37\xe6\x11\x00\x00\xff\xff\x35\x52\x53\x63\xce\x00\x00\x00")

func _02_add_directory_tableUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__02_add_directory_tableUpSql,
		"02_add_directory_table.up.sql",
	)
}

func _02_add_directory_tableUpSql() (*asset, error) {
	bytes, err := _02_add_directory_tableUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "02_add_directory_table.up.sql", size: 206, mode: os.FileMode(420), modTime: time.Unix(1529575924, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __03_add_curse_release_game_version_tableDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\xf0\x74\x53\x70\x8d\xf0\x0c\x0e\x09\x56\x48\x2e\x2d\x2a\x4e\x8d\x2f\x4a\xcd\x49\x4d\x2c\x4e\x8d\x4f\x4f\xcc\x4d\x8d\x2f\x4b\x2d\x2a\xce\xcc\xcf\xb3\xe6\x02\x04\x00\x00\xff\xff\x1b\xd9\x29\x54\x31\x00\x00\x00")

func _03_add_curse_release_game_version_tableDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__03_add_curse_release_game_version_tableDownSql,
		"03_add_curse_release_game_version_table.down.sql",
	)
}

func _03_add_curse_release_game_version_tableDownSql() (*asset, error) {
	bytes, err := _03_add_curse_release_game_version_tableDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "03_add_curse_release_game_version_table.down.sql", size: 49, mode: os.FileMode(420), modTime: time.Unix(1529783424, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __03_add_curse_release_game_version_tableUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\xce\xcd\xaa\xc2\x30\x10\xc5\xf1\x7d\x9e\xe2\x2c\x1b\xe8\x1b\xdc\x55\x6f\x39\x2d\xc1\x9a\xca\x34\x82\x5d\x95\x62\x07\x29\xf8\x01\x09\xfa\xfc\xa2\x20\x56\x71\x3b\xfc\xfe\xc3\x29\x85\x45\x20\x42\xf1\xdf\x10\xae\x82\x6f\x03\xb8\x73\x5d\xe8\xb0\xbf\xc6\xa4\x43\xd4\xa3\x8e\x49\x87\xc3\x78\xd2\xe1\xa6\x31\xcd\x97\x73\x66\x80\xd7\x7d\x9e\xe0\x7c\x60\x4d\x79\xb6\x7e\xdb\x34\xb9\x01\x96\xfc\x27\xd8\x88\x5b\x17\xd2\x63\xc5\x1e\xd9\xfb\x59\xfe\x51\x5a\x03\x54\xad\xd0\xd5\xfe\x01\x17\xce\x42\x58\x51\xe8\x4b\x7e\x2d\xcd\xe6\xc9\x1a\xfb\x67\xee\x01\x00\x00\xff\xff\x33\x5d\xe3\x14\xda\x00\x00\x00")

func _03_add_curse_release_game_version_tableUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__03_add_curse_release_game_version_tableUpSql,
		"03_add_curse_release_game_version_table.up.sql",
	)
}

func _03_add_curse_release_game_version_tableUpSql() (*asset, error) {
	bytes, err := _03_add_curse_release_game_version_tableUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "03_add_curse_release_game_version_table.up.sql", size: 218, mode: os.FileMode(420), modTime: time.Unix(1529783538, 0)}
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
	"01_init_database.down.sql": _01_init_databaseDownSql,
	"01_init_database.up.sql": _01_init_databaseUpSql,
	"02_add_directory_table.down.sql": _02_add_directory_tableDownSql,
	"02_add_directory_table.up.sql": _02_add_directory_tableUpSql,
	"03_add_curse_release_game_version_table.down.sql": _03_add_curse_release_game_version_tableDownSql,
	"03_add_curse_release_game_version_table.up.sql": _03_add_curse_release_game_version_tableUpSql,
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
	"01_init_database.down.sql": &bintree{_01_init_databaseDownSql, map[string]*bintree{}},
	"01_init_database.up.sql": &bintree{_01_init_databaseUpSql, map[string]*bintree{}},
	"02_add_directory_table.down.sql": &bintree{_02_add_directory_tableDownSql, map[string]*bintree{}},
	"02_add_directory_table.up.sql": &bintree{_02_add_directory_tableUpSql, map[string]*bintree{}},
	"03_add_curse_release_game_version_table.down.sql": &bintree{_03_add_curse_release_game_version_tableDownSql, map[string]*bintree{}},
	"03_add_curse_release_game_version_table.up.sql": &bintree{_03_add_curse_release_game_version_tableUpSql, map[string]*bintree{}},
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

