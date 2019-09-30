package main

import (
	"annie-master/config"
	"flag"
	"go-demo/process"
)

func init() {
	flag.IntVar(
		&config.ThreadNumber, "n", 6, "The number of download thread (only works for multiple-parts video)",
	)
}

func main() {
	url := "https://mirrors.aliyun.com/centos/8/BaseOS/x86_64/os/Packages/"
	// 下载文件到XXX下
	process.MultipleDownload(process.Urls(url), url, "D:\\rpms\\")

	//downloadBilibili(process.TypeDemon)

}

func downloadBilibili(videoType int) {
	process.DownloadVideoTopTen(videoType)
}
