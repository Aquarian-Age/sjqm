package sjqm

import (
	"liangzi.local/nongli/ganzhi"
	"liangzi.local/nongli/solar"
	"liangzi.local/sjqm/qm"
	"time"
)

//日紫白起法
func ZiBaiD(dgz string, st time.Time, sy int) (dx int) {
	dzt, xzt := findT(st, sy)
	_, yuanN := YuanN(dgz)            //元 0上元 1中元 2下元
	yinYangN := YinYang(st, dzt, xzt) //1阳局 0阴局
	N := findN(dgz)                   //柱数

	switch yinYangN {
	case 1: //阳局
		switch yuanN {
		case 0:
			dx = (N + 0) % 9
		case 1:
			dx = (N + 6) % 9
		case 2:
			dx = (N + 3) % 9
		}
	case 0: //阴局
		switch yuanN {
		case 0:
			dx = (64 - N) % 9
		case 1:
			dx = (67 - N) % 9
		case 2:
			dx = (61 - N) % 9
		}
	}
	if dx == 0 {
		dx = 9
	}
	return
}

//时辰紫白起法
func ZiBaiH(dgz, hgz string, st time.Time, sy int) (hx int) {
	dzt, xzt := findT(st, sy)
	//_, yuanN := YuanN(dgz)            //元 0上元 1中元 2下元
	yinYangN := YinYang(st, dzt, xzt) //1阳局 0阴局

	hn := ganzhi.ConvHGZToNumber(hgz) //时辰数字
	sw := qm.ShengWang(dgz)           //生 旺(正) 墓

	switch yinYangN {
	case 1:
		switch sw {
		case "旺": //四正时
			hx = (hn + 0) % 9
		case "墓":
			hx = (hn + 3) % 9
		case "生":
			hx = (hn + 6) % 9
		}
	case 0:
		switch sw {
		case "旺": //四正时
			hx = (19 - hn) % 9
		case "墓":
			hx = (16 - hn) % 9
		case "生":
			hx = (13 - hn) % 9
		}
	}
	if hx == 0 {
		hx = 9
	}

	return
}

//柱数 适用年及日
func findN(gz string) int {
	gn, zn := qm.FindgZhiN(gz)
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
	return N
}

//冬至夏至时间
//这里时间精确到日
func findT(st time.Time, sy int) (time.Time, time.Time) {
	jqt := solar.JQT(sy)
	JQ := solar.NewJQ(jqt)
	_, dzt := JQ.Q冬至()
	_, xzt := JQ.Q夏至()
	dzt = time.Date(dzt.Year(), dzt.Month(), dzt.Day(), 0, 0, 0, 0, time.Local)
	xzt = time.Date(xzt.Year(), xzt.Month(), xzt.Day(), 0, 0, 0, 0, time.Local)
	st = time.Date(st.Year(), st.Month(), st.Day(), 0, 0, 0, 0, time.Local)

	return dzt, xzt
}
