package sjqm

import (
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

//九星: 天蓬 天任 天冲 天辅 天英 天芮 天柱 天心 天禽(中宫)寄坤二宫即天芮星
//排值符九星
func ZhiFuStar(xunShouNumber int, dun string, sqly map[int]string) (string, map[int]string, map[int]string) {
	//根据时辰旬首 排值符九星
	zf, _ := 旬遁原始宫位值符值使(dun, sqly)
	starArr := 九星排序(zf)
	xArr := 时辰配宫(xunShouNumber) //值符随时干 如果xunShouNumber==5即值符落中宫 在delE函数实现让它寄坤二宫
	xArr = del(xArr)            //中宫天禽配九星的天禽顺序

	//fmt.Printf("-->delE:时辰配宫:%v\n", xArr)
	//九星配时辰(九宫)
	l1 := len(starArr)
	l2 := len(xArr)
	var starmap = make(map[int]string)
	if l1 == l2 {
		for i := 0; i < len(xArr); i++ {
			for j := i; j < len(starArr); j++ {
				starmap[xArr[i]] = starArr[j]
				break
			}
		}
	}
	//fmt.Printf("-->值使九星配九宫:%v\n", starmap)

	///地盘原始宫位的三奇六仪
	starDun := 地盘三奇六仪配九星(starArr, sqly)
	return zf, starmap, starDun
}

//剥离5
func del(x []int) []int {
	for i := 0; i < len(x); i++ {
		if x[i] == 5 && i == 0 { //5在首位
			//fmt.Printf("-->i=%d x[]i]=%d\n", i, x[i])
			head := x[i+1:]
			x = head
			//fmt.Printf("-->del5:%d\n", x)
			//找坤二宫数字 再排序
			for j := 0; j < len(x); j++ {
				if x[j] == 2 {
					xx1 := x[:j]
					xx2 := x[j:]
					x = append(xx2, xx1...)
					//fmt.Printf("--del5re:%d\n", x)
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
	//fmt.Printf("-->x5:%v\n", x)
	return x
}

//地盘三奇六仪配九宫 九星配宫的数字对应地盘三奇六仪的宫位信息
func 地盘三奇六仪配九星(starArr []string, sqly map[int]string) map[int]string {
	var starDun = make(map[int]string)
	for i := 0; i < len(starArr); i++ {
		for k, v := range sqly {
			if i+1 == k {
				starDun[k] = v
			}
		}
	}
	//fmt.Printf("-->starDun:%v\n", starDun)
	return starDun
}

//根据时辰所在旬遁找出原宫位的值符值使
func 旬遁原始宫位值符值使(dun string, sqly map[int]string) (zf string, dunN int) {
	//确定六甲旬遁落那个宫位
	for n, name := range sqly {
		if strings.EqualFold(dun, name) {
			//fmt.Printf("六甲旬遁:%s 落三奇六仪%d宫位\n", dun, n)
			//根据三奇六仪旬遁宫位数找本宫位原始值符值使
			for number, g := range JGMap() {
				if n == number {
					zf = g.ZF //原始宫位值符
					dunN = n  //旬遁落宫数字
					break
				}
			}
			break
		}
	}
	//fmt.Printf("-->值符:%s 六甲旬遁:%v\n", zf, dunN) //六甲原始宫位旬遁
	return
}

//九星排序
func 九星排序(zf string) []string {
	//从当前值符对九星重新排序
	//jiu := []string{"天蓬", "天任", "天冲", "天辅", "天英", "天芮", "天柱", "天心", "天禽"}
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
	starArr := append(jiux, "天禽") //放到最后 寄坤二宫?
	//fmt.Printf("-->九星排序: %v\n", starArr)
	return starArr
}

//时辰配宫
//当前时辰配九宫
func 时辰配宫(xunShouNumber int) []int {
	//九宫原始位顺序
	x := []int{1, 8, 3, 4, 9, 2, 7, 6, 5}
	for i := 0; i < len(x); i++ {
		if x[i] == xunShouNumber {
			x1 := x[:i]
			x2 := x[i:]
			x = append(x2, x1...)
			//fmt.Println("-->时辰重排数字(即时辰对应的九宫数字):", x)
			break
		}
	}
	return x
}
