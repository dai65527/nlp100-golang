/*
08. 暗号文
与えられた文字列の各文字を，以下の仕様で変換する関数cipherを実装せよ．

英小文字ならば(219 - 文字コード)の文字に置換
その他の文字はそのまま出力
この関数を用い，英語のメッセージを暗号化・復号化せよ．
*/

package main

import (
	"fmt"
)

func cipher(s string) string {
	res := []byte(s)

	for i := range res {
		if res[i] >= 'a' && res[i] <= 'z' {
			res[i] = 219 - res[i]
		}
	}
	return string(res)
}

func main() {
	original := "I'm D.Nakano. I was born in Noto."
	encrypted := cipher(original)
	again := cipher(encrypted)
	fmt.Println(" original:", original)
	fmt.Println("encrypted:", encrypted)
	fmt.Println("    again:", again)
}
