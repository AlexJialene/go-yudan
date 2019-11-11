package main

import (
	"github.com/alexjialene/go-yudan/bullet"
	"fmt"
)

func main() {
	bullet.Start("24422", "-9999", func(msg map[string]string) {

		fmt.Println("=========================================")
		fmt.Println(msg["type"])
		//fmt.Println(msg[""])
		//TODO
	})




}
