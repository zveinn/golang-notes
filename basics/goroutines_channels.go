package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("Worker %d started job %d\n", id, j)
		time.Sleep(time.Second)
		results <- j * 2
	}
}

func goroutinesAndChannels() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	for a := 1; a <= 5; a++ {
		result := <-results
		fmt.Printf("Result received: %d\n", result)
	}

	done := make(chan bool)

	go func() {
		fmt.Println("Goroutine doing some work...")
		time.Sleep(3 * time.Second)
		done <- true
	}()

	<-done
	fmt.Println("Goroutine has finished.")

	result := make(chan bool, 1)
	go func() {
		time.Sleep(1 * time.Second)
		result <- true
	}()

	select {
	case <-result:
		fmt.Println("Success!")
	case <-time.After(1500 * time.Millisecond):
		fmt.Println("Timeout!")
	}

	jobChannel := make(chan string, 1)
	wg := sync.WaitGroup{}

	go func() {
		for job := range jobChannel {
			fmt.Println(job)
			wg.Done()
		}
	}()

	for i := 0; i < 10; i++ {
		wg.Add(1)
		jobChannel <- "job number " + strconv.Itoa(i)
	}

	wg.Wait()
}
