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
	FlowLimitTime int
	FlowLimitCount int
	FlowSum map[string]FlowLimit
)


type FlowLimit struct {
	Ip string
	StartUnixTime int64
	UpdateUnixTime int64
	FlowCount int
}



func init() {
	flag.BoolVar(&h,"h",false,"This help")
	flag.StringVar(&LitenAddr,"ListenAddr","0.0.0.0:93","Set http server listen address")
	flag.StringVar(&ListenRoute,"ListenRoute","/4u6385IP","Set http server listen Route")
	flag.IntVar(&FlowLimitTime,"LimitTime",1,"Set flow limit time , (Second)")
	flag.IntVar(&FlowLimitCount,"LimitCount",10000,"Set the  flow limit throughput within the time")
}

func OutIPAddress(w http.ResponseWriter, r *http.Request) {
	slice01 := strings.Split(r.RemoteAddr,":")
	log.Println( "RemoteAddr:" , slice01[0],"URL:" , r.URL , "UserAgent:" , r.UserAgent())

	IpsFlowLimit := FlowLimit{Ip:slice01[0]}

	// 系统中有值，可以取出来
	if _, ok := FlowSum[slice01[0]] ; ok {
		IpsFlowLimit.Ip = slice01[0]
		IpsFlowLimit.FlowCount = FlowSum[slice01[0]].FlowCount
		IpsFlowLimit.StartUnixTime = FlowSum[slice01[0]].StartUnixTime
		IpsFlowLimit.UpdateUnixTime = FlowSum[slice01[0]].UpdateUnixTime
	}

	fmt.Println(IpsFlowLimit)
	IpsFlowLimit.CoreCount()

	// 赋值存进系统
	FlowSum[slice01[0]] = IpsFlowLimit
	fmt.Fprintf(w,slice01[0])
}

func OutIPAddressMore(w http.ResponseWriter , r *http.Request) {
	log.Println("/more" , r)
	fmt.Fprintln(w,r)
}

func main() {
	flag.Parse()


	// 申请内存
	FlowSum = make(map[string]FlowLimit)


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

	log.Println("LimitTime:" , FlowLimitTime , "'s" , "  LimitCount:" , FlowLimitCount)
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

func (F *FlowLimit)CoreCount() (bool) {
	if 0 == F.StartUnixTime {
		F.StartUnixTime = time.Now().Unix()
	}

	F.FlowCount++

	F.UpdateUnixTime = time.Now().Unix()

	if 10 < F.FlowCount {
		if 60 > F.UpdateUnixTime - F.StartUnixTime {
			fmt.Println(time.Now(),F.FlowCount,": 限流")
			return false
		} else if 60 < F.UpdateUnixTime - F.StartUnixTime {
			F.StartUnixTime = time.Now().Unix()
			F.FlowCount = 1
		}
	}
	fmt.Println(time.Now(),F.FlowCount,": 正常")

	return true
}