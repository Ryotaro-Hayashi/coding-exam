package leetcode

import "sort"

// sort + binary search + cache
func smallerNumbersThanCurrent(nums []int) []int {
	arr := make([]int, len(nums))
	cache := make(map[int]int, len(nums))

	sortedNums := make([]int, len(nums))
	copy(sortedNums, nums)
	sort.Ints(sortedNums)

	for i, num := range nums {
		// キャッシュを利用
		if value, ok := cache[num]; ok {
			arr[i] = value
			continue
		}

		// ソートされたスライスのインデックスを返せば良い
		arr[i] = sort.Search(len(nums), func(j int) bool {
			return num <= sortedNums[j]
		})

		cache[num] = arr[i] // キャッシュをセット
	}

	return arr
}

// Brute force with cache
//func smallerNumbersThanCurrent(nums []int) []int {
//	arr := make([]int, len(nums))
//	cache := make(map[int]int, len(nums))
//	for i, num := range nums {
//		// キャッシュに同じ値があれば利用
//		if value, ok := cache[num]; ok {
//			arr[i] = value
//			continue
//		}
//
//		for _, otherNum := range nums {
//			if num > otherNum {
//				arr[i]++
//			}
//		}
//
//		cache[num] = arr[i] // キャッシュにセット
//	}
//
//	return arr
//}
