package main

import (
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
	"github.com/ying32/govcl/vcl/types/colors"
)

func (f *TMainForm) draw() {
	//画
	const (
		top = 10
		l1  = 10
		r1  = 130 //170 //220
		b1  = 150 //215
		t3  = 2 * b1
		b3  = 3 * b1
	)
	f.SetOnPaint(func(sender vcl.IObject) {
		canvas := f.Canvas()
		//4
		r := types.TRect{l1, top, r1, b1}
		canvas.FillRect(r)
		//2
		r = types.TRect{300, top, 430, b1}
		canvas.FillRect(r)
		//8
		r = types.TRect{Left: l1, Top: t3, Right: r1, Bottom: b3}
		canvas.FillRect(r)
		//6
		r = types.TRect{300, t3, 430, b3}
		canvas.FillRect(r)

	})

}
func (f *TMainForm) OnFormCreatG12() {
	const top1 = 100
	const left = 90
	const font = 16
	const color = colors.ClBlack

	f.kan = vcl.NewLabel(f)
	f.kan.SetParent(f)
	f.kan.SetTop(4*top1 + 20)
	f.kan.SetLeft(3*left - 20)
	f.kan.Font().SetSize(font)
	f.kan.Font().SetColor(color)
	f.kan.SetTextBuf("坎1")

	f.gen = vcl.NewLabel(f)
	f.gen.SetParent(f)
	f.gen.SetTop(4*top1 + 20)
	f.gen.SetLeft(left)
	f.gen.Font().SetSize(font)
	f.gen.Font().SetColor(color)
	f.gen.SetTextBuf("艮8")

	f.zhen = vcl.NewLabel(f)
	f.zhen.SetParent(f)
	f.zhen.SetTop(2*top1 + 40)
	f.zhen.SetLeft(left)
	f.zhen.Font().SetSize(font)
	f.zhen.Font().SetColor(color)
	f.zhen.SetTextBuf("震3")

	f.xun = vcl.NewLabel(f)
	f.xun.SetParent(f)
	f.xun.SetTop(top1)
	f.xun.SetLeft(left)
	f.xun.Font().SetSize(font)
	f.xun.Font().SetColor(color)
	f.xun.SetTextBuf("巽4")

	f.li = vcl.NewLabel(f)
	f.li.SetParent(f)
	f.li.SetTop(top1)
	f.li.SetLeft(3*left - 20)
	f.li.Font().SetSize(font)
	f.li.Font().SetColor(color)
	f.li.SetTextBuf("离9")

	f.kun = vcl.NewLabel(f)
	f.kun.SetParent(f)
	f.kun.SetTop(top1)
	f.kun.SetLeft(4*left + 15)
	f.kun.Font().SetSize(font)
	f.kun.Font().SetColor(color)
	f.kun.SetTextBuf("坤2")

	f.dui = vcl.NewLabel(f)
	f.dui.SetParent(f)
	f.dui.SetTop(2*top1 + 40)
	f.dui.SetLeft(4*left + 15)
	f.dui.Font().SetSize(font)
	f.dui.Font().SetColor(color)
	f.dui.SetTextBuf("兑7")

	f.qian = vcl.NewLabel(f)
	f.qian.SetParent(f)
	f.qian.SetTop(4*top1 + 20)
	f.qian.SetLeft(4*left + 15)
	f.qian.Font().SetSize(font)
	f.qian.Font().SetColor(color)
	f.qian.SetTextBuf("乾6")

}
