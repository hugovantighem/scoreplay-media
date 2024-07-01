package common

import (
	"context"
	"fmt"
	"time"
)

func WithRetry(ctx context.Context,
	fn func(ctx context.Context) (any, error),
	doRetry func(ctx context.Context, result interface{}, err error) (bool, error),
	retries int,
	delay time.Duration,
) (interface{}, error) {
	more := true
	var err error
	var result interface{}
	for i := 0; i < retries && more == true; i++ {
		result, err = fn(ctx)
		more, err = doRetry(ctx, result, err)
		if more {
			fmt.Printf("attempt %d failed, retrying...\n", i+1)
		}
		<-time.After(delay)
	}
	return result, err
}
