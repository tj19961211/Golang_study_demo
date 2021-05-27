package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

//第一题
//主函数
func findtarget(x int, nums []int) []int {
	//处理slice长度过小的问题
	if len(nums)-1 < 0 {
		return nil
	}
	right, left := len(nums)-1, 0

	index := binarySearch(nums, x, left, right)
	//处理传入的 x 不存在slice里面
	if index != x {
		return nil
	}
	return more(index, nums)
}

//搜索target
func binarySearch(nums []int, target int, loIndex, hiIndex int) int {
	if loIndex > hiIndex {
		return -1
	}

	mid := int((loIndex + hiIndex) / 2)
	if nums[mid] < target {
		return binarySearch(nums, target, mid+1, hiIndex)
	} else if nums[mid] > target {
		return binarySearch(nums, target, loIndex, mid-1)
	} else {
		return mid
	}
}

//对相等index左右与index位置相同的数均进行添加到返回的slice
func more(index int, nums []int) []int {
	var leftIndex = index - 1
	var rightIndex = index + 1
	if leftIndex < 0 || rightIndex > len(nums)-1 {
		return nil
	}
	targetList := make([]int, 0)
	if nums[leftIndex] == nums[index] {
		leftIndex -= 1
	}
	if nums[rightIndex] == nums[index] {
		rightIndex += 1
	}
	for i := leftIndex + 1; i < rightIndex; i++ {
		targetList = append(targetList, i)
	}
	return targetList
}

//------------------------------------------------------------------------------------------------
//第二题
//主函数
func QuerySalesData(ctx context.Context) (int, error) {
	var sum int
	var str string
	userName, ok := ctx.Value("user").([]string)
	str, ok = ctx.Value("url").(string)
	if !ok {
		return 0, nil
	}
	for i := 0; i < len(userName); i++ {
		count, msg, err := getCount(str, userName[i])
		sum += count
		if msg != "" || err != nil {
			return 0, err
		}
	}
	return sum, nil
}

//获取单个Count
func getCount(str, name string) (int, string, error) {
	var respmsg respMsg
	fmt.Println(str + name)
	resp, err := http.Get(str + name)
	if err != nil {
		return 0, "", err
	}

	err = json.NewDecoder(resp.Body).Decode(&respmsg)
	if err != nil {
		return 0, "", err
	}
	if respmsg.Code == 100 || respmsg.Msg != "" {
		fmt.Println(respmsg.Msg)
		return 0, respmsg.Msg, nil
	}
	defer resp.Body.Close()
	return respmsg.Count, "", nil
}

//json返回的数据接收结构
type respMsg struct {
	Code  int    `json:"code"`
	Msg   string `json:"msg"`
	Count int    `json:"count"`
}

func main() {
	//fmt.Println(findtarget(7, []int{1, 2, 2, 4, 5, 5, 6, 6, 9, 9, 10}))

	ctx := context.Background()

	ctx = context.WithValue(ctx, "user", []string{"a", "b", "c", "d", "e"})
	ctx = context.WithValue(ctx, "url", "https://interview.moreless.io/questions/async_workers/sales/")

	fmt.Println(QuerySalesData(ctx))

}
