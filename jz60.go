package sjqm

//############################################
//六十甲子
//############################################
func GetJZArr() []string {
	jiazi, jiaxu, jiashen, jiawu, jiachen, jiayin := liuXun()
	return jzArr(jiazi, jiaxu, jiashen, jiawu, jiachen, jiayin)
}

/*
甲戌旬:  [甲戌 乙亥 丙子 丁丑 戊寅 己卯 庚辰 辛巳 壬午 癸未]
甲申旬:  [甲申 乙酉 丙戌 丁亥 戊子 己丑 庚寅 辛卯 壬辰 癸巳]
甲午旬:  [甲午 乙未 丙申 丁酉 戊戌 己亥 庚子 辛丑 壬寅 癸卯]
甲辰旬:  [甲辰 乙巳 丙午 丁未 戊申 己酉 庚戌 辛亥 壬子 癸丑]
甲寅旬:  [甲寅 乙卯 丙辰 丁巳 戊午 己未 庚申 辛酉 壬戌 癸亥]
甲子旬:  [甲子 乙丑 丙寅 丁卯 戊辰 己巳 庚午 辛未 壬申 癸酉]
*/
//生成六旬干支数组 甲子 甲戌 甲申 甲午 甲辰 甲寅
func liuXun() ([]string, []string, []string, []string, []string, []string) {
	gan := []string{"甲", "乙", "丙", "丁", "戊", "己", "庚", "辛", "壬", "癸"}
	zhi := []string{"子", "丑", "寅", "卯", "辰", "巳", "午", "未", "申", "酉", "戌", "亥"}

	var jiaxu, jiashen, jiawu, jiachen, jiayin, jiazi []string
	var newZhi []string
	for x := 0; x < 6; x++ {
		for i := 0; i <= len(gan); i++ {
			for j := 0; j < len(zhi); j++ {
				switch x {
				case 0:
					if i == j && i == 10 {
						head := zhi[i:]
						end := zhi[:i]
						newZhi = append(head, end...)
						jiaxu = makeArr(gan, newZhi)
					}
				case 1:
					if i == j && i == 10 {
						head := newZhi[i:]
						end := newZhi[:i]
						newZhi = append(head, end...)
						jiashen = makeArr(gan, newZhi)
					}
				case 2:
					if i == j && i == 10 {
						head := newZhi[i:]
						end := newZhi[:i]
						newZhi = append(head, end...)
						jiawu = makeArr(gan, newZhi)
					}
				case 3:
					if i == j && i == 10 {
						head := newZhi[i:]
						end := newZhi[:i]
						newZhi = append(head, end...)
						jiachen = makeArr(gan, newZhi)
					}
				case 4:
					if i == j && i == 10 {
						head := newZhi[i:]
						end := newZhi[:i]
						newZhi = append(head, end...)
						jiayin = makeArr(gan, newZhi)
					}
				case 5:
					if i == j && i == 10 {
						head := newZhi[i:]
						end := newZhi[:i]
						newZhi = append(head, end...)
						jiazi = makeArr(gan, newZhi)
					}
				}
			}
		}
	}
	return jiazi, jiaxu, jiashen, jiawu, jiachen, jiayin
}

//按顺序合并六旬甲子数组
func jzArr(arr ...[]string) []string {
	var arrx []string
	for i := 0; i < len(arr); i++ {
		arrx = append(arrx, arr[i]...)
	}
	return arrx
}

//生成干支数组
func makeArr(gan, newZhi []string) []string {
	var arr []string
	for i := 0; i < len(gan); i++ {
		for j := 0; j < len(newZhi); j++ {
			if i == j {
				gz := gan[j] + newZhi[j]
				arr = append(arr, gz)
				break
			}
		}
	}
	return arr
}
