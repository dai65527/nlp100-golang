package main

import (
	"bufio"
	"chap05/kkr"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("../ai.ja.txt.parsed")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// pass to intro
	count := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() == "EOS" {
			count++
		}
		if count == 2 {
			break
		}
	}

	// scan
	morphs := []kkr.Morph{}
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		if scanner.Text() == "EOS" {
			break
		}

		morph, err := kkr.ParseLine(scanner.Text())
		if err != nil {
			continue
		}

		morphs = append(morphs, *morph)
	}

	for i := 0; i < len(morphs); i++ {
		fmt.Println(morphs[i])
	}
}
