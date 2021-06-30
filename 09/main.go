/*
09. TypoglycemiaPermalink
スペースで区切られた単語列に対して，各単語の先頭と末尾の文字は残し，それ以外の文字の順序をランダムに並び替えるプログラムを作成せよ．ただし，長さが４以下の単語は並び替えないこととする．適当な英語の文（例えば”I couldn’t believe that I could actually understand what I was reading : the phenomenal power of the human mind .”）を与え，その実行結果を確認せよ．
*/

package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	str := "I couldn’t believe that I could actually understand what I was reading : the phenomenal power of the human mind ."

	words := strings.Fields(str)
	for i := range words {
		words[i] = shuffle(words[i])
	}
	fmt.Println(strings.Join(words, " "))
}

func shuffle(s string) string {
	r := []rune(s)
	if len(r) > 4 {
		sub := r[1 : len(r)-1]
		rand.Shuffle(len(sub), func(i, j int) { sub[i], sub[j] = sub[j], sub[i] })
	}
	return string(r)
}
