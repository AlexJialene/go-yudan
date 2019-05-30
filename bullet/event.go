package bullet

import (
	"net"
	"fmt"
	"bytes"
	"encoding/binary"
	"strings"
	"io"
)

type Event interface {
	receive(f func(msg map[string]string))

	send(reqData []byte)
}

type EventImpl struct {
	tcpConn *net.TCPConn
}

func (e EventImpl) receive(f func(msg map[string]string)) {
	defer tcpConn.Close()

	for {
		recvByte := make([]byte, GetMaxBufferLen())
		recvLen, err := tcpConn.Read(recvByte)

		if err == io.EOF {
			fmt.Println("tcp已关闭")
		}

		buff := bytes.NewBuffer(recvByte)
		buff.Next(12)

		realByte := make([]byte, recvLen)
		binary.Read(buff, binary.LittleEndian, realByte)
		dataStr := string(realByte)

		if -1!=strings.Index(dataStr , "type@="){
			for 5 < strings.LastIndex(dataStr, "type@=") {
				mapMsg := substring(dataStr , strings.LastIndex(dataStr, "type@="))
				fmt.Println(mapMsg)
				fmt.Println(".....................")

				resMsg := message(mapMsg)
				f(resMsg)
				dataStr = dataStr[0: strings.LastIndex(dataStr, "type@=")]
			}

			fmt.Println("||||||:"+dataStr)
			resMsg := message(dataStr)

			//Callback
			f(resMsg)
		}else{
			fmt.Println("不符合的消息格式")
		}
	}
}

func (e EventImpl) send(reqData []byte) {
	_, err := tcpConn.Write(reqData)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func message(msg string) map[string]string {

	message := make(map[string]string)
	if msg[len(msg)-1:] == "/" {
		msg = msg[0 : len(msg)-1]
	}else{
		msg = msg[0 : len(msg)-14]
	}
	fmt.Println("||||||:"+msg)
	str := strings.Split(msg, "/")
	for _, v := range str {
		fmt.Println(v)
		key := v[ 0:strings.Index(v, "@=")]
		val := v[strings.Index(v, "@=")+len("@="):len(v)]
		if strings.Contains(val, "@A") {
			fmt.Println("***********************")
			val = strings.Replace(val, "@S", "/", -1)
			val = strings.Replace(val, "@A", "@", -1)

		}
		message[key] = val
	}
	return message
}

func substring(str string, start int) string {
	if str == "" {
		return str
	} else {
		if start < 0 {
			start += len(str)
		}
		if start < 0 {
			start = 0
		}
		if start > len(str) {
			return ""
		} else {
			return str[start:len(str)]
		}
	}

	return str

}
