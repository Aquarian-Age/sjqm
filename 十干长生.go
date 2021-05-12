package sjqm

import "strings"

//########################################
//十干长生
//########################################
//阳生阴死（甲阳木的长生地亥 是乙阴木的死地)或者说 甲木的死地是乙木的生地
//阳顺 阴逆
//甲长生在亥死在午  乙长生在午死在亥
//丙长生在寅死在酉  丁长生在酉死在寅
//戊长生在寅死在酉  己长生在酉死在寅
//庚长生在巳死在子  辛长生在子死在巳
//壬长生在申死在卯  癸长生在卯死在申
//k:长生位置 v:十二长生名称 传入干支名称
func ChangSheng(gz string) map[string]string {
	gan := []string{"甲", "乙", "丙", "丁", "戊", "己", "庚", "辛", "壬", "癸"}
	zhi := []string{"子", "丑", "寅", "卯", "辰", "巳", "午", "未", "申", "酉", "戌", "亥"}
	arr := []string{"长生", "沐浴", "冠带", "临官", "帝旺", "衰", "病", "死", "墓", "绝", "胎", "养"}
	chmap := map[string]string{
		"甲": "亥", "丙": "寅", "戊": "寅", "庚": "巳", "壬": "申", //阳干长生位
		"乙": "午", "丁": "酉", "己": "酉", "辛": "子", "癸": "卯", //阴干长生位
	}
	var indexG int //干的索引
	var gmap = make(map[string]string)
	var tmp []string
	for k, v := range chmap {
		for g := 0; g < len(gan); g++ {
			if strings.ContainsAny(gz, gan[g]) {
				indexG = g
				break
			}
		}
		if strings.ContainsAny(gz, k) {
			for i := 0; i < len(zhi); i++ {
				if strings.EqualFold(v, zhi[i]) {
					//找到长生位索引值
					head := zhi[i:]
					end := zhi[:i]
					zhi = append(head, end...)
					//阳顺阴逆排
					x := indexG % 2
					if x == 0 { //阳顺
						for zi := 0; zi < len(zhi); zi++ {
							for a := 0; a < len(arr); a++ {
								if zi == a {
									gmap[zhi[a]] = arr[a]
									break
								}
							}
						}
					}
					if x == 1 { //阴逆
						head = zhi[1:]
						for ii := len(head) - 1; ii >= 0; ii-- {
							tmp = append(tmp, head[ii])
						}
						end = zhi[:1]
						zhi = append(end, tmp...) //长生位保持在第一
						for zi := 0; zi < len(zhi); zi++ {
							for a := 0; a < len(arr); a++ {
								if zi == a {
									gmap[zhi[a]] = arr[a]
									break
								}
							}
						}
					}
					break
				}
			}
			break
		}
	}
	return gmap
}
