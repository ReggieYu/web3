package main

import (
	"fmt"
	"sync"
	"time"
)

type Task struct {
	id       int
	function func() error
}

type TaskResult struct {
	taskId    int
	duration  time.Duration
	startTime time.Time
	endTime   time.Time
	err       error
}

type Scheduler struct {
	tasks      []Task
	concurrent int
	wg         sync.WaitGroup
	results    chan TaskResult
}

func newScheduler(concurrent int) *Scheduler {
	return &Scheduler{
		concurrent: concurrent,
		results:    make(chan TaskResult, 100),
	}
}

func (s *Scheduler) addTask(task Task) {
	s.tasks = append(s.tasks, task)
}

func (s *Scheduler) worker(taskChan <-chan Task) {
	defer s.wg.Done()
	for task := range taskChan {
		start := time.Now()
		err := task.function()
		end := time.Now()
		s.results <- TaskResult{
			taskId:    task.id,
			startTime: start,
			endTime:   end,
			err:       err,
			duration:  end.Sub(start),
		}
	}
}

func (s *Scheduler) run() {
	taskChan := make(chan Task, len(s.tasks))
	defer close(s.results)

	for i := 0; i < s.concurrent; i++ {
		s.wg.Add(1)
		go s.worker(taskChan)
	}

	go func() {
		for _, task := range s.tasks {
			taskChan <- task
		}

		close(taskChan)
	}()

	s.wg.Wait()
}

func (s *Scheduler) printRes() {
	for res := range s.results {
		status := "success"
		if res.err != nil {
			status = "failed"
		}

		fmt.Printf("taskId: %d status: %s startTime: %v endTime: %v duration: %v \n", res.taskId, status, res.startTime,
			res.endTime, res.duration)
	}
}

func main() {
	task1 := func() error {
		time.Sleep(1 * time.Second)
		return nil
	}

	task2 := func() error {
		time.Sleep(2 * time.Second)
		return nil
	}

	task3 := func() error {
		time.Sleep(500 * time.Millisecond)
		return fmt.Errorf("模拟错误")
	}

	scheduler := newScheduler(2)
	scheduler.addTask(Task{
		id:       1,
		function: task1,
	})
	scheduler.addTask(Task{
		id:       2,
		function: task2,
	})
	scheduler.addTask(Task{
		id:       3,
		function: task3,
	})

	go scheduler.run()

	scheduler.printRes()
}
