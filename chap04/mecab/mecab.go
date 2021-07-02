package mecab

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type Morpheme struct {
	Surface string
	Base    string
	Pos     string
	Pos1    string
}

func ParseReader(r io.Reader) ([][]Morpheme, error) {
	scanner := bufio.NewScanner(r)

	result := [][]Morpheme{}
	sentence := []Morpheme{}

	linecount := 0
	for scanner.Scan() {
		linecount++
		data := scanner.Text()

		if len(data) == 0 {
			// after new line
			if !scanner.Scan() {
				return nil, fmt.Errorf("line %d: syntax error: missing data after surface", linecount)
			}
			linecount++
			data += scanner.Text()

		} else if data == "EOS" {
			// end of sentence
			result = append(result, sentence)
			sentence = []Morpheme{}

		} else {
			// parse sentence
			words := strings.Split(data, "\t")
			if len(words) != 2 {
				return nil, fmt.Errorf("line %d: syntax error: no or extra tabs", linecount)
			}
			datas := strings.Split(words[1], ",")
			if len(datas) < 7 {
				return nil, fmt.Errorf("line %d: syntax error: no or extra commas", linecount)
			}

			sentence = append(sentence, Morpheme{
				Surface: words[0],
				Base:    datas[6],
				Pos:     datas[0],
				Pos1:    datas[1],
			})
		}
	}
	return result, nil
}

func ParseFile(filepath string) ([][]Morpheme, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, fmt.Errorf("mecab.ParseFile: %v", err)
	}
	return ParseReader(file)
}

func FindPosFromText(posName string, text [][]Morpheme) []Morpheme {
	res := []Morpheme{}
	for i := range text {
		for _, morpheme := range text[i] {
			if morpheme.Pos == posName {
				res = append(res, morpheme)
			}
		}
	}
	return res
}
