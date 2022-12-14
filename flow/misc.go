package flow

import (
	"context"
	"errors"
)

var ErrAbort = errors.New("stream aborted")

func Take[T any](ctx context.Context, cnt int, src <-chan T) <-chan T {
	dest := make(chan T)
	go func() {
		defer close(dest)
		for i := 0; i < cnt; i++ {
			select {
			case <-ctx.Done():
				return
			case dest <- <-src:
			}
		}
	}()
	return dest
}
