package qm

import (
	"fmt"
	"liangzi.local/nongli/ganzhi"
	"strings"
)

//五符 属日奇门
/*
从本日干禄上起五符顺行
一五行（符），二天曹，三地府（符），四风伯，五雷公，六雨师，七风灵（云），八唐符，九国印，十天关，十一地轴，十二天贼
視風伯所到之方，主有風，雨師所到之方，主有雨
*/
func WuFu(dgz string) []string {
	wfs := []string{"五符", "天曹", "地符", "风伯", "雷公", "雨师", "风云", "唐符", "国印", "天官", "地轴", "天贼"}
	zhi := []string{"子", "丑", "寅", "卯", "辰", "巳", "午", "未", "申", "酉", "戌", "亥"}

	rlu := ganzhi.GZ禄(dgz) //日禄

	for i := 0; i < len(zhi); i++ {
		if strings.EqualFold(rlu, zhi[i]) {
			s1 := zhi[i:]
			s2 := zhi[:i]
			zhi = append(s1, s2...)
			break
		}
	}
	var wfmap = make(map[string]string)
	for k := 0; k < len(zhi); k++ {
		for v := k; v < len(wfs); v++ {
			wfmap[zhi[v]] = wfs[v]
			break
		}
	}
	//fmt.Printf("-->五符:%v\n", wfmap)
	wufus := mapToArr(wfmap)
	return wufus
}

func mapToArr(mapx map[string]string) []string {
	var arr = []string{}
	for k, v := range mapx {
		tmp := fmt.Sprintf("%s:%s", k, v)
		arr = append(arr, tmp)
	}
	return arr
}
