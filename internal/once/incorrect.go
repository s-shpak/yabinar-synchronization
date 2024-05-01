package once

import "sync/atomic"

type IncorrectOnce struct {
	done atomic.Uint32
}

func (o *IncorrectOnce) Do(f func()) {
	if o.done.CompareAndSwap(0, 1) {
		f()
	}
}
