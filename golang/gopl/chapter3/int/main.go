package main

import (
	"fmt"
	"reflect"
	"math"
	//"strings"
)

func main() {	


	// 查看默认的整形变量
	var i int
	var int_8 int8
	var int_16 int16
	var int_32 int32
	var int_64 int64
	var uint_8 uint8
	var uint_16 uint8
	var uint_32 uint32
	var uint_64 uint64

	int_8 = -128 // 取值范围 2 n-1
	
	fmt.Println("i ", i, reflect.TypeOf(i))
	fmt.Println("int_8", int_8, reflect.TypeOf(int_8))
	fmt.Println("uint_8", uint_8, reflect.TypeOf(uint_8))
	fmt.Println("int_16", int_16, reflect.TypeOf(int_16))
	fmt.Println("uint_16", uint_16, reflect.TypeOf(uint_16))
	fmt.Println("int_32", int_32, reflect.TypeOf(int_32))
	fmt.Println("uint_32", uint_32, reflect.TypeOf(uint_32))
	fmt.Println("int_64", int_64, reflect.TypeOf(int_64))
	fmt.Println("uint_64", uint_64, reflect.TypeOf(uint_64))

	//
	m := math.Pow(2, 7)

	fmt.Println(m, reflect.TypeOf(m))
}


