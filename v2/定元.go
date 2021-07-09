package sjqm

import "strings"

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
