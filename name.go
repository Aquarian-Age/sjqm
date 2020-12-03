package sjqm

import "liangzi.local/nongli/ganzhi"

//六十甲子 0:甲子 1:乙丑 2:丙寅...59:癸亥
var jz60 = ganzhi.MakeJZ60()

//九宫原始信息
type JG struct {
	Name   string //八卦名
	Number int    //九宫数字 戴九履一 二四为肩 六八为足 左三右七
	Star   string //值符九星
	Door   string //值使八门
	WX     string //五行属性
}

//九宫排盘结果信息
type G struct {
	JieQi   string //节气
	YinYang string //阴阳遁
	N       int    //定宫数字
	YUAN    string //元
	XS      string //旬首
	ZHIFU   string //值符
	ZHISHI  string //值使

	G1 []string //一宫 九星 八门 暗干支 奇仪配九星 八神 六仪三奇
	G2 []string //二宫 九星 八门 暗干支 奇仪配九星 八神 六仪三奇
	G3 []string //三宫 九星 八门 暗干支 奇仪配九星 八神 六仪三奇
	G4 []string //四宫 九星 八门 暗干支 奇仪配九星 八神 六仪三奇
	G5 []string //五宫 九星 八门 暗干支 奇仪配九星 八神 六仪三奇
	G6 []string //六宫 九星 八门 暗干支 奇仪配九星 八神 六仪三奇
	G7 []string //七宫 九星 八门 暗干支 奇仪配九星 八神 六仪三奇
	G8 []string //八宫 九星 八门 暗干支 奇仪配九星 八神 六仪三奇
	G9 []string //九宫 九星 八门 暗干支 奇仪配九星 八神 六仪三奇
}

//一宫
type G1 struct {
	Number     int    //原始宫位数字 九宫 二四为肩(左4右2)；六八为足(左8右6) 戴九履一(492 816) 左三右七 五中宫
	Dun        string //六甲旬遁(排三奇六仪的旬遁)
	LiuJia     string //六甲即是暗干支
	LiuYiSanQi string //六仪三奇 阳顺阴逆
	ZhiFu      string //值符即是九星 值符随时辰天干
	ZhiShi     string //值使即是八门 值使随时辰地支

}

//二宫
type G2 struct {
	Number     int    //原始宫位数字 九宫 二四为肩(左4右2)；六八为足(左8右6) 戴九履一(492 816) 左三右七 五中宫
	Dun        string //六甲旬遁(排三奇六仪的旬遁)
	LiuJia     string //六甲即是暗干支
	LiuYiSanQi string //六仪 阳顺阴逆
	ZhiFu      string //值符即是九星 值符随时辰天干
	ZhiShi     string //值使即是八门 值使随时辰地支
}

//三宫
type G3 struct {
	Number     int    //原始宫位数字 九宫 二四为肩(左4右2)；六八为足(左8右6) 戴九履一(492 816) 左三右七 五中宫
	Dun        string //六甲旬遁(排三奇六仪的旬遁)
	LiuJia     string //六甲即是暗干支
	LiuYiSanQi string //六仪 阳顺阴逆
	ZhiFu      string //值符即是九星 值符随时辰天干
	ZhiShi     string //值使即是八门 值使随时辰地支
}

//四宫
type G4 struct {
	Number     int    //原始宫位数字 九宫 二四为肩(左4右2)；六八为足(左8右6) 戴九履一(492 816) 左三右七 五中宫
	Dun        string //六甲旬遁(排三奇六仪的旬遁)
	LiuJia     string //六甲即是暗干支
	LiuYiSanQi string //六仪 阳顺阴逆
	ZhiFu      string //值符即是九星 值符随时辰天干
	ZhiShi     string //值使即是八门 值使随时辰地支
}

//五宫
type G5 struct {
	Number     int    //原始宫位数字 九宫 二四为肩(左4右2)；六八为足(左8右6) 戴九履一(492 816) 左三右七 五中宫
	Dun        string //六甲旬遁(排三奇六仪的旬遁)
	LiuJia     string //六甲即是暗干支
	LiuYiSanQi string //六仪 阳顺阴逆
	ZhiFu      string //值符即是九星 值符随时辰天干
	ZhiShi     string //值使即是八门 值使随时辰地支
}

//六宫
type G6 struct {
	Number     int    //原始宫位数字 九宫 二四为肩(左4右2)；六八为足(左8右6) 戴九履一(492 816) 左三右七 五中宫
	Dun        string //六甲旬遁(排三奇六仪的旬遁)
	LiuJia     string //六甲即是暗干支
	LiuYiSanQi string //六仪 阳顺阴逆
	ZhiFu      string //值符即是九星 值符随时辰天干
	ZhiShi     string //值使即是八门 值使随时辰地支
}

//七宫
type G7 struct {
	Number     int    //原始宫位数字 九宫 二四为肩(左4右2)；六八为足(左8右6) 戴九履一(492 816) 左三右七 五中宫
	Dun        string //六甲旬遁(排三奇六仪的旬遁)
	LiuJia     string //六甲即是暗干支
	LiuYiSanQi string //六仪 阳顺阴逆
	ZhiFu      string //值符即是九星 值符随时辰天干
	ZhiShi     string //值使即是八门 值使随时辰地支
}

//八宫
type G8 struct {
	Number     int    //原始宫位数字 九宫 二四为肩(左4右2)；六八为足(左8右6) 戴九履一(492 816) 左三右七 五中宫
	Dun        string //六甲旬遁(排三奇六仪的旬遁)
	LiuJia     string //六甲即是暗干支
	LiuYiSanQi string //六仪 阳顺阴逆
	ZhiFu      string //值符即是九星 值符随时辰天干
	ZhiShi     string //值使即是八门 值使随时辰地支
}

//九宫
type G9 struct {
	Number     int    //原始宫位数字 九宫 二四为肩(左4右2)；六八为足(左8右6) 戴九履一(492 816) 左三右七 五中宫 九宫 二四为肩(左4右2)；六八为足(左8右6) 戴九履一(492 816) 左三右七 五中宫
	Dun        string //六甲旬遁(排三奇六仪的旬遁)
	LiuJia     string //六甲即是暗干支
	LiuYiSanQi string //六仪 阳顺阴逆
	ZhiFu      string //值符即是九星 值符随时辰天干
	ZhiShi     string //值使即是八门 值使随时辰地支
}
