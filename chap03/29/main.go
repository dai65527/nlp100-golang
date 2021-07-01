/*
28. MediaWikiマークアップの除去Permalink
27の処理に加えて，テンプレートの値からMediaWikiマークアップを可能な限り除去し，国の基本情報を整形せよ．
*/

package main

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
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
	reEmp := regexp.MustCompile(`\'{2,5}`)
	for k := range info {
		info[k] = reEmp.ReplaceAllString(info[k], "")
	}

	reInLink := regexp.MustCompile(`\[\[(.*?)\]\]`)
	for k := range info {
		info[k] = reInLink.ReplaceAllString(info[k], "$1")
	}

	reTag := regexp.MustCompile(`\<(.*?)\>`)
	for k := range info {
		info[k] = reTag.ReplaceAllString(info[k], "")
	}

	reTemp := regexp.MustCompile(`\{\{.*?\|(.*?\|)*(.*?)\}\}`)
	for k := range info {
		info[k] = reTemp.ReplaceAllString(info[k], "$2")
	}

	reH := regexp.MustCompile(`\{\{.*?\}\}`)
	for k := range info {
		info[k] = reH.ReplaceAllString(info[k], "")
	}

	res, err := http.Get(
		fmt.Sprintf(
			"https://www.mediawiki.org/w/api.php?action=query&prop=imageinfo&titles=File:%s&format=json&iiprop=url",
			strings.Replace(url.QueryEscape(info["国旗画像"]), "+", "%20", -1)))
	if err != nil {
		fmt.Fprintln(os.Stderr, "http.Get:", err)
		os.Exit(1)
	}
	defer res.Body.Close()
	jsonData, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Fprintln(os.Stderr, "io.ReadAll:", err)
		os.Exit(1)
	}

	reJSONUrl := regexp.MustCompile(`\"url\":\"(.+?)\"`)
	sub := reJSONUrl.FindSubmatch(jsonData)
	url := string(sub[1])
	fmt.Println(url)
}
