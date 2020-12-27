package sjqm

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
	"time"
)

var (
	jqNameArr = []string{
		"冬至", "小寒", "大寒", "立春", "雨水", "惊蛰", "春分", "清明", "谷雨", "立夏", "小满", "芒种",
		"夏至", "小暑", "大暑", "立秋", "处暑", "白露", "秋分", "寒露", "霜降", "立冬", "小雪", "大雪",
		"冬至", "小寒", "大寒", "立春"}
	jzArr = []string{
		"甲子", "乙丑", "丙寅", "丁卯", "戊辰", "己巳", "庚午", "辛未", "壬申", "癸酉", //0-9  甲子干支时辰
		"甲戌", "乙亥", "丙子", "丁丑", "戊寅", "己卯", "庚辰", "辛巳", "壬午", "癸未", //10-19 甲戌干支时辰
		"甲申", "乙酉", "丙戌", "丁亥", "戊子", "己丑", "庚寅", "辛卯", "壬辰", "癸巳", //20-29 甲申干支时辰
		"甲午", "乙未", "丙申", "丁酉", "戊戌", "己亥", "庚子", "辛丑", "壬寅", "癸卯", //30-39 甲午干支时辰
		"甲辰", "乙巳", "丙午", "丁未", "戊申", "己酉", "庚戌", "辛亥", "壬子", "癸丑", //40-49 甲辰干支时辰
		"甲寅", "乙卯", "丙辰", "丁巳", "戊午", "己未", "庚申", "辛酉", "壬戌", "癸亥", //50-59 甲寅干支时辰
	}

	xsArr = []string{"甲子", "甲戌", "甲申", "甲午", "甲辰", "甲寅"} //旬首数组
)

func TestFindJQ(t *testing.T) {
	st := time.Date(2020, time.Month(12), 21, 19, 0, 0, 0, time.Local)
	jqnames := []string{"冬至", "小寒", "大寒", "立春", "雨水", "惊蛰", "春分", "清明", "谷雨", "立夏",
		"小满", "芒种", "夏至", "小暑", "大暑", "立秋", "处暑", "白露", "秋分", "寒露", "霜降", "立冬",
		"小雪", "大雪", "冬至", "小寒", "大寒", "立春"}
	jqt := []time.Time{
		time.Date(2019, time.Month(12), 22, 12, 0, 0, 0, time.Local),
		time.Date(2020, time.Month(1), 06, 05, 0, 0, 0, time.Local),
		time.Date(2020, time.Month(1), 20, 22, 0, 0, 0, time.Local),
		time.Date(2020, time.Month(2), 04, 17, 0, 0, 0, time.Local),
		time.Date(2020, time.Month(2), 19, 12, 0, 0, 0, time.Local),
		time.Date(2020, time.Month(3), 05, 10, 0, 0, 0, time.Local),
		time.Date(2020, time.Month(3), 20, 11, 0, 0, 0, time.Local),
		time.Date(2020, time.Month(4), 04, 15, 0, 0, 0, time.Local),
		time.Date(2020, time.Month(4), 19, 22, 0, 0, 0, time.Local),
		time.Date(2020, time.Month(5), 05, 8, 0, 0, 0, time.Local),
		time.Date(2020, time.Month(5), 20, 21, 0, 0, 0, time.Local),
		time.Date(2020, time.Month(6), 05, 12, 0, 0, 0, time.Local),
		time.Date(2020, time.Month(6), 21, 05, 0, 0, 0, time.Local),
		time.Date(2020, time.Month(7), 06, 23, 0, 0, 0, time.Local),
		time.Date(2020, time.Month(7), 22, 16, 0, 0, 0, time.Local),
		time.Date(2020, time.Month(8), 7, 9, 0, 0, 0, time.Local),
		time.Date(2020, time.Month(8), 22, 23, 0, 0, 0, time.Local),
		time.Date(2020, time.Month(9), 07, 12, 0, 0, 0, time.Local),
		time.Date(2020, time.Month(9), 22, 21, 0, 0, 0, time.Local),
		time.Date(2020, time.Month(10), 8, 03, 0, 0, 0, time.Local),
		time.Date(2020, time.Month(10), 23, 07, 0, 0, 0, time.Local),
		time.Date(2020, time.Month(11), 07, 07, 0, 0, 0, time.Local),
		time.Date(2020, time.Month(11), 22, 04, 0, 0, 0, time.Local),
		time.Date(2020, time.Month(12), 07, 00, 0, 0, 0, time.Local),
		time.Date(2020, time.Month(12), 21, 18, 0, 0, 0, time.Local),
		time.Date(2021, time.Month(01), 05, 11, 0, 0, 0, time.Local),
		time.Date(2021, time.Month(01), 20, 04, 0, 0, 0, time.Local),
		time.Date(2021, time.Month(02), 03, 22, 0, 0, 0, time.Local),
	}
	windex := 24
	wjmc := "冬至"

	index, jmc := FindJQ(st, jqt, jqnames)
	if index != windex {
		t.Errorf("FindJQ(%v,%v,%v)=%d 期望值 %d\n", st, jqt, jqnames, index, windex)
	}
	if !strings.EqualFold(jmc, wjmc) {
		t.Errorf("FindJQ(%v,%v,%v)=%s 期望值 %s\n", st, jqt, jqnames, jmc, wjmc)
	}
}

