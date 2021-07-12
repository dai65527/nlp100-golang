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

func TestParseStringChunks(t *testing.T) {
	s := intro
	res, err := kkr.ParseStringChunks(s)
	assert.NoError(t, err)
	assert.Equal(t, expected, res)
}

const intro = "* 0 17D 1/1 0.388993\n" +
	"人工	名詞,一般,*,*,*,*,人工,ジンコウ,ジンコー\n" +
	"知能	名詞,一般,*,*,*,*,知能,チノウ,チノー\n" +
	"* 1 17D 2/3 0.613549\n" +
	"（	記号,括弧開,*,*,*,*,（,（,（\n" +
	"じん	名詞,一般,*,*,*,*,じん,ジン,ジン\n" +
	"こうち	名詞,一般,*,*,*,*,こうち,コウチ,コーチ\n" +
	"のう	助詞,終助詞,*,*,*,*,のう,ノウ,ノー\n" +
	"、	記号,読点,*,*,*,*,、,、,、\n" +
	"、	記号,読点,*,*,*,*,、,、,、\n" +
	"* 2 3D 0/0 0.758984\n" +
	"AI	名詞,一般,*,*,*,*,*\n" +
	"* 3 17D 1/5 0.517898\n" +
	"〈	記号,括弧開,*,*,*,*,〈,〈,〈\n" +
	"エーアイ	名詞,固有名詞,一般,*,*,*,*\n" +
	"〉	記号,括弧閉,*,*,*,*,〉,〉,〉\n" +
	"）	記号,括弧閉,*,*,*,*,）,）,）\n" +
	"と	助詞,格助詞,引用,*,*,*,と,ト,ト\n" +
	"は	助詞,係助詞,*,*,*,*,は,ハ,ワ\n" +
	"、	記号,読点,*,*,*,*,、,、,、\n" +
	"* 4 5D 2/2 1.035972\n" +
	"「	記号,括弧開,*,*,*,*,「,「,「\n" +
	"『	記号,括弧開,*,*,*,*,『,『,『\n" +
	"計算	名詞,サ変接続,*,*,*,*,計算,ケイサン,ケイサン\n" +
	"* 5 9D 0/3 1.243687\n" +
	"（	記号,括弧開,*,*,*,*,（,（,（\n" +
	"）	記号,括弧閉,*,*,*,*,）,）,）\n" +
	"』	記号,括弧閉,*,*,*,*,』,』,』\n" +
	"という	助詞,格助詞,連語,*,*,*,という,トイウ,トユウ\n" +
	"* 6 9D 0/1 0.691934\n" +
	"概念	名詞,一般,*,*,*,*,概念,ガイネン,ガイネン\n" +
	"と	助詞,並立助詞,*,*,*,*,と,ト,ト\n" +
	"* 7 8D 1/1 1.048596\n" +
	"『	記号,括弧開,*,*,*,*,『,『,『\n" +
	"コンピュータ	名詞,一般,*,*,*,*,コンピュータ,コンピュータ,コンピュータ\n" +
	"* 8 9D 0/3 1.540775\n" +
	"（	記号,括弧開,*,*,*,*,（,（,（\n" +
	"）	記号,括弧閉,*,*,*,*,）,）,）\n" +
	"』	記号,括弧閉,*,*,*,*,』,』,』\n" +
	"という	助詞,格助詞,連語,*,*,*,という,トイウ,トユウ\n" +
	"* 9 10D 0/1 2.129047\n" +
	"道具	名詞,一般,*,*,*,*,道具,ドウグ,ドーグ\n" +
	"を	助詞,格助詞,一般,*,*,*,を,ヲ,ヲ\n" +
	"* 10 12D 0/1 1.363632\n" +
	"用い	動詞,自立,*,*,一段,連用形,用いる,モチイ,モチイ\n" +
	"て	助詞,接続助詞,*,*,*,*,て,テ,テ\n" +
	"* 11 12D 1/3 2.086987\n" +
	"『	記号,括弧開,*,*,*,*,『,『,『\n" +
	"知能	名詞,一般,*,*,*,*,知能,チノウ,チノー\n" +
	"』	記号,括弧閉,*,*,*,*,』,』,』\n" +
	"を	助詞,格助詞,一般,*,*,*,を,ヲ,ヲ\n" +
	"* 12 13D 1/1 0.529353\n" +
	"研究	名詞,サ変接続,*,*,*,*,研究,ケンキュウ,ケンキュー\n" +
	"する	動詞,自立,*,*,サ変・スル,基本形,する,スル,スル\n" +
	"* 13 14D 2/2 0.314025\n" +
	"計算	名詞,サ変接続,*,*,*,*,計算,ケイサン,ケイサン\n" +
	"機	名詞,接尾,一般,*,*,*,機,キ,キ\n" +
	"科学	名詞,一般,*,*,*,*,科学,カガク,カガク\n" +
	"* 14 15D 0/2 4.311133\n" +
	"（	記号,括弧開,*,*,*,*,（,（,（\n" +
	"）	記号,括弧閉,*,*,*,*,）,）,）\n" +
	"の	助詞,連体化,*,*,*,*,の,ノ,ノ\n" +
	"* 15 16D 1/3 2.275064\n" +
	"一	名詞,数,*,*,*,*,一,イチ,イチ\n" +
	"分野	名詞,一般,*,*,*,*,分野,ブンヤ,ブンヤ\n" +
	"」	記号,括弧閉,*,*,*,*,」,」,」\n" +
	"を	助詞,格助詞,一般,*,*,*,を,ヲ,ヲ\n" +
	"* 16 17D 0/0 2.467123\n" +
	"指す	動詞,自立,*,*,五段・サ行,基本形,指す,サス,サス\n" +
	"* 17 34D 0/0 0.791804\n" +
	"語	名詞,一般,*,*,*,*,語,カタリ,カタリ\n" +
	"。	記号,句点,*,*,*,*,。,。,。\n" +
	"* 18 20D 1/2 0.583152\n" +
	"「	記号,括弧開,*,*,*,*,「,「,「\n" +
	"言語	名詞,一般,*,*,*,*,言語,ゲンゴ,ゲンゴ\n" +
	"の	助詞,連体化,*,*,*,*,の,ノ,ノ\n" +
	"* 19 20D 0/1 2.217661\n" +
	"理解	名詞,サ変接続,*,*,*,*,理解,リカイ,リカイ\n" +
	"や	助詞,並立助詞,*,*,*,*,や,ヤ,ヤ\n" +
	"* 20 21D 0/0 0.257647\n" +
	"推論	名詞,サ変接続,*,*,*,*,推論,スイロン,スイロン\n" +
	"、	記号,読点,*,*,*,*,、,、,、\n" +
	"* 21 22D 1/3 2.083214\n" +
	"問題	名詞,ナイ形容詞語幹,*,*,*,*,問題,モンダイ,モンダイ\n" +
	"解決	名詞,サ変接続,*,*,*,*,解決,カイケツ,カイケツ\n" +
	"など	助詞,副助詞,*,*,*,*,など,ナド,ナド\n" +
	"の	助詞,連体化,*,*,*,*,の,ノ,ノ\n" +
	"* 22 24D 1/2 0.693399\n" +
	"知的	名詞,一般,*,*,*,*,知的,チテキ,チテキ\n" +
	"行動	名詞,サ変接続,*,*,*,*,行動,コウドウ,コードー\n" +
	"を	助詞,格助詞,一般,*,*,*,を,ヲ,ヲ\n" +
	"* 23 24D 0/1 2.190897\n" +
	"人間	名詞,一般,*,*,*,*,人間,ニンゲン,ニンゲン\n" +
	"に	助詞,格助詞,一般,*,*,*,に,ニ,ニ\n" +
	"* 24 26D 0/1 1.505879\n" +
	"代わっ	動詞,自立,*,*,五段・ラ行,連用タ接続,代わる,カワッ,カワッ\n" +
	"て	助詞,接続助詞,*,*,*,*,て,テ,テ\n" +
	"* 25 26D 0/1 2.525808\n" +
	"コンピューター	名詞,一般,*,*,*,*,コンピューター,コンピューター,コンピューター\n" +
	"に	助詞,格助詞,一般,*,*,*,に,ニ,ニ\n" +
	"* 26 27D 0/1 3.383079\n" +
	"行わ	動詞,自立,*,*,五段・ワ行促音便,未然形,行う,オコナワ,オコナワ\n" +
	"せる	動詞,接尾,*,*,一段,基本形,せる,セル,セル\n" +
	"* 27 34D 3/3 1.275360\n" +
	"技術	名詞,一般,*,*,*,*,技術,ギジュツ,ギジュツ\n" +
	"」	記号,括弧閉,*,*,*,*,」,」,」\n" +
	"、	記号,読点,*,*,*,*,、,、,、\n" +
	"または	接続詞,*,*,*,*,*,または,マタハ,マタワ\n" +
	"、	記号,読点,*,*,*,*,、,、,、\n" +
	"* 28 29D 2/2 0.253432\n" +
	"「	記号,括弧開,*,*,*,*,「,「,「\n" +
	"計算	名詞,サ変接続,*,*,*,*,計算,ケイサン,ケイサン\n" +
	"機	名詞,接尾,一般,*,*,*,機,キ,キ\n" +
	"* 29 31D 1/3 1.353133\n" +
	"（	記号,括弧開,*,*,*,*,（,（,（\n" +
	"コンピュータ	名詞,一般,*,*,*,*,コンピュータ,コンピュータ,コンピュータ\n" +
	"）	記号,括弧閉,*,*,*,*,）,）,）\n" +
	"による	助詞,格助詞,連語,*,*,*,による,ニヨル,ニヨル\n" +
	"* 30 31D 0/1 1.098217\n" +
	"知的	名詞,形容動詞語幹,*,*,*,*,知的,チテキ,チテキ\n" +
	"な	助動詞,*,*,*,特殊・ダ,体言接続,だ,ナ,ナ\n" +
	"* 31 33D 1/2 0.917915\n" +
	"情報処理	名詞,一般,*,*,*,*,情報処理,ジョウホウショリ,ジョーホーショリ\n" +
	"システム	名詞,一般,*,*,*,*,システム,システム,システム\n" +
	"の	助詞,連体化,*,*,*,*,の,ノ,ノ\n" +
	"* 32 33D 0/1 1.201759\n" +
	"設計	名詞,サ変接続,*,*,*,*,設計,セッケイ,セッケイ\n" +
	"や	助詞,並立助詞,*,*,*,*,や,ヤ,ヤ\n" +
	"* 33 34D 0/1 3.229985\n" +
	"実現	名詞,サ変接続,*,*,*,*,実現,ジツゲン,ジツゲン\n" +
	"に関する	助詞,格助詞,連語,*,*,*,に関する,ニカンスル,ニカンスル\n" +
	"* 34 35D 1/4 0.791804\n" +
	"研究	名詞,サ変接続,*,*,*,*,研究,ケンキュウ,ケンキュー\n" +
	"分野	名詞,一般,*,*,*,*,分野,ブンヤ,ブンヤ\n" +
	"」	記号,括弧閉,*,*,*,*,」,」,」\n" +
	"と	助詞,格助詞,引用,*,*,*,と,ト,ト\n" +
	"も	助詞,係助詞,*,*,*,*,も,モ,モ\n" +
	"* 35 -1D 0/1 0.000000\n" +
	"さ	動詞,自立,*,*,サ変・スル,未然レル接続,する,サ,サ\n" +
	"れる	動詞,接尾,*,*,一段,基本形,れる,レル,レル\n" +
	"。	記号,句点,*,*,*,*,。,。,。\n"

