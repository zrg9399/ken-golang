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
	err := http.ListenAndServe(":801", mux)
	if err != nil {
		log.Fatal(err)
	}

}

func healthz(w http.ResponseWriter, r *http.Request) {
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
	for k, v := range r.Header {
		io.WriteString(w, fmt.Sprintf("%s=%s\n", k, v))

	}
	io.WriteString(w, "===================GOROOT VERSION:============\n")

	fmt.Println("GOROOT is:", os.Getenv("GOROOT"))
	fmt.Println("GOVERSION is:", os.Getenv("GOVERSION"))
	io.WriteString(w, os.Getenv("GOROOT"))
	io.WriteString(w, os.Getenv("GOVERSION"))

	log.Println("/ ClientAddress is " + r.RemoteAddr + " Status Code is " + http.StatusText(200))

}
