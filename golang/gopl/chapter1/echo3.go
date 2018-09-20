// echo 3 版本

package main

import (
	"strings"
	"fmt"
	"os"
)

func main() {
	fmt.Println(strings.Join(os.Args[0:]," "))
}