package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main4() {
	// 绑定路由/hello
	http.HandleFunc("/hello", helloHandle)
	// 绑定路由到/test
	http.HandleFunc("/test", testHandle)
	http.HandleFunc("/baidu", baiduHandle)
	http.HandleFunc("/html", htmlHandle)
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
	//我们也可以设置响应头的信息
	w.Header().Set("Content-Type", "Handler5Json")
	user := User{
		ID:   10,
		Name: "Xiao Ming",
		Age:  23,
	}
	//将user序列化
	data, err := json.Marshal(user)
	if err != nil {
		fmt.Println("json.Marshal(user) err=", err)
	}
	//Marshal返回的就是一个[]byte
	w.Write(data)
}

func baiduHandle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Location", "https:www.baidu.com")
	w.WriteHeader(302)
}

// “中的内容会安照原格式的方式
func htmlHandle(w http.ResponseWriter, r *http.Request) {
	html := `<html>
		<head>
			<title>html页</title>
		</head>
		<body>
			html方式的响应
		</body>
		<html>
	`
	w.Write([]byte(html))
}

type User struct {
	ID   int
	Name string
	Age  int
}

// 访问: http://127.0.0.1:5000/test?a=1001
// 访问: http://127.0.0.1:5000/hello?b=990
