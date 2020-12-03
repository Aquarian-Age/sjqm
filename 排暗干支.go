package sjqm

import (
	"strings"
)

//值符原始宫位
func FindZFNumber(zf string) (zfn int) {
	for _, g := range JGMap() {
		if strings.EqualFold(g.Star, zf) {
			//fmt.Printf("-->值符:%s原始宫位:%d\n", zf, g.Number)
			zfn = g.Number
			break
		}
	}
	return
}

//旬首干支数组
func FindXSGZArr(xunshou string) []string {
	var gzarr []string
	for i := 0; i < len(jz60); i++ {
		if strings.EqualFold(xunshou, jz60[i]) {
			gzarr = jz60[i : i+10]
			break
		}
	}
	//fmt.Printf("-->旬首干支数组gzarr:%s\n", gzarr)
	return gzarr
}

//值符原始宫位 起 时辰旬首 阳顺阴逆
func AnGZ(zfn int, gzarr []string, yinyang int) (agz map[int]string) {
	//yinyang:0->阴 1->阳

	var aGZmap = make(map[int]string) //暗干支
	switch yinyang {
	case 0: //阴 逆
		x := []int{9, 8, 7, 6, 5, 4, 3, 2, 1} //阴遁顺序
		for j := 0; j < len(x); j++ {
			if x[j] == zfn {
				x1 := x[:j]
				x2 := x[j:]
				x = append(x2, x1...)
			}
		}
		///干支配值符原始宫
		for gi := 0; gi < len(gzarr); gi++ {
			for xi := gi; xi < len(x); xi++ {
				aGZmap[x[xi]] = gzarr[gi]
				break
			}
		}
		aGZmap[x[0]] = gzarr[0] + " " + gzarr[9] //最后一个暗干支加到起始宫位上(值符原始宫位)
		agz = aGZmap
	case 1: //阳 顺
		x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9} //阳遁顺序
		for j := 0; j < len(x); j++ {
			if x[j] == zfn {
				x1 := x[:j]
				x2 := x[j:]
				x = append(x2, x1...)
			}
		}
		///干支配值符原始宫
		for gi := 0; gi < len(gzarr); gi++ {
			for xi := gi; xi < len(x); xi++ {
				aGZmap[x[xi]] = gzarr[gi]
				break
			}
		}
		aGZmap[x[0]] = gzarr[0] + " " + gzarr[9]
		agz = aGZmap
	}
	return
}
