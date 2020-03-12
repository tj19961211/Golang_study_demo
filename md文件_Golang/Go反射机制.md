## TypeOf 和 ValueOf

在Go的反射定义中，任何接口都由两部分组成，一个是接口的具体类型，一个是具体类型对应的值。
比如var i int = 3 ，因为interface{}可以表示任何类型，所以变量i可以转为interface{}，所以可以把变量i当成一个接口，那么这个变量在Go反射中的表示就是<Value,Type>，其中Value为变量的值3,Type变量的为类型int。

在Go反射中，标准库为我们提供两种类型来分别表示他们reflect.Value和reflect.Type，并且提供了两个函数来获取任意对象的Value和Type。

> 简单的说reflect.Value是当前接口的值，而reflect.Type是获取到接口值的类型

```go
func main() {
	u:= User{"张三",20}
	t:=reflect.TypeOf(u)
	fmt.Println(t)
}

type User struct{
	Name string
	Age int
}
```

- reflect.TypeOf可以获取任意对象的具体类型，这里通过打印输出可以看到是main.User这个结构体类型。
- reflect.TypeOf函数接受一个空接口interface{}作为参数，所以这个方法可以接受任何类型的对象。

接着上面的例子，我们看下如何反射获取一个对象的Value。

```go
v:=reflect.ValueOf(u)
fmt.Println(v)
```
和TypeOf函数一样，也可以接受任意对象，可以看到打印输出为{张三 20}。对于以上这两种输出，Go语言还通过fmt.Printf函数为我们提供了简便的方法。

```go
fmt.Printf("%T\n",u)
fmt.Printf("%v\n",u)
```
这个例子和以上的例子中的输出一样。


## reflect.Value转原始类型

上面的例子我们可以通过reflect.ValueOf函数把任意类型的对象转为一个reflect.Value，那我们如果我们想逆向转过回来呢，其实也是可以的，reflect.Value为我们提供了Inteface方法来帮我们做这个事情。继续接上面的例子：
```go
	u1:=v.Interface().(User)
	fmt.Println(u1)
```

这样我们就又还原为原来的User对象了,通过打印的输出就可以验证。这里可以还原的原因是因为在Go的反射中，把任意一个对象分为reflect.Value和reflect.Type，而reflect.Value又同时持有一个对象的reflect.Value和reflect.Type,所以我们可以通过reflect.Value的Interface方法实现还原。现在我们看看如何从一个reflect.Value获取对应的reflect.Type。

```go
	t1:=v.Type()
	fmt.Println(t1)
```

如上例中，通过reflect.Value的Type方法就可以获得对应的reflect.Type。


