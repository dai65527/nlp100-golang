/*
02. 「パトカー」＋「タクシー」＝「パタトクカシーー」Permalink
「パトカー」＋「タクシー」の文字を先頭から交互に連結して文字列「パタトクカシーー」を得よ．
https://nlp100.github.io/ja/ch01.html
*/

package main

import "fmt"

func main() {
	s1 := "パトカー"
	s2 := "タクシー"

	r := []rune{}
	r1 := []rune(s1)
	r2 := []rune(s2)
	for i := 0; i < len(r1) && i < len(r2); i++ {
		r = append(r, r1[i])
		r = append(r, r2[i])
	}
	fmt.Println(string(r))
}
