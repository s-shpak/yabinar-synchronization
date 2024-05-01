package context

import (
	"context"
	"time"
)

func TODO_WithCancel(ctx context.Context) (context.Context, func()) {
	panic("Это надо реализовать")
}

type todo_cancelContext struct {
	// Можно менять эту структуру
}

func (cc *todo_cancelContext) TODO_Deadline() (time.Time, bool) {
	panic("Это реализовывать не надо")
}

func (cc *todo_cancelContext) Value(key any) any {
	panic("Это реализовывать не надо")
}

func (cc *todo_cancelContext) Done() <-chan struct{} {
	panic("Это надо реализовать")
}

func (cc *todo_cancelContext) Err() error {
	panic("Это можно реализовать")
}
