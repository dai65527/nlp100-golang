/*
25. テンプレートの抽出Permalink
記事中に含まれる「基礎情報」テンプレートのフィールド名と値を抽出し，辞書オブジェクトとして格納せよ．
*/

package main

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type Wiki struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

func getWikiData(filepath, title string) (*Wiki, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, fmt.Errorf("os.Open: %v", err)
	}
	defer file.Close()

	reader, err := gzip.NewReader(file)
	if err != nil {
		return nil, fmt.Errorf("gzip.NewReader: %v", err)
	}
	defer reader.Close()

	decoder := json.NewDecoder(reader)

	var wiki Wiki
	for decoder.More() {
		decoder.Decode(&wiki)
		if wiki.Title == title {
			return &wiki, nil
		}
	}
	return nil, fmt.Errorf("%s: not found", title)
}

func getInfo(s string) map[string]string {
	info := map[string]string{}

	pattern := regexp.MustCompile(`\{\{基礎情報\s国\n((.*)\n)+?\}\}\n`)
	res := pattern.FindString(s)
	infoString := strings.Split(res, "\n|")
	for _, infoStr := range infoString[1 : len(infoString)-1] {
		kv := strings.Split(infoStr, "=")
		info[strings.Trim(kv[0], " ")] = strings.Trim(infoStr[len(kv[0])+1:], " ")
	}
	return info
}

func main() {
	wiki, err := getWikiData("../jawiki-country.json.gz", "イギリス")
	if err != nil {
		fmt.Fprintln(os.Stderr, "getWikiData:", err)
		os.Exit(1)
	}

	info := getInfo(wiki.Text)
	for k, v := range info {
		fmt.Printf("%s: %s\n", k, v)
	}
}
