package leetcode

func runningSum(nums []int) []int {
	// nums[0]はそのまま
	// 無駄な変数を用意するとパフォーマンスに影響する
	for i:=1; i<len(nums); i++ {
		nums[i] += nums[i-1]
	}
	return nums
}

//func runningSum(nums []int) []int {
//	runningSum := make([]int, len(nums))
//	var sum int
//	for i, num := range nums {
//		sum += num
//		runningSum[i] = sum
//	}
//
//	return runningSum
//}
