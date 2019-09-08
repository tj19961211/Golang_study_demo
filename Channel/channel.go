package main

import (
	"fmt"
	"time"
)

//channel也是一等公民，channel也能作为参数和返回值

//chan<-:可以表示为该channel只能接收外面传来的数据
//<-chan:表示为该channel只能取出里面的数据
func createWorker(id int) chan<- int {
	c := make(chan int)
	go func() {
		fmt.Printf("Worker %d received %c\n",
			id, <-c)
	}()
	return c
}

func worker(id int, c chan int) {
	for {
		//n := <-c //从channel中收数据
		fmt.Printf("Worker %d received %c\n",
			id, <-c)
	}
}

func chanDemo() {
	//var c chan int  要用channel可以选择使用make创建   // c == nil
	var channels [10]chan int
	var channels_1 [10]chan<- int
	for i := 0; i < 10; i++ {
		channels_1[i] = createWorker(i)
		channels[i] = make(chan int)
		go worker(i, channels[i])
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}

	//c <- 1  //发数据,发送数据后没有接受者会产生死锁
	//c <- 2
	time.Sleep(time.Microsecond) //如果不sleep的话2刚传过去，主协程就执行完毕关掉了，
	// 导致2无法在创建的channel中打印出来
}

func main() {
	chanDemo()
}
