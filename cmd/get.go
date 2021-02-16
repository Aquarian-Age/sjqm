package cmd

import (
	"flag"
	"fmt"
	"log"
	"strconv"
)

//命令行輸入內容
type Input struct {
	Y int //年數字
	M int //月數字
	D int //日數字
	H int //時辰數字　1子時　2丑時...１２亥時
	//	B bool //閏月判斷　f非閏月　t當月爲閏月
}

func GetInPut() *Input {
	flag.Usage = func() {
		fmt.Printf("Usage sjqm -i 2020050602f\n")
		flag.PrintDefaults()
	}
	s := flag.String("i", " ", " ")
	flag.Parse()

	rs := []rune(*s)

	ys := string(rs[:4])  //年數字
	ms := string(rs[4:6]) //月數字
	ds := string(rs[6:8])
	hs := string(rs[8:10]) //時辰數字1子時　2丑時...１２亥時
	//bs := string(rs[10:11])

	y, err := strconv.Atoi(ys)
	if err != nil {
		log.Fatal("年份時間解析:", err)
	}

	m, err := strconv.Atoi(ms)
	if err != nil {
		log.Fatal("月份時間解析:", err)
	}
	d, err := strconv.Atoi(ds)
	if err != nil {
		log.Fatal("日期時間解析:", err)
	}
	h, err := strconv.Atoi(hs)
	if err != nil {
		log.Fatal("時辰解析:", err)
	}
	//b, err := strconv.ParseBool(bs)
	//if err != nil {
	//	log.Fatal("閏月解析:", err)
	//}
	return &Input{
		Y: y,
		D: d,
		M: m,
		H: h,
		//	B: b,
	}
}
