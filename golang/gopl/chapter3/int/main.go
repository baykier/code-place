package main

import (
	"fmt"
	"reflect"
)

func main() {
	var i int8
	var ui uint8

	i = 34
	ui = 89

	fmt.Println("ini i is ", i, reflect.TypeOf(i), "byte is ", fmt.Sprintf("%b", i))
	fmt.Println("uini ui is ", ui, reflect.TypeOf(ui))
}
