/*
Once is an object that will perform exactly one action.

A Once must not be copied after first use.

In the terminology of the Go memory model,
the return from f “synchronizes before”
the return from any call of once.Do(f).

The expeted signature of once.Do is `func (o *Once) Do(f func())`
*/
package once

// type Once struct{}

// func (o *Once) Do(f func()) {

// }
