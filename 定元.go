package sjqm

import (
	"strings"
)

/*
当日日干支往前推 一直推到甲日或者己日 看甲日或己日对应的地支
子午卯酉为上元 寅申巳亥为中元 辰戌丑未为下元
*/

//符头定元
func YuanN(dgz string) (string, int) {
	futou := findFT(dgz) //符头 根据符头地支定元
	zhi := [13]string{"err", "子", "丑", "寅", "卯", "辰", "巳", "午", "未", "申", "酉", "戌", "亥"}
	var n int
	for i := 0; i < len(zhi); i++ {
		if strings.ContainsAny(futou, zhi[i]) {
			n = i
			break
		}
	}
	return infoY(n)
}

//符头
func findFT(dgz string) (futou string) {
	var i int
	for i = 0; i < len(Jz60); i++ {
		if strings.EqualFold(dgz, Jz60[i]) {
			break
		}
	}

	for j := i; j >= i-5; j-- {
		if strings.ContainsAny(Jz60[j], "甲") || strings.ContainsAny(Jz60[j], "己") {
			futou = Jz60[j] //符头
			break
		}
	}
	return
}
func infoY(n int) (name string, number int) {
	switch n {
	case 1, 7, 4, 10: //子1 午7 卯4 酉10　上元
		name = "上元"
		number = 0
	case 3, 9, 6, 12: //寅3 申9 巳6 亥12　中元
		name = "中元"
		number = 1
	case 5, 11, 2, 8: //辰5 戌11 醜2 未8　下元
		name = "下元"
		number = 2
	}
	return
}

//////////////////////////////////////下面弃用

/*//符头定元元
func YuanX(dg, dz int) (yuan string, xday int) {
	//dg:當日天干數字 dz:當日地支數字
	fg, offg := 符头天干(dg)
	fz := 符头地支(dz, offg)

	switch fg {
	case 1: //符頭甲
		yuan = yuanx(fz)
	case 6: //符頭己
		yuan = yuanx(fz)
	}
	xday = offg + 1
	return
}

//符頭天干數字  甲1 乙2 丙3 丁4...
func 符头天干(g int) (fg, offg int) {
	//offg:天干偏移數字 天干偏移數字用來計算地支　同時計算三元定局本日和符頭的偏移天數(offg+1)
	//g:天干数子 甲1 乙2 丙3 丁4...
	if g <= 5 {
		fg = 1 //符頭甲
		offg = g - 1
	} else {
		fg = 6 //符頭己
		offg = g - 6
	}
	return
}

//符頭地支數字
func 符头地支(dz, offg int) (fz int) {
	//dz:日支數字　offg：FG符头天干函數返回的差值
	x := 12 - offg + dz
	if x > 12 {
		fz = x - 12
	} else {
		fz = x
	}
	return
}

//符头三元
func yuanx(fz int) (yuanx string) {
	switch fz {
	case 1, 7, 4, 10: //子1 午7 卯4 酉10　上元
		yuanx = "上元"
	case 3, 9, 6, 12: //寅3 申9 巳6 亥12　中元
		yuanx = "中元"
	case 5, 11, 2, 8: //辰5 戌11 醜2 未8　下元
		yuanx = "下元"
	}
	return
}*/
