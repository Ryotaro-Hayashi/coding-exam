package others

func winnerPairOfCards(player1 []string, player2 []string) string {
	var max int
	for i, card := range player1 {
		//log.Println(i, card)
		num := card[1:]
		for _, otherCard := range player1[i+1:] {
			otherNum := otherCard[1:]
			if num == otherNum {
				max++
			}
		}
	}
	return "player1"
}
