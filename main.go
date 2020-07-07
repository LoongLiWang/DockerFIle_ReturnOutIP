package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
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
	log.Println( "RemoteAddr:" , slice01[0],"URL:" , r.URL , "UserAgent:" , r.UserAgent())
	fmt.Fprintf(w,slice01[0])
}

func OutIPAddressMore(w http.ResponseWriter , r *http.Request) {
	log.Println("/more" , r)
	fmt.Fprintln(w,r)
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
	http.HandleFunc("/",OutIPAddress)
	http.HandleFunc("/more",OutIPAddressMore)

	log.Println("Server running on http://" + LitenAddr)
	log.Println("Server running on http://" + LitenAddr + ListenRoute)

	s := &http.Server{
		Addr:	LitenAddr,
		//ReadTimeout:10*time.Second,
		//WriteTimeout:10*time.Second,
		//MaxHeaderBytes:1<<20,
	}
	log.Fatal(s.ListenAndServe())
}