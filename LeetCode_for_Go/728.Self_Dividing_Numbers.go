package main

func selfDividingNumbers(left int, right int) []int {
	ret := []int{}
	var tmp int
	for x := left; x <= right; x++ {
		for tmp = x; tmp > 0; tmp = tmp / 10 {
			s := tmp % 10
			if s == 0 || x%s != 0 {
				break
			}
		}
		if tmp == 0 {
			ret = append(ret, x)
		}
	}
	return ret
}
