## 进程、线程(内核级别)和协程(用户级别)的区别

对于进程、线程都是由内核进行调度的，有 CPU 时间片的概念，进行 抢占式调度（有多种调度算法）

对于 协程(用户级线程)，这是对内核透明的，也就是系统并不知道有协程的存在，是完全由用户自己的程序进行调度的，因为是由用户程序自己控制，那么就很难像抢占式调度那样做到强制的 CPU 控制权切换到其他进程/线程，通常只能进行 协作式调度，需要协程自己主动把控制权转让出去之后，其他协程才能被执行到。

### Goroutine和协程的区别

本质上，goroutine就是协程，不同的是，golang在runtime、系统调用等多方面对goroutine调度进行封装和处理，当遇到长时间执行或者进行系统调用时，会主动把当前 goroutine 的CPU (P) 转让出去，让其他 goroutine 能被调度并执行，也就是 Golang 从语言层面支持了协程。Golang 的一大特色就是从语言层面原生支持协程，在函数或者方法前面加 go关键字就可创建一个协程。

> 其他方面的比较：
- 内存消耗方面  
    - 每个goroutine(协程)默认占用内存远比 Java 、C 的线程少。  
    - goroutine：2KB  
    - 线程：8MB
- 线程和 goroutine 切换调度开销方面  
    - 线程/goroutine 切换开销方面，goroutine 远比线程小  
    - 线程：涉及模式切换(从用户态切换到内核态)、16个寄存器、PC、SP…等寄存器的刷新等。  
    - goroutine：只有三个寄存器的值修改 – PC / SP / DX.