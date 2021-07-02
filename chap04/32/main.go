/*
32. 動詞の基本形
動詞の基本形をすべて抽出せよ．
*/

package main

import (
	"fmt"
	"nlp100c4/mecab"
)

func main() {
	neko, err := mecab.ParseFile("../neko.txt.mecab")
	if err != nil {
		panic(err)
	}

	res := mecab.FindPosFromText("動詞", neko)
	fmt.Println("Found ", len(res))

	for i := 0; i < 10; i++ {
		fmt.Println(res[i].Base)
	}
	fmt.Println(".\n.\n.")
	for i := 3; i > 0; i-- {
		fmt.Println(res[len(res)-i].Base)
	}
}
