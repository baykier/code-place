//short var
package main

import (
	"fmt"
)

func main() {
	i := 12
	m, n := 45,"sss"

	fmt.Printf("i: %d m :%d n:%s \n",i, m, n)

	x, y := 2, 6

	x, y = y, x //交换

	fmt.Printf("x,y %d,%d\n",x,y)
}