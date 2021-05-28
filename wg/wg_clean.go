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
		go func() {
			wg.Add(1)
			defer wg.Done()

			printer(rand.Int())
		}()
		time.Sleep(1 * time.Second)
	}
	wg.Wait()
}

func printer(number int) {
	fmt.Println(number)
}
