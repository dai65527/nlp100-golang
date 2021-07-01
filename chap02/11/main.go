/*
11. タブをスペースに置換
タブ1文字につきスペース1文字に置換せよ．確認にはsedコマンド，trコマンド，もしくはexpandコマンドを用いよ．
*/

package main

import (
	"fmt"
	"io"
	"os"
)

const bufsize = 1024

func main() {
	file, err := os.Open("../popular-names.txt")
	if err != nil {
		fmt.Fprintln(os.Stderr, "os.Open:", err)
		os.Exit(1)
	}
	defer file.Close()

	buf := make([]byte, bufsize)
	for {
		n, err := file.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintln(os.Stderr, "file.Read:", err)
			os.Exit(1)
		}
		for i := 0; i < n; i++ {
			if buf[i] == '\t' {
				buf[i] = ' '
			}
		}
		m, err := os.Stdout.Write(buf[:n])
		if err != nil || n != m {
			fmt.Fprintln(os.Stderr, "os.Stdout.Write", err)
			os.Exit(1)
		}
	}
}
