package reflect_study

import "reflect"

/*
这段代码 非常不安全，也非常幼稚，但请记住，当我们处于「红色」状态（测试失败）时，我们的目标是编写尽可能少的代码。然后我们编写更多的测试来解决我们的问题。
我们需要使用反射来查看 x 并尝试查看它的属性。

反射包 (https://godoc.org/reflect) 有一个函数 ValueOf，该函数值返回一个给定变量的 Value。这为我们提供了检查值的方法，包括我们在下一行中使用的字段。
然后我们对传入的值做了一些非常乐观的假设：

 - 我们只看第一个也是唯一的字段，可能根本就没有字段会引起 panic

 - 然后我们调用 String()，它以字符串的形式返回底层值，但是我们知道，如果这个字段不是字符串，程序就会出错。
*/

/* 重构 1： 当你多次对相同的值进行比较时，通常情况下，将代码重构为 switch 会提高可读性，使代码更易于扩展*/

/* 重构 2：

if val.Kind() == reflect.Slice {
        for i:=0; i< val.Len(); i++ {
            walk(val.Index(i).Interface(), fn)
        }
        return
    }

这招很管用，但很恶心。不过不用担心，我们有测试支持的工作代码，所以我们可以随意修改我们喜欢的代码。

    如果你抽象地想一下，我们想要针对下面的对象调用 walk

        - 结构体中的每个字段

        - 切片中的每一项

我们目前的代码可以做到这一点，但反射用得不太好。我们只是在一开始检查它是否是切片（通过 return 来停止执行剩余的代码），如果不是，我们就假设它是 struct。

让我们重新编写代码，先检查类型，再执行我们的逻辑代码。*/

/*
重构 3：


*/

func walk(x interface{}, fn func(input string)) {

	// 合并取值方法使其成为一个函数
	val := getValue(x)

	numberOfValues := 0
	var getField func(int) reflect.Value

	// val := reflect.ValueOf(x)
	//field := val.Field(0)
	//fn("I still can't believe South Korea beat Germany 2-0 to put them last in their group")
	//fn(field.String())

	// 指针类型的 Value 不能使用 NumField 方法，在执行此方法前需要调用 Elem() 提取底层值。
	// if val.Kind() == reflect.Ptr {
	// 	val = val.Elem()
	// }

	switch val.Kind() {
	case reflect.String:
		fn(val.String())
	case reflect.Struct:
		numberOfValues = val.NumField() // 获取 struct 长度
		getField = val.Field
	case reflect.Slice, reflect.Array:
		numberOfValues = val.Len()
		getField = val.Index
	case reflect.Map:
		for _, key := range val.MapKeys() {
			walk(val.MapIndex(key).Interface(), fn)
		}
	}

	for i := 0; i < numberOfValues; i++ {
		walk(getField(i).Interface(), fn)
	}

	// if val.Kind() == reflect.Slice {
	// 	for i := 0; i < val.Len(); i++ {
	// 		walk(val.Index(i).Interface(), fn)
	// 	}
	// 	return
	// }

	// value 有一个方法 NumField，它返回值中的字段数。这让我们遍历字段并调用 fn 通过我们的测试
	//for i := 0; i < val.NumField(); i++ {
	//	field := val.Field(i)
	//
	//	switch field.Kind() {
	//	case reflect.String:
	//		fn(field.String())
	//	case reflect.Struct:
	//		walk(field.Interface(), fn)
	//	}

	// if field.Kind() == reflect.String {
	// 	fn(field.String())
	// }

	// if field.Kind() == reflect.Struct {
	// 	walk(field.Interface(), fn)
	// }

	// fn(field.String())
	//}
}

func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	return val
}
