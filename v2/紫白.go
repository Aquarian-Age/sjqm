package sjqm

import (
	"strings"
)

//一白 二黑 三碧 四綠 五黃 六白 七赤 八白 九紫
//年紫白起法
func ZiBaiY(n int, ygz string) (yx int) {
	gn, zn := FindgZhiN(ygz)
	ex0, ex1 := 0, 60 //预备数
	diff := gn - zn

	var N int //柱数 N
	if diff >= 0 {
		N0 := ((gn-zn)*5 + ex0) + gn //当(天干-地支)≥0时,则加预备数0;
		N = N0
	} else if diff < 0 {
		N1 := ((gn-zn)*5 + ex1) + gn //当(天干-地支)<0时,则加预备数60
		N = N1
	}

	//暂定1上元 0中元 -1下元
	switch n {
	case 1:
		yx = (65 - N) % 9
	case 0:
		yx = (68 - N) % 9
	case -1:
		yx = (62 - N) % 9
	}
	return
}

//月紫白起法
func ZiBaiM(ygz, mgz string) (mx int) {
	lm := ConvMGZToNumber(mgz) //lm:农历月的数字

	switch ShengWang(ygz) {
	case "生":
		mx = (12 - lm) % 9
	case "旺":
		mx = (18 - lm) % 9
	case "墓":
		mx = (15 - lm) % 9
	}
	if mx == 0 {
		mx = 9
	}
	return
}

//按《⻓生墓绝表》生 寅申⺒亥,旺 子午卯酉。墓 辰戌丑未
func ShengWang(ygz string) (sw string) {
	zhi := Zhi
	var n int
	for i := 0; i < len(zhi); i++ {
		if strings.ContainsAny(ygz, zhi[i]) {
			n = i
			break
		}
	}
	switch n {
	case 3, 6, 9, 12: //生 寅申⺒亥
		sw = "生"
	case 1, 4, 7, 10: //旺 子午卯酉
		sw = "旺" //正
	case 2, 5, 8, 11: //墓 辰戌丑未
		sw = "墓"
	}
	return
}

//天干数子 地支数字
func FindgZhiN(ygz string) (gn, zn int) {
	gan := Gan //十天干 甲1 乙2 丙3 丁4 戊5 己6 庚7 辛8 壬9 癸10
	zhi := Zhi //十二地支 子1 丑2 寅3 卯4 辰5 巳6 午7 未8 申9 酉10 戌11 亥12
	for i := 0; i < len(gan); i++ {
		if strings.ContainsAny(ygz, gan[i]) {
			gn = i
			break
		}
	}
	for j := 0; j < len(zhi); j++ {
		if strings.ContainsAny(ygz, zhi[j]) {
			zn = j
			break
		}
	}
	return
}
