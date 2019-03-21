package main

import (
	"./encoder"
)

var (
	hostName      = "openbarrage.douyutv.com"
	port          = "8601"
	maxBufferLen  = 4096
	messageClient = 689
)

func GetHostName() string {
	return hostName
}

func GetPort() string {
	return port
}

func GetMaxBufferLen() int {
	return maxBufferLen
}

func GetMessageClient() int {
	return messageClient
}

func JoinRoomData(roomId string) [] byte {
	e := encoder.GetEncoder()
	encoder.AddItem(e, "type", "loginreq")
	encoder.AddItem(e, "roomid", roomId)
	return encoder.ByteResult(e, GetMessageClient())
}
