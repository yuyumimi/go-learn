package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	dir(".txt")
}

/**
目录操作
*/
func dir(suffix string) {

	//des := "f:/tmp"
	fmt.Println("输入目录(例如：f:\\dir\\mp4)")
	var path string
	//fmt.Scanln(&path)
	//fmt.Scanf(" %s", &path)
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		path = scanner.Text()
	}
	fmt.Println("输入的目录为：", path)
	file, err := os.OpenFile(path, os.O_RDONLY, os.ModeDir)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	fileInfo, err := file.Readdir(-1)

	if err != nil {
		fmt.Println(err)
		return
	}
	count := 0
	for _, info := range fileInfo {
		if info.IsDir() {
			//fmt.Println("目录：", info.Name())
		} else {
			name := info.Name()
			if checkFile(name, suffix) {
				//fmt.Println("拷贝文件：", name)
				//copyFileToPath(path+"/"+name, des+"/"+name)
				fmt.Println("统计文件：", name)

				count += countWordInFile("love", path+"/"+name)
			}
		}
	}
	fmt.Println(count)
}

/**
判断文件后缀类型
*/
func checkFile(name string, suffix string) bool {
	return strings.HasSuffix(name, suffix)
}

/**
拷贝文件到指定目录
*/
func copyFileToPath(src, des string) {
	f_r, err := os.OpenFile(src, os.O_RDWR, 6)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("open succ")

	defer f_r.Close()

	// 创建写文件
	copyF, err := os.Create(des)
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

/**
统计文件里面的词
*/
func countWordInFile(word, fileName string) int {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		return -1
	}
	reader := bufio.NewReader(file)
	count := 0
	for {

		bytes, err := reader.ReadBytes('\n')
		if err != nil && err == io.EOF {
			break
		}
		row := string(bytes[:])
		fmt.Println(row)
		count += wordCount(row)
	}
	return count
}

func wordCount(row string) int {
	fields := strings.Fields(row)
	count := 0
	for _, field := range fields {
		if field == "Love" {
			count += 1
		}
	}
	return count
}
