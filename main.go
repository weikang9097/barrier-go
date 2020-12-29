package main

import (
	"fmt"
	"sync"

	"gitlab.moebius.com/trainning-room/barrier"
)

func main() {
	bucket := 10
	b := barrier.New(bucket)

	var wg sync.WaitGroup
	wg.Add(bucket)





	for i := 0; i < bucket; i++ {

		go func() {
			left := b.Obtain()
			fmt.Printf("barrier left %d \n", left)
			wg.Done()
		}()
	}

	wg.Wait()
	b.Obtain()

}
