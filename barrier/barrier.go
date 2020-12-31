package barrier

import (
	"errors"
	"fmt"
	"sync/atomic"
)

// IBarrier ...
type IBarrier interface {
	Wait()
}

type barrier struct {
	n   int32
	m   int32
	chn chan struct{}
}

// New ...
func New(bucket int) (IBarrier, error) {
	if bucket < 1 {
		return nil, errors.New("argument should be greater than 1")
	}
	return &barrier{
		n:   int32(bucket),
		chn: make(chan struct{}),
	}, nil
}

// Arrive ...
func (b *barrier) Wait() {
	count := atomic.AddInt32(&b.m, 1)
	fmt.Printf("barrier: %d arrived\n", count)

	if count == b.n {
		fmt.Printf("barrier: release blocking\n")
		close(b.chn)
		return
	}

	<-b.chn
}
