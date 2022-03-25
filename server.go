package main

import (
	"fmt"
	"net"
)

type Server struct {
	Ip   string
	Port int
}

// NewServer Create Server
func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip:   ip,
		Port: port,
	}

	return server
}

func (this *Server) Handler(conn net.Conn) {
	println("...当前链接的业务")
}

func (this *Server) Start() {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", this.Ip, this.Port))
	if err != nil {
		println("net.Listen err:", err)
		return
	}

	defer func(listener net.Listener) {
		err := listener.Close()
		if err != nil {
			println("listener.Close() err: ", err)
		}
	}(listener)

	for {
		conn, err := listener.Accept()
		if err != nil {
			println("listener accept err:", err)
			continue
		}

		go this.Handler(conn)
	}
}
