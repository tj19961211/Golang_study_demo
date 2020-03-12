# Go数组(array)和切片(slice)

        数组是我们最常用的线性结构，比如在python中我们最常使用的是list。在Go中提供了两种常见的线性结构：数组(array)和切片(slice)。
        数组就是固定长度的某种类型的序列，而切片更加灵活，它的长度是可以变化的，所以在业务中最常使用到的是切片


## 数组

        数组是一个包含相同类型的固定长度的序列，不像动态语言那样可以在list中存储不同类型的值，Go语言中数组的类型需要一致。

```go
//声明数组： [length]Type
//初始化： [N]Type{value1, value2, ..., valueN}
//省略长度： [...]Type{value1, value2, ..., vlaueN}
//二维数组： [M][N]Type
```

以下是一个例子
```go
package main

improt "fmt"

func testArray() {
    // 声明一个 int64数组，声明之后每个元素是该类型默认的『零值』
    var arrayIn64 [3]int64
    arrayIn64[0], arrayIn64[1] = 1, 2
    fmt.Println(arrayIn64)

    //声明并且初始化
    arrayString := [3]string{"tan", "tang", "dfsagjhy"}
    fmt.Println(arrayString)

    // 也可以省略长度，让 go 自动计算。这个时候你需要是使用省略号 ...
    // 创建一个长度为 3 的 float64 数组
    arrayFloat := [...]float64{1.5, 8.8, 6.6}
    fmt.Println(arrayFloat)

    //二维数组
    matrix := [2][2]int64{
        {0, 1},
        {2, 3}
    }
    fmt.Println(matrix)
}

func main() {
    testArray()
}
```

一般对于数组的操作也就是获取长度(len函数， 0到n-1)，获取指定下标的元素([index]),给数组第i个元素赋值等

```go
func testArrayOperation() {
    names := [4]string{"tan", "li", "ds", "dsa"}
    fmt.Printf("names has %d elements\n", len(names))
    fmt.Println(names[1]) //NOTE: 注意如果下标越界即超过范围会发生panic
    name[3] = "lao zhao"
    fmt.Println(name[3])
}
```

## 切片

        切片比数组更加灵活，它的长度是可以变化的。切片的容量实可以自动扩容的，每次扩容是当前容量的一倍。
        你可以简单地理解为切片是一个指向数组的指针，这个数组有它的总容量(capacity)，和目前使用使用的长度(length)。创建一个切片我们可以使用构造方式或者内置的 make 函数。

```go
// 创建一个类型为 Type, 长度为 length, 容量为 capacity 的 slice。一般我们不太关心容量而是关心长度
make([]Type, length, c capacity)
// 创建一个类型为 Type, 长度为 length 的 slice，一般我们不太关心容量，而是让 go 帮我们自动处理扩容问题
 make([]Type, length) // 最常用的一种方式, 如果实现知道 slice 的长度，可以避免 slice 扩容操作，性能更好
// 创建一个 Type 类型 slice
[]Type{} // 和 make([]Type, 0) 等价
// 创建并且初始化一个 slice。注意和数组的区别是 [] 里边没有省略号 ...
[]Type{value1, value2, ... , valueN}
```


以下例子可以快速了解slice
```go
func testSlice() {
    // 创建并且初始化一个 slice
    names := []string{"zhang", "sda", "dsad", "li"}
    // 打印 names, 长度和容量
    fmt.Println(names, len(names), cap(names))
    names2 := names[:3] // 获取子切片 0,1,2 三个元素，注意左闭右开区间
    fmr.Println(names2)
    // 尝试修改一下 names2 ，注意 names 也会跟着改变么？
    names2[0] = "lao zhang"
    fmt.Println(names, names2) // 你会发现names也变了，这里起始它们共用了底层结构，注意这个问题

    // 遍历一个 slice 我们使用 for/range 语法
    for idx, name := range names { // 如果没有用到下标 idx，可以写成下划线 _ 作为占位符，但是不能省略
        fmt.Println(idx, name)
    }

    // 修改切片主要通过赋值和 append 操作。使用 append 修改切片
    vals := make([]int, 0)
    for i := 0; i < 3; i++ {
        vals = append(vals, i)
    }
    fmt.Println(vals)
    vals2 := []int{3, 4, 5}
    newVals := append(vals, vals2...) // 可以使用省略号的方式『解包』一个 slice 来连接两个 slice
    fmt.Println(newVals)
}
```

>>>>Tip:    
    如果在创建一个 slice 之前预先知道了它的长度，make 函数最好传递长度进去，防止 append 操作可能导致重新分配内存降低效率。 比如下边这个例子，使用第二种方式效率更高一些：



```go
package main

import "fmt"

func main() {
    manyInts := make([]int, 1000000)

    // bad way
    a := make([]int, 0)
    for _, val := range manyInts {
        a = append(a, val+val) // 扩容 a 会导致重新分配内存
    }
    fmt.Println(a)

    // good way
    b := make([]int, len(manyInts))
    for i, val := range manyInts {
        b[i] = val + val // 注意这里是赋值了，不是 append
    }
    fmt.Println(b)
}
```


### 如何给一个切片排序？


        切片操作跟python list 比较相似，但是也要注意一些区别。子切片与原切片共享底层结构，如果需要深拷贝你得自己去复制一个新的。
        另外 go 只支持正数的索引，你需要保证 slice 索引值必须要在 0 到 length-1，否则会出现 panic 导致程序退出。


        这里介绍一下如何来排序和搜索一个 slice，除了自己写排序算法之外，标准库提供了 sort 包来帮助我们处理排序问题。 常用的有几个函数，go 标准库文档已经有非常好的示例（好好学英语啊）：

```go
sort.Ints(a []int) // Ints sorts a slice of ints in increasing order.
sort.Float64s(a []float64) // Float64s sorts a slice of float64s in increasing order (not-a-number values are treated as less than other values).
sort.Search(n int, f func(int) bool) int // Search uses binary search to find and return the smallest index i in [0, n) at which f(i) is true
```
