package sjqm

import "time"

//st:当前输入时间戳 dzt:冬至时间戳　xzt:夏至时间戳
//时间精确到小时
// 判断阴阳遁 0:阴遁 1:阳遁
// 夏(含)至以后用阴遁 冬至(含)以后用阳遁
func YinYang(st, dzt, xzt time.Time) (yy int) {
	st = time.Date(st.Year(), st.Month(), st.Day(), st.Hour(), 0, 0, 0, time.Local)
	dzt = time.Date(dzt.Year(), dzt.Month(), dzt.Day(), dzt.Hour(), 0, 0, 0, time.Local)
	xzt = time.Date(xzt.Year(), xzt.Month(), xzt.Day(), xzt.Hour(), 0, 0, 0, time.Local)

	if st.After(xzt) || st.Equal(xzt) {
		yy = 0 //阴
	} else if st.Before(xzt) {
		yy = 1
	}
	if st.After(dzt) || st.Equal(dzt) {
		yy = 1 //阳
	}
	return
}
