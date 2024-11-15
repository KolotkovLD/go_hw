package hw06pipelineexecution

import (
	"log"
)

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{} // нужно заменить на context
)

// Stage является исполнителем задач на шаге
// слушает канал задач и кладет результат в канал результатов.
type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	out := make(Bi)

	// инициализация стейджей
	for _, stage := range stages {
		in = stage(in)
		log.Println("Stage is started")
	}

	go func() {
		for {
			select {
			case <-done:
				log.Println("is done")
				close(out)
				return
			case val, ok := <-in:
				log.Println("Read input value")
				if !ok {
					log.Println("In chanel is empty")
					close(out)
					return
				}
				out <- val
			}
		}
	}()

	return out
}

//
//
// func tryRunPipe() {
//	in := make(Bi)
//	done := make(Bi)
//	stages := []Stage{}
//	out := ExecutePipeline(in, done, stages...)
//	go func() {
//		for result := range out {
//			fmt.Println(result)
//		}
//		log.Println("Pipe is done")
//	}()
//
//	in <- 1
//	in <- 45
//	done <- struct{}{}
//
//}
