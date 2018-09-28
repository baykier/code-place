//dup3
package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"strings"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		fmt.Printf("至少输入一个文件")
	} else {
		for _, file := range files {
			data, err := ioutil.ReadFile(file)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup3 err:%v \n",err)
				continue
			}
			for _, line := range strings.Split(string(data), "\n") {
				counts[line] ++
			}
		}
		for line, n := range counts {
			if n > 1 {
				fmt.Printf("%s count is :%d\n", line, n)
			}
		}
	}
}