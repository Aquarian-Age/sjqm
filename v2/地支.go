package sjqm

// ZHI 地支属性
type ZHI struct {
	Name      string `json:"name"`       //名称
	FangXiang string `json:"fang_xiang"` //方向(后天八卦)
	WuXing    string `json:"wu_xing"`    //五行属性
	Gua       string `json:"gua"`        //八卦名称
	YinYang   bool   `json:"yin_yang"`   //阴阳 true:阳　false:阴
	ShengXiao string `json:"sheng_xiao"` //十二生肖
}

//地支属性信息
func NewZHI(zhi string) (z *ZHI) {
	//地支
	zhiArr := []string{"子", "丑", "寅", "卯", "辰", "巳", "午", "未", "申", "酉", "戌", "亥"}
	//生肖
	sx := []string{"鼠", "牛", "虎", "兔", "龙", "蛇", "马", "羊", "猴", "鸡", "狗", "猪"}
	//八方
	bf := []string{"北", "东北", "东", "东南", "南", "西南", "西", "西北"}
	//八卦
	bg := []string{"坎", "艮", "震", "巽", "离", "坤", "兑", "乾"}
	//五行
	wx := []string{"木", "火", "土", "金", "水"}

	switch zhi {
	case zhiArr[0]: //子
		z = &ZHI{
			Name:      zhiArr[0],
			FangXiang: bf[0],
			WuXing:    wx[4],
			Gua:       bg[0],
			YinYang:   true,
			ShengXiao: sx[0],
		}
	case zhiArr[1]: //丑
		z = &ZHI{
			Name:      zhiArr[1],
			FangXiang: bf[1],
			WuXing:    wx[2],
			Gua:       bg[1],
			YinYang:   false,
			ShengXiao: sx[1],
		}
	case zhiArr[2]: //寅
		z = &ZHI{
			Name:      zhiArr[2],
			FangXiang: bf[1],
			WuXing:    wx[0],
			Gua:       bg[1],
			YinYang:   true,
			ShengXiao: sx[2],
		}
	case zhiArr[3]: //卯
		z = &ZHI{
			Name:      zhiArr[3],
			FangXiang: bf[2],
			WuXing:    wx[0],
			Gua:       bg[2],
			YinYang:   false,
			ShengXiao: sx[3],
		}
	case zhiArr[4]: //辰土龙
		z = &ZHI{
			Name:      zhiArr[4],
			FangXiang: bf[3],
			WuXing:    wx[2],
			Gua:       bg[3],
			YinYang:   true,
			ShengXiao: sx[4],
		}
	case zhiArr[5]: //巳火蛇
		z = &ZHI{
			Name:      zhiArr[5],
			FangXiang: bf[3],
			WuXing:    wx[1],
			Gua:       bg[3],
			YinYang:   false,
			ShengXiao: sx[5],
		}
	case zhiArr[6]: //午
		z = &ZHI{
			Name:      zhiArr[6],
			FangXiang: bf[4],
			WuXing:    wx[1],
			Gua:       bg[4],
			YinYang:   true,
			ShengXiao: sx[6],
		}
	case zhiArr[7]: //未土羊
		z = &ZHI{
			Name:      zhiArr[7],
			FangXiang: bf[5],
			WuXing:    wx[2],
			Gua:       bg[5],
			YinYang:   false,
			ShengXiao: sx[7],
		}
	case zhiArr[8]: //申金猴
		z = &ZHI{
			Name:      zhiArr[8],
			FangXiang: bf[5],
			WuXing:    wx[3],
			Gua:       bg[5],
			YinYang:   true,
			ShengXiao: sx[8],
		}
	case zhiArr[9]: //酉金鸡
		z = &ZHI{
			Name:      zhiArr[9],
			FangXiang: bf[6],
			WuXing:    wx[3],
			Gua:       bg[6],
			YinYang:   false,
			ShengXiao: sx[9],
		}
	case zhiArr[10]: //戌土狗
		z = &ZHI{
			Name:      zhiArr[10],
			FangXiang: bf[7],
			WuXing:    wx[2],
			Gua:       bg[7],
			YinYang:   true,
			ShengXiao: sx[10],
		}
	case zhiArr[11]: //亥水猪
		z = &ZHI{
			Name:      zhiArr[11],
			FangXiang: bf[7],
			WuXing:    wx[4],
			Gua:       bg[7],
			YinYang:   false,
			ShengXiao: sx[11],
		}
	}
	return
}
