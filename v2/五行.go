package sjqm

import "strings"

//五行生克 木火土金水 临位相生隔位相克
//比和n=0 前者生后者n=1 前者克后者n=-1 后者生前者n=2 后者克前者n=-2
func Wxsk(wx1, wx2 string) (n int) {
	wxs := []string{"木", "火", "土", "金", "水", "木", "火"}
	var w1 int
	for i := 0; i < 5; i++ {
		if strings.EqualFold(wx1, wxs[i]) {
			w1 = i
		}
	}
	var w2 int
	for j := 0; j < 5; j++ {
		if strings.EqualFold(wx2, wxs[j]) {
			w2 = j
			break
		}
	}
	for x := 0; x < len(wxs); x++ {
		//前者生/克后者
		k := w1 + 2
		s := w1 + 1
		if strings.EqualFold(wx1, wx2) {
			n = 0
			break
		}
		if strings.EqualFold(wx2, wxs[k]) {
			n = -1
			break
		}
		if strings.EqualFold(wx2, wxs[s]) {
			n = 1
			break
		}
		//后者生/克前者
		rk := w2 + 2
		rs := w2 + 1
		if strings.EqualFold(wx1, wxs[rk]) {
			n = -2
			break
		}
		if strings.EqualFold(wx1, wxs[rs]) {
			n = 2
			break
		}

	}
	return
}
