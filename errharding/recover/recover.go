package main

import (
	"fmt"
)

//一下代码式panic与recover的工作流程

func tryRecover() {
	defer func() {
		r := recover()
		if err, ok := r.(error); ok {
			fmt.Println("Error occurred:", err)
		} else {
			panic(r)
		}
	}()
	//panic(errors.New("This is an error"))
	b := 0
	a := 5 / b
	fmt.Println(a)
}

func main() {
	tryRecover()
}
