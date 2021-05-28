package main

import (
	"fmt"
	"time"
)

func main() {
	jobs := make(chan int)
	results := make(chan int)

	go publisher(20, jobs)
	go worker(jobs, results)

	for r := range results {
		fmt.Println(r)
	}
}

func publisher(num int, jobs chan<- int) {
	for i := 0; i < 20; i++ {
		jobs <- i
		time.Sleep(100 * time.Millisecond)
	}
	close(jobs)
}

func worker(jobs <-chan int, results chan<- int) {
	for j := range jobs {
		results <- fib(j)
	}
	close(results)
}

func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}
