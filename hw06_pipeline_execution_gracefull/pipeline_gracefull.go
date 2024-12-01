package hw06pipelineexecutiongracefull

import (
	"context"
	"log"
	"sync"
	"time"
)

type (
	In  = chan interface{}
	Out = chan interface{}
)

// Stage описывает функцию для выполнения задачи на определенном этапе пайплайна.
type Stage func(ctx context.Context, in In) Out

// ExecutePipeline запускает пайплайн, состоящий из нескольких этапов.
func ExecutePipeline(ctx context.Context, in In, stages ...Stage) Out {
	out := in
	for _, stage := range stages {
		log.Printf("Starting stage %v\n", stage)
		out = stage(ctx, out)
	}
	return out
}

func gracefulShutdown(ctx context.Context, wg *sync.WaitGroup, timeout time.Duration) {
	for {
		select {
		case <-ctx.Done():
		case <-time.After(timeout):
		}
	}
	wg.Wait()
}

// ----------------------------------------------------------------------------------------------------------

//package main
//
//import (
//	"context"
//	"log"
//	"sync"
//	"time"
//)
//
//type (
//	In         = chan any
//	InClosable = chan any
//	Out        = In
//)
//
//type Stage interface {
//	// Do MUST run tasks inside one or multiple goroutines
//	// When In is closed, Do MUST close Out
//	Do(ctx context.Context, in In) Out
//}
//
//type StageFunc func(ctx context.Context, in In) Out
//
//func (f StageFunc) Do(ctx context.Context, in In) Out {
//	return f(ctx, in)
//}
//
//func foo() {
//	ExecutePipeline(nil, nil, StageFunc(func(ctx context.Context, in In) Out {
//
//	}))
//
//	//	http.Handle("GET /api/v1/users", )
//	//	http.ListenAndServe("localhost:8000", nil)
//}
//
//func ExecutePipeline(ctx context.Context, initial chan any, stages ...Stage) Out {
//	in := In(initial)
//	for _, stage := range stages {
//		in = stage(ctx, in)
//	}
//
//	return in
//}
//
//func listenAndServe(ctx context.Context, wg *sync.WaitGroup, requests chan any) {
//	for request := range requests {
//		wg.Add(1)
//		go handleRequest(ctx, wg, request)
//	}
//}
//
//func gracefulShutdown(ctx context.Context, wg *sync.WaitGroup, timeout time.Duration, requests chan any) {
//	close(requests)
//	//timeout := 30 * time.Second
//	for {
//		select {
//		case <-ctx.Done():
//		case <-time.After(timeout):
//		}
//	}
//	wg.Wait()
//}
//
//func handleRequest(ctx context.Context, wg *sync.WaitGroup, request any) {
//	defer wg.Done()
//	// Обработка запроса
//	log.Printf("Handling request: %+v\n", request)
//}

// ----------------------------------------------------------------------------------------------------------

//func main() {
//	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
//	defer cancel()
//
//	requests := make(chan any)
//
//	// Определим стадии пайплайна
//	stages := []Stage{
//		func(ctx context.Context, in In) Out {
//			out := make(Out)
//			go func() {
//				defer close(out)
//				for v := range in {
//					select {
//					case <-ctx.Done():
//						return
//					case out <- v:
//					}
//				}
//			}()
//			return out
//		},
//		func(ctx context.Context, in In) Out {
//			out := make(Out)
//			go func() {
//				defer close(out)
//				for v := range in {
//					select {
//					case <-ctx.Done():
//						return
//					case out <- v:
//					}
//				}
//			}()
//			return out
//		},
//	}
//
//	// Запустим пайплайн
//	pipeline := ExecutePipeline(ctx, requests, stages...)
//
//	// Слушатель запросов
//	wg := &sync.WaitGroup{}
//	go listenAndServe(ctx, wg, pipeline)
//
//	// Имитация отправки запросов
//	go func() {
//		for i := 0; i < 10; i++ {
//			requests <- i
//		}
//		close(requests)
//	}()
//
//	// Завершение работы
//	gracefulShutdown(ctx, wg, 30*time.Second, requests)
//}
