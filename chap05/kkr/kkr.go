package kkr

import (
	"fmt"
	"strings"
)

type Morph struct {
	Surface string
	Base    string
	Pos     string
	Pos1    string
}

type ChunkInfo struct {
	Index    int     // 文節のindex
	Dst      int     // 係先の文節index
	HeadWord int     // 主辞のindex
	FuncWord int     // 機能語のindex
	Score    float64 // 係り関係のスコア
}

type Chunk struct {
	Morphs []Morph
	Dst    int
	Srcs   []int
}

func ParseLineMorph(line string) (*Morph, error) {
	var surface, feature string

	n, err := fmt.Sscanf(line, "%s\t%s", &surface, &feature)
	if n != 2 || err != nil {
		return nil, fmt.Errorf("invalid format")
	}
	features := strings.Split(feature, ",")
	if len(features) < 7 {
		return nil, fmt.Errorf("invalid format")
	}
	return &Morph{
		Surface: surface,
		Base:    features[6],
		Pos:     features[0],
		Pos1:    features[1],
	}, nil
}

func ParseLineChunkInfo(line string) (*ChunkInfo, error) {
	var index, dst, headWord, funcWord int
	var score float64
	n, err := fmt.Sscanf(line, "* %d %dD %d/%d %f", &index, &dst, &headWord, &funcWord, &score)
	if n != 5 || err != nil {
		return nil, fmt.Errorf("invalid format")
	}
	return &ChunkInfo{
		Index:    index,
		Dst:      dst,
		HeadWord: headWord,
		FuncWord: funcWord,
		Score:    score,
	}, nil
}

// func ParseLine

// func ParseCaboCha(r io.Reader) ([]Morph, error) {

// 	return []Morph{}, nil
// }

// func ParseCaboChaFile(filepath string) ([]Morph, error) {
// 	file, err := os.Open(filepath)
// 	if err != nil {
// 		return nil, fmt.Errorf("os.Open: %v", err)
// 	}
// 	defer file.Close()
// 	return ParseCaboCha(file)
// }
