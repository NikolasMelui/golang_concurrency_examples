package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	once := new(sync.Once)
	for i := 0; i < 10; i++ {
		once.Do(printerOnce)
		printer(rand.Int())
		time.Sleep(1 * time.Second)
	}
}

func printer(number int) {
	fmt.Println(number)
}

func printerOnce() {
	fmt.Println("Will print once!")
}
