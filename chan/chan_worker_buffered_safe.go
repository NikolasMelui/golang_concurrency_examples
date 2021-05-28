package main

import (
	"fmt"
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
		jobs <- i
		fmt.Println(<-results)
	}
}

func publisher(num int, jobs chan<- int) {
	for i := 0; i < num; i++ {
		jobs <- i
	}
}

func worker(jobs <-chan int, results chan<- int) {
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
