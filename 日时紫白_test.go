package sjqm

import (
	"fmt"
	"strings"
	"testing"
	"time"
)

//年紫白
func TestZiBaiY(t *testing.T) {
	zbn := []struct {
		ygz string
		n   int
	}{{ygz: "辛卯", n: 1}, {ygz: "壬申", n: 0}, {ygz: "癸亥", n: -1}}
	w := []int{1, 5, 2}

	for i := 0; i < len(zbn); i++ {
		x := ZiBaiY(zbn[i].n, zbn[i].ygz)
		if x != w[i] {
			t.Errorf("func ZiBaiY(%d %s)=%d 期望值:%d\n", zbn[i].n, zbn[i].ygz, x, w[i])
		}
	}
}

//月紫白
func TestZiBaiM(t *testing.T) {
	ygz := "庚子"
	mgz := "戊子"
	want := 7
	mx := ZiBaiM(ygz, mgz)
	if want != mx {
		t.Errorf("func ZiBaiM(%s %s)=%d want:%d\n", ygz, mgz, mx, want)
	}
	////
	zbm := []struct {
		ygz, mgz string
	}{{ygz: "庚子", mgz: "丁亥"}, {ygz: "辛丑", mgz: "己亥"}, {ygz: "壬寅", mgz: "辛亥"}, {ygz: "壬寅", mgz: "癸丑"}}
	w := []int{8, 5, 2, 9}

	for i := 0; i < len(zbm); i++ {
		mx := ZiBaiM(zbm[i].ygz, zbm[i].mgz)
		if w[i] != mx {
			t.Errorf("func ZiBaiM(%s %s)=%d want:%d\n", zbm[i].ygz, zbm[i].mgz, mx, w[i])
		}
	}
}

