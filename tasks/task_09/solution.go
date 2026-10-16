package main

import (
	"context"
	"errors"
	"sync"
)

type Result[R any] struct {
	Value R
	Err   error
}

var ErrInvalidWorkers = errors.New("workers must be positive")

func ParallelMap[T any, R any](ctx context.Context, workers int, in []T, fn func(context.Context, T) (R, error)) ([]Result[R], error) {
	if workers <= 0 {
		return nil, ErrInvalidWorkers
	}
	if ctx.Err() != nil {
		return nil, ctx.Err()
	}

	if len(in) == 0 {
		return []Result[R]{}, nil
	}

	arrResult := make([]Result[R], len(in))

	for i := 0; i < len(in); {
		var wg sync.WaitGroup
		select {
		case <-ctx.Done():
			wg.Wait()
			return nil, ctx.Err()
		default:
			for g := 0; g < min(workers, len(in)-i); g++ {
				wg.Go(func() {
					val, err := fn(ctx, in[i+g])
					if err != nil {
						arrResult[i+g].Err = err
						return
					}
					arrResult[i+g].Value = val
				})
			}
			wg.Wait()
			i += min(workers, len(in)-i)
		}
	}
	return arrResult, nil
}
