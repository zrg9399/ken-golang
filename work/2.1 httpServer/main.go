package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	//注册 handle 处理函数
	http.HandleFunc("/healthz", healthz)
	//Use the default DefaultServeMux.
	//• ListenAndService
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal(err)
	}

}

//• 定义 handle 处理函数
func healthz(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "ok")
}
