package main

import (
	"fmt"
	"regexp"
)

/**
?s:单行模式
*/
func main() {
	compile := regexp.MustCompile("<div>(?s:(.*?))</div>")

	s := `
	<div>123</div>
<div>abc</div>
<div>et3</div>
<div>
	hahah
123

</div>
`

	submatch := compile.FindAllStringSubmatch(s, -1)
	for _, item := range submatch {
		fmt.Println("aaa:", item[0])
		fmt.Println("bbb:", item[1])
	}

}
