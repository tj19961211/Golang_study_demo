package reflect_study

import (
	"reflect"
	"testing"
)

/*
什么是 interface？

  由于函数使用已知的类型，例如 string，int 以及我们自己定义的类型，如 BankAccount，我们享受到了 Go 为我们提供的类型安全。

  这意味着我们可以免费获得一些文档，如果你试图向函数传递错误的类型，编译器就会报错。

  但是，你可能会遇到这样的情况，即你不知道要编写的函数参数在编译时是什么类型的。

  Go 允许我们使用类型 interface{} 来解决这个问题，你可以将其视为 任意 类型。

  所以 walk(x interface{}, fn func(string)) 的 x 参数可以接收任何的值。
*/

/*
那么为什么不通过将所有参数都定义为 interface 类型来得到真正灵活的函数呢？

  - 作为函数的使用者，使用 interface 将失去对类型安全的检查。如果你想传入 string 类型的 Foo.bar 但是传入的是 int 类型的 Foo.baz，编译器将无法通知你这个错误。你也搞不清楚函数允许传递什么类型的参数。知道一个函数接收什么类型，例如 UserService，是非常有用的。

  - 作为这样一个函数的作者，你必须检查传入的 所有 参数，并尝试断定参数类型以及如何处理它们。这是通过 反射 实现的。这种方式可能相当笨拙且难以阅读，而且一般性能比较差（因为程序必须在运行时执行检查）。

简而言之，除非真的需要否则不要使用反射。

如果你想实现函数的多态性，请考虑是否可以围绕接口（不是 interface 类型，这里容易让人困惑）设计它，以便用户可以用多种类型来调用你的函数，这些类型实现了函数工作所需要的任何方法。

我们的函数需要能够处理很多不同的东西。和往常一样，我们将采用迭代的方法，为我们想要支持的每一件新事物编写测试，并一路进行重构，直到完成。
*/

/*
最后一个问题

记住，Go 中的 map 不能保证顺序一致。因此，你的测试有时会失败，因为我们断言对 fn 的调用是以特定的顺序完成的。

为了解决这个问题，我们需要将带有 map 的断言移动到一个新的测试中，在这个测试中我们不关心顺序。
*/

func TestWalk(t *testing.T) {

	cases := []struct {
		Name         string
		Input        interface{}
		ExectedCalls []string
	}{
		{
			"struct with one string field",
			struct {
				Name string
				City string
			}{"Chris", "London"},
			[]string{"Chris", "London"},
		},
		{
			"Newsted fields",
			Person{
				"Chris",
				Profile{33, "London"},
			},
			[]string{"Chris", "London"},
		},
		{
			"Pointers to things",
			&Person{
				"Chris",
				Profile{33, "London"},
			},
			[]string{"Chris", "London"},
		},
		{
			"Slices",
			[]Profile{
				{33, "London"},
				{34, "Beijing"},
			},
			[]string{"London", "Beijing"},
		},
		{
			"Arrays",
			[2]Profile{
				{33, "London"},
				{34, "Beijing"},
			},
			[]string{"London", "Beijing"},
		},
		{
			"Maps",
			map[string]string{
				"Foo":  "Bar",
				"Bar2": "Baz",
			},
			[]string{"Bar", "Boz"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExectedCalls) {
				t.Errorf("got %v , want %v", got, test.ExectedCalls)
			}
		})
	}

	t.Run("With maps", func(t *testing.T) {
		aMap := map[string]string{
			"Foo": "Bar",
			"Baz": "Boz",
		}

		var got []string
		walk(aMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "Bar")
		assertContains(t, got, "Boz")
	})
}

func assertContains(t *testing.T, haystack []string, needle string) {
	contain := false

	for _, x := range haystack {
		if x == needle {
			contain = true
		}
	}

	if !contain {
		t.Errorf("expected %+v to contain '%s' but it didnt", haystack, needle)
	}
}

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

/*
总结

介绍了 reflect 包中的一些概念。

使用递归遍历任意数据结构。

在回顾中做了一个糟糕的重构，但不用对此感到太沮丧。通过迭代地进行测试，这并不是什么大问题。

这只是 reflection 的一个小方面。Go 博客上有一篇精彩的文章介绍了更多细节。(https://blog.golang.org/laws-of-reflection)

现在你已经了解了反射，请尽量避免使用它。
*/