//定元 测试
func Test定元(t *testing.T) {
	//上元 二十天
	s上元 := []string{"甲子", "乙丑", "丙寅", "丁卯", "戊辰", "己卯", "庚辰", "辛巳", "壬午", "癸未", "甲午", "乙未", "丙申", "丁酉", "戊戌", "己酉", "庚戌", "辛亥", "壬子", "癸丑"}
	//中元 二十天
	s中元 := []string{"己巳", "庚午", "辛未", "壬申", "癸酉", "甲申", "乙酉", "丙戌", "丁亥", "戊子", "己亥", "庚子", "辛丑", "壬寅", "癸卯", "甲寅", "乙卯", "丙辰", "丁巳", "戊午"}
	//下元 二十天
	s下元 := []string{"甲戌", "乙亥", "丙子", "丁丑", "戊寅", "己丑", "庚寅", "辛卯", "壬辰", "癸巳", "甲辰", "乙巳", "丙午", "丁未", "戊申", "己未", "庚申", "辛酉", "壬戌", "癸亥"}

	//预期值
	expected1 := "上元"
	expected2 := "中元"
	expected3 := "下元"

	n1, n2, n3 := 0, 1, 2

	//上元预期值测试
	for i := 0; i < len(s上元); i++ {
		actual, nx1 := YuanN(s上元[i])
		if !strings.EqualFold(actual, expected1) {
			t.Errorf("YuanN(%s)=%s 预期结果:%s", s上元[i], actual, expected1)
		}
		if nx1 != n1 {
			t.Errorf("YuanN(%s)=%d 预期结果:%d\n", s上元[i], nx1, n1)
		}
	}

	//中元预期值测试
	for j := 0; j < len(s中元); j++ {
		actual, nx2 := YuanN(s中元[j])
		if !strings.EqualFold(actual, expected2) {
			t.Errorf("YuanN(%s)=%s 预期结果:%s", s中元[j], actual, expected2)
		}
		if nx2 != n2 {
			t.Errorf("YuanN(%s)=%d 预期结果:%d\n", s中元[j], nx2, n2)
		}
	}

	//下元预期值测试
	for k := 0; k < len(s下元); k++ {
		actual, nx3 := YuanN(s下元[k])
		if !strings.EqualFold(actual, expected3) {
			t.Errorf("YuanN(%s)=%s 预期结果:%s\n", s下元[k], actual, expected3)
		}

		if nx3 != n3 {
			t.Errorf("YuanN(%s)=%d 预期结果:%d\n", s上元[k], nx3, n3)
		}
	}
}

