package main

import (
	"github.com/AlexJialene/go-yudan/bullet"
	"fmt"
)

func main() {
	bullet.Start("24422", "-9999", func(msg map[string]string) {

		fmt.Println(msg["type"])
		//TODO

	})




}
