package sjqm

import (
	"fmt"
	"testing"
)

//天三门
func TestTianSanMen(t *testing.T) {
	yj := "亥"   //"巳"
	hgz := "丙午" //"戊辰"
	want := map[string]string{"从魁": "辰", "太冲": "戌", "小吉": "寅"}
	smmap := TianSanMen(yj, hgz)

	fmt.Println(smmap)
	fmt.Println(want)
	//if !reflect.DeepEqual(smmap, want) {
	//	t.Errorf("func TianSanMen(%s,%s)=%v 期望值:%v\n", yj, hgz, smmap, want)
	//}
}
