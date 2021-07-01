/*
17. １列目の文字列の異なり
1列目の文字列の種類（異なる文字列の集合）を求めよ．確認にはcut, sort, uniqコマンドを用いよ．
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	set := map[string]struct{}{}
	for scanner.Scan() {
		words := strings.Fields(scanner.Text())
		if len(words) > 0 {
			set[words[0]] = struct{}{}
		}
	}

	for k := range set {
		fmt.Println(k)
	}
}
