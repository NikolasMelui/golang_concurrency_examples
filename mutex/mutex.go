package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Object struct {
	MapMutex *sync.Mutex
	Map      map[string]int
}

func writeToMap(object *Object) {
	object.MapMutex.Lock()
	// defer object.MapMutex.Unlock()

	randInt := rand.Int()
	fmt.Println("Trying to set value: ", randInt)
	object.Map["value"] = randInt
	time.Sleep(1 * time.Second)
}

func main() {
	object := Object{
		MapMutex: new(sync.Mutex),
		Map:      make(map[string]int, 0),
	}

	go func() {
		for i := 0; i < 10; i++ {
			go writeToMap(&object)
			time.Sleep(1 * time.Second)
		}
	}()

	ticker := time.NewTicker(1 * time.Second)
	for {
		select {
		case <-ticker.C:
			fmt.Println("Current value is: ", object.Map["value"])
		}
	}
}
