package sjqm

import (
	"fmt"
	"strings"
)

//地四户：除定危开 时支上加月建
//建除满平一顺流,定执破危相接去,成收开闭掌中周,除定危开为四户,此方有难来逃避
func diSiHu(hgz string) string {
	zhi := []string{"子", "丑", "寅", "卯", "辰", "巳", "午", "未", "申", "酉", "戌", "亥"}
	jc := []string{"建", "除", "满", "平", "定", "执", "破", "危", "成", "收", "开", "闭"}
	//时辰排序
	for i := 0; i < len(zhi); i++ {
		if strings.ContainsAny(hgz, zhi[i]) {
			//月支重排
			mx1 := zhi[i:]
			mx2 := zhi[:i]
			zhi = append(mx1, mx2...)
			break
		}
	}
	//fmt.Printf("地支:%s\n建除:%s\n", zhi, jc)
	//地支:[巳 午 未 申 酉 戌 亥 子 丑 寅 卯 辰]
	//建除:[建 除 满 平 定 执 破 危 成 收 开 闭]
	//生成全部地支:建除的map键值对
	var shmap = make(map[string]string)
	for ji := 0; ji < len(jc); ji++ {
		for zi := 0; zi < len(zhi); zi++ {
			if ji == zi {
				shmap[jc[ji]] = zhi[zi]
				break
			}
		}
	}

	//找地四户
	disihu := []string{"除", "定", "危", "开"}
	var sihumap = make(map[string]string)
	for k, v := range shmap {
		for di := 0; di < len(disihu); di++ {
			if strings.EqualFold(k, disihu[di]) {
				sihumap[k] = v
			}
		}
	}
	dshs := fmt.Sprintf("除在:%s 定在:%s 危在:%s 开在:%s", sihumap["除"], sihumap["定"], sihumap["危"], sihumap["开"])
	return dshs
}
