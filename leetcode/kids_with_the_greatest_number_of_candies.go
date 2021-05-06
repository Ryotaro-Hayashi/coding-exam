package leetcode

func kidsWithCandies(candies []int, extraCandies int) []bool {
	var max int
	for _, candy := range candies { // 最高値を算出
		if candy > max {
			max = candy
		}
	}

	boolSlice := make([]bool, len(candies))
	for i, candy := range candies {
		if candy + extraCandies >= max {
			boolSlice[i] = true
		}
	}

	return boolSlice
}
