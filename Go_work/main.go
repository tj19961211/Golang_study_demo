package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

const (
	MaxQueue  = 200 // 随便设置值
	MaxWorker = 100 // 随便设置值
)

var JobQueue chan Job

func init() {
	JobQueue = make(chan Job, MaxQueue)
}

type Payload struct {
}

type Job struct {
	PayLoad Payload
}

type Worker struct {
	WorkerPool chan chan Job
	JobChannel chan Job
	quit       chan bool
}

func NewWorker(workerPool chan chan Job) Worker {
	return Worker{
		WorkerPool: workerPool,
		JobChannel: make(chan Job),
		quit:       make(chan bool),
	}
}

// Start 方法开启一个 worker 循环，监听退出 channel ， 可按需停止这个循环
func (w Worker) Start() {
	go func() {
		for {
			// 将当前的 worker 注册到 worker 队列中
			w.WorkerPool <- w.JobChannel
			select {
			case job := <-w.JobChannel:
				// 真正业务的地方
				// 模拟操作耗时
				time.Sleep(500 * time.Millisecond)
				fmt.Printf("上传成功：%v\n", job)
			case <-w.quit:
				return
			}
		}
	}()
}

func (w Worker) stop() {
	go func() {
		w.quit <- true
	}()
}

// 初始化操作

type Dispathcher struct {
	// 注册到 dispathcher 的 worker channel 池
	WorkerPool chan chan Job
}

func NewDispathcher(maxWorkers int) *Dispathcher {
	pool := make(chan chan Job, maxWorkers)
	return &Dispathcher{WorkerPool: pool}
}

func (d *Dispathcher) Run() {
	// 开始运行 n 个 worker
	for i := 0; i < MaxWorker; i++ {
		worker := NewWorker(d.WorkerPool)
		worker.Start()
	}
	go d.dispatch()
}

func (d *Dispathcher) dispatch() {
	for {
		select {
		case job := <-JobQueue:
			go func(job Job) {
				// 尝试获取一个可用的 worker job channel，阻塞直到有可用的 worker
				jobChannel := <-d.WorkerPool
				// 分发任务到 worker job channel 中
				jobChannel <- job
			}(job)
		}
	}
}

func payloadHandler(w http.ResponseWriter, r *http.Request) {
	work := Job{PayLoad: Payload{}}
	JobQueue <- work
	_, _ = w.Write([]byte("操作成功"))
}

func main() {
	// 通过调度器创建 worker， 监听来自 JobQueue 的任务
	d := NewDispathcher(MaxWorker)
	d.Run()
	http.HandleFunc("/payload", payloadHandler)
	log.Fatal(http.ListenAndServe(":8099", nil))
}

// func (p *Payload) UpdateToS3() error {
// 	// 存储逻辑，模拟操作耗时
// 	time.Sleep(500 * time.Millisecond)
// 	fmt.Println("上传成功")
// 	return nil
// }

// func payloadHandler(w http.ResponseWriter, r *http.Request) {
// 	// 业务过滤
// 	// 请求 body 解析
// 	var p Payload
// 	//go p.UpdateToS3()
// 	Queue <- p
// 	w.Write([]byte("操作成功"))
// }

// // 处理任务
// func StartProcessor() {
// 	for {
// 		select {
// 		case payload := <-Queue:
// 			payload.UpdateToS3()
// 		}
// 	}
// }

// func main() {
// 	http.HandleFunc("/payload", payloadHandler)
// 	// 单独开一个 g 接收与处理任务
// 	go StartProcessor()
// 	log.Fatal(http.ListenAndServe(":8099", nil))
// }
