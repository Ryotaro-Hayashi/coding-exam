package leetcode

import "strings"

func arrayStringsAreEqual(word1 []string, word2 []string) bool {
	return strings.Join(word1, "") == strings.Join(word2, "")
}

//func arrayStringsAreEqual(word1 []string, word2 []string) bool {
//	var string1, string2 string
//	for _, word1 := range word1 {
//		string1 += word1
//	}
//
//	for _, word2 := range word2 {
//		string2 += word2
//	}
//
//	return string1==string2
//}
