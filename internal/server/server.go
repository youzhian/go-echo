package server

import (
	"log" //日志包
	"net" //网络库
)

type Server struct{
	addr string //服务器地址
	listener net.Listener //监听器
}
//构造器 创建一个server，避免外部访问我们的成员变量
func NewServer(addr string) *Server{
	return &Server{addr: addr}
}
//第一步，创建监听
func (s *Server) Listen() error{
	listener,err :=net.Listen("tcp",s.addr)
	if err != nil{
		return err
	}
	s.listener = listener
	log.Printf("server listen on %v",s.addr)

	return nil
}
//第二步，运行
func (s *Server) Run() {
	for{
		conn,err := s.listener.Accept()
		if err != nil{
			log.Printf("Accept with err:%v",err)
			continue
		}
		log.Printf("new connection form %v",conn.RemoteAddr())
		s.handleEcho(conn)
	}
}
//第三步，处理一个消息
func (s *Server) handleEcho(conn net.Conn){
	//创建一个协程，处理消息
	go func() {
		//定义一个缓冲区
		buf := make([]byte,1024)
		for{
			_,readErr := conn.Read(buf)
			if readErr != nil{
				log.Printf("read with error: %v",readErr)
				break
			}
			_,writeErr := conn.Write(buf)
			if writeErr != nil{
				log.Printf("write with error: %v",writeErr)
				break
			}
		}
	}()
}