//定局 测试
func Test定局(t *testing.T) {
	/*
		expy1, expy2, expy3 := 0, 1, 2 //定元数字
		index0 := 0                    //

		name0 := "冬至"
		dz1, dz2, dz3 := 1, 7, 4

		if strings.EqualFold(name0, jqNameArr[index0]) || strings.EqualFold(name0, jqNameArr[24]) {

			n1 := FindJU(expy1, index0, jqNameArr)
			if n1 != dz1 {
				t.Errorf("FindJu(%d %d %s)=%d 预期值:%d\n", expy1, index0, jqNameArr, n1, dz1)
			}

			n2 := FindJU(expy2, index0, jqNameArr)
			if n2 != dz2 {
				t.Errorf("FindJu(%d %d %s)=%d 预期值:%d\n", expy2, index0, jqNameArr, n2, dz2)
			}

			n3 := FindJU(expy3, index0, jqNameArr)
			if n3 != dz3 {
				t.Errorf("FindJu(%d %d %s)=%d 预期值:%d\n", expy3, index0, jqNameArr, n3, dz3)
			}
		}
	*/
	//////////////////
	expy := []int{0, 1, 2} //三元数组
	//index := []int{0, 1, 2} ==>i           //节气索引数组
	name := []string{
		"冬至", "小寒", "大寒", "立春", "雨水", "惊蛰", "春分", "清明", "谷雨", "立夏", "小满", "芒种",
		"夏至", "小暑", "大暑", "立秋", "处暑", "白露", "秋分", "寒露", "霜降", "立冬", "小雪", "大雪",
		"冬至", "小寒", "大寒", "立春"} //节气名称

	//局数字
	type jqn struct {
		n1, n2, n3 int //n1:上元 n2:中元 n3:下元
	}
	//节气三元定局数字
	exptYuan := []jqn{
		{1, 7, 4}, //冬至
		{2, 8, 5}, //小寒
		{3, 9, 6}, //大寒
		{8, 5, 2}, //立春
		{9, 6, 3}, //雨水
		{1, 7, 4}, //惊蛰
		{3, 9, 6}, //春分
		{4, 1, 7}, //清明
		{5, 2, 8}, //谷雨
		{4, 1, 7}, //立夏
		{5, 2, 8}, //小满
		{6, 3, 9}, //芒种

		{9, 3, 6}, //夏至
		{8, 2, 5}, //小暑
		{7, 1, 4}, //大暑
		{2, 5, 8}, //立秋
		{1, 4, 7}, //处暑
		{9, 3, 6}, //白露
		{7, 1, 4}, //秋分
		{6, 9, 3}, //寒露
		{5, 8, 2}, //霜降
		{6, 9, 3}, //立冬
		{5, 8, 2}, //小雪
		{4, 7, 1}, //大雪
	}

	for i := 0; i < len(exptYuan); i++ {
		if !strings.EqualFold(name[i], jqNameArr[i]) {
			t.Errorf("FindJU(_,%d,_) 节气:%s 预期节气:%s\n", i, name[i], jqNameArr[i])
		} else if strings.EqualFold(name[i], jqNameArr[i]) {
			for j := 0; j < len(expy); j++ {
				switch j {
				case 0:
					jx := FindJU(expy[j], i, jqNameArr)
					if jx != exptYuan[i].n1 {
						t.Errorf("FindJu(%d %d %s)=%d 预期值:%d\n", expy[j], i, jqNameArr, jx, exptYuan[i].n1)
					}
				case 1:
					jx := FindJU(expy[j], i, jqNameArr)
					if jx != exptYuan[i].n2 {
						t.Errorf("FindJu(%d %d %s)=%d 预期值:%d\n", expy[j], i, jqNameArr, jx, exptYuan[i].n2)
					}
				case 2:
					jx := FindJU(expy[j], i, jqNameArr)
					if jx != exptYuan[i].n3 {
						t.Errorf("FindJu(%d %d %s)=%d 预期值:%d\n", expy[j], i, jqNameArr, jx, exptYuan[i].n3)
					}
				}
			}
			break
		}
	}
}

