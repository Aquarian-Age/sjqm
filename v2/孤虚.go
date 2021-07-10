package sjqm

import (
	"fmt"
	"strings"
)

//时孤虚
func guXuH(dgz, hgz string) (gxs string) {
	guxu := HGuXu(dgz, hgz)
	if _, ok := guxu["孤"]; ok {
		gxs = fmt.Sprintf("孤:%s 虚:%s", guxu["孤"], guxu["虚"])
	}
	return
}

//孤虚法
func HGuXu(dgz, hgz string) (hgu map[string]string) {
	hgu = make(map[string]string)
	for i := 0; i < len(Jz60); i++ {
		if strings.EqualFold(dgz, Jz60[i]) { //尋找日干支在六十甲子中的索引號
			switch i {
			case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9: //甲子旬　孤:戌亥　虚:辰巳"
				if strings.ContainsAny(hgz, "戌") {
					hgu = map[string]string{"孤": "戌", "虚": "辰"}
				} else if strings.ContainsAny(hgz, "亥") {
					hgu = map[string]string{"孤": "亥", "虚": "巳"}
				}
			case 10, 11, 12, 13, 14, 15, 16, 17, 18, 19: //甲戌旬　孤:申酉　虚:辰巳
				if strings.ContainsAny(hgz, "申") {
					hgu = map[string]string{"孤": "申", "虚": "寅"}
				} else if strings.ContainsAny(hgz, "酉") {
					hgu = map[string]string{"孤": "酉", "虚": "卯"}
				}
			case 20, 21, 22, 23, 24, 25, 26, 27, 28, 29: //甲申旬　孤:午未　虚:辰巳
				//当前时辰正好为为孤虚时辰
				if strings.ContainsAny(hgz, "午") {
					hgu = map[string]string{"孤": "午", "虚": "子"}
				} else if strings.ContainsAny(hgz, "未") {
					hgu = map[string]string{"孤": "未", "虚": "丑"}
				}
			case 30, 31, 32, 33, 34, 35, 36, 37, 38, 39: //甲午旬　孤:辰巳　虚:辰巳
				//当前时辰正好为为孤虚时辰
				if strings.ContainsAny(hgz, "辰") {
					hgu = map[string]string{"孤": "辰", "虚": "戌"}
				} else if strings.ContainsAny(hgz, "巳") {
					hgu = map[string]string{"孤": "巳", "虚": "亥"}
				}
			case 40, 41, 42, 43, 44, 45, 46, 47, 48, 49: //甲辰旬　孤:寅卯　虚:辰巳
				//当前时辰正好为为孤虚时辰
				if strings.ContainsAny(hgz, "寅") {
					hgu = map[string]string{"孤": "寅", "虚": "申"}
				} else if strings.ContainsAny(hgz, "卯") {
					hgu = map[string]string{"孤": "卯", "虚": "酉"}
				}
			case 50, 51, 52, 53, 54, 55, 56, 57, 58, 59: //甲寅旬　孤:子丑　虚:辰巳
				//当前时辰正好为为孤虚时辰
				if strings.ContainsAny(hgz, "子") {
					hgu = map[string]string{"孤": "子", "虚": "午"}
				} else if strings.ContainsAny(hgz, "丑") {
					hgu = map[string]string{"孤": "丑", "虚": "未"}
				}
			}
		}
	}
	return
}
