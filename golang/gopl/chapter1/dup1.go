//打印给定的文件
package main

import (
	"fmt"
	"os"
	"bufio"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {		
		counts[input.Text()]++
		if len(counts) > 10 {
			break
		}
	}
	for line,n := range counts {
		if n > 0 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}

}