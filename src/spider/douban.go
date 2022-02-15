package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strconv"
)

//http://movie.douban.com/top250
//http://movie.douban.com/top250?start=25&filter=
//http://movie.douban.com/top250?start=50&filter=

var titleReg = "<img width=\"100\" alt=\"(?s:(.*?))\""
var numReg = "<span class=\"rating_num\" property=\"v:average\">(?s:(.*?))</span>"
var commentReg = "<span>(\\d*?)人评价</span>"
var titleCompile = regexp.MustCompile(titleReg)
var numCompile = regexp.MustCompile(numReg)
var commentCompile = regexp.MustCompile(commentReg)

func match(s string, compile *regexp.Regexp) (item [][]string) {
	item = compile.FindAllStringSubmatch(s, -1)
	return
}

func parseData(s string, compile *regexp.Regexp) (data []string) {
	items := match(s, compile)
	data = make([]string, 25)
	for i, item := range items {
		data[i] = item[1]
	}

	return
}

func main() {

	var start, end int
	fmt.Println("开始页：")
	fmt.Scan(&start)
	fmt.Println("结束页：")

	fmt.Scan(&end)

	doWorkForDouban(start, end)

	fmt.Println("抓取任务完成 ")
}

func doWorkForDouban(start int, end int) {
	fmt.Printf("正在抓取网页从%d到%d\n", start, end)
	page := make(chan int)
	for i := start; i <= end; i++ {
		go spiderPageDouban(i, page)

	}

	for i := start; i <= end; i++ {
		fmt.Printf("第%d页抓取完成\n", <-page)
	}

}

func spiderPageDouban(i int, page chan int) {

	url := "http://movie.douban.com/top250?start=" + strconv.Itoa((i-1)*25) + "&filter="
	result, err := httpGetDouban(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	//解析标题
	titles := parseData(result, titleCompile)
	//解析评分
	nums := parseData(result, numCompile)
	//解析评论数量
	comments := parseData(result, commentCompile)

	fmt.Println(titles)
	fmt.Println(nums)
	fmt.Println(comments)

	saveFile(i, titles, nums, comments)

	page <- i
}

func saveFile(id int, titles []string, nums []string, comments []string) {
	f, err := os.Create("f://tmp/第" + strconv.Itoa(id) + "页.txt")
	if err != nil {
		fmt.Println("Create err:", err)
		return
	}
	defer f.Close()

	l := len(titles)

	//标题 电影名称	评分	评分认数
	f.WriteString("电影名称\t\t\t评分\t\t\t评分人数\n")

	for i := 0; i < l; i++ {

		sprintf := fmt.Sprintf("%s\t\t\t%s\t\t\t%s\n", titles[i], nums[i], comments[i])
		f.WriteString(sprintf)
	}
}

func httpGetDouban(url string) (result string, err error) {

	request, _ := http.NewRequest("GET", url, nil)
	//豆瓣防抓取
	request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/98.0.4758.82 Safari/537.36")
	client := http.Client{}
	resp, err1 := client.Do(request)

	//resp, err1 := http.Get(url)
	if err1 != nil {
		err = err1
	}
	defer resp.Body.Close()
	buf := make([]byte, 4096)
	for {
		n, err2 := resp.Body.Read(buf)
		if n == 0 {
			//fmt.Println("抓取网页结束")
			break
		}
		if err2 != nil && err2 != io.EOF {
			err = err2
			return
		}

		result += string(buf[:n])
	}
	return
}
