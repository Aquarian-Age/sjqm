package sjqm

import "strings"

//阴阳贵人诀 返回干支对应的阳贵人　阴贵人
func GuiRenJue(dgz string) (yang, yin string) {
	yiJi := []string{"申", "子"}   //乙申阳　乙子阴　己申阴　己子阳
	jianXu := []string{"未", "丑"} //甲未阳　甲丑阴　戊未阴　戊丑阳
	bingDing := []string{"酉", "亥"}
	renGui := []string{"卯", "巳"}
	geng := []string{"丑", "未"}
	xin := []string{"寅", "午"}
	gan := []string{"甲", "乙", "丙", "丁", "戊", "己", "geng", "xin", "壬", "癸"}

	var i int
	for i = 0; i < len(gan); i++ {
		if strings.ContainsAny(dgz, gan[i]) {
			break
		}
	}
	switch i {
	case 0: //甲
		yang = jianXu[0]
		yin = jianXu[1]
	case 4: //戊
		yang = jianXu[1]
		yin = jianXu[0]
	case 1: //乙
		yang = yiJi[0]
		yin = yiJi[1]
	case 5: //己
		yang = yiJi[1]
		yin = yiJi[0]
	case 2: //丙
		yang = bingDing[0]
		yin = bingDing[1]
	case 3: //丁
		yang = bingDing[1]
		yin = bingDing[0]
	case 6: //geng
		yang = geng[0]
		yin = geng[1]
	case 7: //xin
		yang = xin[0]
		yin = xin[1]
	case 8: //壬
		yang = renGui[0]
		yin = renGui[1]
	case 9: //癸
		yang = renGui[1]
		yin = renGui[0]
	}
	return
}
