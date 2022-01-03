package main

import (
	"fmt"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello  world")
}
func htmlHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	html := `<!doctype  html>  
    <META  http-equiv="Content-Type"  content="text/html"  charset="utf-8">    
    <html  lang="zhCN">  
            <head>   
                    <title>Golang</title>  
                    <meta  name="viewport"  content="width=device-width,  initial-scale=1.0,  maximum-scale=1.0,  user-scalable=0;"  />   
            </head>     
            <body>        
            <div  id="app">Welcome!</div>       
            </body>   
    </html>`
	fmt.Fprintf(w, html)
}
func main() {
	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(indexHandler))
	mux.HandleFunc("/welcome", htmlHandler)
	http.ListenAndServe(":8001", mux)
}

//ServerMux结构体
//type ServeMux struct {
//	mu    sync.RWMutex
//	m     map[string]muxEntry
//	hosts bool
//}
//
//type muxEntry struct {
//	explicit bool
//	h        Handler
//	pattern  string
//}
//server结构体
//func ListenAndServe(addr string, handler Handler) error {
//	server := &Server{Addr: addr, Handler: handler}
//	return server.ListenAndServe()
//}

//type Server struct {
//	Addr         string
//	Handler      Handler
//	ReadTimeout  time.Duration
//	WriteTimeout time.Duration
//	TLSConfig    *tls.Config
//
//	MaxHeaderBytes int
//
//	TLSNextProto map[string]func(*Server, *tls.Conn, Handler)
//
//	ConnState func(net.Conn, ConnState)
//	ErrorLog *log.Logger
//	disableKeepAlives int32     nextProtoOnce     sync.Once
//	nextProtoErr      error
//}
