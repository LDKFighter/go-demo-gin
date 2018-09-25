package factory

import (
	"fmt"
)

type Dispatcher struct {
	WorkerPool chan chan Job
}

func NewDispatcher(maxWorkers int) *Dispatcher {
	pool := make(chan chan Job, maxWorkers)
	return &Dispatcher{WorkerPool: pool}
}

func (self *Dispatcher) dispatch() {
	for {
		select {
		case job := <-JobQueue: //每当有新工作需要处理时，从空闲员工池内获取流水线
			go func(job Job) {
				jobChannel := <-self.WorkerPool
				jobChannel <- job
			}(job)
		}
	}
}

//调度器开始工作
func (self *Dispatcher) Run() {
	fmt.Sprintln("disptch run %d", MAX_WORKERS)
	for i := 0; i < MAX_WORKERS; i++ {
		worker := NewWorker(self.WorkerPool)
		worker.Start()
	}
	go self.dispatch()
}
