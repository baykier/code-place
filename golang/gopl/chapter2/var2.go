// test for go type
package main 
import (
	"fmt"
	"reflect"
	"github.com/baykier/code-place/golang/gopl/pkg/temp"
)

type baby string
type name string
type age int

var n name
var m age

func main() {
	n = "哈哈哈"
	m = 50
	fmt.Printf("n type is %s\n", reflect.TypeOf(n))
	fmt.Printf("n value is %s\n", n)
	fmt.Printf("%s\n", temp.CToF(60))
	k := temp.CToK(100)
	fmt.Println(k)
}