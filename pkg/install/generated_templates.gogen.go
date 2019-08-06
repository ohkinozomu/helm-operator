// Code generated by vfsgen; DO NOT EDIT.

package install

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

// templates statically implements the virtual filesystem provided to vfsgen.
var templates = func() http.FileSystem {
	fs := vfsgen۰FS{
		"/": &vfsgen۰DirInfo{
			name:    "/",
			modTime: time.Date(2019, 8, 6, 14, 34, 3, 563392085, time.UTC),
		},
		"/flux-helm-operator-account.yaml.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "flux-helm-operator-account.yaml.tmpl",
			modTime:          time.Date(2019, 8, 6, 14, 33, 53, 494182202, time.UTC),
			uncompressedSize: 948,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x92\x39\x6f\xdc\x30\x10\x85\x7b\xfe\x8a\x01\x5c\x38\x09\x2c\x05\xee\x02\x75\xb6\x8b\x14\x09\x52\x28\x47\x13\xa4\x18\x92\x4f\x59\xc6\x5c\x8e\x30\x24\x37\x87\xb0\xff\x3d\x90\xb4\x06\xbc\x8e\xed\x34\xdb\x8d\xe6\xd2\x9b\xc7\xaf\x69\x1a\x73\x46\x9f\x36\xa0\x0c\xdd\x05\x07\x62\xe7\xa4\xa6\x72\x41\x2e\xd6\x5c\xa0\xa4\x12\x91\x2f\x88\x93\x3f\x4a\x91\x0d\xc9\x87\xf4\x9d\x58\x61\xce\x48\x52\xfc\x4d\x09\xf0\xf0\x34\x88\xd2\xbb\x6a\xa1\x09\x05\x99\x7e\x86\xb2\x59\x46\x1a\xcb\x19\x7e\xfe\x03\x72\x26\x27\xa9\xa8\x44\x7a\xd1\x5f\x5f\xdd\xbc\x6c\x0d\x8f\xe1\x0b\x34\x07\x49\x1d\xed\x2e\xcd\x6d\x48\xbe\xa3\x8f\xab\xaa\xab\x55\x94\xd9\xa2\xb0\xe7\xc2\x9d\x21\x8a\x6c\x11\xf3\x1c\x11\x25\xde\xa2\xa3\x21\xd6\x5f\xcd\x06\x71\xdb\xc8\x08\xe5\x22\x6a\x9e\x2e\x4d\x13\x85\x81\xda\x0f\xbc\x45\x1e\xd9\x81\xf6\xfb\x43\xf7\xf2\xd9\xd1\x34\x1d\x57\xa7\x89\x90\xfc\x7e\x6f\x66\xcf\xee\x8b\x55\xcb\xae\xe5\x5a\x36\xa2\xe1\x0f\x97\x20\xa9\xbd\x7d\x93\xdb\x20\xaf\x77\x97\x16\x85\xef\x6e\xb9\x59\xdd\xeb\x25\xe2\x94\x87\x18\xad\x11\xcb\x78\x43\x3c\x86\xb7\x2a\x75\xcc\x1d\x7d\x3d\x7f\x75\xfe\x6d\xd9\xa9\xc8\x52\xd5\xe1\x28\xb9\x83\xda\x7b\x89\x86\x92\xa4\xfe\xd0\xf8\xb9\x7f\xff\x74\xef\x09\xae\xbf\x5e\xc9\x39\xad\x09\x12\xd1\x63\x98\x17\xdc\x99\xf0\x8c\x36\x43\xf4\xef\x9b\x3c\xb3\x3d\x57\xfb\x03\xae\x1c\x5c\x7e\x14\xcd\xff\x08\x7f\x88\xd6\x43\xf6\x1e\xa3\x2d\xe6\x39\xf2\x18\xb8\xc6\xb2\xe2\x37\x53\xfa\x37\x00\x00\xff\xff\xad\xec\xff\x2b\xb4\x03\x00\x00"),
		},
		"/flux-helm-release-crd.yaml.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "flux-helm-release-crd.yaml.tmpl",
			modTime:          time.Date(2019, 8, 6, 14, 33, 53, 502755756, time.UTC),
			uncompressedSize: 4007,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x57\xdf\x6f\xfb\x34\x10\x7f\xcf\x5f\x61\x0d\xa4\x02\x5a\x0a\x13\x08\x41\x5e\xd0\xc4\x34\x0d\x0d\xd8\xb4\x6a\x7b\x99\x86\x74\x4d\xae\xa9\xa9\x63\x9b\xf3\x25\x5b\x41\xfc\xef\xc8\x49\xd3\x2d\x59\x7e\x74\xed\x26\x5e\xbe\x7e\x6a\x7d\xbe\x8f\xef\x3e\xf7\xc3\x97\x30\x0c\x03\xb0\xf2\x0e\xc9\x49\xa3\x23\x01\x56\xe2\x13\xa3\xf6\xff\xdc\x74\xf5\x83\x9b\x4a\xf3\x75\x71\x32\x47\x86\x93\x60\x25\x75\x12\x89\x9f\x73\xc7\x26\xbb\x41\x67\x72\x8a\xf1\x0c\x17\x52\x4b\x96\x46\x07\x19\x32\x24\xc0\x10\x05\x42\x68\xc8\x30\x12\x4b\x54\x19\xa1\x42\x70\xe8\xa6\x0b\x95\x3f\x4d\x1f\x11\x0a\x9c\x3e\x1a\x5a\xb9\xc0\x59\x8c\xfd\xd9\x94\x4c\x6e\x23\xf1\x4a\x5e\xa1\x38\x7f\x44\x88\xea\xee\x0b\x54\xd9\x4d\x05\x58\xee\x2a\xe9\xf8\xb2\x2d\xf9\x55\x3a\x2e\xa5\x56\xe5\x04\xaa\x69\x46\x29\x70\x4b\x43\xfc\xfb\x33\x78\x28\x96\x14\x08\xe1\x62\x63\x31\x12\xa5\xc0\x42\x8c\x89\xdf\xcb\xe7\xb4\x71\x75\x73\xd8\x31\x70\xee\x22\xf1\xcf\xbf\x81\x10\x45\x4d\x5c\xcd\xd1\x76\x6b\x0b\x5d\x51\xf1\x2c\x2f\x31\x90\x0a\x4c\x22\xc1\x94\x63\xbd\xc5\x86\x20\xc5\xed\x5e\x01\x4a\x26\xe0\x89\xad\x80\x8c\x45\x7d\x7a\xfd\xcb\xdd\xb7\xb3\x78\x89\x19\x44\x1b\x35\x4b\xc6\x22\xb1\xac\xad\x2b\xa1\x36\xc4\xd6\x8b\xf0\xaf\x5c\x92\xbf\xef\x7e\x12\x2f\x81\x78\xf2\xf0\x42\xda\x85\x50\x69\x95\x8c\x79\x36\x9a\x02\x21\x78\xed\x79\x72\x4c\x52\xa7\x2d\x91\x05\x66\x24\x1d\x89\xa3\x3f\xee\x21\xfc\xfb\x9b\xf0\xc7\x87\x2f\xee\xc3\xcd\xaf\xaf\xea\xad\x2f\x7f\xfa\xfc\xa8\xa1\xc8\x32\x43\x93\x73\xf7\x45\x52\x33\xa6\x48\x2d\xd9\xc2\x50\x06\x5c\x4a\xbf\xff\xae\x65\xb9\x43\xbe\x03\x95\xb7\x5d\xaa\x01\xe7\xc6\x28\x04\x1d\xb4\xe0\x62\xbc\xb5\x29\x41\xd2\xe3\x6f\x97\x16\x19\xa5\xe6\x10\xaf\xba\x35\xcc\xfc\x4f\x8c\xb9\xcd\x50\x0f\xdf\x7e\xa1\x86\xb9\x7a\x75\xfd\xb0\x09\x5b\xe3\xdf\xae\x46\x18\x13\x02\xef\xa1\x99\x48\xe7\x2d\xbd\x30\x66\xd5\xe1\xc6\x98\x76\x4f\xb4\xc5\x48\xc4\xc5\x60\xd4\xfd\x7a\x04\x39\x80\xda\x65\x4e\xe1\xd3\xe4\x5c\x2a\x9c\x79\x2e\xb8\x27\x61\x80\x08\xd6\x2d\x89\x64\xcc\x3a\x7c\x1f\x88\x7c\xb3\x12\x7d\x5f\x68\x14\x62\xb5\x86\xd2\x63\xd3\x57\x3b\xf6\x07\x6a\xb2\xf4\xd0\x9d\x93\xc9\x3e\xd6\xb7\x61\xc3\x63\xa3\x17\x32\xfd\x0d\xec\x25\xae\x6f\x70\x31\xe4\x43\x0f\xbe\xd8\x8d\xbf\x71\x53\xc4\x20\x8f\x62\xb8\xbf\xd5\x6b\x85\xeb\x83\xf4\x8d\xf5\xad\x1d\xd4\x18\x48\x5f\x09\xf9\x37\xc4\x27\xec\x27\x3a\xcb\x75\x38\x9d\x7e\xec\x21\x0d\x6a\x56\xbe\xf5\xef\xc3\x69\x4e\x6a\x6f\x4a\x73\x1a\x75\xe6\x83\x19\x29\x47\x05\xdf\x1a\xdf\x87\x0c\x0b\xbc\xdc\x9b\x0d\xaf\xfc\xbf\xd2\x61\x34\x5e\x75\xb0\x10\x36\x87\xab\x66\x93\xeb\xf0\xb6\x79\xfe\x65\x09\x8f\x1e\x7e\x95\xa0\xa3\x1a\x2f\x03\xd8\x3a\x5c\x0c\x4c\x47\x1d\xf1\x2c\x91\xda\xa7\x3b\x29\x69\x5a\x90\x4a\x9e\x1c\x8b\xbe\xd0\x0f\x87\x3d\xed\x7e\xca\x77\x88\x76\x3d\x23\xa4\x92\xc5\x67\x42\x1b\x16\x89\xff\x4a\xc1\x44\xcc\xd7\xe2\xea\x74\xd6\xa1\xd4\x9f\x5f\x23\xb7\xd1\x70\x6d\xf4\xea\xb9\x95\xb4\x67\x68\x6f\x6d\xd2\x33\x7f\x0d\x27\x64\x93\x66\x42\x6b\x9c\x64\x43\x6b\xcf\x76\xd9\xc9\x8f\xc5\x64\xf3\x21\xf2\x66\xe2\x9f\xd1\x0e\xe4\x3f\x27\xb5\x2b\xff\x7b\xcc\x35\xd5\xaa\x3f\xc0\x0e\xb3\xd4\x61\x56\x20\xed\x6a\x6c\x59\x0d\xd7\xb9\x52\xd5\xcc\xd8\x7d\xf7\xbb\x3e\x9a\xff\x05\x00\x00\xff\xff\x50\xad\xd4\x1b\xa7\x0f\x00\x00"),
		},
		"/helm-operator-deployment.yaml.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "helm-operator-deployment.yaml.tmpl",
			modTime:          time.Date(2019, 8, 6, 14, 33, 53, 503124839, time.UTC),
			uncompressedSize: 3085,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x56\x4f\x8f\xdb\xc6\x0e\xbf\xfb\x53\x10\xce\x21\x17\x5b\x72\x90\xbc\x77\x50\x4e\xef\x21\x48\xb3\x40\x77\xbb\x68\x82\x02\x39\x25\xf4\x88\xb6\xa6\x1e\x0d\xd5\x19\xca\xae\x60\xa4\x9f\xbd\xe0\x48\x96\xa5\xec\x9f\x20\x7b\xf2\x52\xe4\x8f\xff\x7e\x24\x67\xbd\x5e\x2f\xb0\xb1\x7f\x50\x88\x96\x7d\x01\xd8\x34\x31\x3f\xbe\x5a\x1c\xac\x2f\x0b\x78\x47\x8d\xe3\xae\x26\x2f\x8b\x9a\x04\x4b\x14\x2c\x16\x00\x1e\x6b\x2a\x60\xe7\xda\xbf\xd7\x15\xb9\x7a\xcd\x0d\x05\x14\x0e\xe7\x33\xd8\x1d\x64\x77\x58\x53\x6c\xd0\x10\x7c\xfb\x36\x68\xa7\x7f\x0b\x38\x9f\xe7\x5f\xcf\x67\x20\x5f\xaa\x5a\x6c\xc8\x28\x74\xa0\xc6\x59\x83\xb1\x80\x57\x0b\x80\x48\x8e\x8c\x70\xd0\x2f\x00\x35\x8a\xa9\x7e\xc5\x2d\xb9\xd8\x0b\x9e\x8e\x44\x6d\x25\xa0\xd0\xbe\xeb\x55\xa5\x6b\xa8\x80\xdf\xc9\x04\x42\xa1\x05\x80\x50\xdd\x38\x14\x1a\xa0\x27\xd9\xe9\x9f\x9b\x79\x79\xd6\x8f\xfe\xa1\xf7\x2c\x28\x96\xfd\xc4\xa6\x09\x5c\x93\x54\xd4\xc6\xcc\x72\x1e\x4d\x40\x0d\x61\x29\xa1\xa5\x65\x52\xba\xe4\x9c\x7e\x53\x38\x5a\x43\xff\x33\x86\x5b\x2f\x77\xcf\xbb\x3b\xb2\x6b\x6b\x1a\x5d\xbd\x80\x4f\x15\xc1\x8e\x9d\xe3\x93\xf5\xfb\xe1\x33\xd8\x08\x3b\x0e\xd0\x46\x95\x21\x98\x36\x0a\xd7\x36\x52\x09\x07\xcf\x27\xff\xa5\xe2\x28\x11\x76\xd6\xd1\x6a\x04\x3a\x55\xd6\x54\xd0\x71\x0b\x27\xeb\x1c\x78\xa2\x12\x84\xa1\x64\x6d\xac\x8a\xd5\x48\x7f\x04\xe0\x93\x87\xbd\x15\xed\x18\x43\x40\xa9\x28\x8c\x30\x52\xa1\x1f\x1c\xef\xad\x54\xed\x16\x38\x80\x54\x04\xce\x1e\x28\x83\xcf\xdc\xbe\x74\x0e\xd0\x45\x1e\x5d\xd4\x9a\x37\x58\x19\x31\xac\x17\x4e\x36\x86\xbd\xa0\xf5\x14\x56\xb0\x25\xc7\xa7\xec\xa2\x32\xaa\x7e\xe6\x16\x6a\xec\x7a\xc0\x93\x8d\x95\x02\x36\x81\x8f\xb6\x24\x40\x0f\x31\x56\x5f\x0c\xfb\x9d\xdd\x7f\x97\xae\x32\xdb\xb2\xd7\x38\x6b\x0e\xd4\xc7\xcd\x9e\xe0\xeb\x4d\xa9\x9f\xa4\x7b\x6f\x1d\x7d\x7d\x9b\x0a\x69\x7d\x14\xf4\x86\x56\x43\x2d\x5e\x06\x1a\x81\xfa\x5c\xe7\x18\xbf\x58\xf9\xd0\x6e\x53\x7d\x32\xb8\xfb\x7f\xca\x85\xbc\x84\x0e\x0e\xd4\x41\xac\xb8\x75\x25\x6c\xaf\x18\xcb\x3e\xc4\xe5\x50\xcc\x1e\x68\x79\x8d\x7d\xa9\x7e\x53\x99\xa8\x04\xeb\xe1\x9f\x3c\x8b\xb1\xca\x1f\x96\x63\x3d\xd0\x35\xc6\xaa\xb4\xd7\xa6\x00\xf4\x38\xb7\xd8\x14\x13\xe1\x8c\xdc\x31\x56\xeb\x5e\x6b\xa6\x51\xd2\x0e\x5b\x27\xb7\x5c\x52\x01\x9b\x37\x9b\xcd\xa3\x2d\x98\x50\xa6\xb2\x71\xa4\xe1\xa5\x58\x23\x13\x47\xd6\x48\x85\xd7\x7e\xab\x61\xd4\x6e\x7d\xfc\xf8\x21\x55\x48\x6b\x8e\xc6\x50\x8c\x6f\x81\xb2\x7d\xb6\x02\xbc\xd4\xb4\x4c\x7b\x49\xb5\x32\xb8\xd9\x8d\x10\x33\x3f\x7f\xb6\x51\x52\x1f\x62\x6b\xaa\xe4\x6f\x95\x5a\x30\xe4\x32\x21\xc5\x68\x8f\x2e\x10\x96\x1d\x34\x6c\xbd\x44\x40\x81\x9c\xc4\xe4\x5a\x99\x32\xd7\x5a\xdb\x81\x15\x80\x11\xf0\xe2\x5e\xdd\x5e\x07\x08\xbd\x28\xfb\xda\x48\xdf\xd1\xe1\x40\xdd\x2a\x45\x38\x99\xab\x0b\x47\x2f\x03\x35\xc2\x4c\x18\x8b\x5b\x3e\xd2\x0a\x4e\x56\x2a\xad\xce\x9c\x99\x03\xa1\xd2\x66\xd4\xa4\x09\x4d\x35\x82\x68\x11\xad\x4f\x49\x47\x5d\x7c\x72\xe1\x3b\x95\x50\x51\xa0\xa7\x99\xb3\xb7\xb2\x3e\x50\x37\xe1\x40\x0f\x30\xe7\x4d\x2f\x9b\xec\x2a\x35\xeb\x5b\xf3\x3c\x7b\xd2\xa2\x21\x3f\x92\x59\xab\xbe\x66\xef\xba\x15\x9c\x08\x4e\xec\x5f\x0a\x6c\x09\x70\xeb\x48\xab\x64\xaa\x9a\xcb\x07\xb1\x7e\xaa\x28\x12\xc8\x89\x2f\xeb\x10\x30\x50\x22\x4d\x82\xed\xa9\xa6\x6d\x8f\x56\x38\x58\x8a\x59\x87\xb5\xfb\x6e\x01\xa0\x2f\x87\x26\x0c\x4b\x12\x8d\x52\xc4\x86\x74\x76\xba\x2c\x6d\x56\x87\x22\x14\x74\x9f\x6a\xdf\x48\x87\xd6\x60\x1b\xaf\x93\x3b\x3a\x94\xb4\x87\x43\x4d\xa1\x1f\x86\x1a\x0f\xd4\xef\x30\xc5\xcd\xaf\xc0\xd7\x94\x9f\x6e\xc2\x34\xf6\xb5\xc6\xfe\xb3\xed\x48\xa7\x63\x8a\xf2\xbc\x8b\x14\xe3\x04\x91\xea\x46\xba\x77\x36\x14\x70\xfe\xb6\x18\xe4\xe3\x2e\x1e\x8f\xcf\xfa\x47\xb7\x71\xe8\x55\xa0\xd4\x1f\xcf\xb0\x2c\xf4\xee\x46\x59\x82\xad\x71\x4f\xfd\x95\x9a\x59\x66\xf0\xde\xfa\x32\x15\xae\xd6\x7b\x13\xc8\xe8\x13\xe4\x8a\x17\xc8\x11\x46\xd2\xab\x92\x30\xe0\xd8\xbf\x5f\x74\x64\x2b\x91\x26\x16\x79\x5e\xb5\xdb\xac\x64\x73\xa0\x90\x19\xae\xf3\x90\x9f\x08\x8f\x74\xe2\x70\x88\xf9\xcc\x5b\x2e\xb8\x8f\x13\x70\xe5\x84\x3e\x43\xf4\x89\xa2\x21\x08\xee\x67\xe3\x02\xbd\xcf\x02\x06\x74\xcb\x69\x43\x98\x72\x0e\x5b\x6c\xb2\x57\x9b\x6c\x33\x37\xba\x6f\x9d\xbb\x67\x67\x4d\x57\xc0\xcd\xee\x8e\xe5\x3e\x50\x9c\xe6\xd6\x70\x90\xc9\x1b\xe2\x52\x5d\x4d\x6a\x14\x4e\xda\x70\xcf\x41\x0a\x78\xbd\x79\x7d\xf5\x13\x28\x72\x1b\x0c\x4d\x50\x54\xf8\x57\x4b\x51\x66\x32\x00\xd3\xb4\x05\xfc\x67\x53\xcf\x84\x35\xd5\x1c\xba\x02\xfe\xfb\xe6\xd6\x8e\x1f\xfa\x11\xbb\x55\x9e\x4f\x30\x5e\xc0\x8d\x37\xae\x2d\xa9\x5f\xf8\xc3\x13\x61\x7e\xd1\x9f\x7c\x78\x70\x78\xb8\x82\x15\x52\x07\xf4\xed\x65\x51\x4e\x9e\x08\x15\x5d\x2e\x4a\x49\xc6\x61\xa0\xb2\x5f\x8d\xd9\xc4\xf6\xd1\xcb\xd7\xb3\x39\x45\x73\x8f\x52\x15\x90\x07\x66\x49\xc7\x73\xa6\xa1\x23\xf9\x9b\x77\x5d\x01\xfa\x44\x7b\x04\x75\xbe\x15\x1f\xc2\xce\x6e\xc5\x23\xf6\x4f\x0d\xf4\x43\xa4\x23\x86\x01\x49\x29\x95\x8f\x86\xdd\x8f\x50\xa7\x33\xfc\x13\xb0\x79\x6f\xf7\x6f\x00\x00\x00\xff\xff\x57\x65\xcf\x4a\x0d\x0c\x00\x00"),
		},
	}
	fs["/"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/flux-helm-operator-account.yaml.tmpl"].(os.FileInfo),
		fs["/flux-helm-release-crd.yaml.tmpl"].(os.FileInfo),
		fs["/helm-operator-deployment.yaml.tmpl"].(os.FileInfo),
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