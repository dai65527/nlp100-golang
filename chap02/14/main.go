/*
14. 先頭からN行を出力
自然数Nをコマンドライン引数などの手段で受け取り，入力のうち先頭のN行だけを表示せよ．確認にはheadコマンドを用いよ．
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

	scanner := bufio.NewScanner(os.Stdin)
	for ; n > 0 && scanner.Scan(); n-- {
		fmt.Println(scanner.Text())
	}
}
