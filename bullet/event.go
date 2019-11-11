package bullet

import (
	"net"
	"strings"
	"log"
	"github.com/alexjialene/go-yudan/encoder"
)

type Event interface {
	receive(f func(msg map[string]string))

	send(reqData []byte)

	getRoom() string
}

type EventImpl struct {
	tcpConn *net.TCPConn
	roomId  string
}

func (e EventImpl) receive(f func(msg map[string]string)) {
	defer tcpConn.Close()
	for {
		header := make([]byte, 12)
		n, err := tcpConn.Read(header)
		if err != nil {
			log.Println("[ERROR]:", err)
		}
		b := make([]byte, encoder.ByteToInt(header[0:4])-8)
		n, err = tcpConn.Read(b)
		if err != nil {
			log.Println("[ERROR]:", err)

		}
		data := string(b[:n])
		f(message(data))
		if err != nil {
			log.Println("[ERROR]:", err)
		}
	}
}

func (e EventImpl) send(reqData []byte) {
	_, err := tcpConn.Write(reqData)
	if err != nil {
		return
	}
}

func (e EventImpl) getRoom() string {
	return e.roomId
}

func message(data string) map[string]string {
	message := make(map[string]string)

	end := 0x0

	if data[len(data)-1:] == string(end) {

		data = data[:len(data)-2]
		array := strings.Split(data, "/")

		if 1 <= len(array) {
			for _, v := range array {
				s := strings.Split(v, "@=")
				if len(s) <= 1 {
					continue
				}
				message[s[0]] = s[1]
			}
		}
	}
	return message
}
