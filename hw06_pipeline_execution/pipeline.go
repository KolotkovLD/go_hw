package hw06pipelineexecution

import "fmt"

//type (
//	In  = <-chan interface{}
//	Out = In
//	Bi  = chan interface{}
//)

//type Stage func(in In)I (out Out)

type Stage func(in In) Out
type In = any
type Out = any

//func stage(in <-chan interface{}) (out <-chan interface{}) {
//	out = make(chan interface{})
//	go func() { /* Some work */ }()
//	return out
//}

func ExecutePipeline(in In, done In, stages ...Stage) Out {

	for _, stage := range stages {
		in = stage(in)
	}

	return nil
}

// Stage - функция, которая принимает канал на чтение и возвращает новый канал,
// из которого будут поступать результаты обработки данных.
func Stage(name string, in <-chan interface{}) <-chan interface{} {
	out := make(chan interface{})
	go func() {
		defer close(out)
		for v := range in {
			fmt.Printf("Stage %s processing value: %v\n", name, v)
			// Здесь выполняется какая-то работа над значением
			out <- v // Отправляем обработанное значение дальше
		}
	}()
	return out
}

// Pipeline запускает пайплайн из нескольких стадий
func Pipeline(stages ...func(in <-chan interface{}) <-chan interface{}) <-chan interface{} {
	var currentStage <-chan interface{}
	currentStage = make(chan interface{})

	// Запускаем все стадии последовательно
	for _, stage := range stages {
		currentStage = stage(currentStage)
	}

	return currentStage
}

func main() {
	stage1 := func(in <-chan interface{}) <-chan interface{} {
		return Stage("stage1", in)
	}
	stage2 := func(in <-chan interface{}) <-chan interface{} {
		return Stage("stage2", in)
	}
	stage3 := func(in <-chan interface{}) <-chan interface{} {
		return Stage("stage3", in)
	}

	input := make(chan interface{}, 10)
	output := Pipeline(stage1, stage2, stage3)(input)

	// Генерация входных данных
	go func() {
		defer close(input)
		for i := 0; i < 5; i++ {
			input <- i
		}
	}()

	// Получение результатов
	for result := range output {
		fmt.Println("Final result:", result)
	}
}
