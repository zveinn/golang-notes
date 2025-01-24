package main

import (
	"fmt"
	"sync"
)

var (
	counter int
	wg      sync.WaitGroup
	once    sync.Once
)

// SafeCounter uses mutex to safely increment counter
type SafeCounter struct {
	mu sync.Mutex
}

// Increment safely increments the counter
func (c *SafeCounter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	counter++
}

// GetCounter returns the current counter value
func (c *SafeCounter) GetCounter() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return counter
}

func syncOperations() {
	safeCounter := SafeCounter{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				safeCounter.Increment()
			}
		}()
	}

	wg.Wait()
	fmt.Printf("Counter after first loop: %d\n", safeCounter.GetCounter())

	counter = 0

	var runOnce func()

	runOnce = func() {
		once.Do(func() {
			fmt.Println("This will only run once no matter how many goroutines call it.")
		})
	}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			runOnce()
			safeCounter.Increment()
		}()
	}

	wg.Wait()
	fmt.Printf("Counter after sync.Once example: %d\n", safeCounter.GetCounter())

	fmt.Println("After delay, attempting to run once again.")
	runOnce()

	fmt.Printf("Final counter value: %d\n", safeCounter.GetCounter())
}
