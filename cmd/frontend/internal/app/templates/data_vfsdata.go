// Code generated by vfsgen; DO NOT EDIT.

// +build dist

package templates

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	pathpkg "path"
	"time"
)

// Data statically implements the virtual filesystem provided to vfsgen.
var Data = func() http.FileSystem {
	fs := vfsgen۰FS{
		"/": &vfsgen۰DirInfo{
			name:    "/",
			modTime: time.Date(2018, 9, 25, 17, 2, 45, 887370446, time.UTC),
		},
		"/styles.css": &vfsgen۰CompressedFileInfo{
			name:             "styles.css",
			modTime:          time.Date(2018, 9, 25, 17, 2, 45, 887370446, time.UTC),
			uncompressedSize: 5262,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x58\x4b\x8f\xe3\xb8\x11\x3e\x5b\xbf\xa2\x60\x20\xe8\x07\x24\xb5\x24\xdb\xfd\x50\x5f\x16\x3b\x9b\x49\x02\xcc\x2e\x82\xe9\x9d\xdc\x29\x89\xb6\x89\xa6\x49\x85\xa4\xda\xf6\x1a\xfe\xef\x41\x51\x2f\xca\xb2\xa7\x7b\x0e\xc1\xcc\xc0\x06\x58\xfc\xea\x5d\xac\xaa\xf6\xe1\x50\xd0\x25\x13\x14\xa6\x2f\x66\xcf\xa9\x9e\x1e\x8f\xde\xdd\xad\x07\xb7\xf0\xe7\x9a\x69\xd0\x96\xb8\xa6\xd4\x40\x83\xab\x29\xb0\x94\x0a\xcc\x9a\x02\xb2\xfb\x8a\x96\xf2\x18\x70\x22\x0a\x26\x56\x50\x92\x15\xd5\xa1\x07\xb7\x77\x9e\xe7\x01\x0a\x9b\xfc\x87\x28\x46\x32\x4e\xb5\x87\xc4\x5b\x38\x78\x93\x20\xc8\x78\x45\x83\x78\x32\x99\x40\x0a\x6a\x95\x5d\x47\x3e\x3c\x26\x3e\xc4\xf3\xe8\xe6\xb9\xbb\x4f\x06\xf7\xf1\xfc\xde\x87\x24\x49\x1c\xc0\xac\x07\xc4\x0b\x1f\x62\x14\x91\xcc\x5d\xc4\xdc\x41\xc4\x4f\x3e\x24\xd1\xd3\x29\x64\xe1\x40\x9e\x50\xc0\x6c\xe1\x43\xb2\x40\x43\x10\x53\x56\xaa\xe4\xd6\xd6\x16\x94\xcc\x7c\x88\xf1\xf3\xd0\x88\x69\x20\x49\x0f\x79\x78\x40\x55\xa8\xe9\x61\x00\x99\xf5\x90\x79\xd2\x40\x16\x8b\x46\x93\x54\x44\xac\x5c\x4d\xc9\x2c\xf2\xe1\xfe\xc9\x87\x26\x28\x0d\xa0\xd7\x93\xcc\x1f\x7d\x78\xba\x47\xcf\x07\x88\x5e\x4d\xb2\xc0\xb8\xdc\xe3\x57\xd2\x7a\xb4\x52\x94\x0a\xab\x06\x1a\x94\xbd\x7e\xb4\x5f\xb5\x9c\x1a\x92\x38\x10\x34\x35\xbe\x47\x55\xf3\x01\x66\xe6\x60\x30\x28\x49\x64\xcd\x99\x35\xba\x14\x2d\xea\x34\x77\x72\x30\x8b\x71\x54\x7f\xd5\x82\x10\x93\x0c\x31\x89\x0f\x0f\x33\xfc\x34\x62\x72\x29\x79\xb0\x52\x64\x1f\xc4\x9d\x67\xf7\x3e\xa0\xfb\x0f\x4d\x84\x7b\x48\xd2\x42\x16\x8d\x98\x38\x4a\x4e\x31\xb3\x41\x55\xc4\xf3\x07\x4c\xe7\x7c\x84\xea\xd4\x91\x21\xd0\x87\x28\x8c\xc7\xe8\xe4\x32\x7a\x64\xc1\xbc\x73\xc4\x86\x2d\xe9\xea\xb2\x8e\x2d\xd9\x4f\x26\x4e\x16\x6d\xa1\x44\x6e\x5d\x1a\xba\x33\xd0\xfe\x3b\x1f\x12\xce\xc4\xeb\x09\xa4\x7b\x47\xf1\xbc\xc7\x04\x6b\xf9\x46\x95\x8b\xe9\xdf\x22\x62\x34\x11\x3a\xd0\x54\xb1\x65\x0a\xd3\x9c\xa9\xbc\xe2\x44\x05\xda\x14\x53\x1f\x02\x52\x62\x61\xeb\xbd\x36\x74\xe3\xc3\xaf\x28\xef\x77\x92\xbf\xd8\xf3\x67\x29\x8c\x0f\xd3\xf0\xe5\xf3\x1f\x2f\x7f\xd2\x9d\x09\xbe\xd2\x15\xf2\x4e\x7d\x98\xbe\x10\x01\x9f\x15\x11\x39\xd3\xb9\xb4\x84\xcf\x7f\xbc\xc0\x6f\x4c\x97\x9c\xec\xa7\x3e\x7c\x95\x99\x34\xd2\x87\xe9\x97\x2a\x67\x05\x81\x7f\x28\x22\x0a\x8a\x40\xf2\x46\x05\x53\x20\xe8\xce\x4c\x7d\xa8\x4f\x3e\xfc\x93\xf2\x37\x6a\x58\x4e\x7c\xf8\x96\x55\xc2\x54\x3e\x4c\x35\x5d\x49\x0a\x15\x43\x98\x62\x84\xfb\xd0\xbb\x62\xbd\xdf\x48\x21\x75\x49\x72\x9a\x76\x7a\x3e\x49\xa1\x25\x47\x45\xbf\x4b\x41\x72\xe9\x43\x07\xaa\xa3\x91\x11\x4d\x83\xa5\x14\x26\xd0\xec\x2f\x9a\x42\xbc\x28\x77\xcf\xdd\x05\x67\x82\x06\x6b\xca\x56\x6b\x93\x42\x1c\x2e\xda\xfa\x2d\x06\x3c\x51\xf8\x44\x37\xcf\xdd\xcd\x09\xd3\x63\xf3\x76\x48\xc1\x2a\x9d\xc2\x0c\xe5\x7b\x47\xcf\xfb\xc5\x4a\x58\x92\x9c\xc2\xc1\x03\x68\x4e\x1b\xc6\xf7\x29\x5c\xb9\x79\xb9\x7a\xf6\x00\xb4\xca\x53\xa8\x14\xbf\xbe\x3a\x1c\x88\xd6\xd4\x7c\xfb\xfa\x05\xa6\x77\xc8\xa5\xef\x5c\x74\xa0\xab\x4c\x53\x63\x68\x71\x87\x86\x18\x19\x74\xb7\x99\x94\xaf\xe1\x56\x2e\x97\xd3\xe3\xf1\xea\x06\x1b\xff\x86\x98\xeb\x2b\xa4\x5c\xdd\x3c\x7b\xc7\x9f\x64\xd2\xbf\x0c\xe1\x2c\xff\x8e\x61\x93\x3a\xd6\x38\xae\x52\x90\x19\x67\xff\xad\xe8\xcf\x32\x97\x17\xef\x1a\xba\x6d\x52\x8f\xe0\x9f\x67\xe6\x07\xa3\x3a\x34\xf6\x52\xa4\x9b\x45\xe2\x93\xe4\x52\x35\xeb\x40\x58\x0f\x7d\x38\x40\x8e\xd4\x14\xde\x88\xba\x6e\x57\x81\x9b\x67\x38\x36\x88\xe4\x2c\x22\x71\x10\xb3\xb3\x88\x99\x45\x78\x61\x3b\xb1\x4f\x41\x2d\xbd\x16\xd4\x0e\xed\x0b\xa8\x64\x80\x1a\x29\x6c\xe9\x8d\xca\xed\x9a\x19\xda\x63\xec\xb1\xbe\x71\x26\xd7\xa9\x0c\xe7\xaa\x56\xe6\x8c\xb0\xcb\xd8\xe4\x14\x3b\xb2\xcd\xb9\x3a\xc5\xce\x2f\x63\xe7\x8d\x2b\xd9\x2a\xe8\xf2\x94\x91\xfc\x75\xa5\x64\x25\x8a\xe0\x72\xca\x1a\x7c\xf2\x1e\x3e\x19\xe2\x67\xef\xe1\x67\xbd\x41\x6d\x78\xc7\xf8\x2e\xd2\x08\x43\x3f\xbe\x23\x15\xaf\x6b\x99\xcd\x92\xbb\x2f\xe5\x4a\x91\x72\xbd\xaf\x0b\x34\x93\xc5\x1e\xb7\xd3\x01\x13\x8e\xd9\xae\xf4\xdb\x37\x58\x5f\xf5\xb3\xa4\xef\x38\xb6\xbb\x37\x4e\x0c\xe6\x04\x42\x06\x6d\xde\x01\x39\x74\x3b\x94\xb7\x34\x7b\x65\xa6\xe1\xdd\x48\x69\xd6\x4c\xac\x52\x20\xc2\x30\xc2\x19\xd1\xd4\x36\x09\xaf\x54\xd4\xf7\x70\x7e\xf8\x98\xe2\x02\x7b\xc5\x39\x33\xbb\xf1\x75\xd6\xca\xe1\x64\xba\x60\xe5\xe9\x90\xfa\xae\x95\x95\x91\xd6\x3c\x32\x8a\x25\xee\x06\xc8\x8a\x31\x0d\x0a\x9a\x4b\x45\x0c\x93\x22\x05\x21\x05\x75\x44\x1a\x45\x84\x66\xf5\x15\xe1\x1c\xa2\x70\xa6\xf1\x7a\x23\xff\xba\x74\x77\x9e\x8c\x56\xa4\xf5\x72\x73\xce\x96\x7a\xef\xb1\x73\xcc\x0b\x2d\x41\x57\x99\xe1\xd4\x01\x33\xb1\xa6\x8a\x99\x73\x46\x57\xa2\xa0\x0a\x83\x32\xe2\x6f\x55\xc2\x79\x47\xb1\x02\x4b\x38\xc0\x86\xa8\x15\x13\x41\x26\x8d\x91\x1b\xbb\x2c\x28\xba\xb1\x25\xcf\xd7\x11\x1c\x60\x90\x88\xc8\x29\xdc\x6f\x86\x71\x66\x18\xd5\x75\xdd\x86\x99\x54\x05\x55\xc3\x4d\x14\xdf\x41\x4b\x3e\xdf\x1b\xba\x07\x79\xca\x3d\x7f\x97\xb7\xed\x15\x7a\x4d\x0a\xb9\xc5\x70\x65\x72\x17\xd4\xa7\x14\x22\x88\xcb\x1d\xdc\x97\x3b\x88\xce\x68\xb5\x6b\xf3\xdd\x6d\x9b\x6b\x87\x11\x33\xac\x5d\x02\x7a\x67\xd5\x94\x4c\xb4\x31\x9d\x10\xc1\x36\x4d\x34\x91\x0e\x51\xb8\xd0\x40\xf1\x19\x31\x11\xc8\xca\x40\x6c\x6d\xfb\xe5\x95\xee\x97\x8a\x6c\xa8\xae\x61\x07\x6f\xb2\x88\xfe\x86\x39\xc1\x4a\xc1\xc9\x96\x82\x92\x86\x18\x7a\x1d\x3f\x46\x05\x5d\xdd\x80\xce\x09\xa7\xd7\x71\x98\xdc\xc0\xd1\x9b\xc4\xd1\xc7\xe0\x08\xee\x32\xf3\x77\x4e\x37\x54\x98\x26\x31\x77\xb7\xf0\x85\x69\xa3\xed\x81\xb3\x2e\xe5\x29\xc4\x98\xe9\x36\xa5\xf0\x6b\x65\x8c\x14\xba\x49\xa6\x11\x75\x40\x6d\x02\xda\x1d\xb0\x8e\x63\x7d\xc2\x00\x0e\xda\x9f\x37\x29\xea\xb5\x19\xeb\xd5\x56\x4d\xc6\x65\xfe\x7a\x61\x60\x97\xa4\x28\xec\x4b\x7d\x2c\x77\x10\xdf\xd7\xcb\xeb\xff\xe3\xe5\xd5\xae\x4c\xc6\x33\x31\x33\xc2\xb6\xf8\x8f\x4d\x80\x0e\xde\x3d\xab\x0f\x8c\x19\x0c\xeb\xbf\x15\x85\x3b\xf8\xad\xda\x64\xf0\x09\xdb\xa3\x16\xac\x2c\xa9\xb1\x61\x2e\x95\x7d\xe5\xbd\xa4\x3e\x96\xef\x44\xbe\xbe\x4e\x6d\x8d\x6b\xc9\x59\x71\xee\x75\x9c\x34\x9b\x76\x5e\x7a\x93\xb3\x6f\xbe\xa7\x1b\x59\x3a\xc4\x2e\x53\xb1\x3d\xd7\x5e\x8d\x5c\x09\xdb\xc3\xc1\x9b\x60\x80\x96\x5c\x6e\x83\x5d\xdb\x88\x01\xdc\x07\xc5\x84\xa6\x06\x82\xd8\xbe\xcd\x79\xb9\xab\xff\x58\x8d\x7c\xfb\x3f\x8c\x17\x6e\x63\xff\x61\x3e\x2c\x92\x1f\x64\x3a\xf6\xe6\xfb\xbd\x23\x61\xc9\xc5\xe0\x58\x0d\x8e\xb2\x1c\x1c\x73\x2e\x47\xdd\xbd\x0b\xb8\xa3\x00\x08\x1c\x2e\x37\x64\x47\x9e\xdc\x8c\xe4\x0d\x56\xaa\xe1\xce\xcb\xec\xd6\x3c\xd4\x14\x72\x36\xf0\x47\x9b\xf1\x00\x6a\x7f\xcc\x39\x35\xd2\xe5\x7b\xdd\x16\xee\xd1\x90\x95\x7b\x5c\x56\x62\x24\xb5\xdf\x5e\x87\x06\x11\x33\x88\x99\xd9\x97\xe7\x63\x36\x62\x2c\x68\xee\x32\xbe\x11\xe5\x0f\xc4\xbe\x8d\xe4\x34\xbf\x0b\x35\x82\xb6\x8a\x58\x55\x5b\xa9\x8a\x00\x0f\x29\x64\x8a\x92\xd7\x00\x09\x16\x72\x38\x50\x51\x1c\x8f\xde\xff\x02\x00\x00\xff\xff\x31\x76\xd8\x1c\x8e\x14\x00\x00"),
		},
		"/ui": &vfsgen۰DirInfo{
			name:    "ui",
			modTime: time.Date(2018, 9, 25, 17, 2, 45, 887370446, time.UTC),
		},
		"/ui/app.html": &vfsgen۰CompressedFileInfo{
			name:             "app.html",
			modTime:          time.Date(2018, 9, 25, 17, 2, 45, 887370446, time.UTC),
			uncompressedSize: 2589,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x56\x4d\x93\xda\x46\x13\x3e\xa3\x5f\xd1\xd6\x05\xf0\x8b\xa4\xf5\x7b\x4a\x2d\x88\x24\x6b\xbb\xf2\x51\x76\x0e\xd9\xcd\x21\xe5\xf2\xa1\x99\x69\xa4\x59\xa4\x19\x79\xa6\x81\xa5\xb0\xfe\x7b\x6a\x66\x96\x05\x36\xd8\xb1\x2f\x20\xf5\x3c\xdd\xfd\xf4\xe7\x68\xf6\x42\x1a\xc1\xbb\x8e\xa0\xe6\xb6\x99\x27\x33\xff\x07\x0d\xea\xaa\x4c\x49\xa7\x20\x1a\x74\xae\x4c\x17\xe8\x28\x9d\x27\xc9\xac\x26\x94\xf3\x64\xb0\xdf\xe7\xbf\xe9\x7b\x12\x4c\x32\xff\x95\x50\xde\x99\xae\xef\x93\xc1\xac\x25\x46\x10\x35\x5a\x47\x5c\xa6\x6b\x5e\x66\x3f\xa4\xf3\x83\xbc\x66\xee\x32\xfa\xb4\x56\x9b\x32\x7d\xc8\xd6\x98\x09\xd3\x76\xc8\x6a\xd1\x50\x0a\xc2\x68\x26\xcd\x65\xaa\xa8\x24\x59\xd1\x51\x4d\x63\x4b\x65\xba\x51\xb4\xed\x8c\xe5\x13\xe4\x56\x49\xae\x4b\x49\x1b\x25\x28\x0b\x2f\x13\x50\x5a\xb1\xc2\x26\x73\x02\x1b\x2a\x5f\x3d\xb7\x52\x19\x53\x9d\x79\xd3\x86\x2d\x6a\xd7\x20\xd3\x65\xa2\xaf\x23\x32\x7b\x87\xba\x5a\x63\x75\xaa\x4b\x3a\x0d\xa9\x00\xb5\x84\x3c\xe0\x1e\x38\xbf\x35\x6b\x2b\xa8\xb2\xd8\xd5\x6f\x0c\xbf\x36\xed\x7b\x23\x09\x8e\xc9\x09\x3c\x86\x87\x68\x86\x4f\xd6\x86\xcf\x98\x4f\xe0\x80\xc9\x96\x8a\x4b\x61\x36\x64\x87\xd1\x1d\x69\x19\x0c\xee\xf7\x5b\xc5\x35\xe4\xef\x89\x51\x22\xe3\x73\x27\x92\x9c\xb0\xaa\x63\x65\xf4\xd1\x4f\xba\xdf\xe7\x6f\x8e\x07\x7d\x9f\x42\x71\x9e\xa4\x21\x6f\x15\x33\xd9\x6b\x56\xdc\xd0\xb9\xe6\x9d\x17\x7d\x4d\x47\xa0\x95\x27\x41\xb9\x75\xdb\xa2\xdd\x0d\xbf\x88\xff\x7e\x8e\x9d\x35\x1d\x59\xde\x95\xa9\xa9\xae\x7d\xe7\x9e\x76\x04\x2d\x9c\x62\xba\x08\x1f\x7a\xf8\x37\x45\x74\xa6\xf2\x3d\x04\xf7\x7b\xd2\x32\x54\x21\x38\x9a\x1f\xcd\xcf\x8a\x28\x49\x06\xb3\x46\xe9\x15\x58\x6a\xca\xd4\xf1\xae\x21\x57\x13\x71\x0a\xb5\xa5\x65\x30\xfb\xb3\x73\xc4\x7f\xfd\xf9\xae\xef\x8b\x78\x1e\xff\xf2\xc5\x5a\xcb\x86\x72\xe1\xdc\x8f\xfb\xfd\x86\xac\x53\x46\x43\xfa\x05\x48\xda\xf7\xe9\x93\x2f\x25\xcb\xa1\x3b\x76\x65\x26\x6a\x6b\x5a\xca\x7c\xae\xd8\x58\xca\x14\x53\x3b\x8c\x8c\x2e\x1d\x1d\xb8\xf9\xa1\x70\xd7\x45\x11\x31\x79\x9c\xa4\x5c\x98\xb6\x38\xc0\x0b\x49\x8c\xaa\x29\x64\x75\x5f\x2f\x4d\x7b\xaf\x08\x51\x76\xa6\xb9\x6f\xb4\x92\xed\xa2\x5a\xc9\xe5\xb2\x43\xb1\x4a\x9f\xa5\x81\xd0\x8a\xfa\xe0\xa6\x30\x1d\xe9\x28\xca\x1f\xda\x26\x05\x5f\xe2\x32\xc5\xae\x6b\x94\x40\x9f\xed\x13\xc4\x49\x71\xfe\x17\xc1\x3e\xcb\x65\x7a\x32\x84\x70\x1b\xcd\x87\xf2\xf8\x41\x8d\x7b\xeb\x04\x71\x67\x51\xac\xc8\x86\xba\x45\x73\xa0\x2a\xed\xc3\x17\xae\x9b\x27\x83\xc1\x14\x46\xcb\xb5\x16\xde\x0b\x8c\x78\x02\x76\x02\x38\x01\x31\x81\xd5\x04\x68\x02\x72\x0c\x61\x05\x8c\x5e\xf0\x87\xd5\x47\xff\xc2\xf9\x2f\x8d\x59\x60\x73\x47\x4d\xa3\x2a\xd2\xfc\x07\xb6\xe4\x3a\x14\x04\xe5\xd7\x0e\x3f\x7f\x86\x0f\x1f\xa7\x5f\x41\xe4\xdd\xda\xd5\xa3\xd5\x78\x0a\xde\x15\x94\x70\xe4\xe5\xfd\x8e\xbc\x34\xff\xe4\x9d\xc4\x87\x60\x70\x1c\xb5\xd0\x56\xeb\x96\x34\xbb\x31\xf4\x53\xb8\x88\x9c\x82\x27\x68\x73\x61\x09\x99\xde\x36\xe4\xf1\x23\x1c\x4f\x41\x06\x79\x45\xfc\x28\x74\x37\xbb\x3b\xac\x3c\xaf\x11\x8e\x3f\x5c\x79\xcd\x1c\xdd\x4e\x0b\x28\xe1\x95\x7f\x71\xd6\x3f\x8a\x29\xc8\xbc\x43\xeb\x63\x30\x92\x72\xa5\x1d\x59\xbe\xa1\xa5\xb1\x34\x8a\xa9\xeb\xa1\x1f\x6d\x95\x96\x66\x3b\x01\x69\x44\xa0\x38\x81\x34\x16\x22\x9d\xc0\x53\xdf\xf9\x16\xc3\xea\xd0\x78\xd8\x29\x17\x9a\x8f\x0f\x49\xca\xd0\xb2\x5a\xa2\x60\x57\x70\xac\xa8\x3f\xcb\xef\x9d\x37\xf2\x84\x4a\xc7\xe3\x69\x32\x98\x15\xd1\xfe\xd9\xcc\x5e\xac\x7d\xa4\x96\x8b\xb8\xdd\xa1\x84\xfd\xfe\xb0\xea\xc3\x0e\x3e\x00\x3a\xac\xe8\xad\xb5\xc6\x46\x48\x7c\x0c\x66\x4f\x5d\x9d\x5f\x9a\x37\x86\xd9\xb4\x7d\x9f\xcc\x8a\x78\xad\x26\xb3\x85\x91\xbb\x67\xc8\x1b\x23\x77\x87\xeb\x55\xaa\x8d\x1f\xe7\xd4\x1a\xc3\xe9\x7c\x56\x48\xb5\xf1\xb3\xa4\xcd\xa3\x8f\xbf\xcd\x1a\x34\x91\x04\x36\x40\x1a\x17\x0d\xc1\xef\xb8\xc1\xdb\x18\x19\x1b\xb0\x6b\x0d\x5c\x2b\x07\xd8\x75\xf9\xac\x78\x52\x3c\x86\xef\xac\xf8\xd7\x1e\x0a\x27\xae\xf0\x3a\x8f\x3b\xe6\xfe\x7c\x0b\x5d\x02\x84\x1d\xf4\x85\xe8\x7d\x4c\x27\xd1\xc7\xa8\x93\x6f\xbd\x49\x67\x61\xdd\x9d\x17\xaa\x78\x09\xb7\xeb\xce\x5f\x97\xb0\xa5\xc5\x4a\x31\xf8\x5e\x63\x88\x5f\x07\x0e\x5e\x16\xc9\xe0\x27\x17\x11\x0e\x46\x1d\x4a\xa9\x74\x75\x0d\x2d\x3e\x8c\xae\xba\x87\xf1\x18\xf6\x49\x32\x18\xe4\x55\x98\xbc\x4c\xe3\x66\x81\x76\xe2\x25\x2d\x2a\x9d\xf9\xf2\x42\x6e\xcd\x36\x88\xe2\xf2\xc9\xac\xd9\xc2\x3e\x19\x0c\x06\x8f\xd6\xb2\x86\x96\x1c\x4d\xbe\xfa\x7f\xf7\x30\x01\xd2\x9b\x91\xc3\x25\x65\x68\x09\xb3\x40\x28\x60\x42\x0f\x1e\xd5\xac\xaa\xea\xff\xd4\x0b\xa0\xa8\xd8\x07\xa6\xad\x91\xd8\x64\x7e\x17\x06\x62\xf1\xe7\x02\x9d\xab\x4b\xbe\xae\x0e\x76\x7c\xfa\x43\x3a\xe7\xc9\xf1\xcb\xc2\x37\xa4\xff\x1a\xfc\x27\x00\x00\xff\xff\x79\x5d\xc7\xe6\x1d\x0a\x00\x00"),
		},
		"/ui/error.html": &vfsgen۰CompressedFileInfo{
			name:             "error.html",
			modTime:          time.Date(2018, 9, 25, 17, 2, 45, 887370446, time.UTC),
			uncompressedSize: 1214,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x54\xc1\x6e\xe3\x46\x0c\x3d\x4b\x5f\xc1\xf5\xa5\x6d\xe0\x48\xe8\xad\x08\x26\x46\x17\xbb\x3d\xf8\xd0\xed\xa2\xf1\xa5\x47\x4a\x43\x59\x83\x8c\x66\x54\x92\xb2\x23\x18\xfa\xf7\x62\x46\x59\x3b\x1b\x74\x4f\xc6\x50\x7c\x7c\x8f\xef\x11\xbe\x5c\xea\xbb\xb2\xd8\xff\xf9\xf5\xaf\xbf\x0f\x1f\xbf\x1c\x3e\xc0\xa1\x27\x26\x40\x26\xd0\x73\x04\x62\x8e\x0c\x4a\xc3\xe8\x51\x49\x1e\xca\xb2\xb8\x5f\x8b\x55\xaf\x83\x4f\xaf\x0e\x43\x3b\xbf\x29\x95\xc5\xed\x01\x4e\xc0\x92\xb8\x63\x20\x0b\x1a\xa1\x21\x40\x81\xc1\x05\x37\xa0\x07\x0c\x16\x98\xbc\xc3\xc6\xe7\xfa\x18\x45\x5c\xe3\xa9\x82\xbd\x82\x8d\x24\x65\x11\xa2\x82\x0b\xad\x9f\x2c\xc1\xc7\x2f\xff\x40\xd4\x9e\xde\xe8\xd9\x26\x98\xf6\x34\xc3\x80\x33\x30\xfd\x3b\x39\x26\xb0\xa8\x08\xda\xa3\x26\xfa\x10\xb5\x2c\x46\x26\xa1\xa0\xf0\x33\x55\xc7\x0a\xee\x3e\xc5\x61\x88\x21\xf7\x6d\xe1\xdc\xbb\xb6\xcf\xf8\x44\x46\x2f\x4e\x14\x5c\x07\xb9\xf3\xe8\x54\x88\x4f\xc4\x10\x79\xe5\x2e\x8b\xa4\x98\x2c\x4c\x63\x0c\xb0\x7e\x94\x6c\x97\xa5\x23\xa3\x25\xfb\x4b\x05\x87\xde\xc9\x55\x24\x48\x1f\x27\x9f\x36\x0d\x96\x18\xe8\x44\x01\xce\x3d\x85\xb2\xd0\x9e\x40\x66\x51\x1a\x92\x50\xa1\x13\x31\xf9\x19\x1a\x8e\xcf\x14\xaa\xb2\x2c\xde\x59\xfb\x3f\x6e\x0e\x91\x09\x26\x21\x86\x8e\x1d\x05\xeb\xe7\x6d\xb6\xf5\x4c\x80\xfe\x8c\xb3\x80\xf2\x0c\x1a\x93\xec\x4c\xef\x14\x3a\xc7\xa2\xd9\xe3\x49\x48\x56\xb7\x3a\x8e\xc3\xd5\x17\xd2\x76\x9d\xd2\x62\xb8\xba\xff\xba\xfd\xd5\x7a\xf0\xee\x99\xa0\x41\xa1\xac\xad\x82\x7d\x48\x49\xe4\xfd\xf4\xd5\x7d\x85\x0e\x9d\x97\xa4\x75\xa5\xdf\x7e\x77\x3b\x2e\x67\x07\x1d\x7a\xdf\x60\xfb\x7c\x8d\x6c\x12\xb2\x55\x79\x57\x2f\x4b\x69\x3e\xd8\xd8\xea\x3c\x12\x24\xc8\xae\x34\xd9\x07\x8f\xe1\xf8\xb8\xa1\xb0\x49\x05\x42\xbb\x2b\x0b\x33\x90\x22\xb4\x3d\xb2\x90\x3e\x6e\x26\xed\xee\x7f\xdb\xa4\xba\x3a\xf5\xb4\xbb\x5c\xaa\x27\x45\x9d\xe4\x53\xb4\xb4\x2c\x70\x7d\x1f\xe8\x45\x97\x05\xee\xe1\x29\x4e\xdc\xa6\x08\xc7\xde\xd4\x2b\xa8\x34\xf5\x3a\xdd\x34\xd1\xce\x69\x98\x75\x27\x68\x3d\x8a\x3c\x6e\xf2\x22\x89\xa1\x30\xfd\xaf\xe9\xa7\x30\x32\x62\xf8\xf6\x59\xf2\xf4\xfb\x36\x5a\xda\xbc\x67\x37\x75\xea\xfc\x21\x46\xe9\x45\xdf\x60\x56\x85\x37\x8c\xa9\x57\xbe\xcb\xc5\x75\x50\xfd\x91\x64\x2c\x4b\x9e\x35\x32\x7d\xa7\xee\x36\xe9\xb5\xcb\xd4\x23\xd3\x8a\xa5\x60\x33\xca\xf4\x0c\x75\x1e\x3b\xee\x9e\x22\xf3\xbc\x4d\x99\x30\xfd\x24\xd0\x10\x05\x40\x18\x39\x36\x9e\x86\x0a\xbe\x7a\x42\x21\x30\x08\x3d\x53\xf7\xb8\x19\xd0\x79\x8d\x0f\x32\x8d\x63\x64\xfd\x5d\x6e\x06\x56\x6d\x1c\x36\xbb\x36\x06\xc5\x36\x5d\x99\xa9\x71\x97\x0f\xea\xdb\x31\xe5\x43\xc9\x7f\x2d\xfb\xcf\x0f\x60\x44\x39\x86\xe3\x55\xe8\xfe\x73\xde\x77\x2d\x9a\x7a\x4c\xce\xd7\xd6\x9d\x52\x20\x6b\x12\xa6\x5e\xcf\xe1\xbf\x00\x00\x00\xff\xff\x84\x89\xd8\xf0\xbe\x04\x00\x00"),
		},
	}
	fs["/"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/styles.css"].(os.FileInfo),
		fs["/ui"].(os.FileInfo),
	}
	fs["/ui"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/ui/app.html"].(os.FileInfo),
		fs["/ui/error.html"].(os.FileInfo),
	}

	return fs
}()

