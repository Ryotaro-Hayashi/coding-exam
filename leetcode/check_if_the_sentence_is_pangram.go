package leetcode

/*
	A pangram is a sentence where every letter of the English alphabet appears at least once.

	Given a string sentence containing only lowercase English letters, return true if sentence is a pangram, or false otherwise.
 */

func checkIfPangram(sentence string) bool {
	var alphabetCount int
	alphabetASCII := make([]int, 26) // alphabetは全部で26個
	for _, word := range sentence {
		i := int(word - 'a') // aのASCIIを0に調整
		if alphabetASCII[i] == 0 { // 新しいalphabetを発見
			alphabetCount++ // 新しく発見したalphabetをカウント
			alphabetASCII[i] = i // 発見したalphabetはスライスに保存
 		}
	}
	return alphabetCount == 26
}

// mapを使って無理やり
//func checkIfPangram(sentence string) bool {
//	alphabetMap := map[string]struct{}{
//		"a": {}, "b": {}, "c": {}, "d": {}, "e": {}, "f": {}, "g": {}, "h": {}, "i": {}, "j": {}, "k": {}, "l": {}, "m": {}, "n": {},
//		"o": {}, "p": {}, "q": {}, "r": {}, "s": {}, "t": {}, "u": {}, "v": {}, "w": {}, "x": {}, "y": {}, "z": {},
//	}
//
//	for _, word := range sentence {
//		if _, ok := alphabetMap[string(word)]; ok {
//			delete(alphabetMap, string(word))
//		}
//	}
//
//	return len(alphabetMap) == 0
//}
