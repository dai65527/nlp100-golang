package main

import (
	"bufio"
	"chap05/kkr"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("../ai.ja.txt.parsed")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// pass to intro
	scanner := bufio.NewScanner(file)
	scanner.Split(func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		if atEOF && len(data) == 0 {
			return 0, nil, nil
		}
		for i := 0; i < len(data); i++ {
			if i+4 < len(data) && strings.Compare("EOS\n", string(data[i:i+4])) == 0 {
				return i + 4, data[:i], nil
			}
		}
		if atEOF {
			return len(data), data, nil
		}
		return 0, nil, nil
	})

	// pass
	scanner.Scan()
	scanner.Scan()

	// intro
	scanner.Scan()
	chunks, err := kkr.ParseStringChunks(scanner.Text())
	if err != nil {
		panic(err)
	}

	for i := range chunks {
		fmt.Println("Dst:", chunks[i].Dst)
		fmt.Println("Srcs:", chunks[i].Srcs)
		fmt.Println("morphs:", chunks[i].Morphs)
	}
}
