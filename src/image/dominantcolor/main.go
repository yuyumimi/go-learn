package main

import (
	"fmt"
	"github.com/cenkalti/dominantcolor"
	"image"
	"image/color"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io/ioutil"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

/**
Find dominant color in images
提供一个rgb颜色，分析出主色与其相似的图片
https://github.com/cenkalti/dominantcolor.git
*/
func FindDomiantColor(fileInput string) (string, color.RGBA, error) {
	f, err := os.Open(fileInput)
	defer f.Close()
	if err != nil {
		fmt.Println("File not found:", fileInput)
		return "", color.RGBA{}, err
	}
	img, _, err := image.Decode(f)
	if err != nil {
		return "", color.RGBA{}, err
	}

	c := dominantcolor.Find(img)
	return dominantcolor.Hex(c), c, nil
}
func Color16ToRGB(colorStr string) (red, green, blue int, err error) {
	color64, err := strconv.ParseInt(strings.TrimPrefix(colorStr, "#"), 16, 32)
	if err != nil {
		return
	}
	colorInt := int(color64)
	return colorInt >> 16, (colorInt & 0x00FF00) >> 8, colorInt & 0x0000FF, nil
}
func distance(a, b color.RGBA) float64 {
	dr := uint32(a.R) - uint32(b.R)
	dg := uint32(a.G) - uint32(b.G)
	db := uint32(a.B) - uint32(b.B)
	return math.Sqrt(float64(dr*dr + dg*dg + db*db))
}

type Result struct {
	color    string
	filePath string
	distance float64
}

func main() {
	rgba := color.RGBA{R: 51, G: 51, B: 255}
	fmt.Println(rgba)

	path := "F:\\yuyu\\workspace\\go-workspace\\go-learn\\src\\image\\pic\\"
	pics := scanDir(path, "")
	//F:\\yuyu\\workspace\\go-workspace\\go-learn\\dominantcolor-master\\firefox.png

	results := make([]Result, len(pics))
	for i, pic := range pics {

		filePath := path + pic
		color, rgb, err := FindDomiantColor(filePath)
		if err != nil {
			fmt.Println("FindDomiantColor err:", err)
			break
		}
		dis := distance(rgb, rgba)
		result := Result{color, filePath, dis}
		results[i] = result
		fmt.Printf("%s pic: %s dis: %f \n ", pic, color, dis)
	}
	//sort.Sort(ResultSlice(results)) // 按照 distance 的逆序排序
	//fmt.Println(results)

	sort.Sort(sort.Reverse(ResultSlice(results))) // 按照 distance 的升序排序
	fmt.Println(results)
	outHtml(results, path)

}

//结构体排序开始
type ResultSlice []Result

func (r ResultSlice) Len() int {
	return len(r)
}

func (r ResultSlice) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

func (r ResultSlice) Less(i, j int) bool { // 重写 Less() 方法， 从大到小排序
	return r[j].distance < r[i].distance
}

//结构体排序结束

func outHtml(results []Result, outputDirectory string) {
	var buff strings.Builder
	buff.WriteString("<html><body><h1>Colors listed in order of dominance: hex color followed by number of entries</h1><table border=\"1\">")

	for _, result := range results {

		buff.WriteString("<tr><td><img src=\"" + result.filePath + "\" width=\"200\" border=\"1\"></td>")
		buff.WriteString("<td style=\"background-color: #%s;width:200px;height:50px;text-align:center;\">")
		buff.WriteString("<h3>" + result.color + "</h3>")

		buff.WriteString("<h3>" + fmt.Sprintf("%f", result.distance) + "</h3>")
		buff.WriteString("</td>")
		buff.WriteString("</tr>")
	}

	buff.WriteString("</table></body><html>")

	// And write it to the disk
	var err error
	if err = ioutil.WriteFile(outputDirectory+"output.html", []byte(buff.String()), 0644); err != nil {
		panic(err)
	}
}

/**
判断文件后缀类型
*/
func checkFile(name string, suffix string) bool {
	return strings.HasSuffix(name, suffix)
}
func scanDir(path string, suffix string) (files []string) {
	file, err := os.OpenFile(path, os.O_RDONLY, os.ModeDir)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	fileInfos, err := file.Readdir(-1)

	if err != nil {
		fmt.Println(err)
		return
	}

	num := len(fileInfos)
	files = make([]string, num)
	for i, info := range fileInfos {
		if info.IsDir() {
			//fmt.Println("目录：", info.Name())
		} else {
			name := info.Name()
			if checkFile(name, suffix) {
				//fmt.Println("拷贝文件：", name)
				//copyFileToPath(path+"/"+name, des+"/"+name)
				//fmt.Println("统计文件：", name)
				files[i] = name
			}
		}
	}

	return
}
