package main

import (
	"fmt"
	"github.com/Aquarian-Age/sjqm"
	"liangzi.local/cal"
	"log"
	"strconv"
	"time"
)

var (
	t                    = time.Now().Local()
	y                    = t.Year()
	m                    = int(t.Month())
	d                    = t.Day()
	h                    = t.Hour()
	ymds, timeS, ganZhiS string
	jieQiS, zhiFS        string
)

func convEdit(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Println("convEdit:", err)
	}
	return i
}

//干支历
func ymd(t time.Time) string {
	gz := cal.NewGanZhiInfo(t.Year(), int(t.Month()), t.Day(), t.Hour())
	ymds = gz.YearGZ + "年 " + gz.MonthGZ + "月 " + gz.DayGZ + "日 " + gz.HourGZ + "时"
	return ymds
}

//奇门
func qiMenS(y, m, d, h int) {
	gz := cal.NewGanZhiInfo(y, m, d, h)
	dgz := gz.DayGZ
	hgz := gz.HourGZ
	cust := time.Date(y, time.Month(m), d, h, 0, 0, 0, time.Local)
	g, _ := sjqm.Result(y, dgz, hgz, cust)

	jieQiS = fmt.Sprintf("\n%s %s %s %d局 旬首:%s", g.JieQi, g.YinYang, g.YUAN, g.N, g.XS)
	zhiFS = "\n值符:" + g.ZHIFU + " 值使:" + g.ZHISHI
	//0九星 1八门  2暗干支 3天盘奇仪 4八神 5地盘奇仪
	s1 = g.G1[0]
	m1 = g.G1[1]
	gzs1 = g.G1[2]
	sqy1 = g.G1[3]
	god1 = g.G1[4]
	eqy1 = g.G1[5]
	///
	s8 = g.G8[0]
	m8 = g.G8[1]
	gzs8 = g.G8[2]
	sqy8 = g.G8[3]
	god8 = g.G8[4]
	eqy8 = g.G8[5]
	////
	s3 = g.G3[0]
	m3 = g.G3[1]
	gzs3 = g.G3[2]
	sqy3 = g.G3[3]
	god3 = g.G3[4]
	eqy3 = g.G3[5]
	////
	s4 = g.G4[0]
	m4 = g.G4[1]
	gzs4 = g.G4[2]
	sqy4 = g.G4[3]
	god4 = g.G4[4]
	eqy4 = g.G4[5]
	///
	s5 = g.G5[0] //九星
	//m5 = g.G5[1]
	gzs5 = g.G5[2]
	eqy5 = g.G5[3]
	///
	s9 = g.G9[0]
	m9 = g.G9[1]
	gzs9 = g.G9[2]
	sqy9 = g.G9[3]
	god9 = g.G9[4]
	eqy9 = g.G9[5]
	///
	s2 = g.G2[0]
	m2 = g.G2[1]
	gzs2 = g.G2[2]
	sqy2 = g.G2[3]
	god2 = g.G2[4]
	eqy2 = g.G2[5]
	///
	s7 = g.G7[0]
	m7 = g.G7[1]
	gzs7 = g.G7[2]
	sqy7 = g.G7[3]
	god7 = g.G7[4]
	eqy7 = g.G7[5]
	///
	s6 = g.G6[0]
	m6 = g.G6[1]
	gzs6 = g.G6[2]
	sqy6 = g.G6[3]
	god6 = g.G6[4]
	eqy6 = g.G6[5]
}
