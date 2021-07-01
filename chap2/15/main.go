/*
15. 末尾のN行を出力Permalink
自然数Nをコマンドライン引数などの手段で受け取り，入力のうち末尾のN行だけを表示せよ．確認にはtailコマンドを用いよ．
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Invalid number of arguments")
		os.Exit(1)
	}

	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, "strconv.Atoi:", err)
		os.Exit(1)
	}

	lines := make([]string, n)

	scanner := bufio.NewScanner(os.Stdin)
	idx := 0
	for scanner.Scan() {
		lines[idx%n] = scanner.Text()
		idx++
	}

	if idx < n {
		for i := 0; i < idx; i++ {
			fmt.Println(lines[i])
		}
	} else {
		for i := 0; i < n; i++ {
			fmt.Println(lines[idx%n])
			idx++
		}
	}
}
