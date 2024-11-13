package hw06pipelineexecution

import "fmt"

type (
	In  = <-chan interface{}
	Out = chan<- interface{}
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func Stage(in In) (out Out) {
	out = make(chan interface{})
	go func() {
		defer close(out)
		for task := range in {
			fmt.Printf("Stage %s processing.")
			result := task()
			out <- result
		}
	}()
	return out
}

func ExecutePipeline(in In, done In, stages ...Stage) Out {

	for _, stage := range stages {
		in = stage(in)
	}

	return nil
}
