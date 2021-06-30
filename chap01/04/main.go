/*
04. 元素記号
“Hi He Lied Because Boron Could Not Oxidize Fluorine. New Nations Might Also Sign Peace Security Clause. Arthur King Can.”という文を単語に分解し，1, 5, 6, 7, 8, 9, 15, 16, 19番目の単語は先頭の1文字，それ以外の単語は先頭の2文字を取り出し，取り出した文字列から単語の位置（先頭から何番目の単語か）への連想配列（辞書型もしくはマップ型）を作成せよ．
https://nlp100.github.io/ja/ch01.html
*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "Hi He Lied Because Boron Could Not Oxidize Fluorine. New Nations Might Also Sign Peace Security Clause. Arthur King Can."

	strs := strings.Fields(s)

	tb := map[string]int{}
	for i, str := range strs {
		tb[getSymbol(i+1, str)] = i + 1
	}
	fmt.Println(tb)
}

func getSymbol(idx int, s string) string {
	if isOne(idx) {
		return string(([]rune(s))[:1])
	}
	return string(([]rune(s))[:2])
}

func isOne(n int) bool {
	a := []int{1, 5, 6, 7, 8, 9, 15, 16, 19}
	for _, ai := range a {
		if n == ai {
			return true
		}
	}
	return false
}
