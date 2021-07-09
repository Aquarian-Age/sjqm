package sjqm

import (
	"github.com/starainrt/astro/basic"
	"github.com/starainrt/astro/calendar"
	"sort"
	"time"
)

var JqName = []string{
	"冬至", "小寒", "大寒", "立春", "雨水", "惊蛰",
	"春分", "清明", "谷雨", "立夏", "小满", "芒种",
	"夏至", "小暑", "大暑", "立秋", "处暑", "白露",
	"秋分", "寒露", "霜降", "立冬", "小雪", "大雪", "冬至",
}

//定节气
func FindJq(t time.Time) (int, string, []time.Time) {
	jqt := jqT(t.Year())
	jqt1 := jqT(t.Year() + 1)
	jqt = append(jqt, jqt1[1:]...)
	//
	var begint, endt time.Time
	var index int
	var jmc string
	//i=2从上年大寒开始  i=3立春  i=26大寒
	for i := 0; i < 27; i++ { //立春<<jqt范围<<大寒
		begint = jqt[i]
		begint = time.Date(begint.Year(), begint.Month(), begint.Day(), begint.Hour(), 0, 0, 0, time.Local)
		endt = jqt[i+1]
		endt = time.Date(endt.Year(), endt.Month(), endt.Day(), endt.Hour(), 0, 0, 0, time.Local)
		if (t.After(begint) || t.Equal(begint)) && t.Before(endt) { //当前时间前的一个节气
			index = i       //节气索引
			jmc = JqName[i] //当前时间所属的节气名称
			break
		}
	}
	jmct := jqt[index].Format("2006-01-02") //2021-07-23 时间显示07-22
	jmc = jmc + ": " + jmct

	return index, jmc, jqt
}
func jqT(year int) []time.Time {
	year -= 1 //k:1-->上一年冬至时间 k:25-->本年冬至时间 k:4--本年立春
	jq := basic.GetOneYearJQ(year)
	var keys []int
	for k := range jq {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	var jqt []time.Time //24节气
	for _, v := range keys {
		xt := calendar.JDE2Date(jq[v])
		xt = time.Date(xt.Year(), xt.Month(), xt.Day(), xt.Hour(), 0, 0, 0, time.Local)
		jqt = append(jqt, xt)
	}
	return jqt
}

//24节气时间戳数组　上一年冬至到本年冬至 arrt:0:上一年冬至 12:本年夏至时间 24:本年冬至
//12节气时间戳数组　0:上一年小寒
//12中气时间戳数组 0:上年冬至
func GetJieQiArr(year int) (time.Time, []time.Time, []time.Time, []time.Time) {
	year -= 1 //k:1-->上一年冬至时间 k:25-->本年冬至时间 k:4--本年立春
	jq := basic.GetOneYearJQ(year)
	var keys []int
	for k := range jq {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	var arr []time.Time    //24节气
	var zqArr []time.Time  //中气(0:上一年冬至 到本年冬至)
	var jieArr []time.Time //12节
	var lct time.Time      //立春

	for _, v := range keys {
		arr = append(arr, calendar.JDE2Date(jq[v]))
		if v%2 == 1 { //中气
			zqArr = append(zqArr, calendar.JDE2Date(jq[v]))
		}
		if v%2 == 0 { //节
			jieArr = append(jieArr, calendar.JDE2Date(jq[v]))
		}
		if v == 4 { //立春
			lct = calendar.JDE2Date(jq[v])
		}
	}

	return lct, arr, jieArr, zqArr
}
