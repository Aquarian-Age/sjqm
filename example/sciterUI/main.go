package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/sciter-sdk/go-sciter"
	"github.com/sciter-sdk/go-sciter/window"
	"liangzi.local/nongli/ccal"
	"liangzi.local/nongli/ganzhi"
	"liangzi.local/nongli/lunar"
	"liangzi.local/sjqm"
)

func main() {
	w, err := window.New(sciter.SW_TITLEBAR|sciter.SW_RESIZEABLE|sciter.SW_CONTROLS|sciter.SW_MAIN,
		&sciter.Rect{Left: 0, Top: 0, Right: 1280, Bottom: 800}, //设置初始窗口大小
	)
	if err != nil {
		log.Fatal(err)
	}
	//加载文件
	//w.LoadFile("ccal.html")
	w.LoadHtml(html, "")
	setWinHandler(w)
	w.Show()
	w.Run()
}

//#################################
func setWinHandler(w *window.Window) {
	w.DefineFunction("qimeninfo", qimeninfo)
	w.DefineFunction("ymdinfo", ymdinfo)
}

//奇门
func qimeninfo(args ...*sciter.Value) *sciter.Value {
	ly, lm, ld, lh, sx, lmb := args[0].String(), args[1].String(), args[2].String(), args[3].String(), args[4].String(), args[5].String()
	y, m, d, h, b := ConvStoInt(ly, lm, ld, lh, lmb)
	err, s, _, g, _ := ccal.Input(y, m, d, h, sx, b)
	if err != nil {
		log.Fatal("ccal-input:", err)
	}
	///时家奇门
	dgzm := fmt.Sprintf("%s%s", g.DayGanM, g.DayZhiM)
	hgzm := g.HourGanZhiM
	//这里的s.SHour是由输入的时辰转换而来
	st := time.Date(s.SYear, time.Month(s.SMonth), s.SDay, s.SHour, 0, 0, 0, time.Local)
	G, _ := sjqm.Result(y, dgzm, hgzm, st)
	byteg, err := json.Marshal(G)
	if err != nil {
		log.Fatal("奇门G", err)
	}
	jsG := string(byteg)
	return sciter.NewValue(jsG)
}

//纪年信息
func ymdinfo(args ...*sciter.Value) *sciter.Value {
	ly, lm, ld, lh, sx, lmb := args[0].String(), args[1].String(), args[2].String(), args[3].String(), args[4].String(), args[5].String()
	y, m, d, h, b := ConvStoInt(ly, lm, ld, lh, lmb)
	err, solar, lu, g, _ := ccal.Input(y, m, d, h, sx, b)
	if err != nil {
		log.Fatal("ccal-input:", err)
	}
	ygz := fmt.Sprintf("%s%s", g.YearGanM, g.YearZhiM) //年干支
	mgz := g.MonthGanZhiM                              //月干支
	dgz := fmt.Sprintf("%s%s", g.DayGanM, g.DayZhiM)   //日干支
	hgz := g.HourGanZhiM                               //时干支
	var aliasM string
	if lu.Leapmb == true {
		aliasM = "是"
	} else {
		aliasM = "否"
	}
	jng := fmt.Sprintf("干支纪年:%s年-%s月-%s日-%s时", ygz, mgz, dgz, hgz)
	jns := fmt.Sprintf("阳历纪年:%d年-%d月-%d日-周%s", solar.SYear, solar.SMonth, solar.SDay, solar.SWeek)
	jnl := fmt.Sprintf("农历纪年: %d年%s月(%s)%s %s时(%d时)",
		lu.LYear, lunar.Ymc[lu.LMonth-1], lu.LYdxs, lunar.Rmc[lu.LDay-1], lu.LaliasHour, lu.LHour)
	jnmb := fmt.Sprintf("本年是否有闰月:%s 闰%d月", aliasM, lu.LeapMonth)
	//纳音
	ygzny := ganzhi.GZ纳音(ygz)
	mgzny := ganzhi.GZ纳音(mgz)
	dgzny := ganzhi.GZ纳音(dgz)
	hgzny := ganzhi.GZ纳音(hgz)
	ny := fmt.Sprintf("干支纳音: %s %s %s %s", ygzny[ygz], mgzny[mgz], dgzny[dgz], hgzny[hgz])

	jn := JN{
		Solar:  jns,
		Lunar:  jnl,
		GanZhi: jng,
		Leapmb: jnmb,
		NaYin:  ny,
	}
	jnjson, err := json.Marshal(jn)
	if err != nil {
		log.Fatal(err)
	}
	return sciter.NewValue(string(jnjson))
}
func ConvStoInt(ys, ms, ds, hs, bs string) (int, int, int, int, bool) {
	y, err := strconv.Atoi(ys)
	if err != nil {
		log.Fatal("年份時間解析:", err)
	}

	m, err := strconv.Atoi(ms)
	if err != nil {
		log.Fatal("月份時間解析:", err)
	}
	d, err := strconv.Atoi(ds)
	if err != nil {
		log.Fatal("日期時間解析:", err)
	}
	h, err := strconv.Atoi(hs)
	if err != nil {
		log.Fatal("時辰解析:", err)
	}
	b, err := strconv.ParseBool(bs)
	if err != nil {
		log.Fatal("閏月bool解析:", err)
	}
	return y, m, d, h, b
}

//纪年信息
type JN struct {
	Solar  string `json:"solar"`
	Lunar  string `json:"lunar"`
	GanZhi string `json:"gan_zhi"`
	Leapmb string `json:"leapmb"`
	NaYin  string `json:"na_yin"`
}
