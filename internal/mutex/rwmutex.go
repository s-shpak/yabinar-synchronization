package mutex

import (
	"math/rand"
	"sync"
)

type MuxCounter struct {
	val int
	mux *sync.Mutex
}

func NewMuxCounter() *MuxCounter {
	return &MuxCounter{
		mux: &sync.Mutex{},
	}
}

func (mc *MuxCounter) FrequentReadsFrequentWrites(c chan<- struct{}) {
	go func() {
		if rand.Intn(2) == 1 {
			mc.mux.Lock()
			defer mc.mux.Unlock()

			mc.val = rand.Intn(100)
			return
		}
		mc.mux.Lock()
		defer mc.mux.Unlock()

		if mc.val == 42 && c != nil {
			c <- struct{}{}
		}
	}()
}

func (mc *MuxCounter) FrequentReadsSpareWrites(c chan<- struct{}) {
	go func() {
		if rand.Intn(1000) < 5 {
			mc.mux.Lock()
			defer mc.mux.Unlock()

			mc.val = rand.Intn(100)
			return
		}
		mc.mux.Lock()
		defer mc.mux.Unlock()

		if mc.val == 42 && c != nil {
			c <- struct{}{}
		}
	}()
}

type RWMuxCounter struct {
	val int
	mux *sync.RWMutex
}

func NewRWMuxCounter() *RWMuxCounter {
	return &RWMuxCounter{
		mux: &sync.RWMutex{},
	}
}

func (mc *RWMuxCounter) FrequentReadsFrequentWrites(c chan<- struct{}) {
	go func() {
		if rand.Intn(2) == 1 {
			mc.mux.Lock()
			defer mc.mux.Unlock()

			mc.val = rand.Intn(100)
			return
		}
		mc.mux.RLock()
		defer mc.mux.RUnlock()

		if mc.val == 42 && c != nil {
			c <- struct{}{}
		}
	}()
}

func (mc *RWMuxCounter) FrequentReadsSpareWrites(c chan<- struct{}) {
	go func() {
		if rand.Intn(1000) < 5 {
			mc.mux.Lock()
			defer mc.mux.Unlock()

			mc.val = rand.Intn(100)
			return
		}
		mc.mux.RLock()
		defer mc.mux.RUnlock()

		if mc.val == 42 && c != nil {
			c <- struct{}{}
		}
	}()
}
