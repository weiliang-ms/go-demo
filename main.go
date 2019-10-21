package main

import (
	"go-demo/process"
)

func main() {
	//url := "https://mirrors.aliyun.com/centos/8/BaseOS/x86_64/os/Packages/"
	//resp, _:= http.Get(url)
	//utils.ReadFileContentAll(resp)
}

func downloadBilibili(videoType int) {
	process.DownloadVideoTopTen(videoType)
}
