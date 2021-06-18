package iteration

//在 Go 中 for 用来循环和迭代，Go 语言没有 while，do，until 这几个关键字，你只能使用 for。

// func Repeat(character string) string {
// 	return ""
// }

// 使用迭代的方式返回五个相同的 character
func Repeat(character string) string {
	var repeated string
	for i := 0; i < 5; i++ {
		repeated += character // += 是自增赋值运算符（Add AND assignment operator），它把运算符右边的值加到左边并重新赋值给左边。
	}
	return repeated
}
