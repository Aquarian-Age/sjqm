package sjqm

import "strings"

//时辰对应的旬首
func XunShou(hgz string) (xunShouGZ string, indexXS int) {
	//当前时辰往前推到甲
	var i int
	for i = 0; i < len(Jz60); i++ {
		if strings.EqualFold(hgz, Jz60[i]) {
			for j := i; j >= i-10; j-- {
				if strings.ContainsAny(Jz60[j], "甲") {
					xunShouGZ = Jz60[j] ///时辰旬首
					indexXS = j
					break
				}
			}
			break
		}
	}

	return
}

//旬首的干支数组
func 旬首干支数组(indexXS int) []string {
	var xsGZArr []string
	for i := indexXS; i < indexXS+10; i++ {
		gz := Jz60[i]
		xsGZArr = append(xsGZArr, gz)
	}

	return xsGZArr
}
