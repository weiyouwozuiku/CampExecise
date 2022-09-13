package chap2

import (
	"fmt"
	"sync"
)

func ManyWait() {
	var wg sync.WaitGroup
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func(j int) {
			defer wg.Done()
			hello(j)
		}(i)
	}
	wg.Wait()
}

func hello(j int) {
	fmt.Print("Hello from goroutine ", j, "\n")
}
