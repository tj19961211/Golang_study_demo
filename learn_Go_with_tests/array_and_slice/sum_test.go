package array

import (
	"reflect"
	"testing"
)

/*
数组的容量是我们在声明它时指定的固定值。我们可以通过两种方式初始化数组：

[N]type{value1, value2, ..., valueN} e.g. numbers := [5]int{1, 2, 3, 4, 5}
[...]type{value1, value2, ..., valueN} e.g. numbers := [...]int{1, 2, 3, 4, 5}

在错误信息中打印函数的输入有时很有用。我们使用 %v（默认输出格式）占位符来打印输入，它非常适用于展示数组。
*/

/*
一个源码版本控制的小贴士

此时如果你正在使用源码的版本控制工具（你应该使用它！），我会在此刻先提交一次代码。因为我们已经拥有了一个有测试支持的程序。
但我 不会 将它推送到远程的 master 分支，因为我马上就会重构它。在此时提交一次代码是一种很好的习惯。因为你可以在之后重构导致的代码乱掉时回退到当前版本。

你总是能够回到这个可用的版本。
*/

/*
质疑测试的价值是非常重要的。

测试并不是越多越好，而是尽可能的使你的代码更加健壮。太多的测试会增加维护成本，因为 维护每个测试都是需要成本的。
*/

/*
Go 有内置的计算测试 覆盖率的工具，它能帮助你发现没有被测试过的区域。

我们不需要追求 100% 的测试覆盖率，它只是一个供你获取测试覆盖率的方式。

只要你严格遵循 TDD 规范，那你的测试覆盖率就会很接近 100%。
*/
func TestSum(t *testing.T) {
	t.Run("test version 1 to sum for numbers", func(t *testing.T) {
		numbers := [5]int{1, 2, 3, 4, 5}
		//numbers2 := [5]int{1, 2, 3, 4, 5}
		got := Sum_v1(numbers)
		got2 := SumV1Refactor(numbers)
		want := 15

		if got != want || got2 != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6, 6, 7, 8, 9}
		got := Sum_v2(numbers)
		want := 51

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

}

/*
在 Go 中不能对切片使用等号运算符。

你可以写一个函数迭代每个元素来检查它们的值。

但是一种比较简单的办法是使用 reflect.DeepEqual，它在判断两个变量是否相等时十分有用。

需要注意的是 reflect.DeepEqual 不是「类型安全」的，所以有时候会发生比较怪异的行为。

为了看到这种行为，暂时将测试修改为： want := "bob"

这里我们尝试比较 slice 和 string。这显然是不合理的，但是却通过了编译！

注意：在使用 reflect 包的时候需要注意，不安全类型在 go 编译时是不会进行类型判断检查的

所以使用 reflect.DeepEqual 比较简洁但是在使用时需多加小心
*/

/*
顺便说一下，切片有容量的概念。如果你有一个容量为 2 的切片，但使用 mySlice[10]=1 进行赋值，会报运行时错误。

不过你可以使用 append 函数，它能为切片追加一个新值。
*/
func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSumAllTails(t *testing.T) {

	checkSums := func(t *testing.T, got, want []int) {
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("make the sums of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}

		checkSums(t, got, want)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}

		checkSums(t, got, want)
	})

}

/*
总结

我们学习了：

	数组
	切片
		多种方式的切片初始化
		切片的容量是 固定 的，但是你可以使用 append 从原来的切片中创建一个新切片
		如何获取部分切片
	使用 len 获取数组和切片的长度
	使用测试代码覆盖率的工具 -----> go test -cover
	reflect.DeepEqual 的妙用和对代码类型安全性的影响

数组和切片的元素可以是任何类型，包括数组和切片自己。如果需要你可以定义 [][]string 。

Go 官网博客中关于切片的文章(https://blog.golang.org/slices-intro) 可以让你更加深入的了解切片。尝试写更多的测试来从中学到东西
*/
