package sjqm

import "strings"

//yn:三元数字 上0 中1 下2 index:节气名称的索引数字 jqnames:节气数组 0:冬至 1小寒 2大寒...
//定局
func FindJU(yn int, index int, jqnames []string) (jn int) {
	for i := 0; i < len(jqnames); i++ {
		if index == i {
			for name, intarr := range JieNumber() {
				if strings.EqualFold(name, jqnames[i]) {
					jn = intarr[yn]
				}
			}
			break
		}
	}
	return
}

func JieNumber() map[string][]int {
	var jqmap = make(map[string][]int)
	//阳遁
	jqmap["冬至"] = []int{1, 7, 4}
	jqmap["小寒"] = []int{2, 8, 5}
	jqmap["大寒"] = []int{3, 9, 6}
	jqmap["立春"] = []int{8, 5, 2}
	jqmap["雨水"] = []int{9, 6, 3}
	jqmap["惊蛰"] = []int{1, 7, 4}
	jqmap["春分"] = []int{3, 9, 6}
	jqmap["清明"] = []int{4, 1, 7}
	jqmap["谷雨"] = []int{5, 2, 8}
	jqmap["立夏"] = []int{4, 1, 7}
	jqmap["小满"] = []int{5, 2, 8}
	jqmap["芒种"] = []int{6, 3, 9}
	/////阴遁
	jqmap["夏至"] = []int{9, 3, 6}
	jqmap["小暑"] = []int{8, 2, 5}
	jqmap["大暑"] = []int{7, 1, 4}
	jqmap["立秋"] = []int{2, 5, 8}
	jqmap["处暑"] = []int{1, 4, 7}
	jqmap["白露"] = []int{9, 3, 6}
	jqmap["秋分"] = []int{7, 1, 4}
	jqmap["寒露"] = []int{6, 9, 3}
	jqmap["霜降"] = []int{5, 8, 2}
	jqmap["立冬"] = []int{6, 9, 3}
	jqmap["小雪"] = []int{5, 8, 2}
	jqmap["大雪"] = []int{4, 7, 1}

	return jqmap
}