type vfsgen۰FS map[string]interface{}

func (fs vfsgen۰FS) Open(path string) (http.File, error) {
	path = pathpkg.Clean("/" + path)
	f, ok := fs[path]
	if !ok {
		return nil, &os.PathError{Op: "open", Path: path, Err: os.ErrNotExist}
	}

	switch f := f.(type) {
	case *vfsgen۰CompressedFileInfo:
		gr, err := gzip.NewReader(bytes.NewReader(f.compressedContent))
		if err != nil {
			// This should never happen because we generate the gzip bytes such that they are always valid.
			panic("unexpected error reading own gzip compressed bytes: " + err.Error())
		}
		return &vfsgen۰CompressedFile{
			vfsgen۰CompressedFileInfo: f,
			gr:                        gr,
		}, nil
	case *vfsgen۰DirInfo:
		return &vfsgen۰Dir{
			vfsgen۰DirInfo: f,
		}, nil
	default:
		// This should never happen because we generate only the above types.
		panic(fmt.Sprintf("unexpected type %T", f))
	}
}

// vfsgen۰CompressedFileInfo is a static definition of a gzip compressed file.
type vfsgen۰CompressedFileInfo struct {
	name              string
	modTime           time.Time
	compressedContent []byte
	uncompressedSize  int64
}

