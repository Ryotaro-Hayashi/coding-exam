package others

import "unicode/utf8"

func hasSameType(user1 string, user2 string) bool {
	// 文字数比較（組み込みのlenじゃダメなの初めて知った）
	if utf8.RuneCountInString(user1) != utf8.RuneCountInString(user2) {
		return false
	}

	for i:= range user1 { // iは, user1のある1文字のindex
		for j := range user1[i+1:] { // jは, 注目文字以外のuser1のある1文字のindex
			// user1のある1文字とその他の文字を全探索で比較
			if user1[i] == user1[i+1+j] { // user1の中で一致しているindexではuser2でも一致しなければいけない
				if user2[i] != user2[i+1+j] {
					return false
				}
			} else { // user1の中で一致していないindexではuser2で一致していてはいけない
				if user2[i] == user2[i+1+j] {
					return false
				}
			}
		}
	}

	return true
}
