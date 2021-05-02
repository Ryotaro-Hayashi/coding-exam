package leetcode

/*
	We have n chips, where the position of the ith chip is position[i].

	We need to move all the chips to the same position. In one step, we can change the position of the ith chip from position[i] to:

	position[i] + 2 or position[i] - 2 with cost = 0.
	position[i] + 1 or position[i] - 1 with cost = 1.
	Return the minimum cost needed to move all the chips to the same position.
 */

/*
	偶数だけ動かすのはコストが0なので全てのコインは位置が1or2になる
	位置が1or2のコインの数の小さい方が最小の移動コスト
 */

func minCostToMoveChips(position []int) int {
	// コインの位置が奇数or偶数のカウント
	var oddCount, evenCount int
	for _, number := range position {
		if number % 2 != 0 {
			oddCount++
		} else {
			evenCount++
		}
	}
	// コストが低い方を返す
	if oddCount > evenCount {
		return evenCount
	}
	return oddCount
}
