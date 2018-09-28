//like curl

package main 

import (
	"fmt"
	"net/http"
	"io"
	"os"
)

func main () {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch url err:%s \n",err)
			os.Exit(1)
		}
		if  _, err := io.Copy(os.Stdout, resp.Body); err != nil {
			fmt.Fprintf(os.Stderr, "fetch reading err:%s \n",err)
			os.Exit(1)
		}
		resp.Body.Close()	
	}
}