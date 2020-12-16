//奇门方法 克应
package sjqm

import (
	"liangzi.local/sjqm/qm"
	"strings"
)

//QM结构体 yj:月将对应的地支
func NewQM(yjZhi, dgz, hgz string) *QM {
	qmap := new(QM)
	siHumap := qm.DiSiHu(hgz)
	diSiMenMap := qm.DiSiMen(yjZhi, dgz, hgz)
	tms := qm.TianMa(yjZhi, hgz)
	tianSanMenMap := qm.TianSanMen(yjZhi, hgz)
	wuFuMap := qm.WuFu(dgz)

	qmap = &QM{
		DiSiHuMap:     siHumap,
		DiSiMenMap:    diSiMenMap,
		TianMaS:       tms,
		TianSanMenMap: tianSanMenMap,
		WuFuMap:       wuFuMap,
	}
	return qmap
}

//地私门克应
func (qm *QM) DiSiMenKY() map[string]string {
	dsm := qm.DiSiMenMap
	kymap := map[string]string{
		"天乙": "行远良人,贵人车马,长者欢欣",
		"螣蛇": "虚惊怪异,半途而回,风雨相阻,若有鸦鸣,有人追捕",
		"朱雀": "遇生色物,远闻鼓声,文书无阻",
		"六合": "路逢车马,阴人彩衣,儿童戏耍",
		"勾陈": "路逢鬪打,作事勾留,谋为参差",
		"青龙": "喜兆则马路逢官吏,锦衣、奇花",
		"天空": "贱物载道,阳增阴道,类聚笑话",
		"白虎": "见死闻悲,官事惊迫,途逢兵革",
		"太常": "酒食巫师,或为优使,彩童神轴",
		"玄武": "盗贼亡失,若非牙偤,即是乞儿",
		"太阴": "小求大得,阴私和合,音乐相随",
		"天后": "童子戏耍,妇人送物,女子还家",
	}
	var dsmkymap = make(map[string]string)
	for name, _ := range dsm {
		for kn, info := range kymap {
			if strings.EqualFold(name, kn) {
				dsmkymap[kn] = info
			}
		}
	}
	return dsmkymap
}
