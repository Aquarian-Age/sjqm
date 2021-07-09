package sjqm

import (
	"sort"
	"strings"
)

//六甲旬遁: 甲子遁戊　甲戌遁己　甲申遁庚　甲午遁辛　甲辰遁壬　甲寅遁癸

//值符动态宫位数字 即x时辰值符所在N宫位
//时辰旬首对应的九宫位
func XunShouHour(xunshou, hgz string, xsGZArr []string, sqly map[int]string) (xunShouNumber int, gzNanme, dun string) {
	//xunshou:旬首 hgz:时辰干支 xsGZArr:旬首为首的十天干数组 sqly:地盘三奇六仪
	switch xunshou {
	case "甲子": //甲子遁戊 (找到地盘三奇六仪的戊在那个宫位数字即可)
		for i := 0; i < len(xsGZArr); i++ {
			if strings.EqualFold(hgz, xsGZArr[i]) { //每个时辰对应的值符
				for k, v := range sqly {
					if strings.EqualFold(hgz, "甲子") && strings.ContainsAny(v, "戊") {
						//fmt.Printf("-->值符宫位(时辰宫位):%d 时辰干支:%s\n", k, v)
						xunShouNumber = k //时辰旬首落九宫位置的数字(不是值符原始宫位数字)
						gzNanme = "甲子"    //时辰干支名称
						dun = "戊"         //六甲旬遁
					}
					if strings.ContainsAny(hgz, v) && !strings.ContainsAny(hgz, "甲") {
						//fmt.Printf("-->值符宫位(时辰宫位):%d 时辰干支:%s\n", k, v) //时辰宫位数字即是值符所在位置也是排九星的起始位置
						xunShouNumber = k
						gzNanme = hgz
						dun = "戊"
						break
					}
				}
				break
			}
		}
	case "甲戌": //甲戌遁己
		for i := 0; i < len(xsGZArr); i++ {
			if strings.EqualFold(hgz, xsGZArr[i]) {
				////fmt.Printf("-->时辰: %s 值符:%s\n", hgz, zhifu)

				for k, v := range sqly {
					if strings.EqualFold(hgz, "甲戌") && strings.ContainsAny(v, "己") {
						//fmt.Printf("-->值符宫位(时辰宫位):%d 时辰干支:%s\n", k, v)
						xunShouNumber = k
						gzNanme = "甲戌"
						dun = "己"
					}
					if strings.ContainsAny(hgz, v) && !strings.ContainsAny(hgz, "甲") {
						//fmt.Printf("-->值符宫位(时辰宫位):%d 时辰干支:%s\n", k, v) //时辰宫位数字即是值符所在位置也是排九星的起始位置
						xunShouNumber = k
						gzNanme = hgz
						dun = "己"
						break
					}
				}
				break
			}
		}
	case "甲申": //甲申遁庚
		for i := 0; i < len(xsGZArr); i++ {
			if strings.EqualFold(hgz, xsGZArr[i]) {
				//fmt.Printf("-->时辰: %s 值符:%s\n", hgz, zhifu)

				for k, v := range sqly {
					if strings.EqualFold(hgz, "甲申") && strings.ContainsAny(v, "庚") {
						//fmt.Printf("-->值符宫位(时辰宫位):%d 时辰干支:%s\n", k, v)
						xunShouNumber = k
						gzNanme = "甲申"
						dun = "庚"
					}
					if strings.ContainsAny(hgz, v) && !strings.ContainsAny(hgz, "甲") {
						//fmt.Printf("-->值符宫位(时辰宫位):%d 时辰干支:%s\n", k, v) //时辰宫位数字即是值符所在位置也是排九星的起始位置
						xunShouNumber = k
						gzNanme = hgz
						dun = "庚"
						break
					}
				}
				break
			}
		}
	case "甲午": //甲午遁辛
		for i := 0; i < len(xsGZArr); i++ {
			if strings.EqualFold(hgz, xsGZArr[i]) {
				//fmt.Printf("-->时辰: %s 值符:%s\n", hgz, zhifu)

				for k, v := range sqly {
					if strings.EqualFold(hgz, "甲午") && strings.ContainsAny(v, "辛") {
						//fmt.Printf("-->值符宫位(时辰宫位):%d 时辰干支:%s\n", k, v)
						xunShouNumber = k
						gzNanme = "甲午"
						dun = "辛"
					}
					if strings.ContainsAny(hgz, v) && !strings.ContainsAny(hgz, "甲") {
						//fmt.Printf("-->值符宫位(时辰宫位):%d 时辰干支:%s\n", k, v) //时辰宫位数字即是值符所在位置也是排九星的起始位置
						xunShouNumber = k
						gzNanme = hgz
						dun = "辛"
						break
					}
				}
				break
			}
		}
	case "甲辰": //甲辰遁壬
		for i := 0; i < len(xsGZArr); i++ {
			if strings.EqualFold(hgz, xsGZArr[i]) {
				//fmt.Printf("-->时辰: %s 值符:%s\n", hgz, zhifu)

				for k, v := range sqly {
					if strings.EqualFold(hgz, "甲辰") && strings.ContainsAny(v, "壬") { //甲辰遁壬
						//fmt.Printf("-->值符宫位(时辰宫位):%d 时辰干支:%s\n", k, v)
						xunShouNumber = k
						gzNanme = "甲辰"
						dun = "壬"
					}
					if strings.ContainsAny(hgz, v) && !strings.ContainsAny(hgz, "甲") {
						//fmt.Printf("-->值符宫位(时辰宫位):%d 时辰干支:%s\n", k, v) //时辰宫位数字即是值符所在位置也是排九星的起始位置
						xunShouNumber = k
						gzNanme = hgz
						dun = "壬"
						break
					}
				}
				break
			}
		}
	case "甲寅": //甲寅遁癸
		for i := 0; i < len(xsGZArr); i++ {
			if strings.EqualFold(hgz, xsGZArr[i]) {
				//fmt.Printf("-->时辰: %s 值符:%s\n", hgz, zhifu)

				for k, v := range sqly {
					if strings.EqualFold(hgz, "甲寅") && strings.ContainsAny(v, "癸") {
						//fmt.Printf("-->值符宫位(时辰宫位):%d 时辰干支:%s\n", k, v)
						xunShouNumber = k
						gzNanme = "甲寅"
						dun = "癸"
					}
					if strings.ContainsAny(hgz, v) && !strings.ContainsAny(hgz, "甲") {
						//fmt.Printf("-->值符宫位(时辰宫位):%d 时辰干支:%s\n", k, v) //时辰宫位数字即是值符所在位置也是排九星的起始位置
						xunShouNumber = k
						gzNanme = hgz
						dun = "癸"
						break
					}
				}
				break
			}
		}
	}
	return
}

