package main

import (
	"fmt"
	"github.com/baykier/code-place/golang/gopl/pkg/endian"
)

func main() {
	isMsb := endian.IsMSB()
	if isMsb {
		fmt.Println("大端序")
	} else {
		fmt.Println("小端序")
	}
}