/*
03. 円周率
“Now I need a drink, alcoholic of course, after the heavy lectures involving quantum mechanics.”という文を単語に分解し，各単語の（アルファベットの）文字数を先頭から出現順に並べたリストを作成せよ．
*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "Now I need a drink, alcoholic of course, after the heavy lectures involving quantum mechanics."

	strs := strings.FieldsFunc(s, func(r rune) bool {
		return (r == ' ' || r == ',' || r == '.')
	})

	lens := []int{}
	for _, str := range strs {
		lens = append(lens, len(str))
	}
	fmt.Println(lens)
}