//根据时辰所在旬遁找出原宫位的值符值使
func SelfZF(dun string, sqly map[int]string) (zf string, dunN int) {
	//确定六甲旬遁落那个宫位
	for n, name := range sqly {
		if strings.EqualFold(dun, name) {
			//根据三奇六仪旬遁宫位数找本宫位原始值符值使
			for number, g := range JGMap() {
				if n == number {
					zf = g.Star //原始宫位值符
					dunN = n    //旬遁落宫数字
					break
				}
			}
			break
		}
	}
	return
}

//九星排序
func SortStar(zf string) []string {
	//从当前值符对九星重新排序
	jiux := []string{"天蓬", "天任", "天冲", "天辅", "天英", "天芮", "天柱", "天心"}
	//如果值符是天禽 天禽寄2宫
	for index := 0; index < len(jiux); index++ {
		if strings.EqualFold(zf, "天禽") {
			if strings.EqualFold("天芮", jiux[index]) {
				s1 := jiux[:index]
				s2 := jiux[index:]
				jiux = append(s2, s1...)
			}
		}
		if strings.EqualFold(zf, jiux[index]) && !strings.EqualFold(zf, "天禽") {
			s1 := jiux[:index]
			s2 := jiux[index:]
			jiux = append(s2, s1...)
		}
	}
	starArr := append(jiux, "天禽") //放到最后 寄坤二宫

	return starArr
}

//九星: 天蓬 天任 天冲 天辅 天英 天芮 天柱 天心 天禽(中宫)寄坤二宫即天芮星
//排值符九星
func ZhiFuStar(xunShouNumber int, starArr []string) map[int]string {
	//根据时辰旬首 排值符九星
	xArr := hGong(xunShouNumber) //值符随时干 如果xunShouNumber==5即值符落中宫 在delE函数实现让它寄坤二宫
	xArr = del(xArr)             //中宫天禽配九星的天禽顺序
	//九星配时辰(九宫)
	l1 := len(starArr)
	l2 := len(xArr)
	var starmap = make(map[int]string)
	if l1 == l2 {
		for i := 0; i < len(xArr); i++ {
			for j := 0; j < len(starArr); j++ {
				if i == j {
					starmap[xArr[i]] = starArr[j]
					break
				}
			}
		}
	}
	return starmap
}

