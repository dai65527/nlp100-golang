/*
19. 各行の1コラム目の文字列の出現頻度を求め，出現頻度の高い順に並べるPermalink
各行の1列目の文字列の出現頻度を求め，その高い順に並べて表示せよ．確認にはcut, uniq, sortコマンドを用いよ
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type Line struct {
	Name  string
	Count int
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	names := []string{}
	for scanner.Scan() {
		words := strings.Fields(scanner.Text())
		if len(words) < 1 {
			continue
		}
		names = append(names, words[0])
	}

	sort.Slice(names, func(i, j int) bool {
		return names[i] < names[j]
	})

	lines := []Line{}
	for i := 0; i < len(names); {
		count := 0
		var j int
		for j = i; j < len(names) && names[i] == names[j]; j++ {
			count++
		}
		lines = append(lines, Line{
			Name:  names[i],
			Count: count,
		})
		i = j
	}

	sort.SliceStable(lines, func(i, j int) bool {
		return lines[i].Count > lines[j].Count
	})

	for i := range lines {
		fmt.Println(lines[i].Name)
	}
}
