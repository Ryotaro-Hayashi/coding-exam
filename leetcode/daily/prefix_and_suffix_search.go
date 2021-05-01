package daily

import (
	"strings"
)

type WordFilter struct {
	words []string
}


func Constructor(words []string) WordFilter {
	return WordFilter{
		words: words,
	}
}

func (this *WordFilter) F(prefix string, suffix string) int {
	index := -1
	for i, word := range this.words {
		// 前方一致と後方一致
		if strings.HasPrefix(word, prefix) && strings.HasSuffix(word, suffix) && i > index {
			index = i
		}
	}
	return index
}
