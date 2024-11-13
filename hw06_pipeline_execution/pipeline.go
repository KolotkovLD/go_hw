package hw06pipelineexecution

import (
	"log"
	"sync"
)

type (
	In  = chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

//func Stage(in In) (out Out) {
//	out = make(chan interface{})
//	go func() {
//		defer close(out)
//		for input := range in {
//			log.Println("Stage processing.")
//
//			result := func(input interface{}) int {
//				log.Println(input)
//				return 0
//			}(input)
//			out <- result
//		}
//	}()
//	return out
//}

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	//
	//for _, stage := range stages {
	//	in = stage(in)
	//}
	//
	//return nil

	wg := &sync.WaitGroup{} // Для синхронизации завершения всех горутин

	// Канал для распространения сигнала завершения
	stop := make(Bi)

	// Функция для закрытия выходных каналов при завершении работы
	closeOutput := func(out Out) {
		if out != nil {
			select {
			case <-done:
				log.Println("Pipeline stopped")
			default:
				close(out)
			}
		}
	}

	// Входной канал передается первому этапу
	currentIn := in

	// Проходим по всем этапам
	for _, stage := range stages {
		currentOut := make(Out)

		// Увеличиваем счетчик WaitGroup перед запуском новой горутины
		wg.Add(1)
		log.Println(" -  Новая горутина")

		// Запускаем горутину для выполнения этапа
		go func(in In, out Out, stop Bi) {
			defer wg.Done()
			defer closeOutput(out)

			// Вызываем этап
			out = stage(in)
			log.Println(" -  Вызов этапов")

			// Передаем результат следующему этапу
			for v := range out {
				select {
				case <-stop:
					log.Println("Закрываем горутину")
					return
				case out <- v:
					log.Println("Передаём результат")
				}
			}
		}(currentIn, currentOut, stop)

		// Следующий этап получает данные из выхода текущего
		currentIn = currentOut
	}

	// Ожидаем завершения всех горутин
	//go func() {
	wg.Wait()
	close(stop)
	//}()

	return currentIn
}
