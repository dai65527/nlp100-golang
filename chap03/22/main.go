/*
22. カテゴリ名の抽出Permalink
記事のカテゴリ名を（行単位ではなく名前で）抽出せよ．
*/

package main

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"os"
	"regexp"
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

func main() {
	wiki, err := getWikiData("../jawiki-country.json.gz", "イギリス")
	if err != nil {
		fmt.Fprintln(os.Stderr, "getWikiData:", err)
		os.Exit(1)
	}

	categoryRE := regexp.MustCompile(`(?m)^\[\[Category:(.*?)(\|.*?)?\]\]$`)
	categoryLines := categoryRE.FindAllStringSubmatch(wiki.Text, -1)

	for i := range categoryLines {
		fmt.Println(categoryLines[i][1])
	}
}
