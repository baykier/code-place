// begme main file
package main 

import (
	"fmt"
)

func main () {
	fmt.Println("begme !!!")
	//getProcess()
	k := killTaskMgr()

	if k {
		fmt.Println("关闭成功")
	} else {
		fmt.Println("关闭失败")
	}
}