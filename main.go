package main

import (
	"fmt"
	"main/driver"
	"main/driver/aliyun"
)

func main() {
	fmt.Println(driver.New(aliyun.NewAuth(aliyun.HuDong, "102", "324")))
}
