package sjqm

import (
	"fmt"
	"strings"
)

//几有紧急事，从天马而出，可避祸
//太冲天马 先用月将加于本时，顺数至卯，即太冲
func tianMa(yj, hgz string) (tms string) {
	zhi := []string{"亥", "子", "丑", "寅", "卯", "辰", "巳", "午", "未", "申", "酉", "戌"}
	var szhi, ezhi []string
	for i := 0; i < len(zhi); i++ {
		if strings.ContainsAny(hgz, zhi[i]) {
			e1 := zhi[i:]
			e2 := zhi[:i]
			ezhi = append(e1, e2...)
			break
		}
	}

	for j := 0; j < len(zhi); j++ {
		if strings.ContainsAny(zhi[j], yj) {
			s1 := zhi[j:]
			s2 := zhi[:j]
			szhi = append(s1, s2...)
			break
		}
	}
	//天月将落宫应都是卯 地月将即是太冲天马的方位
	for s := 0; s < len(szhi); s++ {
		for e := 0; e < len(ezhi); e++ {
			if s == e {
				if strings.EqualFold(szhi[e], "卯") {
					tms = ezhi[e]
				}
				break
			}
		}
	}
	tms = fmt.Sprintf("天马:%s", tms)
	return
}
