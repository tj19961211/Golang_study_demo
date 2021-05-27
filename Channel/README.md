# Channel

作为Go核心的数据结构和Goroutine之间的通信方式，Channel是支撑Go语言高性能并发编程模型的重要结构，我们首先需要了解Channel背后的设计原理以及它的底层数据结构

## 设计原理

Go语言中最常见的、也是经常被人提及的设计模式就是 -- 不要通过共享内存的方式进行通信，而是应该通过通信方式来共享内存。在很多主流的编程语言中，多个线程传递数据的方式一般都是共享内存，为了解决线程冲突的问题，我们需要限制同一时间能够读写这些变量的线程数量，这与Go语言鼓励的方式并不相同。

![多线程使用共享内存传递数据](https://img.draveness.me/2020-01-28-15802171487042-shared-memory.png)

<center>多线程使用共享内存传递数据</center>

虽然在Go语言中也能使用共享内存加互斥锁进行通信，但是Go语言提供了一种不同的并发模型，也就是通信顺序进程(Communicating sequential processes, CSP )。Goroutine和Channel分别对应CSP中的实体和传递消息的媒介，Go语言中的Goroutine会通过Channel传递数据。

![Goroutine 使用 Channel 传递数据](https://img.draveness.me/2020-01-28-15802171487080-channel-and-goroutines.png)

<center>Goroutine 使用 Channel 传递数据</center>

上图中的两个Goroutine，一个会向Channel中发送数据，另一个会从Channel中接收数据，它两者能够独立运行并不存在直接关联，但能通过Channel间接完成通信。

### 先入先出

目前Channel收发操作均遵循先进先出(FIFO)的设计，具体规则如下：

- 先从Channel读取数据的Goroutine会先接收到数据；

- 先向Channel发送数据的Goroutine会得到先发送数据的权利

