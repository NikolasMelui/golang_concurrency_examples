package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {

	wg := new(sync.WaitGroup)
	for i := 0; i < 10; i++ {
		go printer(rand.Int(), wg)
		time.Sleep(1 * time.Second)
	}
	wg.Wait()
}

func printer(number int, wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()

	fmt.Println(number)
}
