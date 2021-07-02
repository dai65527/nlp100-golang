/*
35. 単語の出現頻度Permalink
文章中に出現する単語とその出現頻度を求め，出現頻度の高い順に並べよ．
*/

package main

import (
	"fmt"
	"nlp100c4/mecab"
	"sort"
)

func main() {
	neko, err := mecab.ParseFile("../neko.txt.mecab")
	if err != nil {
		panic(err)
	}

	keys := []string{}
	words := map[string]int{}
	for i := 0; i < len(neko); i++ {
		for j := 0; j < len(neko[i]); j++ {
			if neko[i][j].Pos != "名詞" || neko[i][j].Base == "*" {
				continue
			}
			words[neko[i][j].Base]++
			if words[neko[i][j].Base] == 1 {
				keys = append(keys, neko[i][j].Base)
			}
		}
	}

	sort.Slice(keys, func(i, j int) bool {
		return words[keys[i]] > words[keys[j]]
	})

	for i := 0; i < 10; i++ {
		fmt.Println(keys[i], words[keys[i]])
	}
}
