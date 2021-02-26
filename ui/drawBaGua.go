package main

import (
	"gioui.org/f32"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"image/color"
)

var (
	black  = color.NRGBA{R: 0, G: 0, B: 0, A: 255}      //黑色
	red    = color.NRGBA{A: 0xff, R: 0xa0}              //离卦颜色
	green  = color.NRGBA{R: 50, G: 205, B: 50, A: 255}  //震巽卦绿色
	yellow = color.NRGBA{R: 255, G: 215, B: 0, A: 255}  //艮卦黄色
	gold   = color.NRGBA{R: 218, G: 165, B: 32, A: 255} //金色
)

//顺序由下向上
var (
	kan1a, kan1b, kan2, kan3a, kan3b           []f32.Point //坎卦
	qian1, qian2, qian3                        []f32.Point //乾卦
	duil1, duil2, duil3a, duil3b               []f32.Point //兑卦
	kun1a, kun1b, kun2a, kun2b, kun3a, kun3b   []f32.Point //坤卦
	lil1, lil2a, lil2b, lil3                   []f32.Point //离卦
	xunl0, xunl1, xunl2, xunl3                 []f32.Point //巽卦
	zhenl0, zhenl1a, zhenl1b, zhenl2a, zhenl2b []f32.Point //震卦
	genl1a, genl1b, genll1a, genll1b, genl     []f32.Point //艮卦
)

//巽
func XunGua(gtx layout.Context) {
	xun3(gtx)
	xun2(gtx)
	xun(gtx)
	xun1(gtx)
}

//震卦
func Zhen(gtx layout.Context) {
	zhenllb(gtx)
	zhenlla(gtx)
	zhenlb(gtx)
	zhenla(gtx)
	zhen0(gtx)
}

//艮卦
func Gen(gtx layout.Context) {
	gen(gtx)
	genlla(gtx)
	genllb(gtx)
	genla(gtx)
	genlb(gtx)
}

//坎卦 x:=175 x<198 Y:495-486-477
func Kan(gtx layout.Context) {
	kanlllb(gtx)
	kanllla(gtx)
	kanll(gtx)
	kanlb(gtx)
	kanla(gtx)
}

//乾卦 x := 310; x < 333
func Qian(gtx layout.Context) {
	qianlll(gtx)
	qianll(gtx)
	qianl(gtx)
}

//兑卦
func Dui(gtx layout.Context) {
	duilllb(gtx)
	duillla(gtx)
	duill(gtx)
	duil(gtx)
}

//坤卦
func Kun(gtx layout.Context) {
	kunlllb(gtx)
	kunllla(gtx)
	kunllb(gtx)
	kunlla(gtx)
	kunlb(gtx)
	kunla(gtx)
}

//离卦
func Li(gtx layout.Context) {
	lill(gtx)
	lilb(gtx)
	lila(gtx)
	lil(gtx)
}

/////////////////////
//艮上
func gen(gtx layout.Context) {
	for x := 60; x <= 83; x++ {
		point := []f32.Point{{X: float32(x), Y: float32(475)}}
		genl = append(genl, point...)
	}
	defer op.Save(gtx.Ops).Load()
	genDx(genl, gtx)
}

//艮中
func genlla(gtx layout.Context) {
	for x := 73; x < 83; x++ {
		point := []f32.Point{{X: float32(x), Y: float32(486)}}
		genll1b = append(genll1b, point...)
	}
	defer op.Save(gtx.Ops).Load()
	genDx(genll1b, gtx)
}
func genllb(gtx layout.Context) {
	for x := 60; x < 70; x++ {
		point := []f32.Point{{X: float32(x), Y: float32(486)}}
		genll1a = append(genll1a, point...)
	}
	defer op.Save(gtx.Ops).Load()
	genDx(genll1a, gtx)
}

//艮下
func genla(gtx layout.Context) {
	for x := 73; x < 83; x++ {
		point := []f32.Point{{X: float32(x), Y: float32(495)}}
		genl1b = append(genl1b, point...)
	}
	defer op.Save(gtx.Ops).Load()
	genDx(genl1b, gtx)
}
func genlb(gtx layout.Context) {
	for x := 60; x < 70; x++ {
		point := []f32.Point{{X: float32(x), Y: float32(495)}}
		genl1a = append(genl1a, point...)
	}
	defer op.Save(gtx.Ops).Load()
	genDx(genl1a, gtx)
}
func genDx(l []f32.Point, gtx layout.Context) {
	var path clip.Path
	path.Begin(gtx.Ops)
	for i, p := range l {
		if i == 0 {
			path.MoveTo(p)
		} else {
			path.LineTo(p)
		}
	}
	style := clip.StrokeStyle{Width: 1, Miter: 1, Cap: clip.RoundCap, Join: clip.RoundJoin}
	paint.FillShape(gtx.Ops, yellow, clip.Stroke{Path: path.End(), Style: style}.Op())
}

