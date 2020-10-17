package sjqm

import (
	"strings"

	"github.com/nongli/ganzhi"
)

/*
地支—天干=旬首常数（甲子0；甲寅2；甲辰4；甲午6；甲申8；甲戌10）
例如求辛巳的旬首，用地支（巳）的序号6减去天干（辛）的序号8（地支小于天干时加12），得10。
所以得旬首甲戌
*/
//时干支对应的旬首
func XS旬首(gz string) (xs string) {
	//旬首的数字和对应的值不能变
	var 旬首 = map[int]string{0: "甲子", 2: "甲寅", 4: "甲辰", 6: "甲午", 8: "甲申", 10: "甲戌"}
	//找出地支天干对应的数字
	var i, j int
	//天干
	for i = 0; i <= 10; i++ {
		if gb := strings.Contains(gz, ganzhi.Gan[i]); gb == true {
			//fmt.Printf("i=%d gan:%s\n", i, ganzhi.Gan[i])
			break
		}
	}
	//地支
	for j = 0; j <= 12; j++ {
		if jb := strings.Contains(gz, ganzhi.Zhi[j]); jb == true {
			//fmt.Printf("j=%d zhi:%s\n", j, ganzhi.Zhi[j])
			break
		}
	}

	n := cal(j, i)
	for k, v := range 旬首 {
		if k == n {
			xs = v //旬首
			//fmt.Printf("旬首:%s\n", xs)
			break
		}
	}
	return
}

//地支减去天干
func cal(j, i int) (res int) {
	if j < i {
		j += 12
		res = j - i
	} else {
		res = j - i
	}
	return
}
