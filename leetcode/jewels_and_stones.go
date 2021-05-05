package leetcode

func numJewelsInStones(jewels string, stones string) int {
	var count int
	jewelMap := make(map[rune]bool, 52)
	for _, jewelRune := range jewels {
		jewelMap[jewelRune] = true
	}

	for _, stoneRune := range stones {
		if jewelMap[stoneRune] {
			count++
		}
	}

	return count
}

// 計算量がn^2
//func numJewelsInStones(jewels string, stones string) int {
//	var count int
//	for _, jewel := range jewels {
//		for _, stone := range stones {
//			if string(jewel) == string(stone) {
//				count++
//			}
//		}
//	}
//
//	return count
//}
