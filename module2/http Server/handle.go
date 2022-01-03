package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	//注册 handle 处理函数
	http.HandleFunc("/index", index)
	http.HandleFunc("/index/test1", test1)
	http.HandleFunc("/index/test2", test2)
	//• ListenAndService
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}

}

//• 定义 handle 处理函数
func index(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello ken are you ok"+"index")
	fmt.Fprintln(w, " ")
	fmt.Fprintln(w, "hello ken are you ok1 ")
}

func test1(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "hello test1 are you ok1 ")
}
func test2(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "hello test2 are you ok1 ")
}

//Handler 处理方法
//type Handler interface {
//	ServeHTTP(ResponseWriter, *Request)
//}