//日紫白
func TestZiBaiD(t *testing.T) {

	gzt := []struct {
		st  time.Time //阳历时间戳
		sy  int       //阳历年数字
		dgz string    //日干支
	}{
		//---->阳遁
		//下元
		{
			st:  time.Date(2020, 12, 27, 0, 0, 0, 0, time.Local),
			sy:  2020,
			dgz: "甲辰",
		},
		{
			st:  time.Date(2020, 12, 28, 0, 0, 0, 0, time.Local),
			sy:  2020,
			dgz: "乙巳",
		},
		{
			st:  time.Date(2020, 12, 29, 0, 0, 0, 0, time.Local),
			sy:  2020,
			dgz: "丙午",
		},
		{
			st:  time.Date(2020, 12, 30, 0, 0, 0, 0, time.Local),
			sy:  2020,
			dgz: "丁未",
		},
		{
			st:  time.Date(2020, 12, 31, 0, 0, 0, 0, time.Local),
			sy:  2020,
			dgz: "戊申",
		},
		//上元
		{
			st:  time.Date(2021, 1, 1, 0, 0, 0, 0, time.Local),
			sy:  2021,
			dgz: "己酉",
		},
		{
			st:  time.Date(2021, 1, 2, 0, 0, 0, 0, time.Local),
			sy:  2021,
			dgz: "庚戌",
		},
		{
			st:  time.Date(2021, 1, 3, 0, 0, 0, 0, time.Local),
			sy:  2021,
			dgz: "辛亥",
		},
		{
			st:  time.Date(2021, 1, 4, 0, 0, 0, 0, time.Local),
			sy:  2021,
			dgz: "壬子",
		},
		{
			st:  time.Date(2021, 1, 5, 0, 0, 0, 0, time.Local),
			sy:  2021,
			dgz: "癸丑",
		},
		//中元
		{
			st:  time.Date(2021, 1, 6, 0, 0, 0, 0, time.Local),
			sy:  2021,
			dgz: "甲寅",
		},
		{
			st:  time.Date(2021, 1, 7, 0, 0, 0, 0, time.Local),
			sy:  2021,
			dgz: "乙卯",
		},
		{
			st:  time.Date(2021, 1, 8, 0, 0, 0, 0, time.Local),
			sy:  2021,
			dgz: "丙辰",
		},
		{
			st:  time.Date(2021, 1, 9, 0, 0, 0, 0, time.Local),
			sy:  2021,
			dgz: "丁巳",
		},
		{
			st:  time.Date(2021, 1, 10, 0, 0, 0, 0, time.Local),
			sy:  2021,
			dgz: "戊午",
		},
		//--->阴遁
		//上元
		{
			st:  time.Date(2020, 6, 22, 0, 0, 0, 0, time.Local),
			sy:  2020,
			dgz: "丙申",
		},
		{
			st:  time.Date(2020, 6, 23, 0, 0, 0, 0, time.Local),
			sy:  2020,
			dgz: "丁酉",
		},
		{
			st:  time.Date(2020, 6, 24, 0, 0, 0, 0, time.Local),
			sy:  2020,
			dgz: "戊戌",
		},
		//中元
		{
			st:  time.Date(2020, 6, 25, 0, 0, 0, 0, time.Local),
			sy:  2020,
			dgz: "己亥",
		},
		{
			st:  time.Date(2020, 6, 26, 0, 0, 0, 0, time.Local),
			sy:  2020,
			dgz: "庚子",
		},
		{
			st:  time.Date(2020, 6, 27, 0, 0, 0, 0, time.Local),
			sy:  2020,
			dgz: "辛丑",
		},
		{
			st:  time.Date(2020, 6, 28, 0, 0, 0, 0, time.Local),
			sy:  2020,
			dgz: "壬寅",
		},
		{
			st:  time.Date(2020, 6, 29, 0, 0, 0, 0, time.Local),
			sy:  2020,
			dgz: "癸卯",
		},
		//下元
		{
			st:  time.Date(2020, 6, 30, 0, 0, 0, 0, time.Local),
			sy:  2020,
			dgz: "甲辰",
		},
		{
			st:  time.Date(2020, 7, 1, 0, 0, 0, 0, time.Local),
			sy:  2020,
			dgz: "乙巳",
		},
		{
			st:  time.Date(2020, 7, 2, 0, 0, 0, 0, time.Local),
			sy:  2020,
			dgz: "丙午",
		},
		{
			st:  time.Date(2020, 7, 3, 0, 0, 0, 0, time.Local),
			sy:  2020,
			dgz: "丁未",
		},
		{
			st:  time.Date(2020, 7, 4, 0, 0, 0, 0, time.Local),
			sy:  2020,
			dgz: "戊申",
		},
	}

	want := []int{8, 9, 1, 2, 3, 1, 2, 3, 4, 5, 3, 4, 5, 6, 7, 4, 3, 2, 4, 3, 2, 1, 9, 2, 1, 9, 8, 7}

	for i := 0; i < len(gzt); i++ {
		dx := ZiBaiD(gzt[i].dgz, gzt[i].st, gzt[i].sy)
		if want[i] != dx {
			t.Errorf("func ZiBaiD(%s %v %d)=%d want:%d\n", gzt[i].dgz, gzt[i].st, gzt[i].sy, dx, want[i])
		}
	}
}

