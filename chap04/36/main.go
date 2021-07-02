/*
36. 頻度上位10語
出現頻度が高い10語とその出現頻度をグラフ（例えば棒グラフなど）で表示せよ．
*/

package main

import (
	"fmt"
	"nlp100c4/mecab"
	"sort"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
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

	pl := plot.New()
	if err != nil {
		panic(err)
	}

	plotValues := plotter.Values{}
	for i := 0; i < 10; i++ {
		plotValues = append(plotValues, float64(words[keys[i]]))
	}

	bar, err := plotter.NewBarChart(plotValues, vg.Points(15))
	if err != nil {
		panic(err)
	}

	pl.Add(bar)
	pl.NominalX([]string{"Word1", "Word2", "Word3", "Word4", "Word5", "Word6", "Word7", "Word8", "Word9", "Word10"}...) // 日本語がプロットできない...
	fmt.Println(keys[:10])

	pl.Title.Text = "Freq. Words"
	pl.Y.Label.Text = "Count"
	pl.X.Label.Position = draw.PosCenter

	if err := pl.Save(15*vg.Centimeter, 10*vg.Centimeter, "graph.png"); err != nil {
		panic(err)
	}
}
