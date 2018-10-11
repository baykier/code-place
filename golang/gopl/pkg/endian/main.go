package main

import (
	"fmt"
)

func main() {
	int_16 := int16(0x1234)
	int_8 := int8(int_16)

	fmt.Println(int_8)
	fmt.Println(&int_16)
	if 0x34 == int_8 {
		fmt.Println("小端序")
	} else {
		fmt.Println("大端序")
	}
}