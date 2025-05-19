package fuse

import (
	"context"

	"github.com/ravi100k/gofs3/backend"

	"bazil.org/fuse"
	"bazil.org/fuse/fs"
)

func Mount(ctx context.Context, mountpoint string, backend backend.StorageBackend) error {
	conn, err := fuse.Mount(
		mountpoint,
		fuse.FSName("gofs3"),
		fuse.Subtype("s3fs"),
		fuse.ReadOnly(),
	)
	if err != nil {
		return err
	}
	defer conn.Close()

	srv := fs.New(conn, nil)
	filesys := &FS{backend: backend}
	if err := srv.Serve(filesys); err != nil {
		return err
	}

	return nil
}

type FS struct {
	backend backend.StorageBackend
}

func (f *FS) Root() (fs.Node, error) {
	return &Dir{backend: f.backend}, nil
}
