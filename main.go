package main

import (
	"go-demo/process"
)

func main() {

	url := "https://mirrors.aliyun.com/centos/8/BaseOS/x86_64/os/Packages/"
	// 下载文件到XXX下
	process.MultipleDownload(process.Urls(url), url, "D:\\rpms\\")
}
