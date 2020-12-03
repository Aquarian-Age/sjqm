package sjqm

import (
	"strings"
)

//时辰旬首原始宫位的值使
func FindZS(xsN int) (zsName string) {
	for n, g := range JGMap() {
		if xsN == n {
			zsName = g.Door
			break
		}
	}
	return
}

//值使(八门)随时辰地支 从时辰地支宫顺时针排
func ZhiShi(hgz string, anGZ map[int]string, zhis string) (zhiShimap map[int]string) {
	//fmt.Printf("-->%s时辰: 值使:%s\n", hgz, zhis)
	x := []int{1, 8, 3, 4, 9, 2, 7, 6, 5} //5最后加
	bm := []string{"休门", "生门", "伤门", "杜门", "景门", "死门", "惊门", "开门"}

	var zsmap = make(map[int]string)
	for n, name := range anGZ {
		if strings.ContainsAny(hgz, name) {
			//fmt.Printf("-->%s时辰 暗干支落%d宫位\n", hgz, n)
			//宫位排序
			for i := 0; i < len(x); i++ {
				if n == x[i] && n == 5 {
					x1 := x[:i]
					x2 := x[i:]
					x = append(x2, x1...)
					//fmt.Printf("--九宫排序:%d\n", x)
					//删除5重新再排序
					x = x[1:]
					for j := 0; j < len(x); j++ {
						if x[j] == 2 { //5宫寄到坤2宫
							xx1 := x[:j]
							xx2 := x[j:]
							x = append(xx2, xx1...)
							//fmt.Printf("--九宫再排序:%d\n", x)
							break
						}
					}
					break
				}
				if n == x[i] && n != 5 {
					//找到5删除
					x = x[:8]
					x1 := x[:i]
					x2 := x[i:]
					x = append(x2, x1...)
					//fmt.Printf("--!5九宫删除5再排序:%d\n", x)
					//x = append(x, 5) //中宫排最后
					break
				}
			}
			break
		}
	}
	//fmt.Printf("--九宫排序:%d\n", x)

	//八门排序
	for j := 0; j < len(bm); j++ {
		if strings.EqualFold(bm[j], zhis) {
			bm1 := bm[:j]
			bm2 := bm[j:]
			bm = append(bm2, bm1...)
			break
		}
	}
	//fmt.Printf("-->八门配九宫:%s\n", bm)
	//八门配宫
	for ix := 0; ix < len(x); ix++ {
		for mx := ix; mx < len(bm); mx++ {
			zsmap[x[ix]] = bm[mx]
			break
		}
	}
	zhiShimap = zsmap
	return
}
