package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	//name := "f:/file.txt"
	name := "e:/a.mp4"
	//createFile(name)
	//openFile(name)
	//readFile(name)
	//文件拷贝

	copyName := "f:/a.mp4"
	copyFile(name, copyName)

}

/**
文件拷贝
*/
func copyFile(file string, copyFile string) {
	f_r, err := os.OpenFile(file, os.O_RDWR, 6)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("open succ")

	defer f_r.Close()

	// 创建写文件
	copyF, err := os.Create(copyFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer copyF.Close()

	//从缓存区读取
	buf := make([]byte, 4096)
	for {
		len, err := f_r.Read(buf)
		if err != nil && err == io.EOF {
			fmt.Printf("read length: %d\n", len)
			break
		}
		copyF.Write(buf[:len])
	}
	fmt.Println("write finish!!!")

}

func createFile(file string) {
	create, err := os.Create(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer create.Close()
	fmt.Println("create succ")
}
func openFile(file string) {
	f, err := os.OpenFile(file, os.O_RDWR, 6)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("open succ")

	defer f.Close()
	seek, err := f.Seek(5, io.SeekStart)
	fmt.Println(seek)
	//writeString, err := f.WriteAt([]byte("44"), seek)
	writeString, err := f.WriteString("33")

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("writeString", writeString)

}

func readFile(file string) {
	f, err := os.OpenFile(file, os.O_RDWR, 6)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("open succ")

	defer f.Close()

	reader := bufio.NewReader(f)

	for {

		bytes, err := reader.ReadBytes('1')

		if err != nil {
			if err == io.EOF {
				fmt.Println("文件结束", err)
				return
			}
			fmt.Println(err)
			return
		}

		fmt.Println(string(bytes))
	}
}
