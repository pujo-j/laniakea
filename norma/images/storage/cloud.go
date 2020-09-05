package main

import "context"

import (
	"gocloud.dev/blob"
	_ "gocloud.dev/blob/azureblob"
	_ "gocloud.dev/blob/gcsblob"
	_ "gocloud.dev/blob/s3blob"
)

type CloudStorage struct {
	bucket *blob.Bucket
}

func initCloudStorage(base string) (Storage, error) {
	b, err := blob.OpenBucket(context.Background(), base)
	if err != nil {
		return nil, err
	}
	return &CloudStorage{bucket: b}, nil
}

func (c *CloudStorage) Exists(ctx context.Context, path string) (bool, error) {
	return c.bucket.Exists(ctx, path)
}

func (c *CloudStorage) Get(ctx context.Context, path string) ([]byte, error) {
	return c.bucket.ReadAll(ctx, path)
}

func (c *CloudStorage) Put(ctx context.Context, path string, data []byte) error {
	return c.bucket.WriteAll(ctx, path, data, nil)
}
