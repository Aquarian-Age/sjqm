package qm

import (
	"strings"
)

//从魁 小吉 太冲 即是天三门
//如得本日贵人到乾亥，就是贵人登天门
func TianSanMen(yj, hgz string) map[string]string {
	ssmap := shenShaMap()
	yjzhi := yjh(yj, hgz)

	var smmap = make(map[string]string)
	for z, name := range ssmap {
		for s, e := range yjzhi {
			if strings.EqualFold(z, s) {
				smmap[name] = e
			}
		}
	}

	return smmap
}

//神煞map k:地月将 v:神煞
func shenShaMap() map[string]string {
	ss := []string{"神后", "大吉", "功曹", "太冲", "天罡", "太乙", "胜光", "小吉", "传送", "从魁", "河魁", "登明"}
	zhi := []string{"子", "丑", "寅", "卯", "辰", "巳", "午", "未", "申", "酉", "戌", "亥"}
	var shenMap = make(map[string]string)
	for k := 0; k < len(zhi); k++ {
		for v := k; v < len(ss); v++ {
			shenMap[zhi[v]] = ss[v]
		}
	}
	//找太冲 小吉 从魁三门
	var smmap = make(map[string]string)
	for z, name := range shenMap {
		if strings.EqualFold(name, "太冲") ||
			strings.EqualFold(name, "小吉") ||
			strings.EqualFold(name, "从魁") {
			smmap[z] = name
		}
	}
	return smmap //shenMap
}

//月将加时辰 k:天月将 v:地月将
func yjh(yj, hgz string) map[string]string {
	zhi := []string{"亥", "子", "丑", "寅", "卯", "辰", "巳", "午", "未", "申", "酉", "戌"}
	//fmt.Println("地月将:  ", zhi)
	var index int
	for j := 0; j < len(zhi); j++ {
		if strings.ContainsAny(hgz, zhi[j]) {
			index = j
			break
		}
	}

	//天月将排序 本月月将居首位
	var yjzhi []string
	for i := 0; i < len(zhi); i++ {
		if strings.ContainsAny(yj, zhi[i]) {
			y1 := zhi[i:]
			y2 := zhi[:i]
			yjzhi = append(y1, y2...)
			break
		}
	}
	//fmt.Printf("-->天月将排序:%s\n", yjzhi)

	end := zhi[index:]     //天月将加时辰 地月将的后半部分
	s1 := yjzhi[:len(end)] //天月将后半部分
	s2 := yjzhi[len(end):] //天月将前半部分

	//天月将加时辰
	var tyj []string
	tyj = append(s2, s1...)

	//天月将map k:天月将 v:地月将
	var tyjmap = make(map[string]string)
	for k := 0; k < len(tyj); k++ {
		for v := k; v < len(zhi); v++ {
			tyjmap[tyj[v]] = zhi[v]
			break
		}
	}
	//fmt.Printf("-->天月将:%v\n", tyjmap)
	return tyjmap
}
