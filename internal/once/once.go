package once

import "sync"

type Once struct {
	done bool
	mux  sync.Mutex
}

func (o *Once) Do(f func()) {
	o.mux.Lock()
	defer o.mux.Unlock()

	if !o.done {
		defer func() {
			o.done = true
		}()

		f()
	}
}