//三奇六仪 测试
func Test三奇六仪(t *testing.T) {
	///////阴遁 阴遁逆布六仪顺布三奇
	yyN := 0
	//jN := 1//戊 乙 丙 丁 癸 壬 辛 庚 己    (1,2,3--->9)
	//jN := 2//己 戊 乙 丙 丁 癸 壬 辛 庚
	//jN := 3//庚 己 戊 乙 丙 丁 癸 壬 辛
	//jN := 4//辛 庚 己 戊 乙 丙 丁 癸 壬
	//jN := 5//壬 辛 庚 己 戊 乙 丙 丁 癸
	//jN := 6//癸 壬 辛 庚 己 戊 乙 丙 丁
	//jN := 7//丁 癸 壬 辛 庚 己 戊 乙 丙
	//jN := 8//丙 丁 癸 壬 辛 庚 己 戊 乙
	//jN := 9//乙 丙 丁 癸 壬 辛 庚 己 戊
	jN := 4
	s4 := map[int]string{1: "辛", 2: "庚", 3: "己", 4: "戊", 5: "乙", 6: "丙", 7: "丁", 8: "癸", 9: "壬"}
	qy := FindSqly(yyN, jN)
	if !reflect.DeepEqual(qy, s4) {
		t.Errorf("FindSqly(%d %d)=%v 预期结果:%v\n", yyN, jN, qy, s4)
	}
	///////阳遁 顺布六仪逆布三奇
	//yjN := 1 1:"戊",2:"己",3:"庚",4:"辛",5:"壬",6:"癸",7:"丁",8:"丙",9:"乙"
	//yjN := 2 1:"乙",2:"戊",3:"己",4:"庚",5:"辛",6:"壬",7:"癸",8:"丁",9:"丙"
	//yjN := 3 1:"丙",2:"乙",3:"戊",4:"己",5:"庚",6:"辛",7:"壬",8:"癸",9:"丁"
	//yjN := 4 1:"丁",2:"丙",3:"乙",4:"戊",5:"己",6:"庚",7:"辛",8:"壬",9:"癸"
	//yjN := 5 1:"癸",2:"丁",3:"丙",4:"乙",5:"戊",6:"己",7:"庚",8:"辛",9:"壬"
	//yjN := 6 1:"壬",2:"癸",3:"丁",4:"丙",5:"乙",6:"戊",7:"己",8:"庚",9:"辛"
	//yjN := 7 1:"辛",2:"壬",3:"癸",4:"丁",5:"丙",6:"乙",7:"戊",8:"己",9:"庚"
	//yjN := 8 1:"庚",2:"辛",3:"壬",4:"癸",5:"丁",6:"丙",7:"乙",8:"戊",9:"己"
	//yjN := 9 1:"己",2:"庚",3:"辛",4:"壬",5:"癸",6:"丁",7:"丙",8:"乙",9:"戊"
	yangN := 1
	yjN := 1
	s1 := map[int]string{1: "戊", 2: "己", 3: "庚", 4: "辛", 5: "壬", 6: "癸", 7: "丁", 8: "丙", 9: "乙"}

	yqy := FindSqly(yangN, yjN)
	if !reflect.DeepEqual(yqy, s1) {
		t.Errorf("FindSqly(%d %d)=%v 预期结果:%v\n", yangN, yjN, yqy, s1)
	}
}

//旬首 测试
func Test旬首(t *testing.T) {

	for i := 0; i < len(jzArr); i++ {
		if i < 10 {
			expHgzArr := jzArr[:10]
			actual := xsArr[0]
			if xs, _ := XunShou(expHgzArr[i]); !strings.EqualFold(xs, actual) {
				t.Errorf("XunShou(%s)=%s 预期值:%s\n", expHgzArr[i], xs, actual)
			}
		}

		if i >= 10 && i < 20 {
			expHgzArr := jzArr[10:20]
			actual := xsArr[1]
			if xs, _ := XunShou(expHgzArr[i-10]); !strings.EqualFold(xs, actual) {
				t.Errorf("XunShou(%s)=%s 预期值:%s\n", expHgzArr[i-10], xs, actual)
			}
		}

		if i >= 20 && i < 30 {
			expHgzArr := jzArr[20:30]
			actual := xsArr[2]
			if xs, _ := XunShou(expHgzArr[i-20]); !strings.EqualFold(xs, actual) {
				t.Errorf("XunShou(%s)=%s 预期值:%s\n", expHgzArr[i-20], xs, actual)
			}
		}

		if i >= 30 && i < 40 {
			expHgzArr := jzArr[30:40]
			actual := xsArr[3]
			if xs, _ := XunShou(expHgzArr[i-30]); !strings.EqualFold(xs, actual) {
				t.Errorf("XunShou(%s)=%s 预期值:%s\n", expHgzArr[i-30], xs, actual)
			}
		}

		if i >= 40 && i < 50 {
			expHgzArr := jzArr[40:50]
			actual := xsArr[4]
			if xs, _ := XunShou(expHgzArr[i-40]); !strings.EqualFold(xs, actual) {
				t.Errorf("XunShou(%s)=%s 预期值:%s\n", expHgzArr[i-40], xs, actual)
			}
		}

		if i >= 50 && i < 60 {
			expHgzArr := jzArr[50:60]
			actual := xsArr[5]
			if xs, _ := XunShou(expHgzArr[i-50]); !strings.EqualFold(xs, actual) {
				t.Errorf("XunShou(%s)=%s 预期值:%s\n", expHgzArr[i-50], xs, actual)
			}
		}
	}
}