///////////////////////

//坎上
func kanllla(gtx layout.Context) {
	for x := 175; x < 185; x++ {
		point := []f32.Point{{X: float32(x), Y: float32(477)}}
		kan3a = append(kan3a, point...)
	}
	defer op.Save(gtx.Ops).Load()
	kanDx(kan3a, gtx)
}
func kanlllb(gtx layout.Context) {
	for x := 188; x < 198; x++ {
		point := []f32.Point{{X: float32(x), Y: float32(477)}}
		kan3b = append(kan3b, point...)
	}
	defer op.Save(gtx.Ops).Load()
	kanDx(kan3b, gtx)
}

//坎中
func kanll(gtx layout.Context) {
	for x := 175; x < 198; x++ {
		point := []f32.Point{{X: float32(x), Y: float32(486)}}
		kan2 = append(kan2, point...)
	}
	defer op.Save(gtx.Ops).Load()
	kanDx(kan2, gtx)
}

//坎下
func kanla(gtx layout.Context) {
	for x := 175; x < 185; x++ {
		point := []f32.Point{{X: float32(x), Y: float32(495)}}
		kan1a = append(kan1a, point...)
	}
	defer op.Save(gtx.Ops).Load()
	kanDx(kan1a, gtx)
}
func kanlb(gtx layout.Context) {
	for x := 188; x < 198; x++ {
		point := []f32.Point{{X: float32(x), Y: float32(495)}}
		kan1b = append(kan1b, point...)
	}
	defer op.Save(gtx.Ops).Load()
	kanDx(kan1b, gtx)
}
func kanDx(l []f32.Point, gtx layout.Context) {
	var path clip.Path
	path.Begin(gtx.Ops)
	for i, p := range l {
		if i == 0 {
			path.MoveTo(p)
		} else {
			path.LineTo(p)
		}
	}
	style := clip.StrokeStyle{Width: 1, Miter: 1, Cap: clip.RoundCap, Join: clip.RoundJoin}
	paint.FillShape(gtx.Ops, black, clip.Stroke{Path: path.End(), Style: style}.Op())
}

///////////////////////

//乾上 Y:495-9-9
func qianlll(gtx layout.Context) {
	for x := 310; x < 333; x++ {
		point := []f32.Point{{X: float32(x), Y: float32(477)}}
		qian3 = append(qian3, point...)
	}
	defer op.Save(gtx.Ops).Load()
	qianDx(qian3, gtx)
}

//乾中 Y:495-9
func qianll(gtx layout.Context) {
	for x := 310; x < 333; x++ {
		point := []f32.Point{{X: float32(x), Y: float32(486)}}
		qian2 = append(qian2, point...)
	}
	defer op.Save(gtx.Ops).Load()
	qianDx(qian2, gtx)
}

//乾下 Y:495
func qianl(gtx layout.Context) {
	for x := 310; x < 333; x++ {
		point := []f32.Point{{X: float32(x), Y: float32(495)}}
		qian1 = append(qian1, point...)
	}
	defer op.Save(gtx.Ops).Load()
	qianDx(qian1, gtx)
}

func qianDx(l []f32.Point, gtx layout.Context) {
	var path clip.Path
	path.Begin(gtx.Ops)
	for i, p := range l {
		if i == 0 {
			path.MoveTo(p)
		} else {
			path.LineTo(p)
		}
	}
	style := clip.StrokeStyle{Width: 1, Miter: 1, Cap: clip.RoundCap, Join: clip.RoundJoin}
	paint.FillShape(gtx.Ops, gold, clip.Stroke{Path: path.End(), Style: style}.Op())
}

///////////////////////

//兑上
func duillla(gtx layout.Context) {
	for x := 310; x < 320; x++ {
		point := []f32.Point{{X: float32(x), Y: float32(332)}}
		duil3a = append(duil3a, point...)
	}
	defer op.Save(gtx.Ops).Load()
	duiDx(duil3a, gtx)
}
func duilllb(gtx layout.Context) {
	for x := 323; x < 333; x++ {
		point := []f32.Point{{X: float32(x), Y: float32(332)}}
		duil3b = append(duil3b, point...)
	}
	defer op.Save(gtx.Ops).Load()
	duiDx(duil3b, gtx)
}

