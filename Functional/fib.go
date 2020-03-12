package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func fibonacci() intGen {
	a, b := 0, 1
	return func() int {
		a, b = b, b+a
		return a
	}
}

//为斐波那契数列函数实现接口
type intGen func() int

func (g intGen) Read(
	p []byte) (n int, err error) {
	next := g() //获取下一个g
	if next > 10000 {
		return 0, io.EOF
	}
	s := fmt.Sprintf("%d\n", next)
	return strings.NewReader(s).Read(p)
}

func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	f := fibonacci()
	//for i := 0; i < 20; i++ {
	//	fmt.Println(f())
	//}
	printFileContents(f)
}
