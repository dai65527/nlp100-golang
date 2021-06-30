/*
00. 文字列の逆順
文字列”stressed”の文字を逆に（末尾から先頭に向かって）並べた文字列を得よ．
https://nlp100.github.io/ja/ch01.html
*/

package main

import "fmt"

func main() {
	s := "stressed"

	b := []byte(s)
	for i := 0; i < len(b)/2; i++ {
		b[i], b[len(b)-i-1] = b[len(b)-i-1], b[i]
	}

	s = string(b)
	fmt.Println(s)
}
