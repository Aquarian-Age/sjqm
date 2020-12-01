package sjqm

import (
	"time"

	"liangzi.local/nongli/solar"
)

func Result(y int, dgz, hgz string, st time.Time) *G {
	jqt := solar.JQT(y)
	jq := solar.NewJQ(jqt)
	jqnames := jq.Name //节气名称数组
	jqarr := jq.Time
	_, dzt := jq.Q冬至()
	_, xzt := jq.Q夏至()
	st = time.Date(st.Year(), st.Month(), st.Day(), st.Hour(), 0, 0, 0, time.Local)
	//定节气
	index, jmc := FindJQ(st, jqarr, jqnames)
	//fmt.Println(jmc)
	//定阴阳遁
	yinyangN := YinYang(st, dzt, xzt)
	var yy string
	if yinyangN == 0 {
		yy = "阴遁"
	}
	if yinyangN == 1 {
		yy = "阳遁"
	}
	//fmt.Println(yy)
	//定元
	yuanx, yn := YuanN(dgz)
	//fmt.Println(yuanx)
	//定局
	juN := FindJU(yn, index, jqnames)
	//fmt.Println(juN)
	sqly := FindSqly(yinyangN, juN)
	//fmt.Printf("地盘三奇六仪:%v\n", sqly)
	//旬首
	xunshou, indexXS := XunShou(hgz)
	//fmt.Printf("旬首:%s\n", xunshou)
	xsGZArr := 旬首干支数组(indexXS)

	//时辰旬首对应的九宫位置数字
	xunShouNumber, _, dun := XunShouHour(xunshou, hgz, xsGZArr, sqly)
	//fmt.Printf("%s时 值符在%d宫位置 遁:%s\n", gzName, xunShouNumber, dun)
	//xsN := FindXSN(xunshou, sqly)
	//fmt.Printf("时辰旬首落%d宫\n", xsN)

	//值符
	zf, starmap, dunmap := ZhiFuStar(xunShouNumber, dun, sqly)
	//fmt.Printf("值符:%s\n九星:%v\n原宫位奇仪:%v\n", zf, starmap, starDun)
	//fmt.Printf("值符:%s\n九星:%v\n", zf, starmap)

	//暗干支
	zfn := FindZFNumber(zf)       //值符原始宫位
	gzarr := FindXSGZArr(xunshou) //旬首干支数组
	agzmap := AnGZ(zfn, gzarr, yinyangN)
	//agzmap := AnGZ(zf, xunshou, yinyangN)
	//fmt.Printf("暗干支:%v\n", agzmap)

	//值使八门
	zhis := FindZS(zfn)
	//fmt.Printf("值使:%s\n", zhis)
	zhishimap := ZhiShi(hgz, agzmap, zhis)
	//fmt.Printf("值使八门:%v\n", zhishimap)

	//八神
	bsmap := BaShen(xunShouNumber, yinyangN)
	//fmt.Printf("八神:%v\n", bsmap)

	////////////////////////////////////////////////结果
	//starmap:值符(九星) zhishimap:值使(八门) agzmap:暗干支 dunmap:六甲旬遁 bsmap:八神 sqly:六仪三奇
	var arr1, arr2, arr3, arr4, arr5, arr6, arr7, arr8, arr9 []string
	arr1 = append(arr1, starmap[1], zhishimap[1], agzmap[1], dunmap[1], bsmap[1], sqly[1]) //一宫信息
	arr2 = append(arr2, starmap[2], zhishimap[2], agzmap[2], dunmap[2], bsmap[2], sqly[2])
	arr3 = append(arr3, starmap[3], zhishimap[3], agzmap[3], dunmap[3], bsmap[3], sqly[3])
	arr4 = append(arr4, starmap[4], zhishimap[4], agzmap[4], dunmap[4], bsmap[4], sqly[4])
	arr5 = append(arr5, starmap[5], zhishimap[5], agzmap[5], dunmap[5], bsmap[5], sqly[5])
	arr6 = append(arr6, starmap[6], zhishimap[6], agzmap[6], dunmap[6], bsmap[6], sqly[6])
	arr7 = append(arr7, starmap[7], zhishimap[7], agzmap[7], dunmap[7], bsmap[7], sqly[7])
	arr8 = append(arr8, starmap[8], zhishimap[8], agzmap[8], dunmap[8], bsmap[8], sqly[8])
	arr9 = append(arr9, starmap[9], zhishimap[9], agzmap[9], dunmap[9], bsmap[9], sqly[9]) //九宫信息

	g := new(G)
	g = &G{
		JieQi:   jmc,
		YinYang: yy,
		N:       juN,
		YUAN:    yuanx,
		XS:      xunshou,
		ZHIFU:   zf,
		ZHISHI:  zhis,
		G1:      arr1,
		G2:      arr2,
		G3:      arr3,
		G4:      arr4,
		G5:      arr5,
		G6:      arr6,
		G7:      arr7,
		G8:      arr8,
		G9:      arr9,
	}
	return g
}
