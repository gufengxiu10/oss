package aliyun

import "github.com/aliyun/aliyun-oss-go-sdk/oss"

type fileManage struct {
	client *oss.Client
}

// 判断文件是否存在
func (f *fileManage) Is() bool {
	return true
}

// 获得文件元信息
func (f *fileManage) Information() {

}

func (f *fileManage) Delete() {

}
