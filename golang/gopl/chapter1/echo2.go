//for eahc 用法

package main

import (
	"fmt"
	"os"
)

func main () {
	// := 等同于 var x = y 只能用于函数内部
	s, sep := "",""

	for _,arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)

}