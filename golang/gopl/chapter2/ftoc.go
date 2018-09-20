//boiling F convert to C
package main

import (
	"fmt"
)

// 温度转换
func fToC (f float64) float64 {
	return (f - 32) * 5 / 9
}

func main () {
	const freezingF, boilingF = 32.0, 212.0
	fmt.Printf("freezing point is %gF or %gC\n", freezingF, fToC(freezingF))
	fmt.Printf("boiling point is %gF or %gC \n", boilingF, fToC(boilingF))
}
