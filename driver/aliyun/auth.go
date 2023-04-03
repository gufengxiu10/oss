package aliyun

import (
	"main/driver"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

var (
	HuDong = "https://oss-cn-hangzhou.aliyuncs.com"
)

type auth struct {
	key      string
	secret   string
	endpoint string
	client   *oss.Client
}

type facede struct {
}

func NewAuth(endpoint, key, secret string) driver.Driver {
	return &auth{
		key:      key,
		secret:   secret,
		endpoint: endpoint,
	}
}

func (a *auth) facede() *facede {
	return &facede{}
}

func (a *auth) OssClient() error {
	client, err := oss.New(a.endpoint, a.key, a.secret)
	if err != nil {
		return err
	}

	a.client = client
	return nil
}

func (a *auth) defaultClient() error {
	if a.client == nil {
		if err := a.OssClient(); err != nil {
			return err
		}
	}

	return nil
}

// oss的client
func (a *auth) Upload(bucketName string) (driver.DriverUploader, error) {

	if err := a.defaultClient(); err != nil {
		return nil, err
	}

	bucket, err := a.client.Bucket(bucketName)
	if err != nil {
		return nil, err
	}

	return newUpload(a.client, bucket), nil
}
