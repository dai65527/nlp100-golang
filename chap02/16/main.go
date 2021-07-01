/*
16. ファイルをN分割する
自然数Nをコマンドライン引数などの手段で受け取り，入力のファイルを行単位でN分割せよ．同様の処理をsplitコマンドで実現せよ．
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type suffix [2]byte

func (suf *suffix) inc() {
	if suf[1] == 'z' {
		suf[0]++
		suf[1] = 'a'
		return
	}
	suf[1]++
}

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

	files := make([]*os.File, n)
	suf := suffix{'a', 'a'}
	for i := 0; i < n; i++ {
		files[i], err = os.Create("mine" + string(suf[:]))
		if err != nil {
			fmt.Fprintln(os.Stderr, "os.Create:", err)
		}
		suf.inc()
	}
	defer func() {
		for i := 0; i < n; i++ {
			files[i].Close()
			if err != nil {
				fmt.Fprintln(os.Stderr, "os.File.Close:", err)
			}
		}
	}()

	scanner := bufio.NewScanner(os.Stdin)
	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	l := len(lines) / n
	if len(lines)%n != 0 {
		l++
	}

	for i := 0; i < l; i++ {
		for j := 0; j < n; j++ {
			if i+l*j < len(lines) {
				fmt.Fprintln(files[j], lines[i+l*j])
			}
		}
	}
}
