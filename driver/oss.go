package driver

import (
	"io"
)

type Driver interface {
	Upload(string) (DriverUploader, error)
}

type DriverUploader interface {
	Obj(string, io.Reader) error
}

type oss struct {
	driver Driver
}

func New(driver Driver) *oss {
	return &oss{
		driver: driver,
	}
}

func (o *oss) Upload(bucket string) (DriverUploader, error) {
	return o.driver.Upload(bucket)
}
