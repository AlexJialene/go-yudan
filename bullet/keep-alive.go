package bullet

import (
	"time"
	"fmt"
	"go-yudan/encoder"
)

type KeepAlive interface {
	keepAlive()
}

type KeepAliveImpl struct {
	event Event
}

func (k KeepAliveImpl) keepAlive() {
	for {
		//心跳包

		fmt.Println("Keep-Alvie send ====>")

		e := encoder.GetEncoder()
		encoder.AddItem(e, "type", "keeplive")
		encoder.AddItem(e, "tick", string(time.Now().Unix()))
		k.event.send(encoder.ByteResult(e, GetMessageClient()))

		time.Sleep(time.Second * 40)
	}

}
