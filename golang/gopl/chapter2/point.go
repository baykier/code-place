//指针

package main

import (
	"fmt"
	"reflect"
)

var num int = 1

func p(p *int) {
	*p ++
}



// func main() {

// 	var pt *int
// 	var ppt **int	

// 	x := 600
// 	pt = &x
// 	ppt = &pt
// 	fmt.Printf("x : %d type is : %s \n", x, reflect.TypeOf(x))
// 	fmt.Printf("pt : %d type is : %s \n", pt, reflect.TypeOf(pt))
// 	fmt.Printf("ppt : %d type is : %s \n", ppt, reflect.TypeOf(ppt))



// 	p( &num)

// 	fmt.Printf("num:%d\n",num)

// 	m := &num;
// 	fmt.Printf("m:%x type: %s\n",&m,reflect.TypeOf(&m))
// 	fmt.Printf("m:%x type: %s\n",m,reflect.TypeOf(m))
// }

func test_f() {
	fmt.Printf("this is a test func\n")
}