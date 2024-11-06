package hw06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

//func Stage(in <-chan interface{}) (out <-chan interface{}) {
//	out = make(chan interface{})
//	go func() { /* Some work */ }()
//	return out
//}

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	// Place your code here.
	return nil
}
