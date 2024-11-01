package hw05parallelexecution

import (
	"errors"
	"log"
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

	taskChan := make(chan Task)

	// Заполняем канал заданий
	go sendTasks(taskChan, tasks, &errorCount, &runTasksCount, n, m)

	for i := 0; i < n; i++ {
		wg.Add(1)
		go runTask(&wg, taskChan, &errorCount, &runTasksCount, i, n, m)
	}

	wg.Wait()
	if stopRule(&errorCount, &runTasksCount, n, m) {
		return ErrErrorsLimitExceeded
	}
	return nil
}

func runTask(wg *sync.WaitGroup,
	taskChan chan Task,
	errorCount *int32,
	runTasksCount *int32,
	workerID int,
	n int,
	m int,
) {
	// Запускает таски из канала taskChan
	defer wg.Done()
	log.Printf("Goroutine %d: started\n", workerID)
	for {
		task, ok := <-taskChan
		atomic.AddInt32(runTasksCount, 1)
		log.Printf("Goroutine %d: received a task, runTasksCount: %d\n", workerID, *runTasksCount)
		if !ok || stopRule(errorCount, runTasksCount, n, m) {
			log.Printf("Goroutine %d: taskChan closed, exiting\n", workerID)
			return
		}

		if err := task(); err != nil {
			atomic.AddInt32(errorCount, 1)
			log.Printf("Goroutine %d: task returned error: %v, errorCount: %d\n", workerID, err, *errorCount)
		}
	}
}

func sendTasks(taskChan chan Task,
	tasks []Task,
	errorCount *int32,
	runTasksCount *int32,
	n int,
	m int,
) {
	// Отправляет таски в канал taskChan
	defer func() {
		close(taskChan)
		log.Printf("sendTasks is done!!!!!!!")
	}()
	for _, task := range tasks {
		if stopRule(errorCount, runTasksCount, n, m) {
			log.Printf("     [sendTasks]  errorCount: %d ; runTasksCount: %d\n", *errorCount, *runTasksCount)
			return
		}
		taskChan <- task
	}
}

func stopRule(errorCount *int32, runTasksCount *int32, n, m int) bool {
	if atomic.LoadInt32(errorCount) >= int32(m) && (int32(n)+int32(m)) <= atomic.LoadInt32(runTasksCount) {
		log.Printf(" [stop_rule]  errorCount: %d ; runTasksCount: %d\n", *errorCount, *runTasksCount)
		return true
	}
	return false
}
