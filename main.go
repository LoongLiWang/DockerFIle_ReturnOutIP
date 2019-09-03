package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

var (
	h bool
	LitenAddr string
	ListenRoute string
)

func init() {
	flag.BoolVar(&h,"h",false,"This help")
	flag.StringVar(&LitenAddr,"ListenAddr","0.0.0.0:93","Set http server listen address")
	flag.StringVar(&ListenRoute,"ListenRoute","/4u6385IP","Set http server listen Route")
}

func OutIPAddress(w http.ResponseWriter, r *http.Request) {
	slice01 := strings.Split(r.RemoteAddr,":")
	log.Println(time.Now(),"-- 4u6385IP --",r)
	fmt.Fprintf(w,slice01[0])
}

func main() {
	flag.Parse()

	if h {
		flag.Usage()
		os.Exit(0)
	}

	if ! strings.HasPrefix(ListenRoute,"/") {
		ListenRoute = "/" + ListenRoute
	}

	http.HandleFunc(ListenRoute,OutIPAddress)

	log.Println("Server running on http://" + LitenAddr + ListenRoute)

	s := &http.Server{
		Addr:	LitenAddr,
		ReadTimeout:10*time.Second,
		WriteTimeout:10*time.Second,
		MaxHeaderBytes:1<<20,
	}
	log.Fatal(s.ListenAndServe())
}