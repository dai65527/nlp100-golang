/*
18. 各行を3コラム目の数値の降順にソートPermalink
各行を3コラム目の数値の逆順で整列せよ（注意: 各行の内容は変更せずに並び替えよ）．確認にはsortコマンドを用いよ（この問題はコマンドで実行した時の結果と合わなくてもよい）．
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Line struct {
	S string
	N int
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	lines := []Line{}
	for scanner.Scan() {
		words := strings.Fields(scanner.Text())
		var n int
		if len(words) >= 3 {
			n, _ = strconv.Atoi(words[2])
		} else {
			n = 0
		}
		lines = append(lines, Line{
			S: scanner.Text(),
			N: n,
		})
	}

	sort.SliceStable(lines, func(i, j int) bool {
		return lines[i].N > lines[j].N
	})

	for _, line := range lines {
		fmt.Println(line.S)
	}
}
