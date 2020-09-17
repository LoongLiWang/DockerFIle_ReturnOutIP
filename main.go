package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
	"gopkg.in/mgo.v2"
)

var (
	h bool
	LitenAddr string
	ListenRoute string
	FlowLimitTime int64
	FlowLimitCount int64
	MongoFlag int64
	MongoHost string
	MongoAuthDB string
	MongoUser string
	MongoPass string
	FlowSum map[string]FlowLimit
	MaxChanLeng = 1000

	MongoChan = make(chan SaveMongoSource,MaxChanLeng)


	MongoSession *mgo.Session
	err error
)


type FlowLimit struct {
	Ip string
	StartUnixTime int64
	UpdateUnixTime int64
	FlowCount int64
}

type SaveMongoSource struct {
	Result string
	Flags int
	UnixTime int64
	Msg string
	Retry int
	TryCount int
	UserAgent string
	FlowLimit
}


func init() {
	flag.BoolVar(&h,"h",false,"This help")
	flag.StringVar(&LitenAddr,"ListenAddr","0.0.0.0:93","Set http server listen address")
	flag.StringVar(&ListenRoute,"ListenRoute","/4u6385IP","Set http server listen Route")
	flag.Int64Var(&FlowLimitTime,"LimitTime",1,"Set flow limit time , (Second)")
	flag.Int64Var(&FlowLimitCount,"LimitCount",10,"Set the  flow limit throughput within the time")
	flag.Int64Var(&MongoFlag,"MongoOn",0,"Mongo Flag (0: Off 1:On)")
	flag.StringVar(&MongoHost,"MongoHost","127.0.0.1","Mongo Host")
	flag.StringVar(&MongoAuthDB,"MongoAuthDB","admin","Mongo Auth DB")
	flag.StringVar(&MongoUser,"MongoUser","none","MongoUser")
	flag.StringVar(&MongoPass,"MongoPass","none","Mongo Password")
}

func OutIPAddress(w http.ResponseWriter, r *http.Request) {
	slice01 := strings.Split(r.RemoteAddr,":")
	if 0 == MongoFlag {
		log.Println( "RemoteAddr:" , slice01[0],"URL:" , r.URL , "UserAgent:" , r.UserAgent())
	}

	IpsFlowLimit := FlowLimit{Ip:slice01[0]}

	// 系统中有值，可以取出来
	if _, ok := FlowSum[slice01[0]] ; ok {
		IpsFlowLimit.Ip = slice01[0]
		IpsFlowLimit.FlowCount = FlowSum[slice01[0]].FlowCount
		IpsFlowLimit.StartUnixTime = FlowSum[slice01[0]].StartUnixTime
		IpsFlowLimit.UpdateUnixTime = FlowSum[slice01[0]].UpdateUnixTime
	}


	Result := IpsFlowLimit.CoreCount()

	// 赋值存进系统
	FlowSum[slice01[0]] = IpsFlowLimit

	if 0 == MongoFlag {
		log.Println("Ips Flow Limit Values: " , IpsFlowLimit)
		log.Println("Maps Values:" , FlowSum)
	}



	if Result {
		fmt.Fprintf(w,slice01[0])

		if 0 != MongoFlag {
			var ss SaveMongoSource
			ss.Result = "正常"
			ss.Flags = 1
			ss.UnixTime = time.Now().Unix()
			ss.FlowLimit = IpsFlowLimit
			ss.UserAgent = r.UserAgent()
			ss.Msg = fmt.Sprintln("正常,阈值:" , FlowLimitTime ,"分钟内最大访问量 ",FlowLimitCount,"次 \n" +
				"实际访问量:","IP地址:",FlowSum[slice01[0]].Ip , " 访问量:" , FlowSum[slice01[0]].FlowCount," 开始时间戳:",FlowSum[slice01[0]].StartUnixTime," 最近访问时间戳:",FlowSum[slice01[0]].UpdateUnixTime)

			if MaxChanLeng <= len(MongoChan)+(MaxChanLeng / 4) {
				for {
					log.Println("Time Out 3's" )
					time.Sleep(3 * time.Second)

					if MaxChanLeng >= len(MongoChan)+(MaxChanLeng / 4) {
						break
					}
				}

			}

			MongoChan <- ss
		}
	} else {
		//fmt.Fprintf(w,"Flow Limit " , r)
		fmt.Fprintln(w,"限流,阈值:" , FlowLimitTime ,"分钟内最大访问量 ",FlowLimitCount,"次 \n" +
			"实际访问量:","IP地址:",FlowSum[slice01[0]].Ip , " 访问量:" , FlowSum[slice01[0]].FlowCount," 开始时间戳:",FlowSum[slice01[0]].StartUnixTime," 最近访问时间戳:",FlowSum[slice01[0]].UpdateUnixTime)

		if 0 != MongoFlag {
			var ss SaveMongoSource
			ss.Result = "异常"
			ss.Flags = 0
			ss.UnixTime = time.Now().Unix()
			ss.FlowLimit = IpsFlowLimit
			ss.UserAgent = r.UserAgent()
			ss.Msg = fmt.Sprintln("限流,阈值:" , FlowLimitTime ,"分钟内最大访问量 ",FlowLimitCount,"次 \n" +
				"实际访问量:","IP地址:",FlowSum[slice01[0]].Ip , " 访问量:" , FlowSum[slice01[0]].FlowCount," 开始时间戳:",FlowSum[slice01[0]].StartUnixTime," 最近访问时间戳:",FlowSum[slice01[0]].UpdateUnixTime)

			if MaxChanLeng <= len(MongoChan)+(MaxChanLeng / 4) {
				for {
					log.Println("Time Out 3's" )
					time.Sleep(3 * time.Second)

					if MaxChanLeng >= len(MongoChan)+(MaxChanLeng / 4) {
						break
					}
				}

			}
			MongoChan <- ss
		}
	}

}

