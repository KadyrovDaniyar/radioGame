package main_test

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestMain(t *testing.T) {
	count := 10
	limit := 3

	var wg sync.WaitGroup
	wg.Add(count)

	concurrentCount := 0
	failed := false

	var mock = func() {
		defer func() {
			wg.Done()
			concurrentCount--
		}()

		concurrentCount++
		if concurrentCount > limit {
			failed = true // test could be failed here without waiting all routines finish
		}

		time.Sleep(100)
	}

	spawn(mock, count, limit)

	wg.Wait()

	if failed {
		fmt.Println("Test failed")
	} else {
		fmt.Println("Test passed")
	}
}

func spawn(fn func(), count int, limit int) {
	limiter := make(chan bool, limit)

	spawned := func() {
		defer func() { <-limiter }()
		fn()
	}

	for i := 0; i < count; i++ {
		limiter <- true
		go spawned()
	}
}
