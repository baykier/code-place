//变量
package main 
import (
	"fmt"
	"reflect"
)

var _1 int

func main() {
	m := 12
	p := &m
	fmt.Println(m)
	fmt.Println(reflect.TypeOf(m))
	fmt.Println(p)
	fmt.Println(*p)
	fmt.Println(reflect.TypeOf(p))
	
	x, y := 9, 8

	n := gcd(x, y)
	fmt.Println(n)
	x, y = x+y, x-y
	fmt.Printf("x:%d\ty:%d\n", x, y)

	// below is slice type variable
	arr := []string{"222", "333", "555"}
	for k, v := range arr {
		fmt.Printf("k:%d v:%s \n", k, v)
	}

}

func gcd(x, y int) int {
	if y != 0 {
		x, y = y, x%y
	}
	return x
}