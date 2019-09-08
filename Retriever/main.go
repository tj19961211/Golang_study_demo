package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"time"
)

type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string,
		form map[string]string) string
}

type RetrieverPoster interface {
	Retriever
	Poster
}

func Post(poster Poster) {
	poster.Post("https://www.imooc.com",
		map[string]string{
			"name":   "Clown",
			"course": "golang",
		})
}

//r必须要实现Get方法，然后再interface里调用r里实现的Get方法
func download(r Retriever) string {
	return r.Get("https://www.imooc.com")
}

type realRetriever struct {
	UserAgent string
	TimeOut   time.Duration
}

//实现interface的Get方法，需要的参数类型是struct的结构
func (r *realRetriever) Get(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	result, err := httputil.DumpResponse(resp, true)

	if err != nil {
		panic(err)
	}

	return string(result)
}

func inspect(r Retriever) {
	switch v := r.(type) {
	case *realRetriever:
		fmt.Println("UserAgent:", v.UserAgent)

	}
	i := r.(*realRetriever)
	fmt.Println(i.UserAgent)
}

func main() {
	//第一种方法
	//先将realRetriever初始化，然后往struct结构里面填上数据
	r := &realRetriever{
		UserAgent: "Mozilla/5.0",
		TimeOut:   time.Minute,
	}
	fmt.Printf("%T, %v\n", r, r)
	inspect(r)

	//第二种方法使用interface：先实现一个interface类型，运用struct的结构往里面填充内容
	//最后调用实现的Get()方法
	var k Retriever
	k = &realRetriever{
		UserAgent: "Mozilla/6.0",
		TimeOut:   time.Minute,
	}

	inspect(k)
	fmt.Printf("%T, %v\n", r, r)
	//fmt.Println(download(k))

	//总结以上两个方法，第一种方法r不是一个interface变量，所以操作起来会很麻烦
	// （就单个文件调用interface变量而言，多文件调用实现者的接口时最好还是赋值一下）。
	//总体来说先定义interface变量后再往里面填值这种效果更好

	//因为realRetriever的结构里完成了interface所需要的Get()方法，
	//所以r可以看作是Retriever interface的类型，可以将r当作参数传入download()
	//由实现者和使用者的理论可以看出，download()是使用者，realRetriever实现的Get()方法是实现者
	//由返回的结果看出，interface有什么作用是使用者决定，并不由实现者决定

	//fmt.Println(download(r))

}
