package sjqm

/*
六仪 戊己庚辛壬癸
三奇 乙丙丁
阳遁顺布六仪逆布三奇
阴遁逆布六仪顺布三奇
甲子遁戊　甲戌遁己　甲申遁庚　甲午遁辛　甲辰遁壬　甲寅遁癸
*/
//排地盘局
func FindSqly(yy int, juN int) (sqly map[int]string) {
	//yy:阴阳数字　n:局数子
	switch yy {
	case 1: //阳
		sqly = SqlyYangDun(juN)
	case 0: //阴
		sqly = SqlyYinDun(juN)
	}
	return
}

//排阴遁地盘三奇六仪
func SqlyYinDun(juN int) map[int]string {
	//juN 定局数字
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	gzarr := []string{"戊", "己", "庚", "辛", "壬", "癸", "丁", "丙", "乙"} //阴遁逆布六仪顺布三奇
	gzmap := make(map[int]string)
	//重新排序
	for j := 0; j < len(x); j++ {
		n := juN - j
		if n < 1 {
			n += 9
		}
		gzmap[n] = gzarr[j]
	}
	return gzmap
}

//排阳遁地盘三奇六仪
func SqlyYangDun(juN int) map[int]string {
	//juN 定局数字
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	gzarr := []string{"戊", "己", "庚", "辛", "壬", "癸", "丁", "丙", "乙"} //阳遁顺布六仪逆布三奇
	gzmap := make(map[int]string)
	//重新排序
	for j := 0; j <= len(x); j++ {
		if x[j] == juN {
			x1 := x[j:]
			x2 := x[:j]
			x = append(x1, x2...)
			break
		}
	}

	for i := 0; i < len(x); i++ {
		gzmap[x[i]] = gzarr[i]
	}

	return gzmap
}
