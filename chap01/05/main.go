/*
05. n-gram
与えられたシーケンス（文字列やリストなど）からn-gramを作る関数を作成せよ．この関数を用い，”I am an NLPer”という文から単語bi-gram，文字bi-gramを得よ．
*/

package main

import (
	"fmt"
	"strings"
	"unicode"
)

func letterNgram(s string, n uint) []string {
	r := []rune(s)
	l := uint(len(r))

	res := []string{}
	for i := uint(0); i+n <= l; i++ {
		res = append(res, string(r[i:i+n]))
	}
	return res
}

func wordNgram(s string, n uint) [][]string {
	words := strings.FieldsFunc(s, func(r rune) bool {
		return (unicode.IsSpace(r) || r == '.' || r == ',')
	})

	l := uint(len(words))
	res := [][]string{}
	for i := uint(0); i+n <= l; i++ {
		res = append(res, words[i:i+n])
	}
	return res
}

func main() {
	s := "I am an NLPer"

	letter := letterNgram(s, 2)
	fmt.Printf("letter: ")
	printstrs(letter)
	fmt.Println()

	word := wordNgram(s, 2)
	fmt.Printf("  word: [ ")
	for _, str := range word {
		printstrs(str)
		fmt.Print(", ")
	}
	fmt.Println("]")
}

func printstrs(strs []string) {
	fmt.Print("[ ")
	for _, s := range strs {
		fmt.Printf("\"%s\", ", s)
	}
	fmt.Print("]")
}
