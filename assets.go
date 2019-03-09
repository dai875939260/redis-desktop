// Code generated by Lorca. DO NOT EDIT.
package main

import (
	"bytes"
	"errors"
	"net/http"
	"os"
	"time"
)

var assets = map[string][]byte{}

var FS = &fs{}

type fs struct {}

func (fs *fs) Open(name string) (http.File, error) {
	if name == "/" {
		return fs, nil;
	}
	b, ok := assets[name]
	if !ok {
		return nil, os.ErrNotExist
	}
	return &file{name: name, size: len(b), Reader: bytes.NewReader(b)}, nil
}

func (fs *fs) Close() error { return nil }
func (fs *fs) Read(p []byte) (int, error) { return 0, nil }
func (fs *fs) Seek(offset int64, whence int) (int64, error) { return 0, nil }
func (fs *fs) Stat() (os.FileInfo, error) { return fs, nil }
func (fs *fs) Name() string { return "/" }
func (fs *fs) Size() int64 { return 0 }
func (fs *fs) Mode() os.FileMode { return 0755}
func (fs *fs) ModTime() time.Time{ return time.Time{} }
func (fs *fs) IsDir() bool { return true }
func (fs *fs) Sys() interface{} { return nil }
func (fs *fs) Readdir(count int) ([]os.FileInfo, error) {
	files := []os.FileInfo{}
	for name, data := range assets {
		files = append(files, &file{name: name, size: len(data), Reader: bytes.NewReader(data)})
	}
	return files, nil
}

type file struct {
	name string
	size int
	*bytes.Reader 
}

func (f *file) Close() error { return nil }
func (f *file) Readdir(count int) ([]os.FileInfo, error) { return nil, errors.New("not supported") }
func (f *file) Stat() (os.FileInfo, error) { return f, nil }
func (f *file) Name() string { return f.name }
func (f *file) Size() int64 { return int64(f.size) }
func (f *file) Mode() os.FileMode { return 0644 }
func (f *file) ModTime() time.Time{ return time.Time{} }
func (f *file) IsDir() bool { return false }
func (f *file) Sys() interface{} { return nil }

