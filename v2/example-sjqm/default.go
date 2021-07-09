package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/Aquarian-Age/sjqm/v2"
	"github.com/Aquarian-Age/xa/pkg/gz"
	"github.com/starainrt/astro/basic"
	"github.com/ying32/govcl/vcl"
)

//默认显示的时间信息
func (f *TMainForm) OnFormDefault(sender vcl.IObject) {
	const (
		height = 800
		b1     = 170
		top    = 10
		l1     = 10
	)
	ys := f.year.Text()[:4]
	y, _ := strconv.Atoi(ys)
	m := f.month.Value()
	d := f.day.Value()
	h := f.hour.Value()
	ct := time.Date(y, time.Month(m), int(d), int(h), 0, 0, 0, time.Local)
	w := int(ct.Weekday())
	wName := WeekName(w)
	solars := fmt.Sprintf("%d年%d月%d日%d点 %s", y, m, d, h, wName)
	gzo := gz.NewGanZhi(y, int(m), int(d), int(h))
	gzs := fmt.Sprintf("%s年%s月%s日%s时", gzo.YGZ, gzo.MGZ, gzo.DGZ, gzo.HGZ)
	_, _, _, moons := basic.GetLunar(y, int(m), int(d))
	ymds := solars + "\n" + gzs + "\n阴历: " + moons
	f.qm = sjqm.NewSjqm(y, int(m), int(d), int(h)) //0九星 1八门 2暗干支 3天盘奇仪 4八神 5地盘奇仪

	g1s := f.qm.G1[0] + " " + f.qm.G1[1] + " " + f.qm.G1[4] + "\n" + " " + f.qm.G1[3] + " " + f.qm.G1[5] + "\n" + f.qm.G1[2]
	g8s := f.qm.G8[0] + " " + f.qm.G8[1] + " " + f.qm.G8[4] + "\n" + " " + f.qm.G8[3] + " " + f.qm.G8[5] + "\n" + f.qm.G8[2]
	g3s := f.qm.G3[0] + " " + f.qm.G3[1] + " " + f.qm.G3[4] + "\n" + " " + f.qm.G3[3] + " " + f.qm.G3[5] + "\n" + f.qm.G3[2]
	g4s := f.qm.G4[0] + " " + f.qm.G4[1] + " " + f.qm.G4[4] + "\n" + " " + f.qm.G4[3] + " " + f.qm.G4[5] + "\n" + f.qm.G4[2]
	g9s := f.qm.G9[0] + " " + f.qm.G9[1] + " " + f.qm.G9[4] + "\n" + " " + f.qm.G9[3] + " " + f.qm.G9[5] + "\n" + f.qm.G9[2]
	g2s := f.qm.G2[0] + " " + f.qm.G2[1] + " " + f.qm.G2[4] + "\n" + " " + f.qm.G2[3] + " " + f.qm.G2[5] + "\n" + f.qm.G2[2]
	g7s := f.qm.G7[0] + " " + f.qm.G7[1] + " " + f.qm.G7[4] + "\n" + " " + f.qm.G7[3] + " " + f.qm.G7[5] + "\n" + f.qm.G7[2]
	g6s := f.qm.G6[0] + " " + f.qm.G6[1] + " " + f.qm.G6[4] + "\n" + " " + f.qm.G6[3] + " " + f.qm.G6[5] + "\n" + f.qm.G6[2]
	g5s := f.qm.G5[0] + " " + f.qm.G5[1] + " " + f.qm.G5[4] + "\n" + " " + f.qm.G5[3] + " " + f.qm.G5[5] + "\n" + f.qm.G5[2]

	s1 := fmt.Sprintf("\n%s\n%s %s %d局 旬首:%s", f.qm.JieQi, f.qm.YinYang, f.qm.YUAN, f.qm.N, f.qm.XS)
	s2 := "值符:" + f.qm.ZHIFU + " 值使:" + f.qm.ZHISHI
	label1s := ymds + "\n" + s1 + "\n" + s2
	//
	const step = 33
	//f.ymds = vcl.NewLabel(f)
	//f.ymds.SetParent(f)
	//f.ymds.SetTop(height - step)
	//f.ymds.SetLeft(9 * step)
	//f.ymds.Font().SetSize(13)
	//f.ymds.SetTextBuf(ymds)

	f.label1 = vcl.NewLabel(f)
	f.label1.SetParent(f)
	f.label1.SetTop(450)
	f.label1.SetLeft(l1)
	f.label1.SetWidth(420)
	f.label1.SetTextBuf(label1s)
	//
	f.g1 = vcl.NewLabel(f)
	f.g1.SetParent(f)
	f.g1.SetTop(2*b1 - 35)
	f.g1.SetLeft(b1)
	f.g1.SetWidth(b1)
	f.g1.SetTextBuf(g1s)

	f.g8 = vcl.NewLabel(f)
	f.g8.SetParent(f)
	f.g8.SetTop(2*b1 - 35)
	f.g8.SetLeft(l1)
	f.g8.SetWidth(b1)
	f.g8.SetTextBuf(g8s)

	f.g3 = vcl.NewLabel(f)
	f.g3.SetParent(f)
	f.g3.SetTop(b1)
	f.g3.SetLeft(l1)
	f.g3.SetWidth(b1)
	f.g3.SetTextBuf(g3s)

	f.g4 = vcl.NewLabel(f)
	f.g4.SetParent(f)
	f.g4.SetTop(top)
	f.g4.SetLeft(l1)
	f.g4.SetWidth(b1)
	f.g4.SetTextBuf(g4s)

	f.g9 = vcl.NewLabel(f)
	f.g9.SetParent(f)
	f.g9.SetTop(top)
	f.g9.SetLeft(b1)
	f.g9.SetWidth(b1)
	f.g9.SetTextBuf(g9s)

	f.g2 = vcl.NewLabel(f)
	f.g2.SetParent(f)
	f.g2.SetTop(top)
	f.g2.SetLeft(300)
	//f.g2.SetWidth(b1)
	f.g2.SetTextBuf(g2s)

	f.g7 = vcl.NewLabel(f)
	f.g7.SetParent(f)
	f.g7.SetTop(b1)
	f.g7.SetWidth(b1)
	f.g7.SetLeft(300)
	f.g7.SetTextBuf(g7s)

	f.g6 = vcl.NewLabel(f)
	f.g6.SetParent(f)
	f.g6.SetTop(2*b1 - 35)
	f.g6.SetLeft(300)
	f.g6.SetWidth(b1)
	f.g6.SetTextBuf(g6s)

	f.g5 = vcl.NewLabel(f)
	f.g5.SetParent(f)
	f.g5.SetTop(b1 + 3*top)
	f.g5.SetLeft(b1 + 4*l1)
	//f.g5.SetWidth(5 * l1)
	f.g5.SetTextBuf(g5s)
}
