package main

import (
	"time"
)

type KeepAlive interface {
	keepAlive()
}

type KeepAliveImpl struct {
	event *Event
}

func (k KeepAliveImpl) keepAlive()  {
	for{

		//心跳包
		time.Sleep(time.Second*40)
	}


}