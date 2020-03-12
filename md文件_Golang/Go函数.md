# Golang函数特性

## 默认函数


golang 并不支持默认参数。但可以通过传递零值并且在代码里判断是否零值来实现，还可以通过传递一个结构体来实现
```go
func Concat1(a string, b int) string {
    if a == "" {
        a = "default-a"
    }
    if b == 0 {
        b = 5
    }
    return fmt.Sprintf("%s%d, a, b)
}
```

## 函数的传参

```go
func changeStr(s string) {
    s = "hehe"
    fmt.Println(s)
}

func main() {
    name := "haha nihao"
    changeStr(name)
    fmt.Println(name) // 打印出来还是 "haha nihao"，没有修改成功，似乎是『值传递』
}
```

从以上代码看来，似乎是值传递，并没有修改传入的值。