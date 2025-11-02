package main

import (
	"fmt"
	"log"
	"time"
)

func main() {

	chans := []chan int{
		make(chan int),
		make(chan int),
	}

	for i := range chans {
		log.Printf("i:%d", i)

		go func(i int, ch chan<- int) {
			for {
				log.Printf("sending data to channel:%d", i)
				time.Sleep(time.Duration(i) * time.Second)
				ch <- i
			}
		}(i+1, chans[i])
	}

	for i := 0; i < 10; i++ {
		select {
		case m0 := <-chans[0]:
			fmt.Println("received:", m0)
		case m1 := <-chans[1]:
			fmt.Println("received:", m1)
		}
	}
}
