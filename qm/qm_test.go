package qm

import (
	"reflect"
	"strings"
	"testing"
)

//天三门
func TestTianSanMen(t *testing.T) {
	yj := "亥"   //"巳"
	hgz := "丙午" //"戊辰"
	want := map[string]string{"从魁": "辰", "太冲": "戌", "小吉": "寅"}
	smmap := TianSanMen(yj, hgz)
	if !reflect.DeepEqual(smmap, want) {
		t.Errorf("func TianSanMen(%s,%s)=%v 期望值:%v\n", yj, hgz, smmap, want)
	}
}

//地四户
func TestDiSiHu(t *testing.T) {
	//mgz := "丙戌"
	hgz := "辛巳"
	want := map[string]string{"危": "子", "定": "酉", "开": "卯", "除": "午"}
	sihumap := DiSiHu(hgz)
	if !reflect.DeepEqual(sihumap, want) {
		t.Errorf("func DiSiHu(%s)=%v 期望值:%v\n", hgz, sihumap, want)
	}

	//mgz1 := "甲寅"
	hgz1 := "庚寅"
	want1 := map[string]string{"危": "酉", "定": "午", "开": "子", "除": "卯"}
	sihumap1 := DiSiHu(hgz1)
	if !reflect.DeepEqual(sihumap1, want1) {
		t.Errorf("func DiSiHu(%s)=%v 期望值:%v\n", hgz1, sihumap, want1)
	}
}

//地私门测试
func TestDiSiMen(t *testing.T) {
	yj := "亥"
	dgz := "甲子"
	hgz := "戊辰"
	want := map[string]string{"六合": "卯", "太常": "申", "太阴": "戌"}
	dsmap := DiSiMen(yj, dgz, hgz)
	if !reflect.DeepEqual(dsmap, want) {
		t.Errorf("func DiSiMen(%s,%s,%s)=%v 期望值:%v\n", yj, dgz, hgz, dsmap, want)
	}

	yj2 := "巳"
	dgz2 := "乙丑"
	hgz2 := "辛巳"
	want2 := map[string]string{"六合": "巳", "太常": "子", "太阴": "戌"}
	dsmap2 := DiSiMen(yj2, dgz2, hgz2)
	if !reflect.DeepEqual(dsmap2, want2) {
		t.Errorf("func DiSiMen(%s,%s,%s)=%v 期望值:%v\n", yj2, dgz2, hgz2, dsmap2, want2)
	}

	yj1 := "亥"
	dgz1 := "甲子"
	hgz1 := "丁卯"
	want1 := map[string]string{"六合": "寅", "太常": "酉", "太阴": "未"}
	dsmap1 := DiSiMen(yj1, dgz1, hgz1)
	if !reflect.DeepEqual(dsmap1, want1) {
		t.Errorf("func DiSiMen(%s,%s,%s)=%v 期望值:%v\n", yj1, dgz1, hgz1, dsmap1, want1)
	}

	yj3 := "午"
	dgz3 := "辛酉"
	hgz3 := "甲午"
	want3 := map[string]string{"六合": "巳", "太常": "戌", "太阴": "子"}
	dsmap3 := DiSiMen(yj3, dgz3, hgz3)
	if !reflect.DeepEqual(dsmap3, want3) {
		t.Errorf("func DiSiMen(%s,%s,%s)=%v 期望值:%v\n", yj3, dgz3, hgz3, dsmap3, want3)
	}
}

//太冲天马
func TestTianMa(t *testing.T) {

	yjh := []struct{ yj, hgz string }{
		{yj: "亥", hgz: "丙寅"},
		{yj: "午", hgz: "甲午"},
		{yj: "子", hgz: "丁卯"},
		{yj: "丑", hgz: "戊辰"},
	}
	want := []string{"午", "卯", "午", "酉"}
	for i := 0; i < len(yjh); i++ {
		tms := TianMa(yjh[i].yj, yjh[i].hgz)
		if !strings.EqualFold(tms, want[i]) {

		}
	}
}

//五符
func TestWuFu(t *testing.T) {
	dgz := "乙丑"
	want := map[string]string{
		"丑": "地轴", "亥": "国印", "午": "风伯", "卯": "五符", "子": "天官", "寅": "天贼",
		"巳": "地符", "戌": "唐符", "未": "雷公", "申": "雨师", "辰": "天曹", "酉": "风云"}
	wfmap := WuFu(dgz)
	if !reflect.DeepEqual(wfmap, want) {
		t.Errorf("func WuFu(%s)=%v 期望值:%v\n", dgz, wfmap, want)
	}
}
