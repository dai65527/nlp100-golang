/*
01. 「パタトクカシーー」
「パタトクカシーー」という文字列の1,3,5,7文字目を取り出して連結した文字列を得よ
https://play.golang.org/p/ItN4LoSeaNn
*/

package main

import "fmt"

func main() {
	s := "パタトクカシー"

	r := []rune{}
	r = append(r, []rune(s)[0])
	r = append(r, []rune(s)[2])
	r = append(r, []rune(s)[4])
	r = append(r, []rune(s)[6])
	fmt.Println(string(r))
}
