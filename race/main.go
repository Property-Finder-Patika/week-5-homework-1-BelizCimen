package main

import (
	"fmt"
	"time"
)

var intVar int = 0
var tempVal int = 0

func Reader() {

	for {
		var val int = intVar
		if val%10 == 0 {
			tempVal += 1
		}
	}
}
func Writer() {
	for {
		intVar += 1
	}
}

func main() {
	go Reader()
	go Writer()
	fmt.Println("intVar = ", intVar)
	time.Sleep(3 * time.Second)
	fmt.Println("intVar = ", intVar)
}
