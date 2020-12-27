package sjqm

import (
	"time"
)

//定节气
func FindJQ(st time.Time, jqt []time.Time, jqnames []string) (index int, jmc string) {
	//st:当前时间戳 jqt:节气时间戳数组 上一年冬至(含)开始 0:冬至　1:小寒　2:大寒...
	var begint, endt time.Time
	//精确度统一到小时
	st = time.Date(st.Year(), st.Month(), st.Day(), st.Hour(), 0, 0, 0, time.Local)
	for i := 2; i < 27; i++ { //i=2从上年大寒开始 因为输入月份为农历年 i=3从立春开始 i=26本年大寒
		begint = jqt[i]
		begint = time.Date(begint.Year(), begint.Month(), begint.Day(), begint.Hour(), 0, 0, 0, time.Local)
		endt = jqt[i+1]
		endt = time.Date(endt.Year(), endt.Month(), endt.Day(), endt.Hour(), 0, 0, 0, time.Local)

		if (st.After(begint) || st.Equal(begint)) && st.Before(endt) { //当前时间前的一个节气
			index = i        //节气索引
			jmc = jqnames[i] //当前时间所属的节气名称
			break
		}
	}
	return
}
