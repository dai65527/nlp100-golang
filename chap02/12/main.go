/*
12. 1列目をcol1.txtに，2列目をcol2.txtに保存
各行の1列目だけを抜き出したものをcol1.txtに，2列目だけを抜き出したものをcol2.txtとしてファイルに保存せよ．確認にはcutコマンドを用いよ．
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("../popular-names.txt")
	if err != nil {
		fmt.Fprintln(os.Stderr, "os.Open:", err)
		os.Exit(1)
	}
	defer file.Close()

	col1txt, err := os.Create("col1.txt")
	if err != nil {
		fmt.Fprintln(os.Stderr, "os.Create:", err)
		os.Exit(1)
	}
	defer func() {
		if err := col1txt.Close(); err != nil {
			fmt.Fprintln(os.Stderr, "os.File.Close", err)
			os.Exit(1)
		}
	}()

	col2txt, err := os.Create("col2.txt")
	if err != nil {
		fmt.Fprintln(os.Stderr, "os.Create:", err)
		os.Exit(1)
	}
	defer func() {
		if err := col2txt.Close(); err != nil {
			fmt.Fprintln(os.Stderr, "os.File.Close", err)
			os.Exit(1)
		}
	}()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		words := strings.Fields(scanner.Text())
		if len(words) < 2 {
			break
		}
		fmt.Fprintln(col1txt, words[0])
		fmt.Fprintln(col2txt, words[1])
	}
}
