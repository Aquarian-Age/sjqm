package main

import (
	"gioui.org/layout"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"image/color"
)

var (
	ed1, ed2, ed3, ed4 widget.Editor
	btn1               widget.Clickable
)

//干支布局
func layoutGanZhiS(gtx layout.Context, th *material.Theme) layout.Dimensions {
	return layout.S.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		gzs := material.H6(th, ganZhiS+jieQiS+zhiFS)
		gzs.TextSize = unit.Dp(14)
		gzs.Font = text.Font{Typeface: "Noto"}
		gzs.Color = black
		return gzs.Layout(gtx)
	})
}

//当前时间
func YmdsLayout(gtx layout.Context, th *material.Theme) layout.Dimensions {
	body := material.Body1(th, timeS)
	//body.Font = text.Font{Typeface: "Noto"}
	body.Alignment = text.Middle
	body.TextSize = unit.Dp(18)
	return body.Layout(gtx)
}

//
func layoutEditAndBtn(th *material.Theme, gtx layout.Context) layout.Dimensions {

	borderWidth := float32(0.5)
	borderColor := color.NRGBA{A: 107}
	wb := widget.Border{
		Color:        borderColor,
		CornerRadius: unit.Dp(2),
		Width:        unit.Dp(borderWidth),
	}
	inset := layout.UniformInset(unit.Dp(3))
	inset.Top = unit.Dp(26)
	return layout.N.Layout(gtx, func(gtx layout.Context) layout.Dimensions { //居中对齐
		return inset.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{}.Layout(gtx,
				//
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return wb.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
						return layout.UniformInset(unit.Dp(9)).Layout(gtx,
							material.Editor(th, &ed1, "year").Layout)
					})
				}),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return wb.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
						return layout.UniformInset(unit.Dp(9)).Layout(gtx,
							material.Editor(th, &ed2, "month").Layout)
					})
				}),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return wb.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
						return layout.UniformInset(unit.Dp(9)).Layout(gtx,
							material.Editor(th, &ed3, "day").Layout)
					})
				}),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return wb.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
						return layout.UniformInset(unit.Dp(9)).Layout(gtx,
							material.Editor(th, &ed4, "hour").Layout)
					})
				}),
				//
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return material.Button(th, &btn1, "show").Layout(gtx)
				}),
			)
		})
	})
}
