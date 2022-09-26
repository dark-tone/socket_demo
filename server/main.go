package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	//创建监听
	ip := "127.0.0.1"
	port := 600

	//.Sprintf函数往address变量中写入IP和端口信息
	address := fmt.Sprintf("%s:%d", ip, port)

	//开始监听 选择TCP协议 和 监听的方向
	listener, err := net.Listen("tcp", address)

	if err != nil {
		fmt.Println("net.Listen err:", err)
		return
	}

	//提取连接
	// Accept会阻塞
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener.Accept err:", err)
			return
		} else {
			handleConn(conn)
		}
	}
}

func handleConn(conn net.Conn) {
	fmt.Println("Accept conn addr:", conn.RemoteAddr())
	//创建容器用于接收读取到的数据
	buf := make([]byte, 1024) //使用make函数创建切片

	//cnt是实际读到的字符数
	cnt, err := conn.Read(buf)
	if err != nil {
		fmt.Println("conn.Read err:", err)
		return
	}

	fmt.Println("Client to Server...")
	fmt.Println("length:", cnt, "data", string(buf))

	//服务器业务实现 小写转成大写
	upperData := strings.ToUpper(string(buf[:cnt])) //只处理前cnt个字符
	cnt, err = conn.Write([]byte(upperData))

	fmt.Println("Server to Client...")
	fmt.Println("length:", cnt, "data", upperData)
}