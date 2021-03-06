// Code generated by go-bindata.
// sources:
// 01_init_database.down.sql
// 01_init_database.up.sql
// 02_add_directory_table.down.sql
// 02_add_directory_table.up.sql
// 03_add_curse_release_game_version_table.down.sql
// 03_add_curse_release_game_version_table.up.sql
// DO NOT EDIT!

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

var __01_init_databaseDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\xf0\x74\x53\x70\x8d\xf0\x0c\x0e\x09\x56\x48\x2e\x2d\x2a\x4e\x8d\x4f\x4c\x49\xc9\xcf\xb3\xe6\xe5\xc2\xa3\xa4\x28\x35\x27\x35\xb1\x38\xd5\x9a\x97\x0b\x10\x00\x00\xff\xff\x0c\xd2\x06\x06\x48\x00\x00\x00")

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

	info := bindataFileInfo{name: "01_init_database.down.sql", size: 72, mode: os.FileMode(438), modTime: time.Unix(1531569018, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __01_init_databaseUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x90\x41\x6b\x84\x30\x10\x85\xef\x82\xff\x61\x8e\x2e\xf4\x1f\xec\xc9\xea\xb8\x84\xda\x58\x62\x0a\xee\x29\x04\x33\x05\x21\x46\x88\x4a\xdb\x7f\x5f\x22\x75\xdd\xc0\xc2\xe2\x69\xe6\x39\xef\xbd\x7c\x85\xc0\x5c\x22\xc8\xfc\xb5\x46\x60\x15\xf0\x46\x02\x76\xac\x95\x2d\xf4\xab\x9f\x49\x69\x63\x26\x07\x59\x9a\x00\x0c\x06\x18\x97\x78\x41\x01\x1f\x82\xbd\xe7\xe2\x0a\x6f\x78\x7d\x09\x92\xd3\x23\x81\xc4\x4e\x6e\xd3\xea\xed\x31\xcc\xeb\x38\x6a\xff\x7b\x2c\xcc\xf4\xed\xec\xa4\x4d\x3f\xad\x6e\xd9\x1d\xd3\xe4\x74\x4e\x93\xf0\x3d\x6d\xe4\xc9\x92\x9e\xe9\x69\xa7\xaf\xc1\xd2\xad\xd7\x66\xc3\x3f\xeb\x7a\x93\x7a\x4f\x7a\x21\xa3\xf4\x2d\x3f\xd6\xf7\x07\xc4\xdb\x0d\x85\xba\x4b\x8c\xd4\x61\x56\xda\x2e\xe4\x9d\x5e\xe8\xf1\x1f\x55\x23\x90\x5d\x78\x28\x98\xed\x5e\x27\x10\x58\xa1\x40\x5e\x60\x44\x3c\x0b\x52\xc3\xa1\xc4\x1a\x25\x42\x91\xb7\x45\x5e\xe2\x4e\xe9\x9f\x11\xe3\x25\x76\xf7\x57\x6a\x70\x86\x7e\xc2\x5d\x84\xea\x48\x3b\xff\x05\x00\x00\xff\xff\x10\xaa\x31\xba\xf0\x01\x00\x00")

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

	info := bindataFileInfo{name: "01_init_database.up.sql", size: 496, mode: os.FileMode(438), modTime: time.Unix(1531588232, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __02_add_directory_tableDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\xf0\x74\x53\x70\x8d\xf0\x0c\x0e\x09\x56\x48\x2e\x2d\x2a\x4e\x8d\x2f\x4a\xcd\x49\x4d\x2c\x4e\x8d\x4f\xc9\x2c\x4a\x4d\x2e\xc9\x2f\xaa\xb4\xe6\xe5\x02\x04\x00\x00\xff\xff\xb3\xe8\x89\x6b\x2f\x00\x00\x00")

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

	info := bindataFileInfo{name: "02_add_directory_table.down.sql", size: 47, mode: os.FileMode(438), modTime: time.Unix(1531569018, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __02_add_directory_tableUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x8f\xc1\x6a\x84\x30\x14\x45\xf7\x81\xfc\xc3\x5b\x2a\xf8\x07\xae\xd2\x78\x95\xd0\x34\x96\xf8\x0a\xba\x92\xa2\x59\x08\xa5\x85\xd8\x42\xe7\xef\x87\x0c\x33\xa3\x23\xcc\xf6\xdd\x73\x1f\xf7\x68\x0f\xc5\x20\x56\x2f\x16\x64\x6a\x72\x2d\x13\x7a\xd3\x71\x47\xd3\x5f\x5c\xc3\x18\xc3\x57\xf8\x5c\xc3\x38\x2f\x31\x4c\xbf\x3f\xf1\x94\x49\x41\x74\xbb\x2e\x33\x19\xc7\x68\xe0\x2f\x4d\xf7\x61\x6d\x91\xf2\x3b\x4d\x8c\x9e\x1f\xb3\x77\x6f\xde\x94\x1f\xe8\x15\x03\x65\xdb\xa3\x62\x2b\xe5\x09\xab\x5b\x0f\xd3\xb8\x84\xed\xa8\x9c\x3c\x6a\x78\x38\x8d\xc3\xc2\x2c\x85\xad\xa3\x0a\x16\x0c\xd2\xaa\xd3\xaa\x82\x14\x79\x29\x85\x14\x57\x4f\xe3\x2a\xf4\x07\xb3\xe5\x7b\x0e\xff\xc9\x2f\xb5\x9f\x49\xef\x16\x94\x52\x9c\x03\x00\x00\xff\xff\x15\x88\x90\x92\x36\x01\x00\x00")

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

	info := bindataFileInfo{name: "02_add_directory_table.up.sql", size: 310, mode: os.FileMode(438), modTime: time.Unix(1531588152, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __03_add_curse_release_game_version_tableDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\xf0\x74\x53\x70\x8d\xf0\x0c\x0e\x09\x56\x48\x2e\x2d\x2a\x4e\x8d\x2f\x4a\xcd\x49\x4d\x2c\x4e\x8d\x4f\x4f\xcc\x4d\x8d\x2f\x4b\x2d\x2a\xce\xcc\xcf\xb3\xe6\xe5\x02\x04\x00\x00\xff\xff\x61\x70\xa7\xf7\x32\x00\x00\x00")

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

	info := bindataFileInfo{name: "03_add_curse_release_game_version_table.down.sql", size: 50, mode: os.FileMode(438), modTime: time.Unix(1531569018, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __03_add_curse_release_game_version_tableUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x8f\x41\x6a\xc3\x30\x10\x45\xf7\x02\xdd\x61\x96\x36\xe4\x06\x59\xa9\xf2\x77\x10\x55\xe5\x22\xab\x90\xac\x44\x88\x87\x20\x68\x63\x90\x5b\xd3\xde\xbe\xa8\xb4\xd4\x36\x59\xbf\xf7\x86\xf9\xda\x43\x05\x50\x50\x0f\x16\x64\x5a\x72\x5d\x20\x1c\x4d\x1f\x7a\xba\x7c\xe4\x89\x63\xe6\x57\x3e\x4f\x1c\xaf\xe7\x37\x8e\x33\xe7\x29\x8d\xb7\x4a\x0a\xa2\x3f\x90\x06\x32\x2e\xe0\x00\xff\x13\xbb\x17\x6b\x77\x85\x2f\x83\xfb\xc6\xb3\x37\x4f\xca\x9f\xe8\x11\x27\xaa\xfe\xcf\xed\x56\x69\x5d\xcc\xb6\xf3\x30\x07\x57\xcc\x85\x58\x93\x47\x0b\x0f\xa7\xb1\xf9\xb6\x2a\xb0\x73\xd4\xc0\x22\x80\xb4\xea\xb5\x6a\x20\x45\xbd\x97\x42\x8a\xdf\xcd\xc6\x35\x38\x6e\x56\xa6\xdb\xc0\x9f\xf1\x3a\x97\x78\x4d\x86\x94\xf9\xf2\x3e\xe6\xaf\xe5\x03\xfb\xef\x00\x00\x00\xff\xff\xd3\x95\xee\x30\x3f\x01\x00\x00")

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

	info := bindataFileInfo{name: "03_add_curse_release_game_version_table.up.sql", size: 319, mode: os.FileMode(438), modTime: time.Unix(1531588153, 0)}
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

