/*
06. 集合
“paraparaparadise”と”paragraph”に含まれる文字bi-gramの集合を，それぞれ, XとYとして求め，XとYの和集合，積集合，差集合を求めよ．さらに，’se’というbi-gramがXおよびYに含まれるかどうかを調べよ．
*/

package main

import "fmt"

func letterNgram(s string, n uint) []string {
	r := []rune(s)
	l := uint(len(r))

	res := []string{}
	for i := uint(0); i+n <= l; i++ {
		res = append(res, string(r[i:i+n]))
	}
	return res
}

func main() {
	s1 := "paraparaparadise"
	s2 := "paragraph"

	bg1 := letterNgram(s1, 2)
	bg2 := letterNgram(s2, 2)

	x := map[string]struct{}{}
	for _, elem := range bg1 {
		x[elem] = struct{}{}
	}
	fmt.Println("X:", x)

	y := map[string]struct{}{}
	for _, elem := range bg2 {
		y[elem] = struct{}{}
	}
	fmt.Println("Y:", y)

	// 和集合: 少なくとも片方に入っているもの
	wa := map[string]struct{}{}
	for k := range x {
		wa[k] = struct{}{}
	}
	for k := range y {
		wa[k] = struct{}{}
	}
	fmt.Println("和集合:     ", wa)

	// 積集合: 両方ともに入っているもの
	seki := map[string]struct{}{}
	for k := range x {
		if _, has := y[k]; has {
			seki[k] = struct{}{}
		}
	}
	fmt.Println("積集合:     ", seki)

	// 差集合: 自分にあって相手にはないもの
	sax := map[string]struct{}{}
	for k := range x {
		if _, has := y[k]; !has {
			sax[k] = struct{}{}
		}
	}
	say := map[string]struct{}{}
	for k := range y {
		if _, has := x[k]; !has {
			say[k] = struct{}{}
		}
	}
	fmt.Println("積集合(X-Y):", sax)
	fmt.Println("積集合(Y-X):", say)

	// "se"はX,Yに含まれるか
	_, xHasSe := x["se"]
	_, yHasSe := y["se"]
	fmt.Println("X has \"se\": ", xHasSe)
	fmt.Println("Y has \"se\": ", yHasSe)
}
