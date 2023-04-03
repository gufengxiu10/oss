package main

import (
	"fmt"
	"oss/driver"
	"oss/driver/aliyun"
)

func main() {
	fmt.Println(driver.New(aliyun.NewAuth(aliyun.HuDong, "102", "324")))
}