//兑中
func duill(gtx layout.Context) {
	for x := 310; x < 333; x++ {
		point := []f32.Point{{X: float32(x), Y: float32(341)}}
		duil2 = append(duil2, point...)
	}
	defer op.Save(gtx.Ops).Load()
	duiDx(duil2, gtx)
}

//兑下
func duil(gtx layout.Context) {
	for x := 310; x < 333; x++ {
		point := []f32.Point{{X: float32(x), Y: float32(350)}}
		duil1 = append(duil1, point...)
	}
	defer op.Save(gtx.Ops).Load()
	duiDx(duil1, gtx)
}
func duiDx(l []f32.Point, gtx layout.Context) {
	var path clip.Path
	path.Begin(gtx.Ops)
	for i, p := range l {
		if i == 0 {
			path.MoveTo(p)
		} else {
			path.LineTo(p)
		}
	}
	style := clip.StrokeStyle{Width: 1, Miter: 1, Cap: clip.RoundCap, Join: clip.RoundJoin}
	paint.FillShape(gtx.Ops, gold, clip.Stroke{Path: path.End(), Style: style}.Op())
}

//////////////////////
//震上
func zhenlla(gtx layout.Context) {
	for x := 60; x < 70; x++ {
		point := []f32.Point{{X: float32(x), Y: float32(332)}}
		zhenl2a = append(zhenl2a, point...)
	}
	defer op.Save(gtx.Ops).Load()
	zhenDx(zhenl2a, gtx)
}
func zhenllb(gtx layout.Context) {
	for x := 73; x < 83; x++ {
		point := []f32.Point{{X: float32(x), Y: float32(332)}}
		zhenl2b = append(zhenl2b, point...)
	}
	defer op.Save(gtx.Ops).Load()
	zhenDx(zhenl2b, gtx)
}

//震中
func zhenla(gtx layout.Context) {
	for x := 60; x < 70; x++ {
		point := []f32.Point{{X: float32(x), Y: float32(341)}}
		zhenl1a = append(zhenl1a, point...)
	}
	defer op.Save(gtx.Ops).Load()
	zhenDx(zhenl1a, gtx)
}
func zhenlb(gtx layout.Context) {
	for x := 73; x < 83; x++ {
		point := []f32.Point{{X: float32(x), Y: float32(341)}}
		zhenl1b = append(zhenl1b, point...)
	}
	defer op.Save(gtx.Ops).Load()
	zhenDx(zhenl1b, gtx)
}

//震下
func zhen0(gtx layout.Context) {
	for x := 60; x < 83; x++ {
		point := []f32.Point{{X: float32(x), Y: float32(350)}}
		zhenl0 = append(zhenl0, point...)
	}
	defer op.Save(gtx.Ops).Load()
	zhenDx(zhenl0, gtx)
}
func zhenDx(l []f32.Point, gtx layout.Context) {
	var path clip.Path
	path.Begin(gtx.Ops)
	for i, p := range l {
		if i == 0 {
			path.MoveTo(p)
		} else {
			path.LineTo(p)
		}
	}
	style := clip.StrokeStyle{Width: 1, Miter: 1, Cap: clip.RoundCap, Join: clip.RoundJoin}
	paint.FillShape(gtx.Ops, green, clip.Stroke{Path: path.End(), Style: style}.Op())
}

///////////////////////

//坤上
func kunllla(gtx layout.Context) {
	for x := 310; x < 320; x++ {
		point := []f32.Point{{X: float32(x), Y: float32(175)}}
		kun3a = append(kun3a, point...)
	}
	defer op.Save(gtx.Ops).Load()
	kunDx(kun3a, gtx)
}
func kunlllb(gtx layout.Context) {
	for x := 323; x < 333; x++ {
		point := []f32.Point{{X: float32(x), Y: float32(175)}}
		kun3b = append(kun3b, point...)
	}
	defer op.Save(gtx.Ops).Load()
	kunDx(kun3b, gtx)
}

//坤中
func kunlla(gtx layout.Context) {
	for x := 310; x < 320; x++ {
		point := []f32.Point{{X: float32(x), Y: float32(186)}}
		kun2a = append(kun2a, point...)
	}
	defer op.Save(gtx.Ops).Load()
	kunDx(kun2a, gtx)
}
func kunllb(gtx layout.Context) {
	for x := 323; x < 333; x++ {
		point := []f32.Point{{X: float32(x), Y: float32(186)}}
		kun2b = append(kun2b, point...)
	}
	defer op.Save(gtx.Ops).Load()
	kunDx(kun2b, gtx)
}

