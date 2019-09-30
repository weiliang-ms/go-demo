package process

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"go-demo/config"
	"go-demo/utils"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

var wgp utils.WaitGroupPool

func MultipleDownload(files []string, url string, storeDiretory string) {
	num := &config.ThreadNumber
	println(num)
	wgp := utils.NewWaitGroupPool(5)
	count := len(files)
	println("下载总个数为：", count)

	t := time.Now()

	// 遍历下载
	for i := 0; i < count; i++ {
		wgp.Add()
		go func(url string, fileName string, storeDiretory string) {
			defer wgp.Done()
			download(url, fileName, storeDiretory)
		}(url, files[i], storeDiretory)
	}

	println("下载耗时：", time.Since(t)/time.Second, "秒")
	wgp.Wait()
}

func download(url string, fileName string, storeDiretory string) {
	//创建目录
	_ = os.MkdirAll(storeDiretory, 0777)

	// 创建文件
	newFile, err := os.Create(storeDiretory + fileName)

	// 返回错误信息
	if err != nil {
		fmt.Println(err.Error())
	}
	defer newFile.Close()
	client := http.Client{}

	println("开始下载文件：" + url + fileName)
	resp, err := client.Get(url + fileName)
	println(fileName+"大小为：", resp.ContentLength/1024, "k")
	defer resp.Body.Close()

	// 写入
	_, err = io.Copy(newFile, resp.Body)
	if err != nil {
		fmt.Println(err.Error())
	}

	println(fileName + "下载完成...")
}

// 解析a标签中的url
func Urls(url string) []string {

	var slice []string

	resp, err := http.Get(url)
	if err != nil {
		// handle error
		println(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	rpms := string(body)

	if err != nil {
		// handle error
	}

	dom, err := goquery.NewDocumentFromReader(strings.NewReader(string(rpms)))
	if err != nil {
		log.Fatalln(err)
	}

	dom.Find("a").Each(func(i int, selection *goquery.Selection) {
		rpm, _ := selection.Attr("href")
		rpmUrl := rpm
		if strings.HasSuffix(rpmUrl, ".rpm") {
			slice = append(slice, rpmUrl)
		}

	})

	return slice
}
