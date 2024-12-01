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
				<-in
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
