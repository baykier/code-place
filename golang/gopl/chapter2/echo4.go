//带参数的echo
package main

import (
	"fmt"
	"flag"
	"strings"
)

var n = flag.Bool("n", false, "是否换新行")
var sep = flag.String("sep", " ", "分割字符串")

func main() {
	flag.Parse()
	fmt.Println(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}