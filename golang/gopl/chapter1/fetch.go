//like curl

package main 

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"os"
)

func main () {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch url err:%s \n",err)
			os.Exit(1)
		}
		body, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch reading err:%s \n",err)
			os.Exit(1)
		}
		fmt.Printf("url:[%s] retuan[%s]", url, body)
	}
}