package main

import (
	"net"
	"fmt"
)

var conn net.Conn

func connect() {
	host, _ := net.LookupHost(GetHostName())
	listen, err := net.Listen("tcp", host[0]+GetPort())
	if nil != err {
		fmt.Println(err)
		return
	}
	for ; ; {
		conn, err := listen.Accept()
		if nil != err {
			fmt.Println(err)
			continue
		}
		process(conn)
	}
}
func process(conn net.Conn) {
	defer conn.Close()
	for ; ; {
		buf := make([]byte, GetMaxBufferLen())
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("read err:", err)
			return
		}
		fmt.Printf(string(buf[0:n]))
	}
}

func main() {
	fmt.Println(JoinRoomData("71415"))

}