func (f *vfsgen۰CompressedFileInfo) Readdir(count int) ([]os.FileInfo, error) {
	return nil, fmt.Errorf("cannot Readdir from file %s", f.name)
}
func (f *vfsgen۰CompressedFileInfo) Stat() (os.FileInfo, error) { return f, nil }

func (f *vfsgen۰CompressedFileInfo) GzipBytes() []byte {
	return f.compressedContent
}

func (f *vfsgen۰CompressedFileInfo) Name() string       { return f.name }
func (f *vfsgen۰CompressedFileInfo) Size() int64        { return f.uncompressedSize }
func (f *vfsgen۰CompressedFileInfo) Mode() os.FileMode  { return 0444 }
func (f *vfsgen۰CompressedFileInfo) ModTime() time.Time { return f.modTime }
func (f *vfsgen۰CompressedFileInfo) IsDir() bool        { return false }
func (f *vfsgen۰CompressedFileInfo) Sys() interface{}   { return nil }

// vfsgen۰CompressedFile is an opened compressedFile instance.
type vfsgen۰CompressedFile struct {
	*vfsgen۰CompressedFileInfo
	gr      *gzip.Reader
	grPos   int64 // Actual gr uncompressed position.
	seekPos int64 // Seek uncompressed position.
}

func (f *vfsgen۰CompressedFile) Read(p []byte) (n int, err error) {
	if f.grPos > f.seekPos {
		// Rewind to beginning.
		err = f.gr.Reset(bytes.NewReader(f.compressedContent))
		if err != nil {
			return 0, err
		}
		f.grPos = 0
	}
	if f.grPos < f.seekPos {
		// Fast-forward.
		_, err = io.CopyN(ioutil.Discard, f.gr, f.seekPos-f.grPos)
		if err != nil {
			return 0, err
		}
		f.grPos = f.seekPos
	}
	n, err = f.gr.Read(p)
	f.grPos += int64(n)
	f.seekPos = f.grPos
	return n, err
}
func (f *vfsgen۰CompressedFile) Seek(offset int64, whence int) (int64, error) {
	switch whence {
	case io.SeekStart:
		f.seekPos = 0 + offset
	case io.SeekCurrent:
		f.seekPos += offset
	case io.SeekEnd:
		f.seekPos = f.uncompressedSize + offset
	default:
		panic(fmt.Errorf("invalid whence value: %v", whence))
	}
	return f.seekPos, nil
}
func (f *vfsgen۰CompressedFile) Close() error {
	return f.gr.Close()
}

