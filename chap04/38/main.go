/*
38. ヒストグラムPermalink
単語の出現頻度のヒストグラムを描け．ただし，横軸は出現頻度を表し，1から単語の出現頻度の最大値までの線形目盛とする．縦軸はx軸で示される出現頻度となった単語の異なり数（種類数）である．
*/

package main

import (
	"nlp100c4/mecab"

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

	words := map[string]int{}
	for i := 0; i < len(neko); i++ {
		for j := 0; j < len(neko[i]); j++ {
			if neko[i][j].Pos == "記号" ||
				neko[i][j].Pos == "助動詞" ||
				neko[i][j].Pos == "助詞" ||
				neko[i][j].Base == "*" {
				continue
			}
			words[neko[i][j].Base]++
		}
	}

	max := 0
	freq := plotter.Values{}
	for _, v := range words {
		freq = append(freq, float64(v))
		if v > max {
			max = v
		}
	}

	pl := plot.New()
	if err != nil {
		panic(err)
	}

	hist, err := plotter.NewHist(freq, max)
	if err != nil {
		panic(err)
	}
	hist.Normalize(1)

	pl.Add(hist)
	pl.Title.Text = "Word Frequency"
	pl.X.Label.Text = "Words"
	pl.X.Label.Position = draw.PosCenter
	pl.Y.Label.Text = "Counts"
	pl.Y.Label.Position = draw.PosCenter
	pl.X.Min = 0
	pl.X.Max = 50

	if err := pl.Save(15*vg.Centimeter, 10*vg.Centimeter, "graph.png"); err != nil {
		panic(err)
	}
}
