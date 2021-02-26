package sjqm

import "liangzi.local/cal"

//六十甲子 0:甲子 1:乙丑 2:丙寅...59:癸亥
var Jz60 = cal.MakeJZ60()

//九宫原始信息
type JG struct {
	Name   string //八卦名
	Number int    //九宫数字 戴九履一 二四为肩 六八为足 左三右七
	Star   string //值符九星
	Door   string //值使八门
	WX     string //五行属性
	ZiBai  string //紫白
}

//九宫排盘结果信息
type G struct {
	JieQi   string `json:"jie_qi"`   //节气
	YinYang string `json:"yin_yang"` //阴阳遁
	N       int    `json:"n"`        //定局数字
	YUAN    string `json:"yuan"`     //元
	XS      string `json:"xs"`       //旬首
	ZHIFU   string `json:"zhifu"`    //值符
	ZHISHI  string `json:"zhishi"`   //值使

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

//九宫排盘结果信息(map)
type GMap struct {
	ZhiShiMap map[int]string `json:"zhi_shi_map"` //排盘结果的值使八门信息
}

//qm包结果信息(map)
type QM struct {
	DiSiHu     string   `json:"di_si_hu"`     //地四户
	DiSiMen    string   `json:"di_si_men"`    //地私门
	TianMa     string   `json:"tian_ma"`      //太冲天马
	TianSanMen string   `json:"tian_san_men"` //map[string]string //天三门
	WuFu       []string `json:"wu_fus"`       //五符
	GuXu       string   `json:"gu_xu"`        //时孤虚
}

//九星休旺
//旺:我生之月 相:月类之月 死:生我之月 囚:克我之月 休:我克之月
type JXXW struct {
	StarName string   //九星名称
	Wang     []string //旺
	Xiang    []string //相
	Si       []string //死
	Qiu      []string //囚
	Xiu      []string //休
}