// vfsgen۰DirInfo is a static definition of a directory.
type vfsgen۰DirInfo struct {
	name    string
	modTime time.Time
	entries []os.FileInfo
}

func (d *vfsgen۰DirInfo) Read([]byte) (int, error) {
	return 0, fmt.Errorf("cannot Read from directory %s", d.name)
}
func (d *vfsgen۰DirInfo) Close() error               { return nil }
func (d *vfsgen۰DirInfo) Stat() (os.FileInfo, error) { return d, nil }

func (d *vfsgen۰DirInfo) Name() string       { return d.name }
func (d *vfsgen۰DirInfo) Size() int64        { return 0 }
func (d *vfsgen۰DirInfo) Mode() os.FileMode  { return 0755 | os.ModeDir }
func (d *vfsgen۰DirInfo) ModTime() time.Time { return d.modTime }
func (d *vfsgen۰DirInfo) IsDir() bool        { return true }
func (d *vfsgen۰DirInfo) Sys() interface{}   { return nil }

// vfsgen۰Dir is an opened dir instance.
type vfsgen۰Dir struct {
	*vfsgen۰DirInfo
	pos int // Position within entries for Seek and Readdir.
}

func (d *vfsgen۰Dir) Seek(offset int64, whence int) (int64, error) {
	if offset == 0 && whence == io.SeekStart {
		d.pos = 0
		return 0, nil
	}
	return 0, fmt.Errorf("unsupported Seek in directory %s", d.name)
}

func (d *vfsgen۰Dir) Readdir(count int) ([]os.FileInfo, error) {
	if d.pos >= len(d.entries) && count > 0 {
		return nil, io.EOF
	}
	if count <= 0 || count > len(d.entries)-d.pos {
		count = len(d.entries) - d.pos
	}
	e := d.entries[d.pos : d.pos+count]
	d.pos += count
	return e, nil
}