//寻找值符值使 测试
func Test值符随时干(t *testing.T) {
	//六甲旬遁: 甲子遁戊　甲戌遁己　甲申遁庚　甲午遁辛　甲辰遁壬　甲寅遁
	xsGZArr1 := jzArr[:10]
	//xsGZArr2 := jzArr[10:20]
	//xsGZArr3 := jzArr[20:30]
	xsGZArr4 := jzArr[30:40]
	//xsGZArr5 := jzArr[40:50]
	//xsGZArr6 := jzArr[50:]

	旬首1 := "甲子"
	//旬首2 := "甲戌"
	//旬首3 := "甲申"
	旬首4 := "甲午"
	//旬首5 := "甲辰"
	//旬首6 := "甲寅"
	hgz := "甲午"
	hgz1 := "甲子"

	////
	name := hgz
	name1 := hgz1
	zfN := 2
	zfN1 := 8
	sqlyY8 := map[int]string{1: "庚", 2: "辛", 3: "壬", 4: "癸", 5: "丁", 6: "丙", 7: "乙", 8: "戊", 9: "己"} //阳遁8局
	zfNumber, zfName, _ := XunShouHour(旬首4, hgz, xsGZArr4, sqlyY8)                                   //阳遁上元八局 甲午时辰

	if zfNumber != zfN || !strings.EqualFold(zfName, name) {
		t.Errorf("ZhiFuHour(%v %v %v %v)=%v,%v 期望结果:%v %v\n", 旬首4, hgz, xsGZArr4, sqlyY8, zfNumber, zfName, zfN, name)
	}

	zfn1, zfname1, _ := XunShouHour(旬首1, hgz1, xsGZArr1, sqlyY8) //阳遁上元八局 甲子时辰
	if zfn1 != zfN1 || !strings.EqualFold(zfname1, name1) {
		t.Errorf("ZhiFuHour(%v %v %v %v)=%v-%v 期望结果:%v-%v\n", 旬首1, hgz1, xsGZArr1, sqlyY8, zfn1, zfname1, zfN1, name1)
	}

}

//原宫位值符
func Test值符(t *testing.T) {

	//2020101906f -->庚子 丁亥 庚辰 辛巳
	//阴遁 上元 5局 旬首:甲戌 值符:天辅
	dun := "己"
	sqly := map[int]string{1: "壬", 2: "辛", 3: "庚", 4: "己", 5: "戊", 6: "乙", 7: "丙", 8: "丁", 9: "癸"}

	wantZF := "天辅"
	wantN := 4

	zf, N := SelfZF(dun, sqly)
	if !strings.EqualFold(zf, wantZF) {
		t.Errorf("SelfZF(%s,%v)=%s 希望结果:%s\n", dun, sqly, zf, wantZF)
	}
	if N != wantN {
		t.Errorf("SelfZF(%s,%v)=%d 希望结果:%d\n", dun, sqly, N, wantN)
	}
}

//九星排序
func Test九星排序(t *testing.T) {
	zhifu := "天辅"
	wantStarArr := []string{"天辅", "天英", "天芮", "天柱", "天心", "天蓬", "天任", "天冲", "天禽"}
	starArr := SortStar(zhifu)

	if !reflect.DeepEqual(starArr, wantStarArr) {
		t.Errorf("SortStar(%s)=%v 期望值:%v\n", zhifu, starArr, wantStarArr)
	}

}

//////////////
//方法
func Test门破(t *testing.T) {
	g := G{
		JieQi:   "小雪",
		YinYang: "阴遁",
		N:       2,
		YUAN:    "下元",
		XS:      "甲寅",
		ZHIFU:   "天心",
		ZHISHI:  "开门",
		G1:      []string{"天柱", "x", "己未", "己", "腾蛇", "己"},
		G2:      []string{"天辅", "杜门", "戊午", "戊", "白虎", "戊"},
		G3:      []string{"天蓬", "x", "丁巳", "乙", "九天", "乙"},
		G4:      []string{"天任", "x", "丙辰", "丙", "九地", "丙"},
		G5:      []string{"天禽", "", "乙卯", "丁", "", "丁"},
		G6:      []string{"天芮", "景门", "甲寅", "癸亥", "癸", "太阴", "癸"},
		G7:      []string{"天英", "x", "壬戌", "壬", "六合", "壬"},
		G8:      []string{"天心", "开门", "辛酉", "辛", "直符", "辛"},
		G9:      []string{"天冲", "x", "庚申", "庚", "玄武", "庚"},
	}
	want := "门破"
	act := g.G门破()
	if !strings.HasPrefix(act, want) {
		t.Errorf(" (g *G) G门破()=%s 期望结果:%s\n", act, want)
	}
}

