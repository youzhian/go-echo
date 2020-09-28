package client

import (
	"fmt"
	"log"
	"net"
)

type Client struct{
	serverAddr string //服务器地址
	conn net.Conn //连接信息
}
//构造器
func NewClient(serverAddr string) *Client{
	return &Client{serverAddr: serverAddr}
}
//第一步，连接
func(c *Client) Connect() error{

	conn,err := net.Dial("tcp",c.serverAddr)
	if err != nil{
		return err
	}
	c.conn = conn
	log.Printf("server connection success!")
	return nil
}
//第二步，运行
func (c *Client) Run(){
	var message string
	buf := make([]byte,1024)
	for{
		//引用传递
		_,inputErr := fmt.Scanln(&message)
		if inputErr != nil{
			log.Printf("input with error %v:",inputErr)
			break
		}
		//向服务器写消息
		_,writeErr := c.conn.Write([]byte(message))
		if writeErr != nil{
			log.Printf("write with error: %v",writeErr)
			break
		}
		//读取服务器返回的消息
		_,readErr := c.conn.Read(buf)
		if readErr != nil{
			log.Printf("read with error :%v",readErr)
			break
		}

		log.Printf("recv %s\n",buf)
	}
}