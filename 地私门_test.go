package sjqm

import (
	"fmt"
	"testing"
)

//地私门测试
func TestDiSiMen(t *testing.T) {
	yj := "亥"
	dgz := "甲子"
	hgz := "戊辰"
	want := map[string]string{"六合": "卯", "太常": "申", "太阴": "戌"}
	dsmap := DiSiMen(yj, dgz, hgz)
	fmt.Println(dsmap, want)

	yj2 := "巳"
	dgz2 := "乙丑"
	hgz2 := "辛巳"
	want2 := map[string]string{"六合": "巳", "太常": "子", "太阴": "戌"}
	dsmap2 := DiSiMen(yj2, dgz2, hgz2)
	fmt.Println(dsmap2, want2)

	yj1 := "亥"
	dgz1 := "甲子"
	hgz1 := "丁卯"
	want1 := map[string]string{"六合": "寅", "太常": "酉", "太阴": "未"}
	dsmap1 := DiSiMen(yj1, dgz1, hgz1)
	fmt.Println(dsmap1, want1)

	yj3 := "午"
	dgz3 := "辛酉"
	hgz3 := "甲午"
	want3 := map[string]string{"六合": "巳", "太常": "戌", "太阴": "子"}
	dsmap3 := DiSiMen(yj3, dgz3, hgz3)
	fmt.Println(dsmap3, want3)

}
