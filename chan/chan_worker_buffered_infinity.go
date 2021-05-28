package main

import (
	"fmt"
	"time"
)

func main() {
	queue := 20

	jobs := make(chan int, queue)
	results := make(chan int, queue)

	go publisher(queue, jobs)

	for i := 0; i < 2; i++ {
		go worker(jobs, results)
	}

	for i := 0; i < queue; i++ {
		fmt.Println(<-results)
	}
}

func publisher(count int, jobs chan<- int) {
	for i := 0; i < count; i++ {
		jobs <- i
		time.Sleep(100 * time.Millisecond)
	}
	close(jobs)
}

func worker(jobs <-chan int, results chan int) {
	for j := range jobs {
		results <- fib(j)
	}
}

func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}
