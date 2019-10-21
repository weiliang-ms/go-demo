package utils

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
)

// 读取系统文件内容
func ReadFileContent(filePath string) (content string) {

	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		println(err)
	}
	return string(b)
}

// 读取系统文件内容
func ReadFileContentAll(resp *http.Response) {
	defer resp.Body.Close()
	//body, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	println(err)
	//}
	//fmt.Println(string(body))

	var b bytes.Buffer
	_, err := io.Copy(&b, resp.Body)
	if err != nil {
		println(err)
	}
	println(string(b.Bytes()))
}
