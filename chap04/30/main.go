/*
30. 形態素解析結果の読み込みPermalink
形態素解析結果（neko.txt.mecab）を読み込むプログラムを実装せよ．ただし，各形態素は表層形（surface），基本形（base），品詞（pos），品詞細分類1（pos1）をキーとするマッピング型に格納し，1文を形態素（マッピング型）のリストとして表現せよ．第4章の残りの問題では，ここで作ったプログラムを活用せよ．

dai65527注) プログラムはmecabパッケージとして作成
*/

package main

import (
	"fmt"
	"nlp100c4/mecab"
)

func main() {
	res, err := mecab.ParseFile("../neko.txt.mecab")
	if err != nil {
		panic(err)
	}

	for i, r := range res[2] {
		fmt.Println("i =", i)
		fmt.Println("Surface:", r.Surface)
		fmt.Println("Base:   ", r.Base)
		fmt.Println("Pos:    ", r.Pos)
		fmt.Println("Pos1:   ", r.Pos1)
	}
}
