package leetcode

func numberOfSteps(num int) int {
	var count int
	number := num
	for number > 0 { // for文でwhile文風に条件を書く
		if number % 2 == 0 { // 偶数のとき
			number /= 2
		} else { // 奇数のとき
			number--
		}
		count++
	}

	return count
}