//func OutIPAddressMore(w http.ResponseWriter , r *http.Request) {
//	log.Println("/more" , r)
//	fmt.Fprintln(w,r)
//}

func main() {
	flag.Parse()


	if 0 != MongoFlag {

		// Mongo 存储
		go ReadSinMongo()

		// Mongo 连接
		dail_info := &mgo.DialInfo{
			Addrs:[]string{MongoHost},
			Source:MongoAuthDB,
			Username:MongoUser,
			Password:MongoPass,
			Timeout:5 * time.Second,
		}

		MongoSession , err = mgo.DialWithInfo(dail_info)
		if err != nil {
			panic(err)
		}

		defer MongoSession.Close()
	}

	// 申请内存
	FlowSum = make(map[string]FlowLimit)

	// 开启协程，删除map值
	go func() {
		for {
			if 30 < len(FlowSum) {
				time.Sleep(1 * time.Second)
			} else if 20 < len(FlowSum) {
				time.Sleep(10 * time.Second)
			} else if 10 < len(FlowSum) {
				time.Sleep(30 * time.Second)
			} else {
				time.Sleep(60 * time.Second)
			}



			for k , v := range FlowSum {
				if time.Now().Unix() > v.UpdateUnixTime + (FlowLimitTime*60) {
					log.Println("回收" , k , "值" , ": " , FlowSum[k])
					delete(FlowSum,k)
				}
			}
		}
	}()


	if h {
		flag.Usage()
		os.Exit(0)
	}

	if ! strings.HasPrefix(ListenRoute,"/") {
		ListenRoute = "/" + ListenRoute
	}

	http.HandleFunc(ListenRoute,OutIPAddress)
	http.HandleFunc("/",OutIPAddress)
	//http.HandleFunc("/more",OutIPAddressMore)

	log.Println("LimitTime:" , FlowLimitTime , "'m" , "  LimitCount:" , FlowLimitCount)
	log.Println("Server running on http://" + LitenAddr)
	log.Println("Server running on http://" + LitenAddr + ListenRoute)

	s := &http.Server{
		Addr:	LitenAddr,
	}
	log.Fatal(s.ListenAndServe())
}

func ReadSinMongo() {
	go func() {
		for {

			Lines := <- MongoChan

			if MaxChanLeng <= len(MongoChan)+(MaxChanLeng / 2) {
				if 1 == Lines.Retry && 3 < Lines.TryCount {
					log.Println("管道存储量不足,丢弃重发大于3次的值:" , Lines.FlowLimit)
					continue
				} else if MaxChanLeng <= len(MongoChan)+(MaxChanLeng / 3) {
					log.Println("管道存储量不足，丢弃重发的值: " , Lines.FlowLimit)
					continue
				} else if MaxChanLeng <= len(MongoChan)+(MaxChanLeng / 4) {
					log.Println("管道存储量不足，丢弃值: " , Lines.FlowLimit)
					continue
				}
			}

			MongoSession.SetMode(mgo.Monotonic,true)
			Collection := MongoSession.DB("RP").C("log")
			if err := Collection.Insert(Lines) ; err != nil {

				if 1 == Lines.Retry {
					if 10 == Lines.TryCount {
						log.Println("MongoDB 重试 10 次,丢掉该连接")
						log.Println("丢弃信息: ", Lines)
						continue
					}
				}

				log.Println("MongoDB Insert Error" , "Info:" , Lines.FlowLimit , err)
				Lines.Retry = 1
				Lines.TryCount = Lines.TryCount + 1

				MongoChan <- Lines

			}
		}
	}()
}

func (F *FlowLimit)CoreCount() (bool) {
	if 0 == F.StartUnixTime {
		F.StartUnixTime = time.Now().Unix()
	}

	F.FlowCount++

	NowUpdateUnixTime := time.Now().Unix()
	F.UpdateUnixTime = NowUpdateUnixTime

	//fmt.Println("相差秒数:" , F.UpdateUnixTime - F.StartUnixTime)


	// 开始判断限流
	if FlowLimitCount < F.FlowCount {
		if (FlowLimitTime * 60) < NowUpdateUnixTime - F.UpdateUnixTime {
			//F.StartUnixTime = time.Now().Unix()
			F.UpdateUnixTime = time.Now().Unix()
			//F.FlowCount = 1
		} else {
			F.UpdateUnixTime = NowUpdateUnixTime
		}

		if (FlowLimitTime * 60) > F.UpdateUnixTime - F.StartUnixTime {
			if 0 == MongoFlag {
				log.Println(time.Now(),F.FlowCount,": 限流")
			}
			return false
		}
	}

	// 开始复判限流
	if F.UpdateUnixTime - F.StartUnixTime >= (FlowLimitTime * 60) && FlowLimitTime < F.FlowCount {
		if 0 == MongoFlag {
			log.Println(time.Now(),F.FlowCount,": 限流")
		}
		return false
	}
	if 0 == MongoFlag {
		log.Println(time.Now(),F.FlowCount,": 正常")
	}

	return true
}