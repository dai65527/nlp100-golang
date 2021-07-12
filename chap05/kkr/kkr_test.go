package kkr_test

import (
	"chap05/kkr"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseLineMorph(t *testing.T) {
	line := "知能	名詞,一般,*,*,*,*,知能,チノウ,チノー"
	res, err := kkr.ParseLineMorph(line)
	assert.NoError(t, err)
	assert.Equal(t, kkr.Morph{Surface: "知能", Base: "知能", Pos: "名詞", Pos1: "一般"}, *res)
}

func TestParseLineInfo(t *testing.T) {
	line := "* 2 3D 1/4 0.758984"
	res, err := kkr.ParseLineChunkInfo(line)
	assert.NoError(t, err)
	assert.Equal(t, kkr.ChunkInfo{Index: 2, Dst: 3, HeadWord: 1, FuncWord: 4, Score: float64(0.758984)}, *res)
}