//坤下
func kunla(gtx layout.Context) {
	for x := 310; x < 320; x++ {
		point := []f32.Point{{X: float32(x), Y: float32(195)}}
		kun1a = append(kun1a, point...)
	}
	defer op.Save(gtx.Ops).Load()
	kunDx(kun1a, gtx)
}
func kunlb(gtx layout.Context) {
	for x := 323; x < 333; x++ {
		point := []f32.Point{{X: float32(x), Y: float32(195)}}
		kun1b = append(kun1b, point...)
	}
	defer op.Save(gtx.Ops).Load()
	kunDx(kun1b, gtx)
}
func kunDx(l []f32.Point, gtx layout.Context) {
	var path clip.Path
	path.Begin(gtx.Ops)
	for i, p := range l {
		if i == 0 {
			path.MoveTo(p)
		} else {
			path.LineTo(p)
		}
	}
	style := clip.StrokeStyle{Width: 1, Miter: 1, Cap: clip.RoundCap, Join: clip.RoundJoin}
	paint.FillShape(gtx.Ops, yellow, clip.Stroke{Path: path.End(), Style: style}.Op())
}

/////////////////

//离上
func lill(gtx layout.Context) {
	for x := 175; x < 198; x++ {
		point := []f32.Point{{X: float32(x), Y: float32(175)}}
		lil3 = append(lil3, point...)
	}
	defer op.Save(gtx.Ops).Load()
	liDx(lil3, gtx)
}

//离中
func lila(gtx layout.Context) {
	for x := 175; x < 185; x++ {
		point := []f32.Point{{X: float32(x), Y: float32(186)}}
		lil2a = append(lil2a, point...)
	}
	defer op.Save(gtx.Ops).Load()
	liDx(lil2a, gtx)
}
func lilb(gtx layout.Context) {
	for x := 188; x < 198; x++ {
		point := []f32.Point{{X: float32(x), Y: float32(186)}}
		lil2b = append(lil2b, point...)
	}
	defer op.Save(gtx.Ops).Load()
	liDx(lil2b, gtx)
}

//离下
func lil(gtx layout.Context) {
	for x := 175; x < 198; x++ {
		point := []f32.Point{{X: float32(x), Y: float32(195)}}
		lil1 = append(lil1, point...)
	}
	defer op.Save(gtx.Ops).Load()
	liDx(lil1, gtx)
}
func liDx(l []f32.Point, gtx layout.Context) {
	var path clip.Path
	path.Begin(gtx.Ops)
	for i, p := range l {
		if i == 0 {
			path.MoveTo(p)
		} else {
			path.LineTo(p)
		}
	}
	style := clip.StrokeStyle{Width: 1, Miter: 1, Cap: clip.RoundCap, Join: clip.RoundJoin}
	paint.FillShape(gtx.Ops, red, clip.Stroke{Path: path.End(), Style: style}.Op())
}

///////////////////////////////
//巽卦上
func xun3(gtx layout.Context) {
	for x3 := 60; x3 < 83; x3++ {
		point := []f32.Point{{X: float32(x3), Y: float32(177)}}
		xunl3 = append(xunl3, point...)
	}
	defer op.Save(gtx.Ops).Load()
	dx(xunl3, gtx)
}

//巽卦中
func xun2(gtx layout.Context) {
	for x1 := 60; x1 < 83; x1++ {
		point := []f32.Point{{X: float32(x1), Y: float32(186)}}
		xunl2 = append(xunl2, point...)
	}
	defer op.Save(gtx.Ops).Load()
	dx(xunl2, gtx)
}

//巽卦下
func xun(gtx layout.Context) {
	for x := 60; x < 70; x++ { //画横线
		point := []f32.Point{{X: float32(x), Y: float32(195)}}
		xunl0 = append(xunl0, point...)
	}
	defer op.Save(gtx.Ops).Load()
	dx(xunl0, gtx)

}
func xun1(gtx layout.Context) {
	for x1 := 73; x1 < 83; x1++ {
		point := []f32.Point{{X: float32(x1), Y: float32(195)}}
		xunl1 = append(xunl1, point...)
	}
	defer op.Save(gtx.Ops).Load()
	dx(xunl1, gtx)
}

func dx(l []f32.Point, gtx layout.Context) {
	var path clip.Path
	path.Begin(gtx.Ops)
	for i, p := range l {
		if i == 0 {
			path.MoveTo(p)
		} else {
			path.LineTo(p)
		}
	}
	style := clip.StrokeStyle{Width: 1, Miter: 1, Cap: clip.RoundCap, Join: clip.RoundJoin}
	paint.FillShape(gtx.Ops, green, clip.Stroke{Path: path.End(), Style: style}.Op())
}
