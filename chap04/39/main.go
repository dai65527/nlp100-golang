/*
39. Zipfの法則
単語の出現頻度順位を横軸，その出現頻度を縦軸として，両対数グラフをプロットせよ．
*/

package main

import (
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
			if neko[i][j].Pos == "記号" ||
				neko[i][j].Pos == "助動詞" ||
				neko[i][j].Pos == "助詞" ||
				neko[i][j].Base == "*" {
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

	points := plotter.XYs{}
	for i, key := range keys {
		points = append(points, plotter.XY{
			X: float64(i + 1),
			Y: float64(words[key]),
		})
	}

	pl := plot.New()
	if err != nil {
		panic(err)
	}

	scatter, err := plotter.NewScatter(points)
	if err != nil {
		panic(err)
	}

	pl.Add(scatter)
	pl.Title.Text = "Word Frequency"
	pl.X.Label.Text = "Frequency Rank"
	pl.X.Label.Position = draw.PosCenter
	pl.Y.Label.Text = "Counts"
	pl.Y.Label.Position = draw.PosCenter
	pl.X.Scale = plot.LogScale{}
	pl.Y.Scale = plot.LogScale{}

	if err := pl.Save(15*vg.Centimeter, 10*vg.Centimeter, "graph.png"); err != nil {
		panic(err)
	}
}
