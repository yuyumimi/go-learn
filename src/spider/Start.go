package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

/**
https://tieba.baidu.com/f?kw=%E7%BB%9D%E5%9C%B0%E6%B1%82%E7%94%9F&ie=utf-8&pn=0
https://tieba.baidu.com/f?kw=%E7%BB%9D%E5%9C%B0%E6%B1%82%E7%94%9F&ie=utf-8&pn=50
https://tieba.baidu.com/f?kw=%E7%BB%9D%E5%9C%B0%E6%B1%82%E7%94%9F&ie=utf-8&pn=150
*/
func main() {

	var start, end int
	fmt.Println("开始页：")
	fmt.Scan(&start)
	fmt.Println("结束页：")

	fmt.Scan(&end)

	doWork(start, end)

	fmt.Println("抓取任务完成 ")
}

func doWork(start int, end int) {
	fmt.Printf("正在抓取网页从%d到%d\n", start, end)
	page := make(chan int)
	for i := start; i <= end; i++ {
		go spiderPage(i, page)

	}

	for i := start; i <= end; i++ {
		fmt.Printf("第%d页抓取完成\n", <-page)
	}

}

func spiderPage(i int, page chan int) {

	url := "https://tieba.baidu.com/f?kw=%E7%BB%9D%E5%9C%B0%E6%B1%82%E7%94%9F&ie=utf-8&pn=" + strconv.Itoa((i-1)*50)
	result, err := httpGet(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	//fmt.Println(result)

	f, err := os.Create("f://tmp/第" + strconv.Itoa(i) + "页.html")
	if err != nil {
		fmt.Println("Create err:", err)
		return
	}
	defer f.Close()

	f.WriteString(result)

	page <- i
}

func httpGet(url string) (result string, err error) {
	resp, err1 := http.Get(url)
	if err1 != nil {
		err = err1
	}
	defer resp.Body.Close()
	buf := make([]byte, 2048)
	for {
		n, err2 := resp.Body.Read(buf)

		if err2 != nil && err2 != io.EOF {
			err = err2
		}

		if n == 0 {
			//fmt.Println("抓取网页结束")
			break
		}
		result += string(buf[:n])
	}
	return
}
