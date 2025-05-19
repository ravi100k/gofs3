package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/ravi100k/gofs3/backend"
	"github.com/ravi100k/gofs3/fuse"
)

func main() {
	bucket := os.Getenv("S3_BUCKET")
	if bucket == "" {
		log.Fatal("S3_BUCKET environment variable is required")
	}

	mountpoint := "/mnt/gofs3" // change as needed
	fmt.Println("Mounting", mountpoint)

	s3Client, err := backend.NewS3Backend(bucket)
	if err != nil {
		log.Fatalf("failed to create S3 backend: %v", err)
	}

	err = fuse.Mount(context.Background(), mountpoint, s3Client)
	if err != nil {
		log.Fatalf("failed to mount FUSE filesystem: %v", err)
	}
}
