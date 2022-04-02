package main

import (
	"fmt"
	"math/rand"
	"time"
)

func hello(message string) {
	fmt.Println(message + "- go routine")
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(2000)))
}

func main() {
	go hello("hello 1")
	go hello("hello 2")
	go hello("hello 3")
	go hello("hello 4")
	go hello("hello 5")
	go hello("hello 6")
	go hello("hello 7")
	go hello("hello 8")
	go hello("hello 9")
	go hello("hello 10")

	time.Sleep(time.Second * 4)
	fmt.Println("chamada normal")

}
