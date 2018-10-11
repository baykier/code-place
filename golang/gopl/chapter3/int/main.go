package main

import (
	"fmt"
	"reflect"
)

func main() {
	var i int8
	var ui uint8

	i = -1
	ui = 89

	fmt.Println("ini i is ", i, reflect.TypeOf(i), "byte is ", fmt.Sprintf("%08b", i), fmt.Sprintf("%08b", ^i), ^i)
	fmt.Println("uini ui is ", ui, reflect.TypeOf(ui))
}
