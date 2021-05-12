package sjqm

//八神 小值符加大值符 阳顺阴逆
//直符 腾蛇 太阴 六合 白虎 玄武 九地 九天
func BaShen(zhif int, yy int) (baShenMap map[int]string) {
	//值符天禽宫数字5这里寄到坤二宫
	if zhif == 5 {
		zhif = 2
	}
	//yy:阴阳遁数字 zhif:值符宫位数字
	bs := []string{"直符", "腾蛇", "太阴", "六合", "白虎", "玄武", "九地", "九天"}
	var bsmap = make(map[int]string)
	switch yy {
	case 0: //阴遁 逆排
		x := []int{1, 6, 7, 2, 9, 4, 3, 8}
		//假定值符在9宫
		for i := 0; i < len(x); i++ {
			if x[i] == zhif {
				//重新排序
				x1 := x[:i]
				x2 := x[i:]
				x = append(x2, x1...)
				//八神配宫
				for xn := 0; xn < len(x); xn++ {
					for sn := 0; sn < len(bs); sn++ {
						if xn == sn {
							bsmap[x[xn]] = bs[sn]
							break
						}
					}
				}
				break
			}
		}
		baShenMap = bsmap
	case 1: //阳遁 顺排
		x := []int{1, 8, 3, 4, 9, 2, 7, 6}
		for i := 0; i < len(x); i++ {
			if x[i] == zhif {
				//重新排序
				x1 := x[:i]
				x2 := x[i:]
				x = append(x2, x1...)
				//八神配宫
				for xn := 0; xn < len(x); xn++ {
					for sn := 0; sn < len(bs); sn++ {
						if xn == sn {
							bsmap[x[xn]] = bs[sn]
							break
						}
					}
				}
				break
			}
		}
		baShenMap = bsmap
	}
	return
}
