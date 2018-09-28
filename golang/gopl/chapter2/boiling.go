//boiling point
package main

import (
	"fmt"
)

const boilingF = 212.0

func main() {
	const ttt = "just a test" //可以不加类型，自动推断
	var f = boilingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("boiling point is %g 。F or %g 。C \n", f, c)
	fmt.Println(ttt)
}