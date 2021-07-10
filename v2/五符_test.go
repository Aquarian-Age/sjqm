package sjqm

import (
	"fmt"
	"testing"
)

//五符
func TestWuFu(t *testing.T) {
	dgz := "乙丑"
	want := map[string]string{
		"丑": "地轴", "亥": "国印", "午": "风伯", "卯": "五符", "子": "天官", "寅": "天贼",
		"巳": "地符", "戌": "唐符", "未": "雷公", "申": "雨师", "辰": "天曹", "酉": "风云"}
	wfmap := wuFu(dgz)
	fmt.Println("五符:", wfmap)
	fmt.Println(want)
	//if !reflect.DeepEqual(wfmap, want) {
	//	t.Errorf("func wuFu(%s)=%v 期望值:%v\n", dgz, wfmap, want)
	//}
}

// [未:雷公 亥:国印 卯:五符 辰:天曹 午:风伯 戌:唐符 子:天官 丑:地轴 寅:天贼 巳:地符 申:雨师 酉:风云]
//[丑:地轴 亥:国印 午:风伯 卯:五符 子:天官 寅:天贼 巳:地符 戌:唐符 未:雷公 申:雨师 辰:天曹 酉:风云]
