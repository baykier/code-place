package main

import (
	"fmt"
	"time"
	"os"
)

func still() {	
	for {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println("task is running...", os.Getpid())
	}
}

func main() {
	still()
}