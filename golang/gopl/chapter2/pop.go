package main 

import (
	"fmt"
	"flag"
	//"reflect"
	"github.com/baykier/code-place/golang/gopl/pkg/pop"
)

var in uint64

func main () {


	flag.Uint64Var(&in, "n", 0, "输入要求值的数字")
	flag.Parse()
	fmt.Printf("%d has %d bit\n", in, pop.Count(in))
	fmt.Println(byte(11))
	
}