var expected []kkr.Chunk = []kkr.Chunk{
	{Morphs: []kkr.Morph{{Surface: "人工", Base: "人工", Pos: "名詞", Pos1: "一般"}, {Surface: "知能", Base: "知能", Pos: "名詞", Pos1: "一般"}}, Dst: 17, Srcs: []int(nil)},
	{Morphs: []kkr.Morph{{Surface: "（", Base: "（", Pos: "記号", Pos1: "括弧開"}, {Surface: "じん", Base: "じん", Pos: "名詞", Pos1: "一般"}, {Surface: "こうち", Base: "こうち", Pos: "名詞", Pos1: "一般"}, {Surface: "のう", Base: "のう", Pos: "助詞", Pos1: "終助詞"}, {Surface: "、", Base: "、", Pos: "記号", Pos1: "読点"}, {Surface: "、", Base: "、", Pos: "記号", Pos1: "読点"}}, Dst: 17, Srcs: []int(nil)},
	{Morphs: []kkr.Morph{{Surface: "AI", Base: "*", Pos: "名詞", Pos1: "一般"}}, Dst: 3, Srcs: []int{3}},
	{Morphs: []kkr.Morph{{Surface: "〈", Base: "〈", Pos: "記号", Pos1: "括弧開"}, {Surface: "エーアイ", Base: "*", Pos: "名詞", Pos1: "固有名詞"}, {Surface: "〉", Base: "〉", Pos: "記号", Pos1: "括弧閉"}, {Surface: "）", Base: "）", Pos: "記号", Pos1: "括弧閉"}, {Surface: "と", Base: "と", Pos: "助詞", Pos1: "格助詞"}, {Surface: "は", Base: "は", Pos: "助詞", Pos1: "係助詞"}, {Surface: "、", Base: "、", Pos: "記号", Pos1: "読点"}}, Dst: 17, Srcs: []int(nil)},
	{Morphs: []kkr.Morph{{Surface: "「", Base: "「", Pos: "記号", Pos1: "括弧開"}, {Surface: "『", Base: "『", Pos: "記号", Pos1: "括弧開"}, {Surface: "計算", Base: "計算", Pos: "名詞", Pos1: "サ変接続"}}, Dst: 5, Srcs: []int{5}},
	{Morphs: []kkr.Morph{{Surface: "（", Base: "（", Pos: "記号", Pos1: "括弧開"}, {Surface: "）", Base: "）", Pos: "記号", Pos1: "括弧閉"}, {Surface: "』", Base: "』", Pos: "記号", Pos1: "括弧閉"}, {Surface: "という", Base: "という", Pos: "助詞", Pos1: "格助詞"}}, Dst: 9, Srcs: []int(nil)},
	{Morphs: []kkr.Morph{{Surface: "概念", Base: "概念", Pos: "名詞", Pos1: "一般"}, {Surface: "と", Base: "と", Pos: "助詞", Pos1: "並立助詞"}}, Dst: 9, Srcs: []int(nil)},
	{Morphs: []kkr.Morph{{Surface: "『", Base: "『", Pos: "記号", Pos1: "括弧開"}, {Surface: "コンピュータ", Base: "コンピュータ", Pos: "名詞", Pos1: "一般"}}, Dst: 8, Srcs: []int{8}},
	{Morphs: []kkr.Morph{{Surface: "（", Base: "（", Pos: "記号", Pos1: "括弧開"}, {Surface: "）", Base: "）", Pos: "記号", Pos1: "括弧閉"}, {Surface: "』", Base: "』", Pos: "記号", Pos1: "括弧閉"}, {Surface: "という", Base: "という", Pos: "助詞", Pos1: "格助詞"}}, Dst: 9, Srcs: []int{6, 7, 9}},
	{Morphs: []kkr.Morph{{Surface: "道具", Base: "道具", Pos: "名詞", Pos1: "一般"}, {Surface: "を", Base: "を", Pos: "助詞", Pos1: "格助詞"}}, Dst: 10, Srcs: []int{10}},
	{Morphs: []kkr.Morph{{Surface: "用い", Base: "用いる", Pos: "動詞", Pos1: "自立"}, {Surface: "て", Base: "て", Pos: "助詞", Pos1: "接続助詞"}}, Dst: 12, Srcs: []int(nil)},
	{Morphs: []kkr.Morph{{Surface: "『", Base: "『", Pos: "記号", Pos1: "括弧開"}, {Surface: "知能", Base: "知能", Pos: "名詞", Pos1: "一般"}, {Surface: "』", Base: "』", Pos: "記号", Pos1: "括弧閉"}, {Surface: "を", Base: "を", Pos: "助詞", Pos1: "格助詞"}}, Dst: 12, Srcs: []int{11, 12}},
	{Morphs: []kkr.Morph{{Surface: "研究", Base: "研究", Pos: "名詞", Pos1: "サ変接続"}, {Surface: "する", Base: "する", Pos: "動詞", Pos1: "自立"}}, Dst: 13, Srcs: []int{13}},
	{Morphs: []kkr.Morph{{Surface: "計算", Base: "計算", Pos: "名詞", Pos1: "サ変接続"}, {Surface: "機", Base: "機", Pos: "名詞", Pos1: "接尾"}, {Surface: "科学", Base: "科学", Pos: "名詞", Pos1: "一般"}}, Dst: 14, Srcs: []int{14}},
	{Morphs: []kkr.Morph{{Surface: "（", Base: "（", Pos: "記号", Pos1: "括弧開"}, {Surface: "）", Base: "）", Pos: "記号", Pos1: "括弧閉"}, {Surface: "の", Base: "の", Pos: "助詞", Pos1: "連体化"}}, Dst: 15, Srcs: []int{15}},
	{Morphs: []kkr.Morph{{Surface: "一", Base: "一", Pos: "名詞", Pos1: "数"}, {Surface: "分野", Base: "分野", Pos: "名詞", Pos1: "一般"}, {Surface: "」", Base: "」", Pos: "記号", Pos1: "括弧閉"}, {Surface: "を", Base: "を", Pos: "助詞", Pos1: "格助詞"}}, Dst: 16, Srcs: []int{16}},
	{Morphs: []kkr.Morph{{Surface: "指す", Base: "指す", Pos: "動詞", Pos1: "自立"}}, Dst: 17, Srcs: []int{1, 2, 4, 17}},
	{Morphs: []kkr.Morph{{Surface: "語", Base: "語", Pos: "名詞", Pos1: "一般"}, {Surface: "。", Base: "。", Pos: "記号", Pos1: "句点"}}, Dst: 34, Srcs: []int(nil)},
	{Morphs: []kkr.Morph{{Surface: "「", Base: "「", Pos: "記号", Pos1: "括弧開"}, {Surface: "言語", Base: "言語", Pos: "名詞", Pos1: "一般"}, {Surface: "の", Base: "の", Pos: "助詞", Pos1: "連体化"}}, Dst: 20, Srcs: []int(nil)},
	{Morphs: []kkr.Morph{{Surface: "理解", Base: "理解", Pos: "名詞", Pos1: "サ変接続"}, {Surface: "や", Base: "や", Pos: "助詞", Pos1: "並立助詞"}}, Dst: 20, Srcs: []int{19, 20}},
	{Morphs: []kkr.Morph{{Surface: "推論", Base: "推論", Pos: "名詞", Pos1: "サ変接続"}, {Surface: "、", Base: "、", Pos: "記号", Pos1: "読点"}}, Dst: 21, Srcs: []int{21}},
	{Morphs: []kkr.Morph{{Surface: "問題", Base: "問題", Pos: "名詞", Pos1: "ナイ形容詞語幹"}, {Surface: "解決", Base: "解決", Pos: "名詞", Pos1: "サ変接続"}, {Surface: "など", Base: "など", Pos: "助詞", Pos1: "副助詞"}, {Surface: "の", Base: "の", Pos: "助詞", Pos1: "連体化"}}, Dst: 22, Srcs: []int{22}},
	{Morphs: []kkr.Morph{{Surface: "知的", Base: "知的", Pos: "名詞", Pos1: "一般"}, {Surface: "行動", Base: "行動", Pos: "名詞", Pos1: "サ変接続"}, {Surface: "を", Base: "を", Pos: "助詞", Pos1: "格助詞"}}, Dst: 24, Srcs: []int(nil)},
	{Morphs: []kkr.Morph{{Surface: "人間", Base: "人間", Pos: "名詞", Pos1: "一般"}, {Surface: "に", Base: "に", Pos: "助詞", Pos1: "格助詞"}}, Dst: 24, Srcs: []int{23, 24}},
	{Morphs: []kkr.Morph{{Surface: "代わっ", Base: "代わる", Pos: "動詞", Pos1: "自立"}, {Surface: "て", Base: "て", Pos: "助詞", Pos1: "接続助詞"}}, Dst: 26, Srcs: []int(nil)},
	{Morphs: []kkr.Morph{{Surface: "コンピューター", Base: "コンピューター", Pos: "名詞", Pos1: "一般"}, {Surface: "に", Base: "に", Pos: "助詞", Pos1: "格助詞"}}, Dst: 26, Srcs: []int{25, 26}},
	{Morphs: []kkr.Morph{{Surface: "行わ", Base: "行う", Pos: "動詞", Pos1: "自立"}, {Surface: "せる", Base: "せる", Pos: "動詞", Pos1: "接尾"}}, Dst: 27, Srcs: []int{27}},
	{Morphs: []kkr.Morph{{Surface: "技術", Base: "技術", Pos: "名詞", Pos1: "一般"}, {Surface: "」", Base: "」", Pos: "記号", Pos1: "括弧閉"}, {Surface: "、", Base: "、", Pos: "記号", Pos1: "読点"}, {Surface: "または", Base: "または", Pos: "接続詞", Pos1: "*"}, {Surface: "、", Base: "、", Pos: "記号", Pos1: "読点"}}, Dst: 34, Srcs: []int(nil)},
	{Morphs: []kkr.Morph{{Surface: "「", Base: "「", Pos: "記号", Pos1: "括弧開"}, {Surface: "計算", Base: "計算", Pos: "名詞", Pos1: "サ変接続"}, {Surface: "機", Base: "機", Pos: "名詞", Pos1: "接尾"}}, Dst: 29, Srcs: []int{29}},
	{Morphs: []kkr.Morph{{Surface: "（", Base: "（", Pos: "記号", Pos1: "括弧開"}, {Surface: "コンピュータ", Base: "コンピュータ", Pos: "名詞", Pos1: "一般"}, {Surface: "）", Base: "）", Pos: "記号", Pos1: "括弧閉"}, {Surface: "による", Base: "による", Pos: "助詞", Pos1: "格助詞"}}, Dst: 31, Srcs: []int(nil)},
	{Morphs: []kkr.Morph{{Surface: "知的", Base: "知的", Pos: "名詞", Pos1: "形容動詞語幹"}, {Surface: "な", Base: "だ", Pos: "助動詞", Pos1: "*"}}, Dst: 31, Srcs: []int{30, 31}},
	{Morphs: []kkr.Morph{{Surface: "情報処理", Base: "情報処理", Pos: "名詞", Pos1: "一般"}, {Surface: "システム", Base: "システム", Pos: "名詞", Pos1: "一般"}, {Surface: "の", Base: "の", Pos: "助詞", Pos1: "連体化"}}, Dst: 33, Srcs: []int(nil)},
	{Morphs: []kkr.Morph{{Surface: "設計", Base: "設計", Pos: "名詞", Pos1: "サ変接続"}, {Surface: "や", Base: "や", Pos: "助詞", Pos1: "並立助詞"}}, Dst: 33, Srcs: []int{32, 33}},
	{Morphs: []kkr.Morph{{Surface: "実現", Base: "実現", Pos: "名詞", Pos1: "サ変接続"}, {Surface: "に関する", Base: "に関する", Pos: "助詞", Pos1: "格助詞"}}, Dst: 34, Srcs: []int{18, 28, 34}},
	{Morphs: []kkr.Morph{{Surface: "研究", Base: "研究", Pos: "名詞", Pos1: "サ変接続"}, {Surface: "分野", Base: "分野", Pos: "名詞", Pos1: "一般"}, {Surface: "」", Base: "」", Pos: "記号", Pos1: "括弧閉"}, {Surface: "と", Base: "と", Pos: "助詞", Pos1: "格助詞"}, {Surface: "も", Base: "も", Pos: "助詞", Pos1: "係助詞"}}, Dst: 35, Srcs: []int{35}}}