func init() {
	assets["/index.html"] = []byte{0x3c, 0x21, 0x44, 0x4f, 0x43, 0x54, 0x59, 0x50, 0x45, 0x20, 0x68, 0x74, 0x6d, 0x6c, 0x3e, 0x3c, 0x68, 0x74, 0x6d, 0x6c, 0x20, 0x6c, 0x61, 0x6e, 0x67, 0x3d, 0x65, 0x6e, 0x3e, 0x3c, 0x68, 0x65, 0x61, 0x64, 0x3e, 0x3c, 0x6d, 0x65, 0x74, 0x61, 0x20, 0x63, 0x68, 0x61, 0x72, 0x73, 0x65, 0x74, 0x3d, 0x75, 0x74, 0x66, 0x2d, 0x38, 0x3e, 0x3c, 0x6d, 0x65, 0x74, 0x61, 0x20, 0x68, 0x74, 0x74, 0x70, 0x2d, 0x65, 0x71, 0x75, 0x69, 0x76, 0x3d, 0x58, 0x2d, 0x55, 0x41, 0x2d, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x74, 0x69, 0x62, 0x6c, 0x65, 0x20, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x3d, 0x22, 0x49, 0x45, 0x3d, 0x65, 0x64, 0x67, 0x65, 0x22, 0x3e, 0x3c, 0x6d, 0x65, 0x74, 0x61, 0x20, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x76, 0x69, 0x65, 0x77, 0x70, 0x6f, 0x72, 0x74, 0x20, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x3d, 0x22, 0x77, 0x69, 0x64, 0x74, 0x68, 0x3d, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2d, 0x77, 0x69, 0x64, 0x74, 0x68, 0x2c, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x2d, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x3d, 0x31, 0x22, 0x3e, 0x3c, 0x6c, 0x69, 0x6e, 0x6b, 0x20, 0x72, 0x65, 0x6c, 0x3d, 0x69, 0x63, 0x6f, 0x6e, 0x20, 0x68, 0x72, 0x65, 0x66, 0x3d, 0x2f, 0x66, 0x61, 0x76, 0x69, 0x63, 0x6f, 0x6e, 0x2e, 0x69, 0x63, 0x6f, 0x3e, 0x3c, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x3e, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x3c, 0x2f, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x3e, 0x3c, 0x6c, 0x69, 0x6e, 0x6b, 0x20, 0x68, 0x72, 0x65, 0x66, 0x3d, 0x2f, 0x63, 0x73, 0x73, 0x2f, 0x61, 0x70, 0x70, 0x2e, 0x35, 0x65, 0x64, 0x66, 0x66, 0x31, 0x34, 0x39, 0x2e, 0x63, 0x73, 0x73, 0x20, 0x72, 0x65, 0x6c, 0x3d, 0x70, 0x72, 0x65, 0x6c, 0x6f, 0x61, 0x64, 0x20, 0x61, 0x73, 0x3d, 0x73, 0x74, 0x79, 0x6c, 0x65, 0x3e, 0x3c, 0x6c, 0x69, 0x6e, 0x6b, 0x20, 0x68, 0x72, 0x65, 0x66, 0x3d, 0x2f, 0x63, 0x73, 0x73, 0x2f, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x2d, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x73, 0x2e, 0x64, 0x65, 0x39, 0x33, 0x66, 0x39, 0x61, 0x37, 0x2e, 0x63, 0x73, 0x73, 0x20, 0x72, 0x65, 0x6c, 0x3d, 0x70, 0x72, 0x65, 0x6c, 0x6f, 0x61, 0x64, 0x20, 0x61, 0x73, 0x3d, 0x73, 0x74, 0x79, 0x6c, 0x65, 0x3e, 0x3c, 0x6c, 0x69, 0x6e, 0x6b, 0x20, 0x68, 0x72, 0x65, 0x66, 0x3d, 0x2f, 0x6a, 0x73, 0x2f, 0x61, 0x70, 0x70, 0x2e, 0x63, 0x64, 0x66, 0x66, 0x31, 0x31, 0x64, 0x30, 0x2e, 0x6a, 0x73, 0x20, 0x72, 0x65, 0x6c, 0x3d, 0x70, 0x72, 0x65, 0x6c, 0x6f, 0x61, 0x64, 0x20, 0x61, 0x73, 0x3d, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x3e, 0x3c, 0x6c, 0x69, 0x6e, 0x6b, 0x20, 0x68, 0x72, 0x65, 0x66, 0x3d, 0x2f, 0x6a, 0x73, 0x2f, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x2d, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x73, 0x2e, 0x30, 0x39, 0x38, 0x30, 0x31, 0x32, 0x65, 0x36, 0x2e, 0x6a, 0x73, 0x20, 0x72, 0x65, 0x6c, 0x3d, 0x70, 0x72, 0x65, 0x6c, 0x6f, 0x61, 0x64, 0x20, 0x61, 0x73, 0x3d, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x3e, 0x3c, 0x6c, 0x69, 0x6e, 0x6b, 0x20, 0x68, 0x72, 0x65, 0x66, 0x3d, 0x2f, 0x63, 0x73, 0x73, 0x2f, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x2d, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x73, 0x2e, 0x64, 0x65, 0x39, 0x33, 0x66, 0x39, 0x61, 0x37, 0x2e, 0x63, 0x73, 0x73, 0x20, 0x72, 0x65, 0x6c, 0x3d, 0x73, 0x74, 0x79, 0x6c, 0x65, 0x73, 0x68, 0x65, 0x65, 0x74, 0x3e, 0x3c, 0x6c, 0x69, 0x6e, 0x6b, 0x20, 0x68, 0x72, 0x65, 0x66, 0x3d, 0x2f, 0x63, 0x73, 0x73, 0x2f, 0x61, 0x70, 0x70, 0x2e, 0x35, 0x65, 0x64, 0x66, 0x66, 0x31, 0x34, 0x39, 0x2e, 0x63, 0x73, 0x73, 0x20, 0x72, 0x65, 0x6c, 0x3d, 0x73, 0x74, 0x79, 0x6c, 0x65, 0x73, 0x68, 0x65, 0x65, 0x74, 0x3e, 0x3c, 0x2f, 0x68, 0x65, 0x61, 0x64, 0x3e, 0x3c, 0x62, 0x6f, 0x64, 0x79, 0x3e, 0x3c, 0x6e, 0x6f, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x3e, 0x3c, 0x73, 0x74, 0x72, 0x6f, 0x6e, 0x67, 0x3e, 0x57, 0x65, 0x27, 0x72, 0x65, 0x20, 0x73, 0x6f, 0x72, 0x72, 0x79, 0x20, 0x62, 0x75, 0x74, 0x20, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x20, 0x64, 0x6f, 0x65, 0x73, 0x6e, 0x27, 0x74, 0x20, 0x77, 0x6f, 0x72, 0x6b, 0x20, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x6c, 0x79, 0x20, 0x77, 0x69, 0x74, 0x68, 0x6f, 0x75, 0x74, 0x20, 0x4a, 0x61, 0x76, 0x61, 0x53, 0x63, 0x72, 0x69, 0x70, 0x74, 0x20, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x2e, 0x20, 0x50, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x20, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x20, 0x69, 0x74, 0x20, 0x74, 0x6f, 0x20, 0x63, 0x6f, 0x6e, 0x74, 0x69, 0x6e, 0x75, 0x65, 0x2e, 0x3c, 0x2f, 0x73, 0x74, 0x72, 0x6f, 0x6e, 0x67, 0x3e, 0x3c, 0x2f, 0x6e, 0x6f, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x3e, 0x3c, 0x64, 0x69, 0x76, 0x20, 0x69, 0x64, 0x3d, 0x61, 0x70, 0x70, 0x3e, 0x3c, 0x2f, 0x64, 0x69, 0x76, 0x3e, 0x3c, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x20, 0x73, 0x72, 0x63, 0x3d, 0x2f, 0x6a, 0x73, 0x2f, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x2d, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x73, 0x2e, 0x30, 0x39, 0x38, 0x30, 0x31, 0x32, 0x65, 0x36, 0x2e, 0x6a, 0x73, 0x3e, 0x3c, 0x2f, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x3e, 0x3c, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x20, 0x73, 0x72, 0x63, 0x3d, 0x2f, 0x6a, 0x73, 0x2f, 0x61, 0x70, 0x70, 0x2e, 0x63, 0x64, 0x66, 0x66, 0x31, 0x31, 0x64, 0x30, 0x2e, 0x6a, 0x73, 0x3e, 0x3c, 0x2f, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x3e, 0x3c, 0x2f, 0x62, 0x6f, 0x64, 0x79, 0x3e, 0x3c, 0x2f, 0x68, 0x74, 0x6d, 0x6c, 0x3e}
}
