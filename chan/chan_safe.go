package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Producer struct {
	ch chan int
}

func (prod *Producer) forever() {
	for {
		prod.ch <- rand.Int()
		time.Sleep(3 * time.Second)
	}
}

func (prod *Producer) limit() {
	for i := 0; i < 10; i++ {
		prod.ch <- rand.Int()
		time.Sleep(1 * time.Second)
	}
}

func (prod *Producer) getReadChan() <-chan int {
	return prod.ch
}

func main() {
	prodOne := Producer{
		ch: make(chan int),
	}

	prodTwo := Producer{
		ch: make(chan int),
	}

	mainTicker := time.NewTicker(20 * time.Second)

	go prodOne.forever()
	go prodTwo.limit()

	for {
		select {
		case i := <-prodOne.getReadChan():
			fmt.Println("ProducerOne gets: ", i)
		case i := <-prodTwo.getReadChan():
			fmt.Println("ProducerTwo receives: ", i)
		case <-mainTicker.C:
			go prodTwo.limit()
		}
	}
}
