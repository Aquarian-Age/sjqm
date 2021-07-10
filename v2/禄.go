package sjqm

import "strings"

//禄
/*
   甲禄在寅　乙禄在卯　甲辰旬寅卯空　甲辰乙巳为无禄
   庚禄在申　辛禄在酉　甲戌旬申酉空　庚辰辛巳为无禄
   丙戊禄在巳　　　　　甲午旬巳空　　丙申戊戌为无禄
   丁己禄在午        甲申旬午空　　丁亥己丑为无禄
   壬禄在亥　　　　　　甲子旬亥空　　壬申为无禄
   癸禄在子          甲寅旬子空　　癸亥为无禄
*/
func Lu(gz string) (lus string) {

	gan := []string{"甲", "乙", "丙", "丁", "戊", "己", "庚", "辛", "壬", "癸"}
	zhi := []string{"子", "丑", "寅", "卯", "辰", "巳", "午", "未", "申", "酉", "戌", "亥"}

	var i int
	for i = 0; i < len(gan); i++ {
		if strings.ContainsAny(gz, gan[i]) {
			break
		}
	}
	switch i {
	case 0: //甲
		//fmt.Printf("%s禄在%s\n", gan[0], zhi[2]) //cli显示
		lus = zhi[2]
	case 1: //乙
		//fmt.Printf("%s禄在%s\n", gan[1], zhi[3])
		lus = zhi[3]
	case 2: //丙
		//fmt.Printf("%s禄在%s\n", gan[2], zhi[5])
		lus = zhi[5]
	case 4: //戊
		//fmt.Printf("%s禄在%s\n", gan[4], zhi[5])
		lus = zhi[5]
	case 3: //丁
		//fmt.Printf("%s禄在%s\n", gan[3], zhi[6])
		lus = zhi[6]
	case 5: //己
		//fmt.Printf("%s禄在%s\n", gan[5], zhi[6])
		lus = zhi[6]
	case 6: //庚
		//fmt.Printf("%s禄在%s\n", gan[6], zhi[8])
		lus = zhi[8]
	case 7: //辛
		//fmt.Printf("%s禄在%s\n", gan[7], zhi[9])
		lus = zhi[9]
	case 8: //壬
		//fmt.Printf("%s禄在%s\n", gan[8], zhi[11])
		lus = zhi[11]
	case 9: //癸
		//fmt.Printf("%s禄在%s\n", gan[9], zhi[0])
		lus = zhi[0]
	}
	return
}
