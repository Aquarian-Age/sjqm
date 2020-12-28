package sjqm

import (
	"liangzi.local/nongli/ganzhi"
	"liangzi.local/nongli/solar"
	"liangzi.local/sjqm/qm"
	"strings"
	"time"
)

//日紫白起法
func ZiBaiD(dgz string, st time.Time, sy int) (dx int) {
	dzt, xzt := FindT(st, sy)
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
	dzt, xzt := FindT(st, sy)
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
func FindT(st time.Time, sy int) (time.Time, time.Time) {
	jqt := solar.JQT(sy)
	JQ := solar.NewJQ(jqt)
	_, dzt := JQ.Q冬至()
	_, xzt := JQ.Q夏至()
	dzt = time.Date(dzt.Year(), dzt.Month(), dzt.Day(), 0, 0, 0, 0, time.Local)
	xzt = time.Date(xzt.Year(), xzt.Month(), xzt.Day(), 0, 0, 0, 0, time.Local)
	st = time.Date(st.Year(), st.Month(), st.Day(), 0, 0, 0, 0, time.Local)

	return dzt, xzt
}

/*#############紫白飞宫#################*/
//k:紫白数 v:落宫数
func ZiBaiGmap(zbn, yy int) (zbGmap map[int]int) {
	//zbn:紫白入中宫数字(紫白起法值) yy:0阴局 1阳局

	//重新排列紫白顺序以配合飞宫数组
	//阳顺排
	zbi := reZbn(zbn)
	//fmt.Printf("阳顺 重排紫白顺序:%d\n", zbi)
	//阴逆排
	rzbi := reArr(zbi)
	//fmt.Printf("阴逆 重排紫白顺序:%d\n", rzbi)

	//567891234:顺飞 	543219876:逆飞
	//阳局顺排飞宫 [5 6 7 8 9 1 2 3 4]
	garr := []int{5, 6, 7, 8, 9, 1, 2, 3, 4}
	//fmt.Printf("阳 顺排飞宫%d\n", garr)
	//阴局逆序排飞宫 [5 4 3 2 1 9 8 7 6]
	rgarr := reArr(garr) //阴 逆排飞宫
	//fmt.Printf("阴 逆排飞宫%d\n", rgarr)

	switch yy {
	case 1:
		zbGmap = zbg(zbi, garr)
	case 0:
		zbGmap = zbg(rzbi, rgarr)
	}
	return
}

//排紫白 k:紫白数 v:原始宫位数
func zbg(zbiarr, garr []int) map[int]int {
	var zbgmap = make(map[int]int)
	for i := 0; i < len(zbiarr); i++ {
		for j := i; j < len(garr); j++ {
			zbgmap[zbiarr[j]] = garr[j]
			break
		}
	}
	return zbgmap
}

//找到紫白数字所在位置重新排列
func reZbn(zbn int) []int {
	zb := []int{1, 2, 3, 4, 5, 6, 7, 8, 9} //紫白原始位置
	var zbi []int
	for i := 0; i < len(zb); i++ {
		if zbn == zb[i] {
			x1 := zb[i:]
			x2 := zb[:i]
			zbi = append(x1, x2...)
			break
		}
	}
	return zbi
}

//逆序排列 中宫5在首位
func reArr(arr []int) []int {
	l := len(arr)
	var rarr []int
	for k := 0; k < l; k++ {
		tmp := arr[l-k-1]
		rarr = append(rarr, tmp)
	}
	rj1 := rarr[:l-1]
	rj2 := rarr[l-1:]
	rarr = append(rj2, rj1...)
	return rarr
}

/*#########紫白克应(生旺退煞死)#########*/
func ZiBaiShengWang(zbn int, zbGmap map[int]int) (gn int, swts string) {
	//zbn:紫白数字 zbGmap:紫白落宫

	gn = findGn(zbn, zbGmap) //gn:紫白落(原始)宫数字
	zbwx := ZiBaiSelf(zbn)   //紫白五行
	gwx := findGnSelf(gn)    //紫白落宫五行

	//比和n=0 前者生后者n=1 前者克后者n=-1 后者生前者n=2 后者克前者n=-2
	wxn := ganzhi.Wxsk(zbwx, gwx)
	switch wxn {
	case 0: //比和为旺
		swts = "旺"
	case 1: //紫白生宫
		swts = "生"
	case -1: //紫白克宫
		swts = "煞"
	case 2: //宫生紫白
		swts = "退"
	case -2: //宫克紫白
		swts = "死"
	}
	return
}

//紫白落宫数字
func findGn(zbn int, zbGmap map[int]int) (gn int) {
	for k, v := range zbGmap {
		if zbn == k {
			gn = v
			break
		}
	}
	return
}

//落宫五行
func findGnSelf(gn int) (gwx string) {
	//原始宫位五行属性
	gwxmap := map[int]string{1: "水", 2: "土", 3: "木", 4: "木", 5: "土", 6: "金", 7: "金", 8: "土", 9: "火"}
	for k, v := range gwxmap {
		if gn == k {
			gwx = v
			break
		}
	}
	return
}

//紫白五行属性
func ZiBaiSelf(zbn int) (zbwx string) {
	//k:紫白 v:五行属性
	//	"一白": "坎水", "二黑": "坤土", "三碧": "震木", "四绿": "巽木", "五黄": "中土",
	//	"六白": "乾金", "七赤": "兑金", "八白": "艮土", "九紫": "离火",
	wx := []string{"木", "火", "土", "金", "水"}
	zbmap := map[int]string{1: "坎水", 2: "坤土", 3: "震木", 4: "巽木", 5: "中土", 6: "乾金", 7: "兑金", 8: "艮土", 9: "离火"}
	for k, v := range zbmap {
		if k == zbn {
			for j := 0; j < len(wx); j++ {
				if strings.ContainsAny(wx[j], v) {
					zbwx = wx[j]
				}
			}
			break
		}
	}
	return
}

//紫白数字转汉字名称
func ConvNToS(zbn int) (name string) {
	zbmap := map[int]string{1: "一白", 2: "二黑", 3: "三碧", 4: "四绿", 5: "五黄", 6: "六白", 7: "七赤", 8: "八白", 9: "九紫"}
	for k, v := range zbmap {
		if zbn == k {
			name = v
			break
		}
	}
	return
}
