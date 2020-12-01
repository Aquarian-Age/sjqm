package sjqm

import (
	"fmt"
	"strings"
)

//门破
//宫制其门是为门破 门制其宫是为宫破
func (g *G) G门破() string {
	//吉门被破
	m3 := g.G3[1] //三宫门信息
	m4 := g.G4[1] //四宫门信息
	m9 := g.G9[1]
	m1 := g.G1[1]
	m7, m6 := g.G7[1], g.G6[1]
	//开门临三宫四宫 是为门破
	if strings.ContainsAny("开门", m3) {
		fmt.Printf("%s门临%d宫 门破\n", m3, 3)
		return "门破 吉事不成 凶灾更凶"
	} else if strings.ContainsAny("开门", m4) {
		fmt.Printf("%s门临%d宫 门破\n", m4, 4)
		return "门破 吉事不成 凶灾更凶"
	}
	//休门临九宫 是为门破
	if strings.ContainsAny("休门", m9) {
		fmt.Printf("%s门临%d宫 门破\n", m9, 9)
		return "门破 吉事不成 凶灾更凶"
	}
	//生门临一宫 门破
	if strings.ContainsAny("生门", m1) {
		fmt.Printf("%s门临%d宫 门破\n", m1, 1)
		return "门破 吉事不成 凶灾更凶"
	}
	//惊门临七宫六宫 门破
	if strings.ContainsAny("景门", m7) {
		fmt.Printf("%s门临%d宫 门破\n", m7, 7)
		return "门破 吉事不成 凶灾更凶"
	} else if strings.ContainsAny("景门", m6) {
		fmt.Printf("%s门临%d宫 门破\n", m6, 6)
		return "门破 吉事不成 凶灾更凶"
	}
	//凶门被破

	return ""
}
