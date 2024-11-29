package hw06pipelineexecutionpoolofworkers

import (
	"context"
	"sync"
	"testing"
	"time"
)

func TestPipeline(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	requests := make(chan any)

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

	// Запустим пайплайн
	pipeline := ExecutePipeline(ctx, requests, stages...)

	// Слушатель запросов
	wg := &sync.WaitGroup{}
	go listenAndServe(ctx, wg, pipeline)

	// Имитация отправки запросов
	go func() {
		for i := 0; i < 10; i++ {
			requests <- i
		}
		close(requests)
	}()

	// Завершение работы

	gracefullShutdown(ctx, wg, 30*time.Second)
}
