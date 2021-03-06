package sjqm

import (
	"time"

	"liangzi.local/cal"
)

func Result(y int, dgz, hgz string, st time.Time) (*G, *GMap) {
	jqobj := cal.NewJQArr(y)
	jqnames := jqobj.Name
	jqarr := jqobj.Time
	dzt := jqobj.Time[24]
	xzt := jqobj.Time[12]
	/////////
	st = time.Date(st.Year(), st.Month(), st.Day(), st.Hour(), 0, 0, 0, time.Local)
	//定节气
	index, jmc := FindJQ(st, jqarr, jqnames)
	//定阴阳遁
	yinyangN := YinYang(st, dzt, xzt)
	var yy string
	if yinyangN == 0 {
		yy = "阴遁"
	}
	if yinyangN == 1 {
		yy = "阳遁"
	}
	//定元
	yuanx, yn := YuanN(dgz)
	//定局
	juN := FindJU(yn, index, jqnames)
	sqly := FindSqly(yinyangN, juN)
	//旬首
	xunshou, indexXS := XunShou(hgz)
	xsGZArr := 旬首干支数组(indexXS)
	//时辰旬首对应的九宫位置数字
	xunShouNumber, _, dun := XunShouHour(xunshou, hgz, xsGZArr, sqly)
	//值符
	zf, dunN := SelfZF(dun, sqly)
	starArr := SortStar(zf)
	//九星配宫
	starmap := ZhiFuStar(xunShouNumber, starArr)
	//六仪三奇配九星
	starQYmap := StarQY(starArr, sqly, dunN, starmap)
	//暗干支
	zfn := FindZFNumber(zf)       //值符原始宫位
	gzarr := FindXSGZArr(xunshou) //旬首干支数组
	agzmap := AnGZ(zfn, gzarr, yinyangN)
	//值使八门
	zhis := FindZS(zfn)
	zhishimap := ZhiShi(hgz, agzmap, zhis)
	//八神
	bsmap := BaShen(xunShouNumber, yinyangN)

	//结果
	//starmap:值符(九星) zhishimap:值使(八门) agzmap:暗干支 starQYmap:天盘奇仪 bsmap:八神 sqly:地盘奇仪
	var arr1, arr2, arr3, arr4, arr5, arr6, arr7, arr8, arr9 []string
	arr1 = append(arr1, starmap[1], zhishimap[1], agzmap[1], starQYmap[1], bsmap[1], sqly[1]) //一宫信息
	arr2 = append(arr2, starmap[2], zhishimap[2], agzmap[2], starQYmap[2], bsmap[2], sqly[2])
	arr3 = append(arr3, starmap[3], zhishimap[3], agzmap[3], starQYmap[3], bsmap[3], sqly[3])
	arr4 = append(arr4, starmap[4], zhishimap[4], agzmap[4], starQYmap[4], bsmap[4], sqly[4])
	arr5 = append(arr5, starmap[5], zhishimap[5], agzmap[5], starQYmap[5], bsmap[5], sqly[5])
	arr6 = append(arr6, starmap[6], zhishimap[6], agzmap[6], starQYmap[6], bsmap[6], sqly[6])
	arr7 = append(arr7, starmap[7], zhishimap[7], agzmap[7], starQYmap[7], bsmap[7], sqly[7])
	arr8 = append(arr8, starmap[8], zhishimap[8], agzmap[8], starQYmap[8], bsmap[8], sqly[8])
	arr9 = append(arr9, starmap[9], zhishimap[9], agzmap[9], starQYmap[9], bsmap[9], sqly[9]) //九宫信息

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

	gmap := new(GMap)
	gmap = &GMap{ZhiShiMap: zhishimap}

	return g, gmap
}
