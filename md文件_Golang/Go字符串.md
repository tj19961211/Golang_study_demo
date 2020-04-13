# 整数

倘若64位的整数仍然满足不了，可以使用大整数big.int和有理数big.Rat类型

# 字符串

在业务中最常用的就是字符串(string)，web开发几乎天天就是和字符串打交道。Go的字符串是使用UTF-8编码的字符序列，这意味着你可以使用任意国家的语言。Go中我们可以使用双引号("")和反引号(``)来创建字符串，区别在于，反引号会忽略转义字符，并且可以创建多行字符串。

```go
func stringDemo() {
    // 如果字符串本身也有双引号，就需要把里边的双引号使用 \ 转义
    s1 := "\"Hello Go\""
    // 使用反斜线就可以直接包含双引号了
    s2 := `"Hello Go"`
    fmt.Println(s1) // 打印出 "Hello Go"
    fmt.Println(s2) // 打印出 "Hello Go"

    s3 := `
你好
`
    s4 := "Golang !"
    fmt.Println(s3 + s4)
}
```

## 字符串和数值类型的转换

在Python里进行这种转换是一件很容易的事情，但是go里面可不方便。Atoi是Ascii to Integer的缩写

```go
func testConvert() { // 测试 int 和 string(decimal) 互相转换的函数
    // https://yourbasic.org/golang/convert-int-to-string/
    // int -> string
    sint := strconv.Itoa(97)
    fmt.Println(sint, sint == "97")

    // byte -> string
    bytea := byte(1)
    bint := strconv.Itoa(int(bytea))
    fmt.Println(bint)

    // int64 -> string
    sint64 := strconv.FormatInt(int64(97), 10)
    fmt.Println(sint64, sint64 == "97")

    // int64 -> string (hex) ，十六进制
    sint64hex := strconv.FormatInt(int64(97), 16)
    fmt.Println(sint64hex, sint64hex == "61")

    // string -> int
    _int, _ := strconv.Atoi("97")
    fmt.Println(_int, _int == int(97))

    // string -> int64
    _int64, _ := strconv.ParseInt("97", 10, 64)
    fmt.Println(_int64, _int64 == int64(97))

    // https://stackoverflow.com/questions/30299649/parse-string-to-specific-type-of-int-int8-int16-int32-int64
    // string -> int32，注意 parseInt 始终返回的是 int64，所以还是需要 int32(n) 强转一下
    _int32, _ := strconv.ParseInt("97", 10, 32)
    fmt.Println(_int32, int32(_int32) == int32(97))

    // int32 -> string, https://stackoverflow.com/questions/39442167/convert-int32-to-string-in-golang
    i := 42
    strconv.FormatInt(int64(i), 10) // fast
    strconv.Itoa(int(i))            // fast
    fmt.Sprint(i)                   // slow

    // int -> int64 ，不会丢失精度
    var n int = 97
    fmt.Println(int64(n) == int64(97))

    // string -> float32/float64  https://yourbasic.org/golang/convert-string-to-float/
    f := "3.14159265"
    if s, err := strconv.ParseFloat(f, 32); err == nil {
        fmt.Println(s) // 3.1415927410125732
    }
    if s, err := strconv.ParseFloat(f, 64); err == nil {
        fmt.Println(s) // 3.14159265
    }
}
```

## 常量和变量

常量顾名思义你没法改变它，在一些全局变量中使用const会更加安全。常量表达式是在编译期计算。
 对于一些被整个模块或者其他模块经常使用的变量来说，最好定义成const防止被意外修改。

```go
   例：
    const (
    Sunday    = 0
    Monday    = 1
    Tuesday   = 2
    Wednesday = 3
    Thursday  = 4
    Friday    = 5
    Saturday  = 6
)
```

## 与字符串相关的包的例子

### strconv包

将字符串与其他类型进行相互的转化和string的应用

```go
//基本数值转换
fmt.Println(strconv.Itoa(10)) // 整型转换成string  Output: 10
fmt.Println(strconv.Atoi("711")) // string转换成整型 Outout: 711 <nil>

//解析
fmt.Println(strconv.ParseBool("false"))  // Output: false <nil>
fmt.Pritnln(strconv.ParseFloat("3.14", 64)) // 传入参数第二个是转换成的bitSize

//解析与格式化是互逆的操作

//格式化
fmt.Println(strconv.FormatBool(True))
fmt.Println(strconv.FormatInt(20, 16))
```
