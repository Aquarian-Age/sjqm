package main

import (
	"gioui.org/f32"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"github.com/wunderkind2k1/gorcle"
	"image"
	"image/color"
)

var points, points2, points3, points4 []f32.Point
var blue = color.NRGBA{R: 30, G: 144, B: 255, A: 255}

//横线
func drawLine1(gtx layout.Context) {
	for x := 10; x < 350; x++ { //画横线
		point := []f32.Point{{X: float32(x), Y: float32(210)}}
		points = append(points, point...)
	}
	//
	defer op.Save(gtx.Ops).Load()
	var path clip.Path
	path.Begin(gtx.Ops)
	for i, p := range points {
		if i == 0 {
			path.MoveTo(p)
		} else {
			path.LineTo(p)
		}
	}
	style := clip.StrokeStyle{Width: 3, Miter: 3, Cap: clip.RoundCap, Join: clip.RoundJoin}
	paint.FillShape(gtx.Ops, blue, clip.Stroke{Path: path.End(), Style: style}.Op())
}
func drawLine2(gtx layout.Context) {
	for x := 10; x < 350; x++ { //画横线
		point := []f32.Point{{X: float32(x), Y: float32(360)}}
		points2 = append(points2, point...)
	}
	//
	defer op.Save(gtx.Ops).Load()
	var path clip.Path
	path.Begin(gtx.Ops)
	for i, p := range points2 {
		if i == 0 {
			path.MoveTo(p)
		} else {
			path.LineTo(p)
		}
	}
	style := clip.StrokeStyle{Width: 3, Miter: 3, Cap: clip.RoundCap, Join: clip.RoundJoin}
	paint.FillShape(gtx.Ops, blue, clip.Stroke{Path: path.End(), Style: style}.Op())
}

//竖线
func drawLine3(gtx layout.Context) {
	for y := 60; y < 500; y++ { //画横线
		point := []f32.Point{{X: float32(120), Y: float32(y)}}
		points3 = append(points3, point...)
	}
	//
	defer op.Save(gtx.Ops).Load()
	var path clip.Path
	path.Begin(gtx.Ops)
	for i, p := range points3 {
		if i == 0 {
			path.MoveTo(p)
		} else {
			path.LineTo(p)
		}
	}
	style := clip.StrokeStyle{Width: 3, Miter: 3, Cap: clip.RoundCap, Join: clip.RoundJoin}
	paint.FillShape(gtx.Ops, blue, clip.Stroke{Path: path.End(), Style: style}.Op())
}
func drawLine4(gtx layout.Context) {
	for y := 60; y < 500; y++ { //画横线
		point := []f32.Point{{X: float32(240), Y: float32(y)}}
		points4 = append(points4, point...)
	}
	//
	defer op.Save(gtx.Ops).Load()
	var path clip.Path
	path.Begin(gtx.Ops)
	for i, p := range points4 {
		if i == 0 {
			path.MoveTo(p)
		} else {
			path.LineTo(p)
		}
	}
	style := clip.StrokeStyle{Width: 3, Miter: 3, Cap: clip.RoundCap, Join: clip.RoundJoin}
	paint.FillShape(gtx.Ops, blue, clip.Stroke{Path: path.End(), Style: style}.Op())
}

//圆
//外圈 绿色  color.RGBA{0x00, 0xff, 0x00, 0xff}
func drawCircle(gtx layout.Context) {
	green := color.RGBA{0x00, 0xff, 0x00, 0xff}
	cirle := gorcle.NewCircle(green, 390)
	theImage := image.NewNRGBA(image.Rect(0, 0, 900, 900))
	_ = cirle.Draw(theImage, 400, 400)
	imgOp := paint.NewImageOp(theImage)
	imgOp.Add(gtx.Ops)
	paint.PaintOp{}.Add(gtx.Ops)
}

//var red = color.NRGBA{A: 0xff, R: 0xa0}
//
//func drawLine(points []f32.Point, gtx layout.Context) {
//	defer op.Save(gtx.Ops).Load()
//	var path clip.Path
//	path.Begin(gtx.Ops)
//	for i, p := range points {
//		if i == 0 {
//			path.MoveTo(p)
//		} else {
//			path.LineTo(p)
//		}
//	}
//	style := clip.StrokeStyle{Width: 20, Miter: 10, Cap: clip.RoundCap, Join: clip.RoundJoin}
//	paint.FillShape(gtx.Ops, red, clip.Stroke{Path: path.End(), Style: style}.Op())
//}