//剥离5
func del(x []int) []int {
	for i := 0; i < len(x); i++ {
		if x[i] == 5 && i == 0 { //5在首位
			head := x[i+1:]
			x = head
			//找坤二宫数字 再排序
			for j := 0; j < len(x); j++ {
				if x[j] == 2 {
					xx1 := x[:j]
					xx2 := x[j:]
					x = append(xx2, xx1...)
					break
				}
			}
			break
		}
		if x[i] == 5 && i == len(x)-1 { //5在末尾
			head := x[:i]
			x = head
			break
		}

		if x[i] == 5 && i != 0 && i != len(x)-1 {
			head := x[:i]
			end := x[i+1:]
			x = append(head, end...)
		}
	}
	x = append(x, 5)
	return x
}

//地盘三奇六仪配九宫 九星配宫的数字对应地盘三奇六仪的宫位信息
func StarQY(starArr []string, sqly map[int]string, dunN int, starmap map[int]string) map[int]string {
	var starQYmap = make(map[string]string)
	//先排序三奇六仪的key
	var keys []int
	for key := range sqly {
		keys = append(keys, key)
	}
	sort.Ints(keys)
	if dunN == 5 {
		x := []int{1, 8, 3, 4, 9, 2, 7, 6, 5}
		for ki := 0; ki < len(keys); ki++ {
			if keys[ki] == dunN {
				xk1 := keys[:ki]
				xk2 := keys[ki:]
				keys = append(xk2, xk1...)
				break
			}
		}
		/////
		for xi := 0; xi < len(x); xi++ {
			//中宫不在首位与末尾...
			if x[xi] == dunN { //定位宫索引
				x1 := x[:xi]
				x2 := x[xi+1:]
				x = append(x2, x1...) //剥离5出来
				for ji := 0; ji < len(x); ji++ {
					if x[ji] == 2 { //寄坤二宫
						xj1 := x[:ji]
						xj2 := x[ji:]
						x = append(xj2, xj1...)
						break
					}
				}
				break
			}
		}
		x = append(x, 5) //附加中宫到末尾
		///
		for stari := 0; stari < len(starArr); stari++ {
			for xi := 0; xi < len(x); xi++ {
				if stari == xi {
					starQYmap[starArr[stari]] = sqly[x[xi]]
					break
				}
			}
		}
	}
	if dunN != 5 {
		x := []int{1, 8, 3, 4, 9, 2, 7, 6}
		for id := 0; id < len(keys); id++ {
			if keys[id] == dunN { //这里的值是5的话要重新考虑计算排序中宫
				keys1 := keys[:id]
				keys2 := keys[id:]
				keys = append(keys2, keys1...)
				break
			}
		}
		//再排宫位
		for xi := 0; xi < len(x); xi++ {
			if x[xi] == dunN { //定位宫索引
				x1 := x[:xi]
				x2 := x[xi:]
				x = append(x2, x1...)
				x = append(x, 5) //附加中宫到末尾
				break
			}
		}
		//六仪三奇配九星
		for stari := 0; stari < len(starArr); stari++ {
			for xi := 0; xi < len(x); xi++ {
				if stari == xi {
					starQYmap[starArr[stari]] = sqly[x[xi]]
					break
				}
			}
		}
	}

	//匹配九宫数字
	var starDunmap = make(map[int]string)
	for sn, starName := range starmap {
		for k, v := range starQYmap {
			if strings.EqualFold(starName, k) {
				starDunmap[sn] = v
			}
		}
	}

	return starDunmap
}

//hGong
//当前时辰配九宫
func hGong(xunShouNumber int) []int {
	//九宫原始位顺序
	x := []int{1, 8, 3, 4, 9, 2, 7, 6, 5}
	for i := 0; i < len(x); i++ {
		if x[i] == xunShouNumber {
			x1 := x[:i]
			x2 := x[i:]
			x = append(x2, x1...)
			break
		}
	}
	return x
}
