package leetcode

// strconvでint→stringにして各桁を取り出して計算する方法もある
func subtractProductAndSum(n int) int {
	sum, product := 0, 1
	for n > 0 {
		surplus := n % 10

		product *= surplus
		sum += surplus

		n /= 10
	}

	return product - sum
}
