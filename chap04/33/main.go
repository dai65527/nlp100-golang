/*
33. 「AのB」
2つの名詞が「の」で連結されている名詞句を抽出せよ．
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
		for j := 0; j < len(neko[i])-2; j++ {
			if neko[i][j].Pos == "名詞" &&
				neko[i][j+1].Surface == "の" &&
				neko[i][j+2].Pos == "名詞" {
				res = append(res, neko[i][j].Surface+
					neko[i][j+1].Surface+
					neko[i][j+2].Surface)
			}
		}
	}

	for i := range res {
		fmt.Println(res[i])
	}
}
