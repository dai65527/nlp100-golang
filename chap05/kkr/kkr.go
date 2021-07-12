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

func ParseLine(line string) (*Morph, error) {
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
