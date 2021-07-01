/*
25. テンプレートの抽出Permalink
記事中に含まれる「基礎情報」テンプレートのフィールド名と値を抽出し，辞書オブジェクトとして格納せよ．
*/

package main

import (
	"bufio"
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
	pattern := regexp.MustCompile(`^\|(.+?)\s=\s(.+)$`)
	scanner := bufio.NewScanner(strings.NewReader(s))
	info := map[string]string{}
	for scanner.Scan() {
		res := pattern.FindAllStringSubmatch(scanner.Text(), 1)
		if len(res) > 0 {
			info[res[0][1]] = res[0][2]
		}
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
	re := regexp.MustCompile(`\'{2,5}`)
	for k := range info {
		info[k] = re.ReplaceAllString(info[k], "")
	}

	for k := range info {
		fmt.Printf("%s: %s\n", k, info[k])
	}
}
