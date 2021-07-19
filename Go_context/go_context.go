package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

/*
注意事项：
    - 不要将上下文存储在结构类型中

	   在 Go 语言中，所有的第三方开源库，业务代码。清一色的都会将 context 放在方法的一个入参参数，作为首位形参

	- 标准要求

	   每一个方法得第一个参数都将 context 作为第一个参数，并使用 ctx 变量名惯用语。

	当然也有极少数是把 context 放在结构体中得。基本常见于：

	    - 底层基础库
		- DDD 结构

	每个请求都是独立得，触发方法得 context 自然每一个都不一样，弄清楚应用的场景很重要，否则遵循 Go 基本规范就好
*/

// 以下展示 go context 的用法
const shortDuration = 1 * time.Millisecond

// 以下是简单的使用 context 方式
func goSimpleUseContext() {
	ctx, cancel := context.WithTimeout(context.Background(), shortDuration)
	defer cancel()

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("clown is rascal")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}

// 正确的使用方式之对第三方的调用
func goCorrectContext() {
	req, err := http.NewRequest("GET", "http://clown.com/", nil)
	if err != nil {
		fmt.Println("http.NewRequest err: %+v", err)
		return
	}

	// 若你发现第三方开源库没支持 context，那建议赶紧跑，换一个。免得在微服务体系下出现级联故障，还没有简单的手段控制，那就很麻烦了。
	ctx, cancel := context.WithTimeout(req.Context(), 50*time.Millisecond)
	defer cancel()

	req = req.WithContext(ctx)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("http.DefaultClient.Do err: %+v", err)
		return
	}

	defer resp.Body.Close()
}

// 在业务场景中，context 传值使用与传必要的业务核心属性，例如： 租户号、小程序ID等。不要将可选参数放到 context 中，否则可能会一团乱
func goWithValueContext() {
	type favContextKey string

	f := func(ctx context.Context, k favContextKey) {
		if v := ctx.Value(k); v != nil {
			fmt.Println("found value: ", v)
			return
		}
		fmt.Println("key not found: ", k)
	}

	k := favContextKey("clown")
	ctx := context.WithValue(context.Background(), k, "Clown")

	f(ctx, k)
	f(ctx, favContextKey("nihao"))
}

func main() {
	goSimpleUseContext()
	goCorrectContext()
	goWithValueContext()
}

/*
总结：
    - 对第三方调用调用要传入 context，用于控制远程调用。
	- 不要将上下文存储在结构类型中，尽可能的作为函数第一位形参传入。
	- 函数调用链必须传播上下文，实现完成链路上的控制。
	- context 的继承和派生，保证父、子级 context 的联动。
	- 不传递 nil context， 不确定的 context 应当使用 TODO
	- context 仅传必要的值，不要让可选参数揉在一起
*/
