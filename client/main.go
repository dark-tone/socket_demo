package main

import (
	"fmt"
	"net"
)

func main()  {
	conn, err := net.Dial("tcp", ":600")
	if err != nil {
		fmt.Println("error occur:", err)
		return
	}

	fmt.Println("Dial success")

	// 构建要发送的数据
	sendData := []byte("hello world")

	// write写socket
	cnt, err := conn.Write(sendData)

	if err != nil {
		fmt.Println("conn.Write err:", err)
		return
	}

	//接收服务器返回的数据
	buf := make([]byte, 1024)

	cnt, err = conn.Read(buf)
	if err != nil {
		fmt.Println("conn.Read err:", err)
		return
	}

	fmt.Println("Server to Client...")
	fmt.Println("length:", cnt, "data", string(buf[:cnt])) //只处理前cnt个字符

	//关闭连接
	conn.Close()
}