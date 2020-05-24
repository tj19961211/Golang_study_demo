# Go变量标识符

给一个 go 变量命名(标识符)的时候，通常使用大骆驼(BigCamel)和小骆驼(smallCamel)命名法，只要保证不和内置的25个关键词冲突就行。

## Go如何声明一个变量

go 中声明一个变量很简单，使用 var 关键字就可以了，声明之后默认使用其类型的零值初始化。 比如int 类型默认是0，字符串就是空串。

当然为了简化 go 还提供了一种使用 := 直接声明并且初始化的方式(:= 只能在函数体里面声明不能在函数体外进行声明)

```go
package main

import "fmt"

func main() {
    var i int64
    var b string
    fmt.Println("i is ", i)
    fmt.Println("b is ", b)

    //同时声明并赋值
    var floatNum float64 = 1.0
    var price1, price2 float64 = 8.8, 9.6
    fmt.Println(floatNum, price1, price2)

    //还有一种简化方式，声明并且赋值，编译器负责推断类型
    ii := 1
    s := "Hello Go"
    fmt.Println("ii is ", ii)
    fmt.Println("s is ", s)
}

```

### bool类型

bool 就是真或者假，一些编程语言使用 0 和非 0 表示。但是 go 里比如 if 语句后边只能是 bool 值或者返回 bool 值的表达式，而不像 c 一样可以使用 int 值。