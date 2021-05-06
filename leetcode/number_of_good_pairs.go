package leetcode

// 計算量O(n)
func numIdenticalPairs(nums []int) int {
	var count int
	seen := make(map[int]int, len(nums)-1) // 対象の数字をキーにして遭遇回数を値に入れるマップ
	for _, num := range nums {
		n, ok := seen[num] // 初見だとn=0, ok=false
		if ok {
			count += n
		}
		seen[num] = n + 1
	}

	return count
}

// 計算量O(n^2)
//func numIdenticalPairs(nums []int) int {
//	var count int
//	for i, num := range nums {
//		for _, otherNum := range nums[i+1:] {
//			if num == otherNum {
//				count++
//			}
//		}
//	}
//
//	return count
//}
