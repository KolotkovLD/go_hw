package hw06pipelineexecutiongracefull

import (
	"context"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

const (
	sleepPerStage = time.Millisecond * 100
	fault         = sleepPerStage / 2
)

func TestPipeline(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	// Определим стадии пайплайна
	stages := []Stage{
		func(ctx context.Context, in In) Out {
			out := make(Out)
			go func() {
				defer close(out)
				for v := range in {
					select {
					case <-ctx.Done():
						return
					case out <- v:
					}
				}
			}()
			return out
		},
		func(ctx context.Context, in In) Out {
			out := make(Out)
			go func() {
				defer close(out)
				for v := range in {
					select {
					case <-ctx.Done():
						return
					case out <- v:
					}
				}
			}()
			return out
		},
	}

	in := make(In)

	// Запустим пайплайн
	result := make([]string, 0, 10)
	start := time.Now()
	for s := range ExecutePipeline(in, nil, stages...) {
		result = append(result, s.(string))
	}
	// Завершение работы

	gracefulShutdown(ctx, wg, 30*time.Second)

	elapsed := time.Since(start)
	require.Less(t,
		int64(elapsed),
		// ~0.8s for processing 5 values in 4 stages (100ms every) concurrently
		int64(sleepPerStage)*int64(len(stages)+len(data)-1)+int64(fault))

}
