package main

import (
	"fmt"
	"net/http"
)

/*
GET /abc HTTP/1.1
User-Agent: PostmanRuntime/7.26.8
Accept: */ /*
Cache-Control: no-cache
Postman-Token: fc89d934-3f32-4f21-875d-f1e6cd828d0a
Host: 127.0.0.1:8080
Accept-Encoding: gzip, deflate, br
Connection: keep-alive

*/
func main() {
	url := "https://movie.douban.com/top250?start=0&filter="
	//url := "https://www.baidu.com"
	//url := "http://127.0.0.1:8080/abc"

	request, _ := http.NewRequest("GET", url, nil)

	request.Header.Set("Content-Type", "text/html; charset=utf-8")
	request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/98.0.4758.82 Safari/537.36")
	//request.Header.Set("Accept", "*/*")
	//request.Header.Set("Cache-Control", "no-cache")
	//request.Header.Set("Postman-Token", "d1a78e3a-e923-4d2a-aca5-a66f5288eee4")
	//request.Header.Set("Host", "movie.douban.com")
	//request.Header.Set("Accept-Encoding", "gzip, deflate, br")
	//request.Header.Set("Connection", "keep-alive")

	//
	//request.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	//request.Header.Set("Accept-Encoding", "gzip, deflate, br")
	//request.Header.Set("Accept-Language", "zh-CN,zh;q=0.9")
	//request.Header.Set("Cache-Control", "max-age=0")
	//request.Header.Set("Connection", "keep-alive")
	//request.Header.Set("Cookie", "douban-fav-remind=1; bid=EZC-I7Q6j84; ll=\"108288\"; _vwo_uuid_v2=DDD26071C028C064DD9930D661BCD87BE|f6dc3175690e662979291026658b62a2; __yadk_uid=MjsldGwa7De4skYx4FQa7lKJJXSIg9RB; __utmz=30149280.1644573101.17.14.utmcsr=baidu|utmccn=(organic)|utmcmd=organic; __utmz=223695111.1644573101.13.4.utmcsr=baidu|utmccn=(organic)|utmcmd=organic; __gads=ID=634d42c8477c2b9f-2292baf38ed000fa:T=1644573097:RT=1644573097:S=ALNI_MaEV7_l3ixmWoqP4V9d3mjRur8Z-Q; __utmc=30149280; __utmc=223695111; ap_v=0,6.0; __utma=30149280.2086929620.1585965417.1644646913.1644649300.22; __utmb=30149280.0.10.1644649300; __utma=223695111.561985530.1585965417.1644646913.1644649300.18; __utmb=223695111.0.10.1644649300; _pk_ref.100001.4cf6=%5B%22%22%2C%22%22%2C1644649300%2C%22https%3A%2F%2Fwww.baidu.com%2Flink%3Furl%3DC1WEt96u5TRvz9ojYhD9Afj57b5LOR92xb4swayIH21MnBobpLKJijY7TkQiOQcT3EwkUUvzXca_vC4EjOr2a_%26wd%3D%26eqid%3Da383a4a20017e6d40000000362063193%22%5D; _pk_ses.100001.4cf6=*; _pk_id.100001.4cf6=2509becdafdeec73.1585965417.18.1644649308.1644647093.")
	//request.Header.Set("Host", "movie.douban.com")
	//request.Header.Set("sec-ch-ua", "\" Not A;Brand\";v=\"99\", \"Chromium\";v=\"98\", \"Google Chrome\";v=\"98\"")
	//request.Header.Set("sec-ch-ua-mobile", "?0")
	//request.Header.Set("sec-ch-ua-platform", "\"Windows\"")
	//request.Header.Set("Sec-Fetch-Dest", "document")
	//request.Header.Set("Sec-Fetch-Mode", "navigate")
	//request.Header.Set("Sec-Fetch-Site", "none")
	//request.Header.Set("Sec-Fetch-User", "?1")
	//request.Header.Set("Upgrade-Insecure-Requests", "1")
	//request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/98.0.4758.82 Safari/537.36")

	client := http.Client{}
	resp, err := client.Do(request)

	if err != nil {
		return
	}
	defer resp.Body.Close()

	buf := make([]byte, 4096)
	var result = []byte{}
	for {
		n, _ := resp.Body.Read(buf)
		if n == 0 {
			break
		}

		copy(result, buf[:n])

	}

	fmt.Println(string(result))
}