//趋三避五
func TestGMap_G趋三避五(t *testing.T) {
	//2020102006f
	//庚子 丁亥 辛巳 癸巳
	//节气:小雪 阴遁 上元 5局 旬首:甲申 值符:天冲 值使:伤门

	zhis := "伤门"
	zsmap := map[int]string{1: "休门", 2: "死门", 3: "伤门", 4: "杜门", 6: "开门", 7: "惊门", 8: "生门", 9: "景门"}
	gmap := &GMap{ZhiShiMap: zsmap}
	want := "值使伤门落3宫 宜趋之吉"
	qbs := gmap.G趋三避五(zhis)
	if !reflect.DeepEqual(want, qbs) {
		t.Errorf("func (gmap *GMap) G趋三避五(%s)=%s 期望值:%s\n", zhis, qbs, want)
	}
}

//天遁
func Test天遁(t *testing.T) {
	g := G{
		JieQi:   "小雪",
		YinYang: "阴遁",
		N:       2,
		YUAN:    "下元",
		XS:      "甲寅",
		ZHIFU:   "天心",
		ZHISHI:  "开门",
		G1:      []string{"天柱", "开门", "己未", "丙", "腾蛇", "丁"},
		G2:      []string{"天辅", "杜门", "戊午", "戊", "白虎", "戊"},
		G3:      []string{"天蓬", "x", "丁巳", "乙", "九天", "乙"},
		G4:      []string{"天任", "x", "丙辰", "丙", "九地", "丙"},
		G5:      []string{"天禽", "", "乙卯", "丁", "", "丁"},
		G6:      []string{"天芮", "景门", "甲寅", "癸亥", "癸", "太阴", "癸"},
		G7:      []string{"天英", "x", "壬戌", "壬", "六合", "壬"},
		G8:      []string{"天心", "x", "辛酉", "辛", "直符", "辛"},
		G9:      []string{"天冲", "x", "庚申", "庚", "玄武", "庚"},
	}

	td := g.G天遁()
	want := "天遁"
	if !strings.EqualFold(td, want) {
		t.Errorf("g.G天遁()=%s 期望值:%s\n", td, want)
	}

}

func TestG九星休旺(t *testing.T) {
	starName := "天芮"
	want := struct {
		name                      string
		wang, xiang, si, qiu, xiu []string
	}{
		name:  starName,
		wang:  []string{"申", "酉"},
		xiang: []string{"辰", "未", "戌", "丑"},
		si:    []string{"巳", "午"},
		qiu:   []string{"寅", "卯"},
		xiu:   []string{"亥", "子"},
	}
	jxxw := NewJXXW(starName)
	jxs := *jxxw
	if !reflect.DeepEqual(jxs.Wang, want.wang) && !reflect.DeepEqual(jxs.StarName, want.name) {
		t.Errorf("func NewJXXW(%s)=%v 期望值:%v\n", starName, jxs, want)
	}
}

//九星吉凶
func TestJX(t *testing.T) {
	/*	ig := &G{
			G1: []string{"天柱", "开门", "己未", "丙", "腾蛇", "丁"},
			G2: []string{"天辅", "杜门", "戊午", "戊", "白虎", "戊"},
			G3: []string{"天蓬", "x", "丁巳", "乙", "九天", "乙"},
			G4: []string{"天任", "x", "丙辰", "丙", "九地", "丙"},
			G5: []string{"天禽", "", "乙卯", "丁", "", "丁"},
			G6: []string{"天芮", "景门", "甲寅", "癸亥", "癸", "太阴", "癸"},
			G7: []string{"天英", "x", "壬戌", "壬", "六合", "壬"},
			G8: []string{"天心", "x", "辛酉", "辛", "直符", "辛"},
			G9: []string{"天冲", "x", "庚申", "庚", "玄武", "庚"},
		}
		jxarr := ig.九星()
		want := []string{"天柱", "天心", "天蓬", "天任", "天冲", "天辅", "天英", "天芮", "天禽"}
		if !reflect.DeepEqual(want, jxarr) {
			t.Errorf("(g *G)九星()=%s 期望值:%s\n", jxarr, want)
		}*/
	tp := NewTianPeng("天蓬")
	fmt.Println(tp.Name, tp.Info)
}
