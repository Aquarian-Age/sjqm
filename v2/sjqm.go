package sjqm

import (
	"github.com/Aquarian-Age/xa/pkg/gz"
	"time"
)

//九宫排盘结果信息
type SJQM struct {
	ygz, mgz, dgz, hgz string
	YjZhi, YjName, Yjt string
	JieQi              string `json:"jie_qi"`   //节气
	YinYang            string `json:"yin_yang"` //阴阳遁
	N                  int    `json:"jn"`       //定局数字
	YUAN               string `json:"yuan"`     //元
	XS                 string `json:"xs"`       //旬首
	ZHIFU              string `json:"zhifu"`    //值符
	ZHISHI             string `json:"zhishi"`   //值使

	G1 []string `json:"g_1"` //一宫 九星 八门 暗干支 天盘奇仪 八神 地盘奇仪
	G2 []string `json:"g_2"` //二宫 九星 八门 暗干支 天盘奇仪 八神 地盘奇仪
	G3 []string `json:"g_3"` //三宫 九星 八门 暗干支 天盘奇仪 八神 地盘奇仪
	G4 []string `json:"g_4"` //四宫 九星 八门 暗干支 天盘奇仪 八神 地盘奇仪
	G5 []string `json:"g_5"` //五宫 九星 八门 暗干支 天盘奇仪 八神 地盘奇仪
	G6 []string `json:"g_6"` //六宫 九星 八门 暗干支 天盘奇仪 八神 地盘奇仪
	G7 []string `json:"g_7"` //七宫 九星 八门 暗干支 天盘奇仪 八神 地盘奇仪
	G8 []string `json:"g_8"` //八宫 九星 八门 暗干支 天盘奇仪 八神 地盘奇仪
	G9 []string `json:"g_9"` //九宫 九星 八门 暗干支 天盘奇仪 八神 地盘奇仪
}

func NewSjqm(y, m, d, h int) *SJQM {
	gzo := gz.NewGanZhi(y, m, d, h)
	ygz := gzo.YGZ
	mgz := gzo.MGZ
	dgz := gzo.DGZ
	hgz := gzo.HGZ
	yjzhi, yjname, yjzt := gzo.GetYueJiangName(y, m, d)
	st := time.Date(y, time.Month(m), d, h, 0, 0, 0, time.Local)

	//定节气
	index, jmc, jqt := FindJq(st)
	xzt := jqt[12] //夏至时间戳
	dzt := jqt[24]

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
	juN := FindJU(yn, index, JqName)
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

	obj := &SJQM{
		ygz:     ygz,
		mgz:     mgz,
		dgz:     dgz,
		hgz:     hgz,
		YjZhi:   yjzhi,
		YjName:  yjname,
		Yjt:     yjzt.Format("2006-01-02"),
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
	return obj
}

//地四户
func (qm *SJQM) DiSiHu() string {
	return diSiHu(qm.hgz)
}

//地私门
func (qm *SJQM) DiSiMen() string {
	return diSiMen(qm.YjZhi, qm.dgz, qm.hgz)
}

//太冲天马
func (qm *SJQM) TaiChongTianMa() string {
	return tianMa(qm.YjZhi, qm.hgz)
}

//天三门
func (qm *SJQM) TianSanMen() string {
	return tianSanMen(qm.YjZhi, qm.hgz)
}

//五符
func (qm *SJQM) WuFu() []string {
	return wuFu(qm.dgz)
}

//时孤虚
func (qm *SJQM) GuXuH() string {
	return guXuH(qm.dgz, qm.hgz)
}
