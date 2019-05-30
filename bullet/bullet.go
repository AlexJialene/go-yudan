package bullet

import (
	"fmt"
	"net"
)

var tcpConn *net.TCPConn
var e Event
var k KeepAlive

func connect() {
	host, _ := net.LookupHost(GetHostName())
	tcpAddr, err := net.ResolveTCPAddr("tcp4", host[0]+":"+GetPort())
	if err != nil {
		fmt.Println(err)
		return
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		fmt.Println(err)
		return
	}

	tcpConn = conn
	e = EventImpl{tcpConn: conn}
	k = KeepAliveImpl{event: &e}
}

func join(conn *net.TCPConn, roomId string) {
	_, err := conn.Write(JoinRoomData(roomId))
	if err != nil {
		fmt.Println(err)
		return
	}
	recvByte := make([]byte, GetMaxBufferLen())
	conn.Read(recvByte)

	b := ParseLoginResult(recvByte)
	if b {
		fmt.Println("join room success")
	}

}

func Start(roomId, groupId string, f func(msg map[string]string)) {
	connect()
	join(tcpConn, roomId)
	joinGroup(roomId, groupId)

	c := make(chan int)

	go e.receive(f)
	go k.keepAlive()
	<-c

}

func joinGroup(roomId, groupId string) {
	e.send(joinGroupData(roomId, groupId))
	fmt.Println("joinGroup")
}
