package process

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
)

var w sync.WaitGroup

func MultipleDownload(files []string, url string, storeDiretory string) {

	// 设置最大协程数量为3（通过信道方式控制）
	ch := make(chan string, 3)
	count := len(files)
	println("下载总个数为：", count)

	// 遍历下载
	for i := 0; i < count; i++ {
		ch <- files[i]
		go download(url, files[i], storeDiretory, ch)
	}

	// 显示关闭
	defer close(ch)
	w.Wait()
}

func download(url string, fileName string, storeDiretory string, ch chan string) {
	w.Add(1)
	defer w.Done()
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
	println(<-ch, "完成")
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
