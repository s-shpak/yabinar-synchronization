package context

import (
	"context"
	"time"
)

type backgroundCtx struct{}

func (backgroundCtx) Deadline() (_ time.Time, _ bool) {
	return
}

func (backgroundCtx) Done() <-chan struct{} {
	return nil
}

func (backgroundCtx) Err() error {
	return nil
}

func (backgroundCtx) Value(key any) any {
	return nil
}

func Background() context.Context {
	return backgroundCtx{}
}
