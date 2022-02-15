package main

import (
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
	"io"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
)

//http://movie.xiaohua.com/top250
//http://movie.xiaohua.com/top250?start=25&filter=
//http://movie.xiaohua.com/top250?start=50&filter=

const articleUrlReg = "<span class=\"article-title\"><a target=\"_blank\" href=\"(?s:(.*?))\""
const articleReg = "<h1 class=\"article-title\">(?s:(.*?))</h1>"
const textReg = "<div class=\"article-text\">(?s:(.*?))</div>"

var articleUrlCompile = regexp.MustCompile(articleUrlReg)
var articleCompile = regexp.MustCompile(articleReg)
var textCompile = regexp.MustCompile(textReg)

const (
	urlPre = "https://xiaohua.zol.com.cn"
)

func match(s string, compile *regexp.Regexp) (item [][]string) {
	item = compile.FindAllStringSubmatch(s, -1)
	return
}

func parseData(s string, compile *regexp.Regexp) (data []string) {
	items := match(s, compile)
	data = []string{}
	for _, item := range items {
		data = append(data, item[1])
	}

	return
}

func main() {

	var start, end int
	fmt.Println("开始页：")
	fmt.Scan(&start)
	fmt.Println("结束页：")

	fmt.Scan(&end)

	doWorkForxiaohua(start, end)

	fmt.Println("抓取任务完成 ")
}

func doWorkForxiaohua(start int, end int) {
	fmt.Printf("正在抓取网页从%d到%d\n", start, end)
	page := make(chan int)
	for i := start; i <= end; i++ {
		go spiderPagexiaohua(i, page)

	}

	for i := start; i <= end; i++ {
		fmt.Printf("第%d页抓取完成\n", <-page)
	}

}

func spiderPagexiaohua(i int, page chan int) {

	url := "https://xiaohua.zol.com.cn/lengxiaohua/" + strconv.Itoa(i) + ".html"
	result, err := httpGetxiaohua(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	//解析
	articlUrls := parseData(result, articleUrlCompile)

	fmt.Println(articlUrls)

	jokes := []Joke{}
	for _, articlUrl := range articlUrls {
		joke, err := spiderPartical(urlPre + articlUrl)
		if err != nil {
			fmt.Println(err)
			continue
		}
		jokes = append(jokes, joke)

	}

	fmt.Println(jokes[1].title)
	fmt.Println("+++++++++++")
	fmt.Println(jokes[1].content)
	//fmt.Println(nums)
	//fmt.Println(comments)

	saveFile(i, jokes)

	page <- i
}

type Joke struct {
	title   string
	content string
}

func spiderPartical(url string) (joke Joke, err error) {
	result, err1 := httpGetxiaohua(url)
	if err != nil {
		err = err1
		return
	}
	//解析标题
	arts := parseData(result, articleCompile)
	//解析内容
	texts := parseData(result, textCompile)

	for i := 0; i < len(arts); i++ {
		//art := strings.Replace(arts[i], "\t", "", -1)
		text := strings.Replace(texts[i], " ", "", -1)
		text = strings.Replace(text, "\n", "", -1)
		text = strings.Replace(text, "\r", "", -1)
		text = strings.Replace(text, "\t", "", -1)
		joke = Joke{arts[i], text}
		//jokes = append(jokes, joke)
		break
	}
	return
}

func saveFile(id int, joeks []Joke) {
	f, err := os.Create("f://tmp/第" + strconv.Itoa(id) + "页.txt")
	if err != nil {
		fmt.Println("Create err:", err)
		return
	}
	defer f.Close()

	l := len(joeks)

	//
	f.WriteString("标题\t\t\t内容\n")

	for i := 0; i < l; i++ {

		sprintf := fmt.Sprintf("%s\t\t\t%s\n", joeks[i].title, joeks[i].content)
		f.WriteString(sprintf)
	}
}

func httpGetxiaohua(url string) (result string, err error) {

	request, _ := http.NewRequest("GET", url, nil)
	//防抓取
	request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/98.0.4758.82 Safari/537.36")
	request.Header.Set("Content-Type", "text/html; charset=gbk")

	client := http.Client{}
	resp, err1 := client.Do(request)

	//resp, err1 := http.Get(url)
	if err1 != nil {
		err = err1
	}
	defer resp.Body.Close()
	buf := make([]byte, 4096)

	/*
		reader := simplifiedchinese.GB18030.NewDecoder().Reader(resp.Body)
		body, err := ioutil.ReadAll(reader)
		tf8Reader := transform.NewReader(resp.Body, simplifiedchinese.GBK.NewDecoder())
		all, err := ioutil.ReadAll(utf8Reader)
	*/
	//转gbk
	reader := simplifiedchinese.GBK.NewDecoder().Reader(resp.Body)
	for {
		n, err2 := reader.Read(buf)
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
