package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	url := "http://ip.wang-li.top:93"

	resp , err := http.Get(url);if err != nil {
		fmt.Println("Http Connect Error", err)
	} else {
		b , err := ioutil.ReadAll(resp.Body); if err != nil {
			fmt.Println("Read Body Error" , err)
		} else {
			fmt.Printf("%s",b)
		}
	}
}