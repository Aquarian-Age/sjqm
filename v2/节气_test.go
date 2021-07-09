package sjqm

import (
	"fmt"
	"testing"
	"time"
)

func TestGetJieQiArr(t *testing.T) {
	year := 2021
	lct, arrt, jiet, zqt := GetJieQiArr(year)
	fmt.Println(len(arrt))
	fmt.Printf("立春时间:%v\n节气:%v\n12节:%v\n12中气:%v\n", lct, arrt, jiet, zqt)
	fmt.Printf("本年夏至时间:%v\n", arrt[12])
	fmt.Printf("本年冬至时间%v\n", arrt[24])
}
func TestFindJq(t *testing.T) {
	tx := time.Now().Local()
	tx = time.Date(2021, time.Month(7), 23, 15, 15, 15, 15, time.Local)
	_, _, jqt := FindJq(tx)
	fmt.Printf("**:%v\n", jqt)
	fmt.Printf("本年夏至时间:%v\n", jqt[12])
	fmt.Printf("本年冬至时间%v\n", jqt[24])
}
