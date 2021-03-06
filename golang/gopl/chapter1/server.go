//server
package main

import (
	"net/http"
	"log"
	"fmt"
)


func main() {
	http.HandleFunc("/", handler)
	fmt.Println("server is running at localhost:8000")	
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.PATH=%q\n", r.URL.Path)
}