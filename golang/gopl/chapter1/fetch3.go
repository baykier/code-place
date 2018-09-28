//like curl

package main 

import (
	"fmt"
	"net/http"
	"io"
	"os"
	"strings"
)

func main () {
	for _, url := range os.Args[1:] {
		if ! strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}
		fmt.Printf("request url is: %s \t", url)
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch url err:%s \n",err)
			os.Exit(1)
		}
		if  _, err := io.Copy(os.Stdout, strings.NewReader(resp.Status)); err != nil {
			fmt.Fprintf(os.Stderr, "fetch reading err:%s \n",err)
			os.Exit(1)
		}
		resp.Body.Close()
		fmt.Println()	
	}
}