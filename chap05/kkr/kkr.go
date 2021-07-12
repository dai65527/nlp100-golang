package kkr

import (
	"bufio"
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

func ParseStringChunks(s string) ([]Chunk, error) {
	scanner := bufio.NewScanner(strings.NewReader(s))

	chunks := []Chunk{}
	var chunk *Chunk = nil
	for scanner.Scan() {
		if scanner.Text() == "" {
			continue
		}
		chunkInfo, err := ParseLineChunkInfo(scanner.Text())

		// new chunk
		if err == nil {
			if chunk != nil {
				chunks = append(chunks, *chunk)
			}
			chunk = &Chunk{Dst: chunkInfo.Dst}
			continue
		}

		// add morph
		morph, err := ParseLineMorph(scanner.Text())
		if err != nil {
			return nil, err
		}
		chunk.Morphs = append(chunk.Morphs, *morph)
	}

	for i := 0; i < len(chunks); i++ {
		if chunks[i].Dst >= 1 {
			chunks[chunks[i].Dst-1].Srcs = append(chunks[chunks[i].Dst-1].Srcs, i+1)
		}
	}

	return chunks, nil
}
