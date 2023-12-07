package main

import (
	"fmt"
	"io"
	"net"
	"sync"
	"time"
)

type Server struct {
	Ip   string
	Port int

	//在线用户的列表
	OnlineMap map[string]*User
	mapLock   sync.RWMutex

	//消息广播的channel
	Message chan string
}

// 创建server接口
func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip:        ip,
		Port:      port,
		OnlineMap: make(map[string]*User),
		Message:   make(chan string),
	}
	return server
}

// 监听Message广播消息channel的goroutine，一旦有消息就发送给全部的在线User
func (s *Server) ListenMessager() {
	for {
		msg := <-s.Message

		//将msg发送给全部User
		s.mapLock.Lock()
		for _, cli := range s.OnlineMap {
			cli.C <- msg
		}
		s.mapLock.Unlock()
	}
}

// 广播消息的方法
func (s *Server) BroadCast(user *User, msg string) {
	sendMsg := "[" + user.Addr + "]" + user.Name + ":" + msg

	s.Message <- sendMsg
}

func (s *Server) Handler(conn net.Conn) {
	//当前连接的业务
	fmt.Println("连接建立成功")

	user := NewUser(conn, s)

	user.Online()

	//监听用户是否活跃的channel
	isLive := make(chan bool)

	//接收客户端传递发送的消息
	go func() {
		buf := make([]byte, 4096) //声明byte数组
		for {
			n, err := conn.Read(buf) //把conn的数据读到buf中
			if n == 0 {
				user.Offline()
				return
			}
			if err != nil && err != io.EOF {
				fmt.Println("conn read err:", err)
			}

			//提取用户的消息（去除'\n'）
			msg := string(buf[:n-1])

			//用户针对msg进行消息处理
			user.DoMessage(msg)

			//用户的任意消息，代表当前用户是活跃的
			isLive <- true
		}
	}()

	//当前handler阻塞
	for {
		select {
		case <-isLive:
			//不做任何事情，为了激活select，更新下面定时器
		case <-time.After(10 * time.Second):
			//已经超时
			//将当前User强制关闭
			user.SendMsg("you are kicked!\n")
			//销毁用的资源
			close(user.C)
			//关闭连接
			conn.Close()
			//退出当前Handler
			return
		}
	}
}

// 启动服务器接口
func (s *Server) Start() {
	//socket listen
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", s.Ip, s.Port))
	if err != nil {
		fmt.Println("net.Listen err:", err)
		return
	}
	//close listen socket
	defer listener.Close()

	//启动监听Message的goroutine
	go s.ListenMessager()

	for {
		//accept
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener accept err:", err)
			continue
		}
		//do handler
		go s.Handler(conn)
	}
}
