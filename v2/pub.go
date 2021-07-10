package sjqm

import "strings"

//转换月干支为数字形式
//寅月1 卯月2 辰月3 ....子月11 丑月12
func ConvMGZToNumber(mgz string) (n int) {
	gzmap := map[string]int{
		"寅": 1, "卯": 2, "辰": 3, "巳": 4, "午": 5, "未": 6, "申": 7, "酉": 8, "戌": 9, "亥": 10, "子": 11, "丑": 12,
	}
	for k, v := range gzmap {
		if strings.ContainsAny(mgz, k) {
			n = v
			break
		}
	}
	return
}

//转换时辰干支为数字形式
func ConvHGZToNumber(hgz string) (hn int) {
	gzmap := map[string]int{
		"子": 1, "丑": 2, "寅": 3, "卯": 4, "辰": 5, "巳": 6, "午": 7, "未": 8, "申": 9, "酉": 10, "戌": 11, "亥": 12,
	}
	for k, v := range gzmap {
		if strings.ContainsAny(hgz, k) {
			hn = v
			break
		}
	}
	return
}

var (
	Gan = [11]string{"err", "甲", "乙", "丙", "丁", "戊", "己", "庚", "辛", "壬", "癸"}
	Zhi = [13]string{"err", "子", "丑", "寅", "卯", "辰", "巳", "午", "未", "申", "酉", "戌", "亥"}
)
