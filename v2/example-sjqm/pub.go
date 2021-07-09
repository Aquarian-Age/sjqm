package main

//周几 0,7周日　1:周一
func WeekName(n int) (w string) {
	switch n {
	case 0, 7:
		w = "周日"
	case 1:
		w = "周一"
	case 2:
		w = "周二"
	case 3:
		w = "周三"
	case 4:
		w = "周四"
	case 5:
		w = "周五"
	case 6:
		w = "周六"
	}
	return
}
