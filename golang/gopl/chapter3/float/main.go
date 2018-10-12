//float test

package main 

import (
	"fmt"
	"reflect"
)

func main () {
	var float_32 float32
	var float_64 float64

	fmt.Println("float_32 var ", float_32, reflect.TypeOf(float_32))
	fmt.Println("float_64 var ", float_64, reflect.TypeOf(float_64))
}