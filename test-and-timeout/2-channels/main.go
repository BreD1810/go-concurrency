package main

import (
	"fmt"
	"time"
)

func main() {
	testPassed := make(chan struct{})

	killThread := make(chan struct{})

	go func(startTime time.Time) {
		for {
			select {
			case <-killThread:
				fmt.Println("received kill")
				return
			case <-time.After(time.Millisecond * 100):
				fmt.Println("waiting")
				// Test if the time now is more than 2 seconds after the goroutine is spawned
				if time.Now().After(startTime.Add(3 * time.Second)) {
					fmt.Println("goroutine reports done")
					close(testPassed) // Altnernatively send struct{}{}
					return
				}
			}
		}
	}(time.Now())

	select {
	case <-testPassed:
		fmt.Println("received positive test result")
	// case <-time.After(10 * time.Second):
	case <-time.After(2 * time.Second):
		fmt.Println("timed out")
		close(killThread) // Alternatively send struct{}{}
	}

	time.Sleep(5 * time.Second)
}
