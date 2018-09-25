package factory

import (
	"fmt"
	"os"
	//"strconv"
)

var (
	MAX_WORKERS int
	MAX_QUEUE   int
)

type Job struct {
	Payload string
}

///定义工作队列(部门工作)
var JobQueue chan Job

type Worker struct {
	WorkerPool chan chan Job //工作队列池
	JObChannel chan Job      //工作队列（流水线）
	quit       chan bool     //结束工作信号
}

func init() {
	MAX_WORKERS = 20 //strconv.Atoi(os.Getenv("GO_MAX_WORKERS"))
	MAX_QUEUE = 10   //strconv.Atoi(os.Getenv("GO_MAX_QUEUE"))
	JobQueue = make(chan Job, MAX_QUEUE)
	fmt.Println("JobQue Length: %s", os.Getenv("GO_MAX_WORKERS"))
}

//参加处理工作池内工作的员工
func NewWorker(jobPool chan chan Job) Worker {
	return Worker{
		WorkerPool: jobPool,
		JObChannel: make(chan Job),  //无缓冲通道，一次处理一个工作
		quit:       make(chan bool), //是否停止工作
	}
}

//开始工作
func (self Worker) Start() {
	go func() {
		self.WorkerPool <- self.JObChannel //注册员工的工作队列到工作池内，等待调度
		select {
		case <-self.JObChannel: //等待工作队列传来任务
			//fmt.Printf("Job %v", Job)
		case <-self.quit:
			return
		}
	}()
}

//停止工作
func (self Worker) Stop() {
	go func() {
		self.quit <- true
	}()
}
