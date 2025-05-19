// fuse/node.go
package fuse

import (
	"context"
	"os"
	"time"

	"bazil.org/fuse"
	"github.com/ravi100k/gofs3/backend"
)

type Dir struct {
	backend backend.StorageBackend
}

func (d *Dir) Attr(ctx context.Context, a *fuse.Attr) error {
	a.Inode = 1
	a.Mode = os.ModeDir | 0555
	a.Mtime = time.Now()
	return nil
}

func (d *Dir) ReadDirAll(ctx context.Context) ([]fuse.Dirent, error) {
	objects, err := d.backend.ListObjects()
	if err != nil {
		return nil, err
	}

	var dirs []fuse.Dirent
	for _, obj := range objects {
		dirs = append(dirs, fuse.Dirent{
			Name: obj,
			Type: fuse.DT_File,
		})
	}
	return dirs, nil
}
