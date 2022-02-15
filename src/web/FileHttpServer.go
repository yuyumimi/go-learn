package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
)

func openFile(w http.ResponseWriter, url string) {
	file, err := os.Open("f:/tmp" + url)

	if err != nil {
		w.Write([]byte("not found"))

		fmt.Println("Open :", err)
		return
	}
	defer file.Close()

	buf := make([]byte, 4096)
	for {

		l, err := file.Read(buf)
		if l == 0 {
			return
		}
		if err != nil {
			fmt.Println("Read :", err)
			return
		}
		_, err = w.Write(buf[:l])
		if err != nil {
			fmt.Println("Write :", err)
			return
		}
	}

}

func requestHandler(w http.ResponseWriter, r *http.Request) {

	openFile(w, r.URL.String())
}
func main() {
	show()

	//http.HandleFunc(
	//	"/",
	//	requestHandler,
	//)
	//
	//http.ListenAndServe("127.0.0.1:8080", nil)
}

func show() {
	listen, _ := net.Listen("tcp", "127.0.0.1:8080")

	defer listen.Close()
	conn, _ := listen.Accept()

	buf := make([]byte, 4096)

	n, _ := conn.Read(buf)

	fmt.Println(string(buf[:n]))
}
