package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
)

const reg = "<img loading=\"lazy\" src=\"(?s:(.*?))\""

const dir = "f://tmp/pic/第%d页.jpg"

var compile = regexp.MustCompile(reg)

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
	page := make(chan int)

	url := "https://www.douyu.com/g_yz"

	result, err := httpGet1(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	pics := parseData(result, compile)

	for i, pic := range pics {
		go saveImg(pic, i, page)
	}
	for i := 0; i < len(pics); i++ {
		fmt.Printf("第%d页抓取完成\n", <-page)
	}

}

func saveImg(url string, i int, page chan int) {
	resp, err1 := http.Get(url)
	if err1 != nil {
		fmt.Println(err1)
		return
	}
	defer resp.Body.Close()

	filename := fmt.Sprintf(dir, i)

	f, err := os.Create(filename)
	if err != nil {
		fmt.Println("Create err:", err)
		return
	}
	defer f.Close()

	buf := make([]byte, 4*1024)
	for {
		n, err1 := resp.Body.Read(buf)
		if n == 0 {
			break
		}
		if err1 != nil {
			fmt.Println(err1)
			return
		}
		f.Write(buf[:n])

	}

	page <- i
}

func httpGet1(url string) (result string, err error) {
	resp, err1 := http.Get(url)
	if err1 != nil {
		err = err1
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	result = string(body)
	return
}
