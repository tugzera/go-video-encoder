package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	workersCount := 5

	channelInput := make(chan int)

	done := make(chan []byte)

	go func() {
		i := 0
		for {
			channelInput <- i
			i++
		}
	}()

	for x := 0; x < workersCount; x++ {
		go processWorker(channelInput, x)
	}
	<-done
}

func processWorker(channelInput chan int, worker int) {
	for value := range channelInput {
		sleepTime := time.Duration(rand.Intn(1000) * int(time.Millisecond))
		time.Sleep(sleepTime)
		fmt.Println("Worker ", worker, ": ", value)
	}
}
