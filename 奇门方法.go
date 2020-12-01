package sjqm

import (
	"fmt"
	"strings"
)

//门破
//宫制其门是为门破 门制其宫是为宫破
func (g *G) G门破() string {
	m3, m4 := g.G3[1], g.G4[1] //三宫 四宫门信息
	m9 := g.G9[1]
	m1 := g.G1[1]
	m7, m6 := g.G7[1], g.G6[1]
	m8, m2 := g.G8[1], g.G2[1]
	//开门临三宫四宫 是为门破 (金克木)(门的五行属性克制了当前宫位的五行属性/作者注解)
	if strings.ContainsAny("开门", m3) || strings.ContainsAny("惊门", m3) {
		fmt.Printf("%s门临%d宫 门破\n", m3, 3)
		return "门破 吉事不成 凶灾更凶"
	} else if strings.ContainsAny("开门", m4) {
		fmt.Printf("%s门临%d宫 门破\n", m4, 4)
		return "门破 吉事不成 凶灾更凶"
	}
	//休门临九宫 是为门破 (水克火)
	if strings.ContainsAny("休门", m9) {
		fmt.Printf("%s门临%d宫 门破\n", m9, 9)
		return "门破 吉事不成 凶灾更凶"
	}
	//生门临一宫 门破 (土克水)
	if strings.ContainsAny("生门", m1) || strings.ContainsAny("死门", m1) {
		fmt.Printf("%s门临%d宫 门破\n", m1, 1)
		return "门破 吉事不成 凶灾更凶"
	}
	//景门临七宫六宫 门破 (火克金)
	if strings.ContainsAny("景门", m7) {
		fmt.Printf("%s门临%d宫 门破\n", m7, 7)
		return "门破 吉事不成 凶灾更凶"
	} else if strings.ContainsAny("景门", m6) {
		fmt.Printf("%s门临%d宫 门破\n", m6, 6)
		return "门破 吉事不成 凶灾更凶"
	}
	//伤门 杜门临八宫二宫 凶门被破 (木克土)
	if strings.ContainsAny("伤门", m8) {
		fmt.Printf("%s门临%d宫 门破\n", m6, 8)
		return "门破 吉事不成 凶灾更凶"
	} else if strings.ContainsAny("伤门", m2) {
		fmt.Printf("%s门临%d宫 门破\n", m6, 2)
		return "门破 吉事不成 凶灾更凶"
	}

	return ""
}
