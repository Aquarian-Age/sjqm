package sjqm

import (
	"fmt"
	"strings"
)

//九星数组 1 8 3 4 9 2 7 6 5
func (g *G) 九星() []string {
	n1 := g.G1[0]
	n8 := g.G8[0]
	n3 := g.G3[0]
	n4 := g.G4[0]
	n9 := g.G9[0]
	n2 := g.G2[0]
	n7 := g.G7[0]
	n6 := g.G6[0]
	n5 := g.G5[0]
	var jxarr []string
	jxarr = append(jxarr, n1, n8, n3, n4, n9, n2, n7, n6, n5)
	return jxarr
}

//门破
//宫制其门是为门破 门制其宫是为宫破
func (g *G) G门破() string {
	m3, m4 := g.G3[1], g.G4[1] //三宫 四宫门信息
	m9 := g.G9[1]
	m1 := g.G1[1]
	m7, m6 := g.G7[1], g.G6[1]
	m8, m2 := g.G8[1], g.G2[1]
	//开门临三宫四宫 是为门破 (金克木)(门的五行属性克制了当前宫位的五行属性/作者注解)
	if strings.EqualFold("开门", m3) || strings.EqualFold("惊门", m3) {
		fmt.Printf("%s临%d宫 门破\n", m3, 3)
		return "门破 吉事不成 凶灾更凶"
	} else if strings.EqualFold("开门", m4) {
		fmt.Printf("%s临%d宫 门破\n", m4, 4)
		return "门破 吉事不成 凶灾更凶"
	}
	//休门临九宫 是为门破 (水克火)
	if strings.EqualFold("休门", m9) {
		fmt.Printf("%s临%d宫 门破\n", m9, 9)
		return "门破 吉事不成 凶灾更凶"
	}
	//生门临宫 门破 (土克水)
	if strings.EqualFold("生门", m1) || strings.EqualFold("死门", m1) {
		fmt.Printf("%s临%d宫 门破\n", m1, 1)
		return "门破 吉事不成 凶灾更凶"
	}
	//景门临七宫六宫 门破 (火克金)
	if strings.EqualFold("景门", m7) {
		fmt.Printf("%s临%d宫 门破\n", m7, 7)
		return "门破 吉事不成 凶灾更凶"
	} else if strings.EqualFold("景门", m6) {
		fmt.Printf("%s临%d宫 门破\n", m6, 6)
		return "门破 吉事不成 凶灾更凶"
	}
	//伤门 杜门临八宫二宫 凶门被破 (木克土)
	if strings.EqualFold("伤门", m8) {
		fmt.Printf("%s临%d宫 门破\n", m6, 8)
		return "门破 吉事不成 凶灾更凶"
	} else if strings.EqualFold("伤门", m2) {
		fmt.Printf("%s临%d宫 门破\n", m6, 2)
		return "门破 吉事不成 凶灾更凶"
	}

	return ""
}

//天遁
//开休生三门 天盘丙奇下临地盘丁奇者是
func (g *G) G天遁() string {
	jm := "开休生"
	bing := "丙" //天盘丙奇
	ding := "丁" //地盘丁奇

	m1 := g.G1[1]
	bing1 := g.G1[3]
	ding1 := g.G1[5]
	if strings.ContainsAny(jm, m1) && strings.EqualFold(bing, bing1) && strings.EqualFold(ding, ding1) {
		fmt.Printf("%s 天盘%s奇 地盘:%s奇 天遁\n", m1, bing1, ding1)
		return "天遁"
	}

	m8 := g.G1[1]
	bing8 := g.G1[3]
	ding8 := g.G1[5]
	if strings.ContainsAny(jm, m8) && strings.EqualFold(bing, bing8) && strings.EqualFold(ding, ding8) {
		fmt.Printf("%s 天盘%s奇 地盘:%s奇 天遁\n", m8, bing8, ding8)
		return "天遁"
	}

	m3 := g.G3[1]
	bing3 := g.G3[3]
	ding3 := g.G3[5]
	if strings.ContainsAny(jm, m3) && strings.EqualFold(bing, bing3) && strings.EqualFold(ding, ding3) {
		fmt.Printf("%s 天盘%s奇 地盘:%s奇 天遁\n", m3, bing3, ding3)
		return "天遁"
	}

	m4 := g.G1[1]
	bing4 := g.G1[3]
	ding4 := g.G1[5]
	if strings.ContainsAny(jm, m4) && strings.EqualFold(bing, bing4) && strings.EqualFold(ding, ding4) {
		fmt.Printf("%s 天盘%s奇 地盘:%s奇 天遁\n", m4, bing4, ding4)
		return "天遁"
	}

	m9 := g.G1[1]
	bing9 := g.G1[3]
	ding9 := g.G1[5]
	if strings.ContainsAny(jm, m9) && strings.EqualFold(bing, bing9) && strings.EqualFold(ding, ding9) {
		fmt.Printf("%s 天盘%s奇 地盘:%s奇 天遁\n", m9, bing9, ding9)
		return "天遁"
	}

	m2 := g.G2[1]
	bing2 := g.G2[3]
	ding2 := g.G2[5]
	if strings.ContainsAny(jm, m2) && strings.EqualFold(bing, bing2) && strings.EqualFold(ding, ding2) {
		fmt.Printf("%s 天盘%s奇 地盘:%s奇 天遁\n", m2, bing2, ding2)
		return "天遁"
	}

	m7 := g.G1[1]
	bing7 := g.G1[3]
	ding7 := g.G1[5]
	if strings.ContainsAny(jm, m7) && strings.EqualFold(bing, bing7) && strings.EqualFold(ding, ding7) {
		fmt.Printf("%s 天盘%s奇 地盘:%s奇 天遁\n", m7, bing7, ding7)
		return "天遁"
	}

	m6 := g.G1[1]
	bing6 := g.G1[3]
	ding6 := g.G1[5]
	if strings.ContainsAny(jm, m6) && strings.EqualFold(bing, bing6) && strings.EqualFold(ding, ding6) {
		fmt.Printf("%s 天盘%s奇 地盘:%s奇 天遁\n", m6, bing6, ding6)
		return "天遁"
	}
	return ""
}

/*
趋三避五
趋三 值使到震宜向之。
避五 值使到巽宫(中宫)宜去之
*/
func (gmap *GMap) G趋三避五(zhis string) (qbs string) {
	zsmap := gmap.ZhiShiMap
	for k, v := range zsmap {
		if strings.EqualFold(zhis, v) {
			if k == 3 {
				//fmt.Printf("-->值使%s落%d宫 宜趋之吉\n", v, k)
				qbs = fmt.Sprintf("值使%s落%d宫 宜趋之吉", v, k)
			}
		}
		if strings.EqualFold(zhis, v) {
			if k == 5 || k == 2 {
				//fmt.Printf("-->值使%s落%d宫 宜趋之吉\n", v, k)
				qbs = fmt.Sprintf("值使%s落%d宫 宜趋之吉", v, k)
			}
		}
	}
	return
}
