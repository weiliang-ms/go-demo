package process

import (
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

type VideoInfo struct {
	Msg    string `json:"msg"`
	Result []struct {
		Arcurl     string `json:"arcurl"`
		RankOffset int    `json:"rank_offset"`
		Title      string `json:"title"`
	} `json:"result"`
}

const (
	TypeDemon = 22 //鬼畜
)

var wg sync.WaitGroup

func DownloadVideoTopTen(videoType int) {

	println("下载类型：", videoType)

	base := "https://s.search.bilibili.com/cate/search?main_ver=v3&search_type=video&view_type=hot_rank&pic_size=160x100&order=click&copy_right=-1"
	today := time.Now().Format("20060102")
	sevenDaysAgo := time.Now().AddDate(0, 0, -7).Format("20060102")
	url := base + "&cate_id=" + strconv.Itoa(videoType) + "&page=1&pagesize=10" + "&time_from=" + sevenDaysAgo + "&time_to=" + today
	for _, v := range getVideoInfo(url).Result {
		go downloadVideo(v.Arcurl, v.Title, "D:\\bilibili\\"+strconv.Itoa(videoType)+"\\")
	}
	wg.Wait()
}

// 获取视频信息（包括url、名称、本周排行）
func getVideoInfo(url string) VideoInfo {

	reps, err := http.Get(url)
	if err != nil {
		println(err)
	}

	body, err := ioutil.ReadAll(reps.Body)
	if err != nil {
		fmt.Printf("read body err, %v\n", err)
	}

	var result VideoInfo
	resultErr := json.Unmarshal(body, &result)
	if resultErr != nil {
		println(resultErr)
	}

	return result

}

func downloadVideo(url string, videoName string, storeDiretory string) {
	wg.Add(1)
	url = strings.Replace(url, "bilibili", "ibilibili", -1)
	//println("待下载url" + url)

	// 创建目录
	_ = os.MkdirAll(storeDiretory, 0777)
	getVideoUrl(url)

	defer wg.Done()
}

// 解析a标签中的url
func getVideoUrl(url string) string {

	var videoUrl string

	resp, err := http.Get(url)
	if err != nil {
		// handle error
		println(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	video := string(body)

	if err != nil {
		// handle error
	}

	dom, err := goquery.NewDocumentFromReader(strings.NewReader(string(video)))
	if err != nil {
		log.Fatalln(err)
	}

	println(dom.Find("script[type=\"application/ld+json\"]").Text())

	return videoUrl
}
