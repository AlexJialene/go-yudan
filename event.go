package main

import (
	"net"
	"fmt"
	"bytes"
	"encoding/binary"
)

type Event interface {
	receive()

	send(reqData []byte)
}

type EventImpl struct {
	tcpConn *net.TCPConn
}

func (e EventImpl) receive() {

	defer  tcpConn.Close()
	for {
		recvByte := make([]byte , GetMaxBufferLen())
		recvLen ,_ := tcpConn.Read(recvByte)

		buff:=bytes.NewBuffer(recvByte)
		buff.Next(12)

		realByte := make([]byte , recvLen)
		binary.Read(buff, binary.LittleEndian, realByte)

		dataStr := string(realByte)
		fmt.Println(dataStr)
		//todo 可在这设置回调func
	}
}

func (e EventImpl) send(reqData []byte) {
	_, err := tcpConn.Write(reqData)
	if err != nil {
		fmt.Println(err)
		return
	}
}