//时辰紫白
func TestZiBaiH(t *testing.T) {
	gzt := []struct {
		st  time.Time //阳历时间戳
		sy  int       //阳历年数字
		dgz string    //日干支
		hgz []string  //时辰干支
	}{
		//---->阳遁
		//下元
		{
			st:  time.Date(2020, 12, 27, 0, 0, 0, 0, time.Local),
			sy:  2020,
			dgz: "甲辰",
			//五鼠遁 甲己还加甲
			hgz: []string{"甲子", "乙丑", "丙寅", "丁卯", "戊辰", "己巳", "庚午", "辛未", "壬申", "癸酉", "甲戌", "乙亥"},
		},
		{
			st:  time.Date(2020, 12, 28, 0, 0, 0, 0, time.Local),
			sy:  2020,
			dgz: "乙巳",
			//乙庚丙作初
			hgz: []string{"丙子", "丁丑", "戊寅", "己卯", "庚辰", "辛巳", "壬午", "癸未", "甲申", "乙酉", "丙戌", "丁亥"},
		},
		{
			st:  time.Date(2020, 12, 29, 0, 0, 0, 0, time.Local),
			sy:  2020,
			dgz: "丙午",
			//丙辛从戊起
			hgz: []string{"戊子", "己丑", "庚寅", "辛卯", "壬辰", "癸巳", "甲午", "乙未", "丙申", "丁酉", "戊戌", "己亥"},
		},
		{
			st:  time.Date(2020, 12, 30, 0, 0, 0, 0, time.Local),
			sy:  2020,
			dgz: "丁未",
			//丁壬庚子居
			hgz: []string{"庚子", "辛丑", "壬寅", "癸卯", "甲辰", "乙巳", "丙午", "丁未", "戊申", "己酉", "庚戌", "辛亥"},
		},
		{
			st:  time.Date(2020, 12, 31, 0, 0, 0, 0, time.Local),
			sy:  2020,
			dgz: "戊申",
			//戊癸从壬起
			hgz: []string{"壬子", "癸丑", "甲寅", "乙卯", "丙辰", "丁巳", "戊午", "己未", "庚申", "辛酉", "壬戌", "癸亥"},
		},
		//上元
		{
			st:  time.Date(2021, 1, 1, 0, 0, 0, 0, time.Local),
			sy:  2021,
			dgz: "己酉",
			hgz: []string{"甲子", "乙丑", "丙寅", "丁卯", "戊辰", "己巳", "庚午", "辛未", "壬申", "癸酉", "甲戌", "乙亥"},
		},
		{
			st:  time.Date(2021, 1, 2, 0, 0, 0, 0, time.Local),
			sy:  2021,
			dgz: "庚戌",
			hgz: []string{"丙子", "丁丑", "戊寅", "己卯", "庚辰", "辛巳", "壬午", "癸未", "甲申", "乙酉", "丙戌", "丁亥"},
		},
		{
			st:  time.Date(2021, 1, 3, 0, 0, 0, 0, time.Local),
			sy:  2021,
			dgz: "辛亥",
			hgz: []string{"戊子", "己丑", "庚寅", "辛卯", "壬辰", "癸巳", "甲午", "乙未", "丙申", "丁酉", "戊戌", "己亥"},
		},
		{
			st:  time.Date(2021, 1, 4, 0, 0, 0, 0, time.Local),
			sy:  2021,
			dgz: "壬子",
			hgz: []string{"庚子", "辛丑", "壬寅", "癸卯", "甲辰", "乙巳", "丙午", "丁未", "戊申", "己酉", "庚戌", "辛亥"},
		},
		{
			st:  time.Date(2021, 1, 5, 0, 0, 0, 0, time.Local),
			sy:  2021,
			dgz: "癸丑",
			hgz: []string{"壬子", "癸丑", "甲寅", "乙卯", "丙辰", "丁巳", "戊午", "己未", "庚申", "辛酉", "壬戌", "癸亥"},
		},
		//中元
		{
			st:  time.Date(2021, 1, 6, 0, 0, 0, 0, time.Local),
			sy:  2021,
			dgz: "甲寅",
			hgz: []string{"甲子", "乙丑", "丙寅", "丁卯", "戊辰", "己巳", "庚午", "辛未", "壬申", "癸酉", "甲戌", "乙亥"},
		},
		{
			st:  time.Date(2021, 1, 7, 0, 0, 0, 0, time.Local),
			sy:  2021,
			dgz: "乙卯",
			hgz: []string{"丙子", "丁丑", "戊寅", "己卯", "庚辰", "辛巳", "壬午", "癸未", "甲申", "乙酉", "丙戌", "丁亥"},
		},
		{
			st:  time.Date(2021, 1, 8, 0, 0, 0, 0, time.Local),
			sy:  2021,
			dgz: "丙辰",
			hgz: []string{"戊子", "己丑", "庚寅", "辛卯", "壬辰", "癸巳", "甲午", "乙未", "丙申", "丁酉", "戊戌", "己亥"},
		},
		{
			st:  time.Date(2021, 1, 9, 0, 0, 0, 0, time.Local),
			sy:  2021,
			dgz: "丁巳",
			hgz: []string{"庚子", "辛丑", "壬寅", "癸卯", "甲辰", "乙巳", "丙午", "丁未", "戊申", "己酉", "庚戌", "辛亥"},
		},

		{
			st:  time.Date(2021, 1, 10, 0, 0, 0, 0, time.Local),
			sy:  2021,
			dgz: "戊午",
			hgz: []string{"壬子", "癸丑", "甲寅", "乙卯", "丙辰", "丁巳", "戊午", "己未", "庚申", "辛酉", "壬戌", "癸亥"},
		},

		//--->阴遁
		//上元
		{
			st:  time.Date(2020, 6, 22, 0, 0, 0, 0, time.Local),
			sy:  2020,
			dgz: "丙申",
			hgz: []string{"戊子", "己丑", "庚寅", "辛卯", "壬辰", "癸巳", "甲午", "乙未", "丙申", "丁酉", "戊戌", "己亥"},
		},

		{
			st:  time.Date(2020, 6, 23, 0, 0, 0, 0, time.Local),
			sy:  2020,
			dgz: "丁酉",
			hgz: []string{"庚子", "辛丑", "壬寅", "癸卯", "甲辰", "乙巳", "丙午", "丁未", "戊申", "己酉", "庚戌", "辛亥"},
		},

		{
			st:  time.Date(2020, 6, 24, 0, 0, 0, 0, time.Local),
			sy:  2020,
			dgz: "戊戌",
			hgz: []string{"壬子", "癸丑", "甲寅", "乙卯", "丙辰", "丁巳", "戊午", "己未", "庚申", "辛酉", "壬戌", "癸亥"},
		},

		//中元
		{
			st:  time.Date(2020, 6, 25, 0, 0, 0, 0, time.Local),
			sy:  2020,
			dgz: "己亥",
			hgz: []string{"甲子", "乙丑", "丙寅", "丁卯", "戊辰", "己巳", "庚午", "辛未", "壬申", "癸酉", "甲戌", "乙亥"},
		},

		{
			st:  time.Date(2020, 6, 26, 0, 0, 0, 0, time.Local),
			sy:  2020,
			dgz: "庚子",
			hgz: []string{"丙子", "丁丑", "戊寅", "己卯", "庚辰", "辛巳", "壬午", "癸未", "甲申", "乙酉", "丙戌", "丁亥"},
		},

		{
			st:  time.Date(2020, 6, 27, 0, 0, 0, 0, time.Local),
			sy:  2020,
			dgz: "辛丑",
			hgz: []string{"戊子", "己丑", "庚寅", "辛卯", "壬辰", "癸巳", "甲午", "乙未", "丙申", "丁酉", "戊戌", "己亥"},
		},

		{
			st:  time.Date(2020, 6, 28, 0, 0, 0, 0, time.Local),
			sy:  2020,
			dgz: "壬寅",
			hgz: []string{"庚子", "辛丑", "壬寅", "癸卯", "甲辰", "乙巳", "丙午", "丁未", "戊申", "己酉", "庚戌", "辛亥"},
		},

		{
			st:  time.Date(2020, 6, 29, 0, 0, 0, 0, time.Local),
			sy:  2020,
			dgz: "癸卯",
			hgz: []string{"壬子", "癸丑", "甲寅", "乙卯", "丙辰", "丁巳", "戊午", "己未", "庚申", "辛酉", "壬戌", "癸亥"},
		},

		//下元
		{
			st:  time.Date(2020, 6, 30, 0, 0, 0, 0, time.Local),
			sy:  2020,
			dgz: "甲辰",
			hgz: []string{"甲子", "乙丑", "丙寅", "丁卯", "戊辰", "己巳", "庚午", "辛未", "壬申", "癸酉", "甲戌", "乙亥"},
		},

		{
			st:  time.Date(2020, 7, 1, 0, 0, 0, 0, time.Local),
			sy:  2020,
			dgz: "乙巳",
			hgz: []string{"丙子", "丁丑", "戊寅", "己卯", "庚辰", "辛巳", "壬午", "癸未", "甲申", "乙酉", "丙戌", "丁亥"},
		},

		{
			st:  time.Date(2020, 7, 2, 0, 0, 0, 0, time.Local),
			sy:  2020,
			dgz: "丙午",
			hgz: []string{"戊子", "己丑", "庚寅", "辛卯", "壬辰", "癸巳", "甲午", "乙未", "丙申", "丁酉", "戊戌", "己亥"},
		},

		{
			st:  time.Date(2020, 7, 3, 0, 0, 0, 0, time.Local),
			sy:  2020,
			dgz: "丁未",
			hgz: []string{"庚子", "辛丑", "壬寅", "癸卯", "甲辰", "乙巳", "丙午", "丁未", "戊申", "己酉", "庚戌", "辛亥"},
		},

		{
			st:  time.Date(2020, 7, 4, 0, 0, 0, 0, time.Local),
			sy:  2020,
			dgz: "戊申",
			hgz: []string{"壬子", "癸丑", "甲寅", "乙卯", "丙辰", "丁巳", "戊午", "己未", "庚申", "辛酉", "壬戌", "癸亥"},
		},
	}

	want := []int{
		//阳 下元
		4, 5, 6, 7, 8, 9, 1, 2, 3, 4, 5, 6,
		7, 8, 9, 1, 2, 3, 4, 5, 6, 7, 8, 9,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 1, 2, 3,
		4, 5, 6, 7, 8, 9, 1, 2, 3, 4, 5, 6,
		7, 8, 9, 1, 2, 3, 4, 5, 6, 7, 8, 9,
		//阳 上元
		1, 2, 3, 4, 5, 6, 7, 8, 9, 1, 2, 3,
		4, 5, 6, 7, 8, 9, 1, 2, 3, 4, 5, 6,
		7, 8, 9, 1, 2, 3, 4, 5, 6, 7, 8, 9,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 1, 2, 3,
		4, 5, 6, 7, 8, 9, 1, 2, 3, 4, 5, 6,
		//阳 中元
		7, 8, 9, 1, 2, 3, 4, 5, 6, 7, 8, 9,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 1, 2, 3,
		4, 5, 6, 7, 8, 9, 1, 2, 3, 4, 5, 6,
		7, 8, 9, 1, 2, 3, 4, 5, 6, 7, 8, 9,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 1, 2, 3,
		//阴 上元
		3, 2, 1, 9, 8, 7, 6, 5, 4, 3, 2, 1,
		9, 8, 7, 6, 5, 4, 3, 2, 1, 9, 8, 7,
		6, 5, 4, 3, 2, 1, 9, 8, 7, 6, 5, 4,
		//阴 中元
		3, 2, 1, 9, 8, 7, 6, 5, 4, 3, 2, 1,
		9, 8, 7, 6, 5, 4, 3, 2, 1, 9, 8, 7,
		6, 5, 4, 3, 2, 1, 9, 8, 7, 6, 5, 4,
		3, 2, 1, 9, 8, 7, 6, 5, 4, 3, 2, 1,
		9, 8, 7, 6, 5, 4, 3, 2, 1, 9, 8, 7,
		//阴 下元
		6, 5, 4, 3, 2, 1, 9, 8, 7, 6, 5, 4,
		3, 2, 1, 9, 8, 7, 6, 5, 4, 3, 2, 1,
		9, 8, 7, 6, 5, 4, 3, 2, 1, 9, 8, 7,
		6, 5, 4, 3, 2, 1, 9, 8, 7, 6, 5, 4,
		3, 2, 1, 9, 8, 7, 6, 5, 4, 3, 2, 1,
	}

	for i := 0; i < len(gzt); i++ {
		for j := 0; j < len(gzt[i].hgz); j++ {
			hx := ZiBaiH(gzt[i].dgz, gzt[i].hgz[j], gzt[i].st, gzt[i].sy)
			index := i*11 + i + j
			if i > 0 {
				//fmt.Printf("%d-", ZiBaiH(gzt[i].dgz, gzt[i].hgz[j], gzt[i].st, gzt[i].sy))
				if want[index] != hx {
					t.Errorf("func ZiBaiD(%s,%s,%v,%d)=%d want:%d\n", gzt[i].dgz, gzt[i].hgz[j], gzt[i].st, gzt[i].sy, hx, want[index])
				}
			} else if i == 0 {
				//fmt.Printf("0-%d-", ZiBaiH(gzt[i].dgz, gzt[i].hgz[j], gzt[i].st, gzt[i].sy))
				if want[j] != hx {
					t.Errorf("func ZiBaiD(%s,%s,%v,%d)=%d want:%d\n", gzt[i].dgz, gzt[i].hgz[j], gzt[i].st, gzt[i].sy, hx, want[j])
				}
			}

		}

	}
}

//紫白飞宫
func TestZiBaiG(t *testing.T) {
	zbobj := struct {
		zbn, yy int
	}{zbn: 3, yy: 1}
	zbGmap := ZiBaiGmap(zbobj.zbn, zbobj.yy)
	fmt.Printf("%v\n", zbGmap) //map[1:3 2:4 3:5 4:6 5:7 6:8 7:9 8:1 9:2]

}

//紫白五行属性
func TestZiBaiSelf(t *testing.T) {
	zbn := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	wxs := []string{"水", "土", "木", "木", "土", "金", "金", "土", "火"}
	for i := 0; i < len(zbn); i++ {
		zbwx := ZiBaiSelf(zbn[i])
		if !strings.EqualFold(zbwx, wxs[i]) {
			t.Errorf("func ZiBaiSelf(%d)=%s want:%s\n", zbn[i], zbwx, wxs[i])
		}
	}

}
