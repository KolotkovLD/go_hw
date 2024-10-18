package hw05parallelexecution

import (
	"errors"
	"sync"
	"sync/atomic"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	var (
		wg            sync.WaitGroup
		errorCount    int32
		runTasksCount int32
	)

	if m <= 0 {
		m = len(tasks) + 1
	}

	errorChan := make(chan error, len(tasks))
	stopChan := make(chan struct{})
	taskChan := make(chan Task, len(tasks))

	// Заполняем канал заданий
	go sendTasks(taskChan, tasks, stopChan)

	// Обрабатываем ошибки
	go checkErr(errorChan, stopChan, m, &errorCount)

	for i := 0; i < n; i++ {
		wg.Add(1)
		go runTask(&wg, taskChan, errorChan, stopChan, m, n, &runTasksCount)
	}
	// i, &errorCount

	wg.Wait()
	close(errorChan)

	if atomic.LoadInt32(&errorCount) >= int32(m) {
		return ErrErrorsLimitExceeded
	}

	return nil
}

func runTask(wg *sync.WaitGroup,
	taskChan chan Task,
	errorChan chan<- error,
	stopChan chan struct{},
	m int, n int,
	runTasksCount *int32,
) {
	// workerID int,
	//	errorCount *int32,
	// Запускает таски из канала taskChan
	defer wg.Done()
	// log.Printf("Goroutine %d: started\n", workerID)
	for {
		select {
		case <-stopChan:
			// log.Printf("Goroutine %d: stopChan closed, exiting\n", workerID)
			return
		case task, ok := <-taskChan:
			if !ok {
				// log.Printf("Goroutine %d: taskChan closed, exiting\n", workerID)
				return
			}
			// log.Printf("Goroutine %d: received a task\n", workerID)
			if err := task(); err != nil {
				// log.Printf("Goroutine %d: task returned error: %v, errorCount: %d\n", workerID, err, *errorCount)
				errorChan <- err
				atomic.AddInt32(runTasksCount, 1)
				if (int32(n) + int32(m)) <= atomic.LoadInt32(runTasksCount) {
					// log.Printf("     Goroutine %d: error \n", workerID)
					<-stopChan
					return
				}
				select {
				case errorChan <- err:
				case <-stopChan:
					return
				}
			}
		}
	}
}

func sendTasks(taskChan chan Task, tasks []Task, stopChan chan struct{}) {
	// Отправляет таски в канал taskChan
	defer close(taskChan)
	for _, task := range tasks {
		select {
		case taskChan <- task:
		case <-stopChan:
			return
		}
	}
}

func checkErr(errorChan chan error, stopChan chan struct{}, m int, errorCount *int32) {
	// Проверяет количество таков с ошибкой и прерывает работу оставшихся
	for err := range errorChan {
		// log.Printf("checkErr: err: %v", err)
		if err != nil {
			if atomic.AddInt32(errorCount, 1) >= int32(m) {
				stopChan <- struct{}{}
				close(stopChan)
				return
			}
		}
	}
}
