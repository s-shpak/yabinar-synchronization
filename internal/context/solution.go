package context

import (
	"context"
	"sync"
	"time"
)

func WithCancel(ctx context.Context) (context.Context, func()) {
	cc := newCancelContext()
	if p, ok := ctx.(*cancelContext); ok {
		p.mux.Lock()
		p.children[cc] = struct{}{}
		p.mux.Unlock()
	}
	return cc, func() { cc.cancel() }
}

type cancelContext struct {
	err      error
	done     chan struct{}
	children map[*cancelContext]struct{}

	mux *sync.Mutex
}

func newCancelContext() *cancelContext {
	return &cancelContext{
		children: make(map[*cancelContext]struct{}),
		done:     make(chan struct{}),
		mux:      &sync.Mutex{},
	}
}

func (cc *cancelContext) Deadline() (time.Time, bool) {
	return time.Time{}, false
}

func (cc *cancelContext) Value(key any) any {
	panic("NYI")
}

func (cc *cancelContext) Done() <-chan struct{} {
	return cc.done
}

func (cc *cancelContext) Err() error {
	cc.mux.Lock()
	defer cc.mux.Unlock()
	return cc.err
}

func (cc *cancelContext) cancel() {
	cc.mux.Lock()
	defer cc.mux.Unlock()

	if cc.err != nil {
		return
	}
	cc.err = context.Canceled
	close(cc.done)
	for child := range cc.children {
		child.cancel()
	}
}
