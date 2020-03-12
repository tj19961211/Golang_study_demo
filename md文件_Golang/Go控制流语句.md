# Go控制流语句

go 语言提供了三种分之方式，if/switch/select，分别来看下它们的使用方式和注意事项。

### if语句
if/else/else if 使用起来跟python的非常像，它的格式如下：
```go
if optionalStatement1; booleanExpression1 {
    block1
}else if optionalStatement2; booleanExpression2 {
    block2
}else {
    block3
}
```

> 需要注意的几点    
- 不像 c 语言，if 条件只能是一个 bool 值或者返回 bool 值得表达式，而不能是 int (c/python 可以)
- if 后边是可以先跟一个表达式，比如之前我们见到的获取 map 值的 if v,ok := m[key]; ok {}

以下几个小示例，快速入门：
```go
func testIf() {
    ok := true
    if ok {
        fmt.Println("ok is true")
    }

    day := "Friday"
    if day == "Friday" {
        fmt.Pirntln("明天不上班")
    }else if day == "Sunday" {
        fmt.Println("周末好快")
    }else {
        fmt.Println("干活啦")
    }

    m := make(map[string]string)
    m["王八"] = "绿豆"
    if v, ok := m["王八"]; ok {
        fmt.Println(v)
    } 
}
```

### switch/case

go 里改善了 c 语言的switch，比如我们不用每一个 case 都要加上 break（很多 bug 都是因为粗心的程序员忘记加上 break 导致的，go 中你再也不用担心啦）。go 语言中 switch 语法如下：

```go
switch optionalStatement; optionalExpression {
    case expressionList1: block1
    …
    case expressionListN: blockN
    default: blockD
}
```

> 需要注意的几点是：
- case 语句不用加上 break，不像 c 不会去自动执行下一个 case，除非你显示使用 fallthrough 语句指定（用的不多)
- default 语句是可选的，如果都没有匹配，可以给一个默认行为

来看几个小的示例快速入门：
```go
func testSwitch() {
    // 常规用法
    day := 0
    switch day {
    case 0, 6:
        fmt.Println("周末")
    case 1, 2, 3, 4, 5:
        fmt.Println("工作日")
    default:
        fmt.Println("不合法")
    }
    // case 后边还可以是表达式
    a, b := 1, 2
    a, b = b, a
    switch {
    case a < b:
        fmt.Println("a < b")
    case a > b:
        fmt.Println("a > b")
    }
}
```

## 循环

和其他语言不同的是，go 只提供了一个关键词 for 来实现循环（没有 while 和 do while），是不是非常吝啬。 不过你会发现其实只要一个 for 就可以满足需求了，当然它提供了几种不同的使用方式，来瞅瞅：

```go
// 常规用法，和其他语言类似
for optionalPreStatement; booleanExpress; optionalPostStatement {
  block
}

// 无限循环，block 会被一直重复执行
for {
  block
}

// 实现while循环，block 一直执行直到 表达式为false
for booleanExpression {
  block
}
```

此外for/range语法还支持让我们去迭代字符串,数组，切片，map或者通道(channel),之前其实已经用过好多次了，请看几个例子

```go
func testFor() {
    intSlice := []int{1, 2, 3}
    for index, item := range intSlice {
        fmt.Println(index, item)
    }
    for index := range intSlice { // 省略 item 之后遍历的是 key，注意不像python 直接遍历值
        fmt.Println(index)
    }

    m := map[string]string{"k1": "v1", "k2": "v2"}
    for k, v := range m {
        fmt.Println(k, v)
    }

    for k := range m {
        fmt.Println(k)
    }
}
```

如何跳出循环呢？和很多编程语言一样，go 使用 break 和 continue 分别跳出或者进入下一个循环，相信你已经非常熟悉了。 其实 break 和 continue 还支持后边跟一个跳转标签，但是一般代码中用得比较少，你只需要知道还有这个功能， 万不得已需要使用的时候再 google 就好，你很可能永远都用不到这个。



## 小结
go 的控制流相比其他语言来说更加简单，比如连 while 都没有提供，直接用一个 for 来解决。本章还没有介绍 select，我们将在后续介绍到 channel 和并发编程的时候介绍 select 的使用。 学会了流程控制之后，下一章我们来看下函数，函数是编写大型项目的基础。