package barrier

type barrier struct {
	N   int
	chn chan int
}

// New ...
func New(bucket int) barrier {
	return barrier{
		N:   bucket,
		chn: make(chan int, bucket),
	}
}

func (b *barrier) Obtain() int {

	b.chn <- 0

	return len(b.chn)
}

func (b *barrier) Release() {
	<-b.chn
}
