package main 

import (
	"fmt"
	"flag"
	"reflect"
	"github.com/baykier/code-place/golang/gopl/pkg/pop"
)

var in uint64

func main () {

	flag.Uint64Var(&in, "n", 0, "输入要求值的数字")
	flag.Parse()
	fmt.Printf("%d has %d bit\n", in, pop.Count(in))

	var p[256]byte
	for i := range p {
		p[i] = p[i/2] + p[i&1]
	}

	for k, v := range p {
		fmt.Printf("k:%d\tv:%b type is %s\n", k, v, reflect.TypeOf(v))
	}
	fmt.Printf("k:%d\tv:%b type is %s\n", 7, p[7/2], reflect.TypeOf(p[7/2]))
	fmt.Println(p[3/2])
	
}
