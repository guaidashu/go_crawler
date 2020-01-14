/**
  create by yy on 2020/1/6
*/

package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"time"
)

// socket 服务端
func main() {
	var (
		err       error
		tcpServer *net.TCPAddr
		listen    *net.TCPListener
		conn      net.Conn
	)

	if tcpServer, err = net.ResolveTCPAddr("tcp4", ":8088"); err != nil {
		panic(err)
	}

	if listen, err = net.ListenTCP("tcp", tcpServer); err != nil {
		panic(err)
	}

	for {
		if conn, err = listen.Accept(); err != nil {
			fmt.Println(err)
			continue
		}

		go handle(conn)
	}

}

func handle(conn net.Conn) {
	var (
		response []byte
		err      error
	)

	defer func() {
		if err = conn.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	go func() {
		if response, err = ioutil.ReadAll(conn); err != nil {
			fmt.Println("err: ", err)
		} else {
			fmt.Println(string(response))
		}
	}()

	// 向客户端发送消息
	time.Sleep(time.Second)
	now := time.Now().String()
	if _, err = conn.Write([]byte(now)); err != nil {
		fmt.Println(err)
	}

}
