package main

import (
	"./encoder"
	"bytes"
	"encoding/binary"
	"strings"
)

var (
	hostName      = "openbarrage.douyutv.com"
	port          = "8601"
	maxBufferLen  = 4096
	messageClient = 689
)

const (
	roomId  = "74751"
	groupId = "-9999"
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

func joinGroupData(roomId, groupId string) []byte {
	e := encoder.GetEncoder()
	encoder.AddItem(e, "type", "joingroup")
	encoder.AddItem(e, "rid", roomId)
	encoder.AddItem(e, "gid", groupId)
	return encoder.ByteResult(e , GetMessageClient())
}

func ParseLoginResult(b []byte) bool {
	buff := bytes.NewBuffer(b)
	b1 := make([]byte, len(b)-12)
	buff.Next(12)
	binary.Read(buff, binary.LittleEndian, b1)

	str := string(b1)
	if strings.Contains(str, "type@=loginres") {
		return true
	}
	return false

}
