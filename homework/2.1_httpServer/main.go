package main

import (
	"flag"
	"fmt"
	"github.com/golang/glog"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"runtime/metrics"
	"strings"
	"time"
)

func main() {
	metrics.Register()
	flag.Set("v", "4")
	glog.V(2).Info("Starting http server...")
	mux := http.NewServeMux()
	mux.HandleFunc("/", rootHandler)
	mux.HandleFunc("/healthz", healthz)
	mux.HandleFunc("/images", images)
	mux.Handle("/metrics", promhttp.Handler())
	err := http.ListenAndServe(":8081", mux)
	if err != nil {
		log.Fatal(err)
	}

}

func healthz(w http.ResponseWriter, r *http.Request) {
	//访问healthz返回200
	log.Println(" Status Code is " + http.StatusText(200))

}
func rootHandler(w http.ResponseWriter, r *http.Request) {

	io.WriteString(w, "===================Details of the http request header:============\n")
	//获取request的header信息并写入response中
	for k, v := range r.Header {
		io.WriteString(w, fmt.Sprintf("%s=%s\n", k, v))
		for _, vv := range v {
			fmt.Printf("Header key: %s, Header value: %s \n", k, v)
			w.Header().Set(k, vv)
		}

	}
	io.WriteString(w, "=================== VERSION:============\n")
	//获取环境变量VERSION信息
	os.Setenv("VERSION", "1.7.3.5")
	w.Header().Get("VERSION")
	fmt.Println("VERSION is:", os.Getenv("VERSION"))

	io.WriteString(w, "VERSION is:"+os.Getenv("VERSION"))
	//记录客户端ip,http返回码、输出到server端的标准输出
	log.Println("/ ClientAddress is " + getCurrentIP(r) + " Status Code is " + http.StatusText(200))

}
func getCurrentIP(r *http.Request) string {
	// 这里也可以通过X-Forwarded-For请求头的第一个值作为用户的ip ,但是要注意的是这两个请求头代表的ip都有可能是伪造的
	ip := r.Header.Get("X-Real-IP")
	if ip == "" {
		//当请求头不存在即不存在代理时直接获取ip
		ip = strings.Split(r.RemoteAddr, ":")[0]
	}
	return ip
}
func images(w http.ResponseWriter, r *http.Request) {
	timer := metrics.NewTimer()
	defer timer.ObserveTotal()
	randInt := rand.Intn(2000)
	time.Sleep(time.Millisecond * time.Duration(randInt))
	w.Write([]byte(fmt.Sprintf("<h1>%d<h1>", randInt)))
}
