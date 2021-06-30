/*
10. 行数のカウントPermalink
行数をカウントせよ．確認にはwcコマンドを用いよ．
*/

package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// file, err := os.Open("../popular-names.txt")
	data, err := ioutil.ReadFile("../popular-names.txt")
	if err != nil {
		fmt.Fprintln(os.Stderr, "os.Open:", err)
		os.Exit(1)
	}

	count := 0
	for _, c := range data {
		if c == '\n' {
			count++
		}
	}
	fmt.Println(count)
}
