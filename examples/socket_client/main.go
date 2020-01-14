/**
  create by yy on 2020/1/6
*/

package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
)

// socket 客户端
func main() {
	var (
		err      error
		addr     *net.TCPAddr
		conn     *net.TCPConn
		response []byte
		input    string
	)

	if len(os.Args) < 2 {
		_, _ = fmt.Fprintf(os.Stderr, "Usage: %s host:port", os.Args[0])
	}

	// 获取命令行参数 socket地址
	server := os.Args[1]
	if addr, err = net.ResolveTCPAddr("tcp4", server); err != nil {
		panic(err)
	}

	// defer func() {
	// 	if err = conn.Close(); err != nil {
	// 		fmt.Println(err)
	// 	}
	// }()

	for {
		if _, err = fmt.Scanln(&input); err != nil {
			fmt.Println(err)
			continue
		}

		// 建立tcp 链接
		if conn, err = net.DialTCP("tcp4", nil, addr); err != nil {
			fmt.Println(err)
		}

		// 向服务端发送数据
		if _, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n")); err != nil {
			fmt.Println(err)
		}

		// 接收服务端的响应
		if response, err = ioutil.ReadAll(conn); err != nil {
			fmt.Println(err)
		}

		fmt.Println(string(response))

		_ = conn.Close()
	}
}
