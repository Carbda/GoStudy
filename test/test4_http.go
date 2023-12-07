package main

import (
	"fmt"
	"net/http"
)

func main() {
	// 绑定路由/hello
	http.HandleFunc("/hello", helloHandle)
	// 绑定路由到/test
	http.HandleFunc("/test", testHandle)
	// 启动服务
	err := http.ListenAndServe(":5000", nil)
	fmt.Println(err)
}

// 处理路由hello
func helloHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("访问路由hello")
	// 解析url参数
	fmt.Println(r.URL.Query())
	w.Write([]byte("hello word!"))
}

// 处理路由test
func testHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("访问路由test")
	// 解析url参数,并输出
	fmt.Println(r.URL.Query())
	//mapQuery := r.URL.Query()
	w.Write([]byte("test doing!"))
}

// 访问: http://127.0.0.1:5000/test?a=1001
// 访问: http://127.0.0.1:5000/hello?b=990
