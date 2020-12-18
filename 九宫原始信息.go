package sjqm

import (
	"liangzi.local/nongli/ganzhi"
	"strings"
)

//九宫原始信息 少了中宫....
func JGMap() map[int]JG {
	var jginfo = make(map[int]JG)
	jginfo = map[int]JG{
		/////////////四阳宫
		1: {Name: "坎", Number: 1, Star: "天蓬", Door: "休门", WX: "水", ZiBai: "一白"}, //坎一宫
		8: {Name: "艮", Number: 8, Star: "天任", Door: "生门", WX: "土", ZiBai: "八白"}, //艮八宫
		3: {Name: "震", Number: 3, Star: "天冲", Door: "伤门", WX: "木", ZiBai: "三碧"}, //震三宫
		4: {Name: "巽", Number: 4, Star: "天辅", Door: "杜门", WX: "木", ZiBai: "四绿"}, //巽四宫
		//////////////四阴宫
		9: {Name: "离", Number: 9, Star: "天英", Door: "景门", WX: "火", ZiBai: "九紫"}, //离九宫
		2: {Name: "坤", Number: 2, Star: "天芮", Door: "死门", WX: "土", ZiBai: "二黑"}, //坤二宫
		7: {Name: "兑", Number: 7, Star: "天柱", Door: "惊门", WX: "金", ZiBai: "七赤"}, //兑七宫
		6: {Name: "乾", Number: 6, Star: "天心", Door: "开门", WX: "金", ZiBai: "六白"}, //乾六宫
		////////////////
		5: {Name: "中", Number: 5, Star: "天禽", Door: "死门", WX: "土", ZiBai: "五黄"}, //中宫
	}
	return jginfo
}

//九星休旺
//旺:我生之月 相:月类之月 死:生我之月 囚:克我之月 休:我克之月
func NewJXXW(starName string) *JXXW {
	zhi := []string{"寅", "卯", "辰", "巳", "午", "未", "申", "酉", "戌", "亥", "子", "丑"}
	jxxw := new(JXXW)
	for _, g := range JGMap() {
		if strings.EqualFold(g.Star, starName) {
			wx1 := g.WX
			wang, xiang, si, qiu, xiu := wxsqx(zhi, wx1)
			//fmt.Printf("-->%s-%s 旺:%s 相:%s 死:%s 囚:%s 休:%s\n", g.Star, g.WX, wang, xiang, si, qiu, xiu)
			jxxw = &JXXW{
				StarName: g.Star,
				Wang:     wang,
				Xiang:    xiang,
				Si:       si,
				Qiu:      qiu,
				Xiu:      xiu,
			}
			break
		}
	}
	return jxxw
}

func wxsqx(zhi []string, wx1 string) (wang, xiang, si, qiu, xiu []string) {
	for i := 0; i < len(zhi); i++ {
		zs := ganzhi.NewZHI(zhi[i])
		if strings.EqualFold(zhi[i], zs.Name) {
			wx2 := zs.WuXing //五行属性
			n := ganzhi.Wxsk(wx1, wx2)
			switch n {
			case 1: //我生 wx1生wx2
				wang = append(wang, zhi[i])
			case 0: //类 wx1比和wx2
				xiang = append(xiang, zhi[i])
			case 2: //生我 wx2生wx1
				si = append(si, zhi[i])
			case -2: //克我 wx2克wx1
				qiu = append(qiu, zhi[i])
			case -1: //我克 wx1克wx2
				xiu = append(xiu, zhi[i])
			}
		}
	}
	return
}
