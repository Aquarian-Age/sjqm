package main

import (
	"fmt"
	"liangzi.local/sjqm/qm"
	ts "liangzi.local/ts/tongshu"
	"log"

	"liangzi.local/nongli/ccal"
	"liangzi.local/sjqm"
	"liangzi.local/sjqm/cmd"
)

func main() {
	input := cmd.GetInPut()

	err, s, l, gz, _ := ccal.Input(input.Y, input.M, input.D, input.H, "马", input.B)
	if err != nil {
		log.Fatalln(err)
	}

	y := l.LYear
	ygz := fmt.Sprintf("%s%s", gz.YearGanM, gz.YearZhiM)
	mgz := gz.MonthGanZhiM
	dgz := fmt.Sprintf("%s%s", gz.DayGanM, gz.DayZhiM)
	hgz := gz.HourGanZhiM
	st := s.SolarDayT
	g, gmap := sjqm.Result(y, dgz, hgz, st)
	fmt.Println(ygz, mgz, dgz, hgz)

	fmt.Printf("节气:%s %s %s %d局 旬首:%s 值符:%s 值使:%s\n"+
		"一宫 ==> 九星:%s 八门:%s 暗干支:%s 天盘奇仪:%s 八神:%s 地盘奇仪:%s\n"+
		"八宫 ==> 九星:%s 八门:%s 暗干支:%s 天盘奇仪:%s 八神:%s 地盘奇仪:%s\n"+
		"三宫 ==> 九星:%s 八门:%s 暗干支:%s 天盘奇仪:%s 八神:%s 地盘奇仪:%s\n"+
		"四宫 ==> 九星:%s 八门:%s 暗干支:%s 天盘奇仪:%s 八神:%s 地盘奇仪:%s\n"+
		"五宫 ==> 九星:%s 八门:%s 暗干支:%s 地盘奇仪:%s\n"+
		"九宫 ==> 九星:%s 八门:%s 暗干支:%s 天盘奇仪:%s 八神:%s 地盘奇仪:%s\n"+
		"二宫 ==> 九星:%s 八门:%s 暗干支:%s 天盘奇仪:%s 八神:%s 地盘奇仪:%s\n"+
		"七宫 ==> 九星:%s 八门:%s 暗干支:%s 天盘奇仪:%s 八神:%s 地盘奇仪:%s\n"+
		"六宫 ==> 九星:%s 八门:%s 暗干支:%s 天盘奇仪:%s 八神:%s 地盘奇仪:%s\n",
		g.JieQi, g.YinYang, g.YUAN, g.N, g.XS, g.ZHIFU, g.ZHISHI,
		g.G1[0], g.G1[1], g.G1[2], g.G1[3], g.G1[4], g.G1[5],
		g.G8[0], g.G8[1], g.G8[2], g.G8[3], g.G8[4], g.G8[5],
		g.G3[0], g.G3[1], g.G3[2], g.G3[3], g.G3[4], g.G3[5],
		g.G4[0], g.G4[1], g.G4[2], g.G4[3], g.G4[4], g.G4[5],
		g.G5[0], g.G5[1], g.G5[2], g.G5[3],
		g.G9[0], g.G9[1], g.G9[2], g.G9[3], g.G9[4], g.G9[5],
		g.G2[0], g.G2[1], g.G2[2], g.G2[3], g.G2[4], g.G2[5],
		g.G7[0], g.G7[1], g.G7[2], g.G7[3], g.G7[4], g.G7[5],
		g.G6[0], g.G6[1], g.G6[2], g.G6[3], g.G6[4], g.G6[5],
	)

	fmt.Println("--------------------------")
	///方法
	g.G门破()
	g.G天遁()
	//孤虚
	guxu := sjqm.HGuXu(dgz, hgz)
	if _, ok := guxu["孤"]; ok {
		fmt.Printf("孤:%s 虚:%s\n", guxu["孤"], guxu["虚"])
	}
	//	天门地户
	jqt := ts.JQT(l.LYear)
	solarT := s.SolarDayT
	yj := ts.NEWZRYLYueJiang(solarT, jqt)
	tmmap := qm.TianSanMen(yj.Zhi, hgz)
	fmt.Printf("天三门:\n月将:%s %s时: 从魁在%s为天门 小吉在%s为天门 太冲在%s为天门\n", yj.Name, hgz, tmmap["从魁"], tmmap["小吉"], tmmap["太冲"])

	sihumap := qm.DiSiHu(hgz)
	fmt.Printf("地四户:\n除在:%s 定在:%s 危在:%s 开在:%s\n", sihumap["除"], sihumap["定"], sihumap["危"], sihumap["开"])

	//趋三避五
	qbs := gmap.G趋三避五(g.ZHISHI)
	fmt.Println(qbs)

	//地私门
	dsmap := qm.DiSiMen(yj.Zhi, dgz, hgz)
	fmt.Printf("地私门: 六合:%s 太常:%s 太阴:%s\n", dsmap["六合"], dsmap["太常"], dsmap["太阴"])

}
