/*
34. 名詞の連接Permalink
名詞の連接（連続して出現する名詞）を最長一致で抽出せよ．
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

	res := []string{}
	for i := 0; i < len(neko); i++ {
		for j := 0; j < len(neko[i]); {
			if neko[i][j].Pos == "名詞" {
				m := ""
				for ; neko[i][j].Pos == "名詞"; j++ {
					m += neko[i][j].Surface
				}
				res = append(res, m)
			} else {
				j++
			}
		}
	}

	for i := range res {
		if len(res[i]) > 25 {
			fmt.Println(res[i])
		}
	}
}
