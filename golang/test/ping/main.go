// ping wanglibao site is avaiable

package main

import (
	"fmt"
	"net"
	"net/http"
	"strings"
	"time"
	"io"
	"io/ioutil"
	//"encoding/json"
)

const (
	WANGLIBAO_URL string = "https://www.wanglibao.com/api.php?v=1" //网利宝接口
)

func main(){
	fmt.Println("校验接口是否部署...")

	check := false
	for {

		if check {
			break
		}
		status, sec, bytes, error := getch(WANGLIBAO_URL)
		if error != nil {
			fmt.Println(error)
			continue
		}
		if status {
			fmt.Println("ping result is ok", sec, "seconds", bytes, "bytes")
			check = true
		} else {
			fmt.Println("ping error!", error)
		}
	}
}

func getch(url string) (bool, int64, int64, error) {

	var defaultTransport http.RoundTripper = &http.Transport{
		Proxy: nil,
		DialContext: (&net.Dialer{
			Timeout:   10 * time.Second,
			KeepAlive: 30 * time.Second,
			DualStack: true,
		}).DialContext,
		MaxIdleConns:          30,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   15 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}
	
	client := &http.Client{Transport: defaultTransport}

	if ! strings.HasPrefix(url, "https://") {
		url = "https://" + url
	}
	start := time.Now()
	resp, err := client.Get(url)
	if err != nil {
		return false, 0, 0, err
	}
	nBytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		return false, 0, 0, err
	}
	secs := time.Since(start).Seconds()
	return  true, int64(secs), nBytes, nil
}