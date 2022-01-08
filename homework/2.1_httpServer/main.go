package main

import (
	"flag"
	"fmt"
	"github.com/golang/glog"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	flag.Set("v", "4")
	glog.V(2).Info("Starting http server...")
	mux := http.NewServeMux()
	mux.HandleFunc("/", rootHandler)
	mux.HandleFunc("/healthz", healthz)
	//	mux.HandleFunc("/healthz", healthz)
	err := http.ListenAndServe(":801", mux)
	if err != nil {
		log.Fatal(err)
	}

}

////创建一个结构体
//type logResposeWriter struct {
//	http.ResposeWriter1
//	statusCode int
//}
//
//func NewLogResponseWriter(w http.ResponseWriter1) *logResposeWriter {
//	return &logResposeWriter{w, http.StatusOK}
//}
//func (lw *logResposeWriter) WriteHeader(code int) {
//	lw.statusCode = code
//	lw.ResposeWriter1.WriteHeader(code)
//
//}
//func AccessLog(next http.Header) http.Header {
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		logW := NewLogResponseWriter(w)
//		log.Println("/healthz ClientAddress is " + r.RemoteAddr + " Status Code is " + logW.statusCode)
//	})
//
//}
func healthz(w http.ResponseWriter, r *http.Request) {
	//访问healthz返回200
	//logW := NewLogResponseWriter(w)
	log.Println("/healthz ClientAddress is " + r.RemoteAddr + " Status Code is " + http.StatusText(200))
	io.WriteString(w, "/healthz ClientAddress is "+r.RemoteAddr+" Status Code is "+http.StatusText(200))
}
func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("entering root handler")
	user := r.URL.Query().Get("user")
	if user != "" {
		io.WriteString(w, fmt.Sprintf("hello [%s]\n", user))
	} else {
		io.WriteString(w, "hello [stranger]\n")
	}
	io.WriteString(w, "===================Details of the http request header:============\n")
	//获取request的header信息并写入response中
	for k, v := range r.Header {
		io.WriteString(w, fmt.Sprintf("%s=%s\n", k, v))

	}
	io.WriteString(w, "=================== VERSION:============\n")
	//获取环境变量VERSION信息
	os.Setenv("VERSION", "1.7.3.5")
	fmt.Println("VERSION is:", os.Getenv("VERSION"))
	io.WriteString(w, "VERSION is:"+os.Getenv("VERSION"))
	//记录客户端ip,http返回码、输出到server端的标准输出
	log.Println("/ ClientAddress is " + r.RemoteAddr + " Status Code is " + http.StatusText(200))

}
