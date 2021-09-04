package ctxutil

import (
	"context"
	"errors"
)

var ErrCtxCancelled = errors.New("context cancelled")

func IsContextCancelled(ctx context.Context) bool {
	select {
	case <-ctx.Done():
		return true

	default:
		return false
	}
}
