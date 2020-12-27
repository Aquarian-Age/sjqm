package sjqm

import (
	"strings"
)

type JiuXingJX struct {
	TianPeng
	TianRen
	TianChong
	TianFu
	TianYing
	TianRui
	TianZhu
	TianXin
	TianQin
}

//讼庭争竞遇天蓬,胜捷名威万里同,春夏用之皆大吉,秋冬用之半为凶;
//嫁娶远行皆不利,修造埋葬亦闲空,须得生门同丙乙,用之万事得昌隆。
type TianPeng struct {
	Name string
	Info string
}
type TianRen struct {
	Name string
	Info string
}
type TianChong struct {
	Name string
	Info string
}
type TianFu struct {
	Name string
	Info string
}
type TianYing struct {
	Name string
	Info string
}
type TianRui struct {
	Name string
	Info string
}
type TianZhu struct {
	Name string
	Info string
}
type TianXin struct {
	Name string
	Info string
}
type TianQin struct {
	Name string
	Info string
}

func NewTianPeng(starName string) (tp *TianPeng) {
	tp = new(TianPeng)
	zfarr := []string{"天蓬", "天任", "天冲", "天辅", "天英", "天芮", "天柱", "天心", "天禽"}
	for i := 0; i < len(zfarr); i++ {
		if strings.EqualFold(zfarr[i], starName) {
			switch i {
			case 0:
				tp = &TianPeng{
					Name: starName,
					Info: "讼庭争竞遇天蓬,胜捷名威万里同,春夏用之皆大吉,秋冬用之半为凶;\n嫁娶远行皆不利,修造埋葬亦闲空,须得生门同丙乙,用之万事得昌隆",
				}
			}
			break
		}
	}
	return
}
