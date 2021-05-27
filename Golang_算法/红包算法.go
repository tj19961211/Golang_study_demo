package main

import (
	"fmt"
	"math/rand"
	"time"
)

// ```
// 思路：
// ```
func main() {
	//初始10红包，10元钱
	count, amount := int64(10), int64(10)
	// //将输入的钱乘上100,方便操作
	// amount = amount * 100
	//剩余金额
	remain := amount * 100
	//验证红包算法的总金额，最后sum应该==amount
	sum := int64(0)
	//发红包
	for i := int64(0); i < count; i++ {
		x := DoubleAverage(count-i, remain)
		//剩余金额
		remain -= x
		//发了多少
		sum += x
		//金额转成元
		fmt.Println(i+1, "=", float64(x)/float64(100))
	}
	fmt.Println()
	fmt.Println("总和 ", float64(sum)/float64(100))
}

func DoubleAverage(count int64, amount int64) int64 {
	//最小钱
	min := int64(1)

	if count == 1 {
		return amount
	}

	//计算最大可用金额,min最小是1分钱,减去的min,下面会加上,避免出现0分钱
	max := amount - min*count
	// //计算最大可用平均值
	// avg := max / count
	// //二倍均值基础加上最小金额,防止0出现,作为上限
	// avg2 := 2*avg + min
	// 加 min 防止0出现,作为上限
	avg := max/3 + min
	//随机红包金额序列元素,把二倍均值作为随机的最大数
	rand.Seed(time.Now().UnixNano())
	//加min是为了避免出现0值,上面也减去了min
	x := rand.Int63n(avg) + min
	return x
}
