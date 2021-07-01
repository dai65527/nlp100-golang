/*
13. col1.txtとcol2.txtをマージ
12で作ったcol1.txtとcol2.txtを結合し，元のファイルの1列目と2列目をタブ区切りで並べたテキストファイルを作成せよ．確認にはpasteコマンドを用いよ．
*/

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	outfile, err := os.Create("merged.txt")
	if err != nil {
		fmt.Fprintln(os.Stderr, "os.Create:", err)
		os.Exit(1)
	}
	defer func() {
		if err := outfile.Close(); err != nil {
			fmt.Fprintln(os.Stderr, "os.File.Close:", err)
			os.Exit(1)
		}
	}()

	col1txt, err := os.Open("col1.txt")
	if err != nil {
		fmt.Fprintln(os.Stderr, "os.Open:", err)
		os.Exit(1)
	}
	defer col1txt.Close()

	col2txt, err := os.Open("col2.txt")
	if err != nil {
		fmt.Fprintln(os.Stderr, "os.Open:", err)
		os.Exit(1)
	}
	defer col2txt.Close()

	scanner1 := bufio.NewScanner(col1txt)
	scanner2 := bufio.NewScanner(col2txt)
	for scanner1.Scan() && scanner2.Scan() {
		if len(scanner1.Text()) > 0 {
			fmt.Fprint(outfile, scanner1.Text(), "\t")
		}
		if len(scanner2.Text()) > 0 {
			fmt.Fprintln(outfile, scanner2.Text())
		} else {
			fmt.Fprintln(outfile)
		}
	}
}
