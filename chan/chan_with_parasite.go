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

func parasite(value int, ch chan int) {
	ch <- value
}

func main() {
	prodOne := Producer{
		ch: make(chan int),
	}

	prodTwo := Producer{
		ch: make(chan int),
	}

	mainTicker := time.NewTicker(20 * time.Second)

	parasiteTicker := time.NewTicker(2 * time.Second)

	go prodOne.forever()
	go prodTwo.limit()

	for {
		select {
		case i := <-prodOne.ch:
			fmt.Println("ProducerOne receives: ", i)
		case i := <-prodTwo.ch:
			fmt.Println("ProducerTwo receives: ", i)
		case <-mainTicker.C:
			go prodTwo.limit()
		case <-parasiteTicker.C:
			fmt.Println("Parasite!!! Sends parasite data to prodTwo channel!!!")
			go parasite(1111111111, prodTwo.ch)
		}
	}
}
