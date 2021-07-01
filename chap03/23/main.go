/*
23. セクション構造
記事中に含まれるセクション名とそのレベル（例えば”== セクション名 ==”なら1）を表示せよ．
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

	categoryRE := regexp.MustCompile(`(?m)^(=+)(.*?)(=+$)`)
	sectionLines := categoryRE.FindAllStringSubmatch(wiki.Text, -1)

	for i := range sectionLines {
		if len(sectionLines[i]) == 4 {
			cnt, err := countEq(sectionLines[i][1])
			if err != nil || sectionLines[i][1] != sectionLines[i][3] {
				continue
			}
			fmt.Println(sectionLines[i][2], cnt-1)
		}
	}
}

func countEq(s string) (int, error) {
	for _, c := range s {
		if c != '=' {
			return 0, fmt.Errorf("%s: contains wrong rune", s)
		}
	}
	return len(s), nil
}
