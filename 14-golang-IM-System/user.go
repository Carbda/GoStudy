package main

import "net"

type User struct {
	Name string
	Addr string
	C    chan string
	Conn net.Conn
}

// 创建一个用户的API
func NewUser(conn net.Conn) *User {
	userAddr := conn.RemoteAddr().String()

	user := &User{
		Name: userAddr,
		Addr: userAddr,
		C:    make(chan string),
		Conn: conn,
	}

	//启动监听当前user channel消息的goroutine
	go user.ListenMessage()

	return user
}

// 监听当前User channel的方法，一旦有消息就发送给对端客户端
func (u *User) ListenMessage() {
	for {
		msg := <-u.C

		u.Conn.Write([]byte(msg + "\n"))
	}
}
