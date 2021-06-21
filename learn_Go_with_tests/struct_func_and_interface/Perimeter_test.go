package struct_func_and_interface

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := &Rectangle{10.0, 10.0}
	got := rectangle.Perimeter()
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {

	// 像其他练习一样我们创建了一个辅助函数，但不同的是我们传入了一个 Shape 类型。
	checkArea := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	}

	t.Run("rectangles", func(t *testing.T) {
		rectangle := &Rectangle{12.0, 6.0}
		checkArea(t, rectangle, 72.0)
		//got := rectangle.Area()
		// want := 100.0

		// if got != want {
		// 	t.Errorf("got %.2f want %.2f", got, want)
		// }
	})

	t.Run("circles", func(t *testing.T) {
		circle := &Circle{Radius: 10}
		checkArea(t, circle, 314.1592653589793)
		// got := circle.Area()
		// want := 314.16

		// if got != want {
		// 	t.Errorf("got %.2f want %.2f", got, want)
		// }
	})
}

/*
现在我们对结构体有一定的理解了，我们可以引入「表格驱动测试」。

表格驱动测试(https://github.com/golang/go/wiki/TableDrivenTests)在我们要创建一系列相同测试方式的测试用例时很有用。
*/

/*
列表驱动测试可以成为你工具箱中的得力武器。

但是确保你在测试中真的需要使用它。

如果你要测试一个接口的不同实现，或者传入函数的数据有很多不同的测试需求，这个武器将非常给力。
*/

/*

 */

func TestArea2(t *testing.T) {
	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{&Rectangle{12, 6}, 72.0},
		{&Circle{10}, 314.1592653589793},
		{&Triangle{12, 6}, 36.0},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("got %.2f want %.2f", got, tt.want)
		}
	}
}

/*
下面是满足要求的最终测试代码：
*/
func TestArea3(t *testing.T) {
	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Circle", shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
		{name: "Rectangle", shape: Rectangle{Radius: 12, Height: 6}, hasArea: 72.0},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, hasArea: 36.0},
	}

	for _, tt := range areaTests {
		// using tt.name from the case to use it as the `t.Run` test name
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("%#v got %.2f want %.2f", tt.shape, got, tt.hasArea)
			}
		})
	}
}

/*
总结

这是进一步的 TDD 实践。我们在对一个基本数学问题的解决方案的迭代中，通过测试学习了语言的新特性。

 - 声明结构体以创建我们自己的类型，让我们把数据集合在一起并达到简化代码的目地

 - 声明接口，这样我们可以定义适合不同参数类型的函数（参数多态）

 - 在自己的数据类型中添加方法以实现接口

 - 列表驱动测试让断言更清晰，这样可以使测试文件更易于扩展和维护

这是重要的一课。因为我们开始定义自己的类型。在像 Go 这样的静态语言中，能定义自己的类型是开发易维护，低耦合，好测试的软件的基础。

接口是把负责从系统的其他部分隐藏起来的伟大工具。在我们的测试中，辅助函数的代码不需要知道具体的几何形状，只需要知道获取它的面积即可。
*/

/*
通用辅助函数的构建：

创建 匿名 struct 列表 (struct 内的结构基本一致，不一致的使用 interface 区分，通常 interface 中接收的是不同类型的 struct) --->

使用 for range 结构迭代调用需要测试的代码
*/
