package sjqm

import (
	"fmt"
	"testing"
)

//地四户
func TestDiSiHu(t *testing.T) {
	//mgz := "丙戌"
	hgz := "辛巳"
	want := map[string]string{"危": "子", "定": "酉", "开": "卯", "除": "午"}
	sihumap := DiSiHu(hgz)
	fmt.Println(sihumap, want)

	//if !reflect.DeepEqual(sihumap, want) {
	//	t.Errorf("func DiSiHu(%s)=%v 期望值:%v\n", hgz, sihumap, want)
	//}

	//mgz1 := "甲寅"
	hgz1 := "庚寅"
	want1 := map[string]string{"危": "酉", "定": "午", "开": "子", "除": "卯"}
	sihumap1 := DiSiHu(hgz1)
	fmt.Println(sihumap1, want1)
	//	if !reflect.DeepEqual(sihumap1, want1) {
	//		t.Errorf("func DiSiHu(%s)=%v 期望值:%v\n", hgz1, sihumap, want1)
	//	}
}
