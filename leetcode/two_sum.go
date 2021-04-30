package leetcode

/*
	Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

	You may assume that each input would have exactly one solution, and you may not use the same element twice.

	You can return the answer in any order.
*/

// Two-pass Hash Table
func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int, len(nums))
	for i, num := range nums {
		numMap[num] = i
	}

	for i, num := range nums {
		if j, ok := numMap[target-num]; ok {
			if i != j {
				return []int{i, j}
			}
		}
	}

	return nil
}

// Brute Force
//func twoSum(nums []int, target int) []int {
//	var returnList []int
//	for i, num := range nums {
//		targetSub := target - num
//		for j, n := range nums[i+1:] {
//			if n == targetSub {
//				returnList = append(returnList, i)
//				returnList = append(returnList, i+1+j)
//				return returnList
//			}
//		}
//	}
//	return returnList
//}
