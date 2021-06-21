package struct_func_and_interface

import "math"

/*
什么是方法？

到目前为止我们只编写过函数但是我们已经使用过方法。当我们调用 t.Errorf 时我们调用了 t(testing.T) 这个实例的方法 ErrorF。

方法和函数很相似但是方法是通过一个特定类型的实例调用的。函数可以随时被调用，比如 Area(rectangle)。不像方法需要在某个事物上调用。

*/

/*
稍等，什么？

这种定义 interface 的方式与大部分其他编程语言不同。通常接口定义需要这样的代码 My type Foo implements interface Bar。

但是在我们的例子里，

 - Rectangle 有一个返回值类型为 float64 的方法 Area，所以它满足接口 Shape

 - Circle 有一个返回值类型为 float64 的方法 Area，所以它满足接口 Shape

 - string 没有这种方法，所以它不满足这个接口

  等等

在 Go 语言中 interface resolution 是隐式的。如果传入的类型匹配接口需要的，则编译正确。

解耦

请注意我们的辅助函数是怎样实现不需要关心参数是矩形，圆形还是三角形的。通过声明一个接口，辅助函数能从具体类型解耦而只关心方法本身需要做的工作。

这种方法使用接口来声明我们仅仅需要的。这种方法在软件设计中非常重要，我们以后在后续部分中还是涉及到更多细节。
*/

func (r *Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func (r *Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (c *Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c *Triangle) Area() float64 {
	return (c.Base * c.Height) * 0.5
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	Base   float64
	Height float64
}

//这样我们就像创建 Rectangle 和 Circle 一样创建了一个新类型，不过这次是 interface 而不是 struct
type Shape interface {
	Area() float64
}
