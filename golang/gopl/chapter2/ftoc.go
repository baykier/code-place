//boiling F convert to C
package main

import (
	"fmt"
	"reflect"
)

// 温度转换
//函数定义 函数命，参数，参数类型，返回值类型
func fToC (f float64) float64 {
	return (f - 32) * 5 / 9
}

func main () {
	//常量const 函数内部只能自己访问
	const freezingF, boilingF = 32.0, 212.0 
	var check_type = 1;
	var check_type2 = 2.0;//默认 float64
	var ct3 float32 = 3.9
	fmt.Printf("check_type value is %d type is %s \n", check_type, reflect.TypeOf(check_type))
	fmt.Printf("check_type2 value is %f type is %s \n", check_type2, reflect.TypeOf(check_type2))
	fmt.Printf("check_type3 value is %f type is %s \n", ct3, reflect.TypeOf(ct3))
	fmt.Printf("freezing point is %gF or %gC\n", freezingF, fToC(freezingF))
	fmt.Printf("boiling point is %gF or %gC \n", boilingF, fToC(boilingF))
}
