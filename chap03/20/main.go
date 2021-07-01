/*
20. JSONデータの読み込み
Wikipedia記事のJSONファイルを読み込み，「イギリス」に関する記事本文を表示せよ．問題21-29では，ここで抽出した記事本文に対して実行せよ．
*/

package main

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"os"
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

	fmt.Println(wiki.Title)
	fmt.Println(wiki.Text)
}
