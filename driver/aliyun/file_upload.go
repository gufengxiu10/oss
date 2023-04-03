package aliyun

import (
	"io"

	"oss/driver"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

type upload struct {
	client *oss.Client
	bucket *oss.Bucket
}

func newUpload(client *oss.Client, bucket *oss.Bucket) driver.DriverUploader {
	return &upload{
		client: client,
		bucket: bucket,
	}
}

func (u *upload) Obj(key string, fd io.Reader) error {
	return u.bucket.PutObject(key, fd)
}
