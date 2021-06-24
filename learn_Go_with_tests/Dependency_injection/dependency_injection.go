package Dependency_injection

import (
	"bytes"
	"fmt"
)

/*
用 writer 把问候发送到我们测试中的缓冲区。

记住 fmt.Fprintf 和 fmt.Printf 一样，只不过 fmt.Fprintf 会接收一个 Writer 参数，用于把字符串传递过去，而 fmt.Printf 默认是标准输出。
*/
func Greet(write *bytes.Buffer, name string) {
	fmt.Fprintf(write, "Hello, %s", name)
}
