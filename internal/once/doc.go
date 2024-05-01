/*
Once is an object that will perform exactly one action.

A Once must not be copied after first use.

In the terminology of the Go memory model,
the return from f “synchronizes before”
the return from any call of once.Do(f).

The expeted signature of once.Do is `func (o *Once) Do(f func())`
*/
package once

import "sync/atomic"

// type Once struct{}

// func (o *Once) Do(f func()) {

// }

type Once struct {
	ctr atomic.Int32
}

func (o *Once) Do(f func()) {
	if o.ctr.Add(1) == 1 {
		f()
	}
}
