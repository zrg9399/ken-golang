package context

import (
	"encoding/json"
	"fmt"
	"httpServer/config"
	"io"
	"log"
	"net/http"
	"strings"
)

const (
	Options = "OPTIONS"
	Get     = "GET"
	Post    = "POST"
	Put     = "PUT"
	Delete  = "DELETE"
)

//定义函数
type Handle func(ctx *Context)

//定义路由结构体
type Router struct {
	router map[string]map[string][]Handle
}

//定义Context结构体
type Context struct {
	Request *http.Request
	Writer  http.ResponseWriter
	status  int
	result  string
}

//定义方法
func (c *Context) Status(status int) {
	c.status = status
}

//定义返回
func (c *Context) response() {
	c.Writer.WriteHeader(c.status)
	i, e := io.WriteString(c.Writer, c.result)
	if e != nil {
		fmt.Println("出错了", i, e)
	}
}

//定义正常的api
func (c *Context) Normal(data interface{}, status ...int) {
	c.status = http.StatusOK
	if len(status) > 0 && status[0] > 0 {
		c.status = status[0]
	}
	switch t := data.(type) {
	case string:
		c.result = t
	default:
		b, e := json.Marshal(data)
		c.result = string(b)
		if e != nil {
			c.Error(e)
			return
		}
	}
}

//定义异常的api
func (c *Context) Error(err interface{}, status ...int) {
	s := ""
	switch err.(type) {
	case error:
		s = err.(error).Error()
	case string:
		s = err.(string)
	default:
		s = fmt.Sprintf("%v", err)
	}
	t := http.StatusInternalServerError
	if len(status) > 0 && status[0] > 0 {
		t = status[0]
	}
	c.status = t
	c.result = s
}

func (R *Router) HandleFunc(method, path string, handle ...Handle) {
	if method == "" {
		panic("Method can not be null!")
	}
	method = strings.ToUpper(method)

	if path == "" {
		panic("path can not be null!")
	}

	if _, ok := R.router[method][path]; ok {
		panic(fmt.Sprintf("panic: handlers are already registered for path '%s'", path))
	}

	if R.router == nil {
		R.router = make(map[string]map[string][]Handle)
	}

	if R.router[method] == nil {
		R.router[method] = make(map[string][]Handle)
	}
	R.router[method][path] = handle
}

func (R *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.UserAgent()
	log.Printf("[%s] %s \t IP:%s", r.Method, r.URL, r.RemoteAddr)
	method := strings.ToUpper(r.Method)
	if method == Options {
		return
	}
	// 防止panic
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("panic: %s\n", err)
			http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		}
	}()
	if f, ok := R.router[r.Method][r.URL.String()]; ok {
		c := &Context{
			Writer:  w,
			Request: r,
			status:  http.StatusOK,
		}
		for _, fc := range f {
			fc(c)
		}

		// 处理请求头，赋值给响应头
		for key, v := range r.Header {
			if key != "Content-Length" {
				w.Header().Set(key, strings.Join(v, ";"))
			}
		}
		w.Header().Set("Server", config.VersionName+"/"+config.Version)
		a := strings.Split(r.RemoteAddr, ":")
		if len(a) > 0 {
			w.Header().Set("Client-Ip", a[0])
		}
		c.response()
	} else {
		w.Header().Set("Server", config.VersionName+"/"+config.Version)
		http.NotFound(w, r)
	}
}

func NewRouter() *Router {
	return new(Router)
}

// 各种请求
// get 请求
func (R *Router) GET(path string, handle ...Handle) {
	R.HandleFunc(Get, path, handle...)

}

//post 请求
func (R *Router) POST(path string, handle ...Handle) {
	R.HandleFunc(Post, path, handle...)
}

//put请求
func (R *Router) PUT(path string, handle ...Handle) {
	R.HandleFunc(Put, path, handle...)
}

//delete请求
func (R *Router) DELETE(path string, handle ...Handle) {
	R.HandleFunc(Delete, path, handle...)
}

// Run 启动
func (R *Router) Run() {
	log.Printf("Start Server\tAddress:%s\tPORT:%d\t Version：%s\tServer name：%s\n", config.Host, config.Port, config.Version, config.VersionName)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%d", config.Host, config.Port), R))
}

func (R *Router) Group(s string) *Router {
	return R
}
