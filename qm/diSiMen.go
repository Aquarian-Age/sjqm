package qm

import (
	"fmt"
	"liangzi.local/cal"
	"strings"
)

//地私门 六合 太阴 太常
func DiSiMen(yj, dgz, hgz string) string {

	yang, yin := cal.GuiRenJue(dgz)
	hyy, index := hgzYinYang(hgz)

	var dsMap = make(map[string]string)
	switch hyy {
	case 1:
		//fmt.Printf("时辰定阴阳:%d 阳\n", hyy)
		rzhi := diYueJiangGong(yj, yang, index)
		//fmt.Printf("-->阳 重排月将:%s\n", rzhi)
		shenSha := ss(rzhi)
		//fmt.Printf("阳顺神煞:%s\n", shenSha)
		dsMap = ssMap(shenSha, rzhi)
		//fmt.Printf("地私门map:%v\n", dsMap)

	case 0:
		//fmt.Printf("时辰定阴阳:%d 阴\n", hyy)
		rzhi := diYueJiangGong(yj, yin, index)
		//fmt.Printf("-->阴 重排月将:%s\n", rzhi)
		shenSha := ss(rzhi)
		//fmt.Printf("阴逆神煞:%s\n", shenSha)
		dsMap = ssMap(shenSha, rzhi)
		//fmt.Printf("地私门map:%v\n", dsMap)
	}
	dsms := fmt.Sprintf("六合:%s 太常:%s 太阴:%s", dsMap["六合"], dsMap["太常"], dsMap["太阴"])
	return dsms
}

//时辰定阴阳
//辰巳午未申酉 白天阳贵
//戌亥子丑寅卯 夜晚阴贵
func hgzYinYang(hgz string) (hyy, index int) {
	zhi := []string{"亥", "子", "丑", "寅", "卯", "辰", "巳", "午", "未", "申", "酉", "戌"}

	for i := 0; i < len(zhi); i++ {
		if strings.ContainsAny(hgz, zhi[i]) {
			index = i
			break
		}
	}

	switch zhi[index] {
	case zhi[5], zhi[6], zhi[7], zhi[8], zhi[9], zhi[10]: //阳
		hyy = 1
	case zhi[11], zhi[0], zhi[1], zhi[2], zhi[3], zhi[4]: //阴
		hyy = 0
	}
	return
}

//月将加时辰找神煞泊地月将宫位 重排
func diYueJiangGong(yj, gr string, index int) []string {
	//时辰在地月将宫位索引值 排序天月将
	zhi := []string{"亥", "子", "丑", "寅", "卯", "辰", "巳", "午", "未", "申", "酉", "戌"}

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

	var tyj []string
	tyj = append(s2, s1...)
	//fmt.Printf("-->地月将:%s\n", zhi)
	//fmt.Printf("-->天月将:%s\n", tyj)

	//生成天月将-地月将map k:天月将 v:地月将
	var yjmap = make(map[string]string)
	for k := 0; k < len(tyj); k++ {
		for v := k; v < len(zhi); v++ {
			yjmap[tyj[k]] = zhi[v]
			break
		}
	}
	//fmt.Printf("-->月将map:%v\n", yjmap)

	//贵人落宫
	var rzhi []string
	for s, e := range yjmap {
		if strings.EqualFold(gr, s) {
			//贵人落地月将宫位
			//fmt.Printf("-->贵人%s落地月将%s宫位\n", gr, e)
			//从地月将落宫位置 从排
			for j := 0; j < len(zhi); j++ {
				if strings.EqualFold(zhi[j], e) {
					r1 := zhi[:j]
					r2 := zhi[j:]
					rzhi = append(r2, r1...)
					break
				}
			}
			break
		}
	}
	//fmt.Printf("-->重排月将:%s\n", rzhi)

	return rzhi
}

//神煞顺逆
//亥子丑寅卯辰 顺排
//巳午未申酉戌 逆排
func ss(rzhi []string) (shenSha []string) {
	dyjg := rzhi[0] //贵人落地月将宫位
	zhi := []string{"亥", "子", "丑", "寅", "卯", "辰", "巳", "午", "未", "申", "酉", "戌"}
	//十二神将原始顺序
	sj := []string{"贵人", "腾蛇", "朱雀", "六合", "勾陈", "青龙", "天空", "白虎", "太常", "玄武", "太阴", "天后"}

	switch dyjg {
	case zhi[0], zhi[1], zhi[2], zhi[3], zhi[4], zhi[5]:
		shenSha = sj
	case zhi[6], zhi[7], zhi[8], zhi[9], zhi[10], zhi[11]:
		//逆排神将
		var rsj []string
		for k := 0; k < len(sj); k++ {
			tmp := sj[len(sj)-k-1]
			rsj = append(rsj, tmp)
		}
		//贵人居首位
		rj1 := rsj[:11]
		rj2 := rsj[11:]
		rsj = append(rj2, rj1...)
		//fmt.Printf("-->神将逆序:%s\n", rsj)
		shenSha = rsj
	}
	return
}

//神煞配地月将 k:神煞 v:地月将
func ssMap(shenSha, rzhi []string) map[string]string {
	var ssmap = make(map[string]string)
	for i := 0; i < len(shenSha); i++ {
		for j := i; j < len(rzhi); j++ {
			ssmap[shenSha[j]] = rzhi[j]
		}
	}
	//找地私门  k:神煞 v:地支
	var dsm = "太阴 六合 太常"
	var dsmap = make(map[string]string)
	for k, v := range ssmap {
		if strings.ContainsAny(k, dsm) {
			dsmap[k] = v
		}
	}
	return dsmap
}
