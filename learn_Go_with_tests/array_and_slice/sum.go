package array

/*
数组和它的类型

数组有一个有趣的属性，它的大小也属于类型的一部分，如果你尝试将 [4]int 作为 [5]int 类型的参数传入函数，是不能通过编译的。它们是不同的类型，就像尝试将 string 当做 int 类型的参数传入函数一样。
因为这个原因，所以数组比较笨重，大多数情况下我们都不会使用它。

Go 的切片（slice）类型不会将集合的长度保存在类型中，因此它的尺寸可以是不固定的。

*/

func Sum_v1(numbers [5]int) (sum int) {
	for i := 0; i < 5; i++ {
		sum += numbers[i]
	}
	return
}

// sum_v1 重构版
func SumV1Refactor(numbers [5]int) (sum int) {
	for _, i := range numbers {
		sum += i
	}
	return
}

func Sum_v2(numbers []int) (sum int) {
	for _, i := range numbers {
		sum += i
	}
	return
}

func SumAll(numberToSum ...[]int) (sums []int) {
	//lenNumber := len(numberToSum)
	//sums = make([]int, lenNumber)

	for _, numbers := range numberToSum {
		//sums[i] = Sum_v2(numbers)            // 会有越界报错的情况
		sums = append(sums, Sum_v2(numbers)) // 重构后修改的方案
	}
	return
}

/*
我们可以使用语法 slice[low:high] 获取部分切片。

如果在冒号的一侧没有数字就会一直取到最边缘的元素。

在我们的函数中，我们使用 numbers[1:] 取到从索引 1 到最后一个元素。

由于以下代码需要做传入 slice 不够切片的情况处理，所以重构后如下
*/
func SumAllTails(numberToSum ...[]int) (sums []int) {
	for _, numbers := range numberToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum_v2(tail))
		}

	}
	return
}
