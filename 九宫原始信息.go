package sjqm

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
