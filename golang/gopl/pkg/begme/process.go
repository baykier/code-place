// this file is part of begme
package main

import (
	"github.com/mitchellh/go-ps"
	"fmt"
	"os"
)

func getProcess() {
	p, err := ps.Processes()


	if err != nil {
		fmt.Println(err)
	}

	for _, process := range p {
		fmt.Println(process.Executable())
	}
}

func killTaskMgr() bool {

	process, err := ps.Processes()
	if err != nil {
		return false
	}
	for _, p := range process {
		//fmt.Println(p.Executable())
		if p.Executable() == "Taskmgr.exe" {
			fmt.Println("关闭任务管理器...")
			proc, err := os.FindProcess(p.Pid())
			if err != nil {
				fmt.Println(err)
			}
			// Kill the process
			proc.Kill()			
			return true
		}
	}
	return false
}