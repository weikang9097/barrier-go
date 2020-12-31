package main

import (
	"fmt"
	"sync"

	"gitlab.moebius.com/trainning-room/barrier"
)

func main() {
	n := 5
	b, _ := barrier.New(n)

	var wg sync.WaitGroup
	wg.Add(int(n))

	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()

			b.Wait()
			fmt.Printf("goroutine went across the barrier\n")
		}()
	}

	wg.Wait()
}
