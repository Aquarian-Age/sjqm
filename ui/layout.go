package main

import (
	"gioui.org/layout"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget/material"
	"image/color"
)

//九星八门八神颜色
var (
	jiuxingC = red   //九星 天盘奇仪 红色
	baMenC   = black //八门黑色
	baShenC  = gold  //八神金色
	diQiYiC  = green //color.NRGBA{R: 34, G: 139, B: 34, A: 255} //地盘奇仪
	gzC      = black //暗干支黑色
)

/**********************
4 9 2
3 5 7
8 1 6
***********************/
//4宫
func show4(th *material.Theme, gtx layout.Context) layout.Dimensions {
	//s4 = `九星4`
	//m4 = `八门4`
	//god4 = "八神4"
	//sqy4 = `s奇仪4`
	//eqy4 = `e奇异4`
	//gzs4 = `暗干支4`

	inset4 := layout.Inset{Top: unit.Dp(60), Left: unit.Dp(30)}
	return inset4.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
			//九星+天盘奇仪
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				star := material.H6(th, s4+" "+sqy4)
				star.TextSize = unit.Dp(14)
				//c := color.NRGBA{30, 144, 255, 255}
				star.Color = jiuxingC
				star.Font = text.Font{Typeface: "Noto"}
				return star.Layout(gtx)
			}),
			//八门
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				door4 := material.H6(th, m4)
				door4.TextSize = unit.Dp(14)
				//c := color.NRGBA{30, 144, 255, 255}
				door4.Color = baMenC
				door4.Font = text.Font{Typeface: "Noto"}
				return door4.Layout(gtx)
			}),
			//八神
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				God4 := material.H6(th, god4)
				God4.Alignment = text.Start
				God4.TextSize = unit.Dp(14)
				//c := color.NRGBA{30, 144, 255, 255}
				God4.Color = baShenC
				God4.Font = text.Font{Typeface: "Noto"}
				return God4.Layout(gtx)
			}),
			//天盘三奇六仪
			//layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			//	sky4 := material.H6(th, sqy4)
			//	sky4.Alignment = text.Start
			//	sky4.TextSize = unit.Dp(14)
			//	c := color.NRGBA{30, 144, 255, 255}
			//	sky4.Color = c
			//	sky4.Font = text.Font{Typeface: "Noto"}
			//	return sky4.Layout(gtx)
			//}),
			//地盘三奇六仪
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				e4 := material.H6(th, eqy4)
				e4.Alignment = text.Start
				e4.TextSize = unit.Dp(14)
				//c := color.NRGBA{30, 144, 255, 255}
				e4.Color = diQiYiC
				e4.Font = text.Font{Typeface: "Noto"}
				return e4.Layout(gtx)
			}),
			//暗干支
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				gz4 := material.H6(th, gzs4)
				gz4.Alignment = text.Start
				gz4.TextSize = unit.Dp(14)
				//c := color.NRGBA{30, 144, 255, 255}
				gz4.Color = gzC
				gz4.Font = text.Font{Typeface: "Noto"}
				return gz4.Layout(gtx)
			}),
			//八卦名称
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				bg4 := material.H6(th, "巽")
				bg4.Alignment = text.Start
				c := color.NRGBA{30, 144, 255, 255}
				bg4.Color = c
				bg4.Font = text.Font{Typeface: "Noto"}
				return bg4.Layout(gtx)
			}),
		)
	})

}

//9宫
func show9(th *material.Theme, gtx layout.Context) layout.Dimensions {
	//s9 = `九星9`
	//m9 = `八门9`
	//god9 = `八神9`
	//sqy9 = `s奇仪9`
	//eqy9 = `e奇异9`
	//gzs9 = `暗干支9`
	inset9 := layout.Inset{Top: unit.Dp(60), Left: unit.Dp(150)}

	return inset9.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
			//九星
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				star9 := material.H6(th, s9+" "+sqy9)
				star9.TextSize = unit.Dp(14)
				///	c := color.NRGBA{30, 144, 255, 255}
				star9.Color = jiuxingC
				star9.Font = text.Font{Typeface: "Noto"}
				return star9.Layout(gtx)
			}),
			//八门
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				door9 := material.H6(th, m9)
				door9.TextSize = unit.Dp(14)
				//c := color.NRGBA{30, 144, 255, 255}
				door9.Color = baMenC
				door9.Font = text.Font{Typeface: "Noto"}
				return door9.Layout(gtx)
			}),
			//八神
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				God9 := material.H6(th, god9)
				God9.Alignment = text.Start
				God9.TextSize = unit.Dp(14)
				//c := color.NRGBA{30, 144, 255, 255}
				God9.Color = baShenC
				God9.Font = text.Font{Typeface: "Noto"}
				return God9.Layout(gtx)
			}),
			//天盘三奇六仪
			//layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			//	sky9 := material.H6(th, sqy9)
			//	sky9.Alignment = text.Start
			//	sky9.TextSize = unit.Dp(14)
			//	c := color.NRGBA{30, 199, 255, 255}
			//	sky9.Color = c
			//	sky9.Font = text.Font{Typeface: "Noto"}
			//	return sky9.Layout(gtx)
			//}),
			//地盘三奇六仪
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				e9 := material.H6(th, eqy9)
				e9.Alignment = text.Start
				e9.TextSize = unit.Dp(14)
				//c := color.NRGBA{30, 199, 255, 255}
				e9.Color = diQiYiC
				e9.Font = text.Font{Typeface: "Noto"}
				return e9.Layout(gtx)
			}),
			//暗干支
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				gz9 := material.H6(th, gzs9)
				gz9.Alignment = text.Start
				gz9.TextSize = unit.Dp(14)
				//c := color.NRGBA{30, 199, 255, 255}
				gz9.Color = gzC
				gz9.Font = text.Font{Typeface: "Noto"}
				return gz9.Layout(gtx)
			}),
			//八卦名称
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				bg9 := material.H6(th, "离")
				bg9.Alignment = text.Start
				c := color.NRGBA{30, 144, 255, 255}
				bg9.Color = c
				bg9.Font = text.Font{Typeface: "Noto"}
				return bg9.Layout(gtx)
			}),
		)
		//star9 := material.H6(th, s9)
		//star9.TextSize = unit.Dp(14)
		//c := color.NRGBA{30, 144, 255, 255}
		//star9.Color = c
		//star9.Font = text.Font{Typeface: "Noto"}
		//return star9.Layout(gtx)
	})
}

//2宫
func show2(th *material.Theme, gtx layout.Context) layout.Dimensions {
	//s2 = `九星2`
	//m2 = `八门2`
	//god2 = `八神-2`
	//sqy2 = `s奇仪2`
	//eqy2 = `e奇异2`
	//gzs2 = `暗干支2`
	inset2 := layout.Inset{Top: unit.Dp(60), Left: unit.Dp(240)}

	return inset2.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
			//九星
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				star2 := material.H6(th, s2+" "+sqy2)
				star2.TextSize = unit.Dp(14)
				star2.Alignment = text.Middle
				//c := color.NRGBA{30, 144, 255, 255}
				star2.Color = jiuxingC
				star2.Font = text.Font{Typeface: "Noto"}
				return star2.Layout(gtx)
			}),
			//八门
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				door2 := material.H6(th, m2)
				door2.Alignment = text.Middle
				door2.TextSize = unit.Dp(14)
				//c := color.NRGBA{30, 144, 255, 255}
				door2.Color = baMenC
				door2.Font = text.Font{Typeface: "Noto"}
				return door2.Layout(gtx)
			}),
			//八神
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				God2 := material.H6(th, god2)
				God2.Alignment = text.Middle
				God2.TextSize = unit.Dp(14)
				//c := color.NRGBA{30, 144, 255, 255}
				God2.Color = baShenC
				God2.Font = text.Font{Typeface: "Noto"}
				return God2.Layout(gtx)
			}),
			//天盘三奇六仪
			//layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			//	sky2 := material.H6(th, sqy2)
			//	sky2.Alignment = text.Middle
			//	sky2.TextSize = unit.Dp(14)
			//	c := color.NRGBA{30, 122, 255, 255}
			//	sky2.Color = c
			//	sky2.Font = text.Font{Typeface: "Noto"}
			//	return sky2.Layout(gtx)
			//}),
			//地盘三奇六仪
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				e2 := material.H6(th, eqy2)
				e2.Alignment = text.Middle
				e2.TextSize = unit.Dp(14)
				//c := color.NRGBA{30, 122, 255, 255}
				e2.Color = diQiYiC
				e2.Font = text.Font{Typeface: "Noto"}
				return e2.Layout(gtx)
			}),
			//暗干支
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				gz2 := material.H6(th, gzs2)
				gz2.Alignment = text.Middle
				gz2.TextSize = unit.Dp(14)
				//	c := color.NRGBA{30, 122, 255, 255}
				gz2.Color = gzC
				gz2.Font = text.Font{Typeface: "Noto"}
				return gz2.Layout(gtx)
			}),
			//八卦名称
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				bg2 := material.H6(th, "坤")
				bg2.Alignment = text.Middle
				c := color.NRGBA{30, 144, 255, 255}
				bg2.Color = c
				bg2.Font = text.Font{Typeface: "Noto"}
				return bg2.Layout(gtx)
			}),
		)
	})
}

/////////////////////////////
//3宫
func show3(th *material.Theme, gtx layout.Context) layout.Dimensions {
	//s3 = `九星3`
	//m3 = `八门3`
	//god3 = `八神3`
	//sqy3 = `s奇仪-3`
	//eqy3 = `e奇异3`
	//gzs3 = `暗干支3`
	inset3 := layout.Inset{Top: unit.Dp(210), Left: unit.Dp(30)}

	return inset3.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
			//九星
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				star3 := material.H6(th, s3+" "+sqy3)
				star3.TextSize = unit.Dp(14)
				//c := color.NRGBA{30, 144, 255, 255}
				star3.Color = jiuxingC
				star3.Font = text.Font{Typeface: "Noto"}
				return star3.Layout(gtx)
			}),
			//八门
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				door3 := material.H6(th, m3)
				door3.TextSize = unit.Dp(14)
				//c := color.NRGBA{30, 144, 255, 255}
				door3.Color = baMenC
				door3.Font = text.Font{Typeface: "Noto"}
				return door3.Layout(gtx)
			}),
			//八神
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				God3 := material.H6(th, god3)
				God3.Alignment = text.Start
				God3.TextSize = unit.Dp(14)
				//	c := color.NRGBA{30, 144, 255, 255}
				God3.Color = baShenC
				God3.Font = text.Font{Typeface: "Noto"}
				return God3.Layout(gtx)
			}),
			//天盘三奇六仪
			//layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			//	sky3 := material.H6(th, sqy3)
			//	sky3.Alignment = text.Start
			//	sky3.TextSize = unit.Dp(14)
			//	c := color.NRGBA{30, 133, 255, 255}
			//	sky3.Color = c
			//	sky3.Font = text.Font{Typeface: "Noto"}
			//	return sky3.Layout(gtx)
			//}),
			//地盘三奇六仪
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				e3 := material.H6(th, eqy3)
				e3.Alignment = text.Start
				e3.TextSize = unit.Dp(14)
				//	c := color.NRGBA{30, 133, 255, 255}
				e3.Color = diQiYiC
				e3.Font = text.Font{Typeface: "Noto"}
				return e3.Layout(gtx)
			}),
			//暗干支
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				gz3 := material.H6(th, gzs3)
				gz3.Alignment = text.Start
				gz3.TextSize = unit.Dp(14)
				//	c := color.NRGBA{30, 133, 255, 255}
				gz3.Color = gzC
				gz3.Font = text.Font{Typeface: "Noto"}
				return gz3.Layout(gtx)
			}),
			//八卦名称
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				bg3 := material.H6(th, "震")
				bg3.Alignment = text.Start
				c := color.NRGBA{30, 144, 255, 255}
				bg3.Color = c
				bg3.Font = text.Font{Typeface: "Noto"}
				return bg3.Layout(gtx)
			}),
		)
	})
}

//5宫
func show5(th *material.Theme, gtx layout.Context) layout.Dimensions {
	inset5 := layout.Inset{Top: unit.Dp(210), Left: unit.Dp(150)}
	return inset5.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
			//九星
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				star5 := material.H6(th, s5)
				star5.TextSize = unit.Dp(14)
				//	c := color.NRGBA{30, 144, 255, 255}
				star5.Color = yellow //c
				star5.Font = text.Font{Typeface: "Noto"}
				return star5.Layout(gtx)
			}),
			//八门
			//layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			//	door5 := material.H6(th, m5)
			//	door5.Alignment = text.Middle
			//	door5.TextSize = unit.Dp(14)
			//	//c := color.NRGBA{70, 144, 255, 255}
			//	door5.Color = baMenC
			//	door5.Font = text.Font{Typeface: "Noto"}
			//	return door5.Layout(gtx)
			//}),
			//无天盘奇仪
			//地盘奇仪
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				e5 := material.H6(th, eqy5)
				e5.TextSize = unit.Dp(14)
				//c := color.NRGBA{70, 127, 255, 255}
				e5.Color = diQiYiC
				e5.Font = text.Font{Typeface: "Noto"}
				return e5.Layout(gtx)
			}),
			//暗干支
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				gz5 := material.H6(th, gzs5)
				gz5.TextSize = unit.Dp(14)
				gz5.Color = gzC
				gz5.Font = text.Font{Typeface: "Noto"}
				return gz5.Layout(gtx)
			}),
		)
	})
}

//7宫
func show7(th *material.Theme, gtx layout.Context) layout.Dimensions {
	//s7 = `九星7`
	//m7 = `八门7`
	//god7 = `八神7`
	//sqy7 = `s奇仪7`
	//eqy7 = `e奇异7`
	//gzs7 = `暗干支7`
	inset7 := layout.Inset{Top: unit.Dp(210), Left: unit.Dp(240)}

	return inset7.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
			//九星
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				star7 := material.H6(th, s7+" "+sqy7)
				star7.Alignment = text.Middle
				star7.TextSize = unit.Dp(14)
				//c := color.NRGBA{70, 144, 255, 255}
				star7.Color = jiuxingC
				star7.Font = text.Font{Typeface: "Noto"}
				return star7.Layout(gtx)
			}),
			//八门
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				door7 := material.H6(th, m7)
				door7.Alignment = text.Middle
				door7.TextSize = unit.Dp(14)
				//c := color.NRGBA{70, 144, 255, 255}
				door7.Color = baMenC
				door7.Font = text.Font{Typeface: "Noto"}
				return door7.Layout(gtx)
			}),
			//八神
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				God7 := material.H6(th, god7)
				God7.Alignment = text.Middle
				God7.TextSize = unit.Dp(14)
				//c := color.NRGBA{70, 144, 255, 255}
				God7.Color = baShenC
				God7.Font = text.Font{Typeface: "Noto"}
				return God7.Layout(gtx)
			}),
			//天盘三奇六仪
			//layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			//	sky7 := material.H6(th, sqy7)
			//	sky7.Alignment = text.Middle
			//	sky7.TextSize = unit.Dp(14)
			//	c := color.NRGBA{70, 127, 255, 255}
			//	sky7.Color = c
			//	sky7.Font = text.Font{Typeface: "Noto"}
			//	return sky7.Layout(gtx)
			//}),
			//地盘三奇六仪
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				e7 := material.H6(th, eqy7)
				e7.Alignment = text.Middle
				e7.TextSize = unit.Dp(14)
				//c := color.NRGBA{70, 127, 255, 255}
				e7.Color = diQiYiC
				e7.Font = text.Font{Typeface: "Noto"}
				return e7.Layout(gtx)
			}),
			//暗干支
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				gz7 := material.H6(th, gzs7)
				gz7.Alignment = text.Middle
				gz7.TextSize = unit.Dp(14)
				//c := color.NRGBA{70, 127, 255, 255}
				gz7.Color = gzC
				gz7.Font = text.Font{Typeface: "Noto"}
				return gz7.Layout(gtx)
			}),
			//八卦名称
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				bg7 := material.H6(th, "兑")
				bg7.Alignment = text.Middle
				c := color.NRGBA{30, 144, 255, 255}
				bg7.Color = c
				bg7.Font = text.Font{Typeface: "Noto"}
				return bg7.Layout(gtx)
			}),
		)
	})
}

//////////////////////////
//8宫
func show8(th *material.Theme, gtx layout.Context) layout.Dimensions {
	//s8 = `九星8`
	//m8 = `八门8`
	//god8 = `八神8`
	//sqy8 = `s奇仪8`
	//eqy8 = `e奇异8`
	//gzs8 = `暗干支8`
	inset8 := layout.Inset{Top: unit.Dp(360), Left: unit.Dp(30)}

	return inset8.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
			//九星
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				star8 := material.H6(th, s8+" "+sqy8)
				star8.TextSize = unit.Dp(14)
				//c := color.NRGBA{80, 144, 255, 255}
				star8.Color = jiuxingC
				star8.Font = text.Font{Typeface: "Noto"}
				return star8.Layout(gtx)
			}),
			//八门
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				door8 := material.H6(th, m8)
				door8.TextSize = unit.Dp(14)
				//c := color.NRGBA{80, 144, 255, 255}
				door8.Color = baMenC
				door8.Font = text.Font{Typeface: "Noto"}
				return door8.Layout(gtx)
			}),
			//八神
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				God8 := material.H6(th, god8)
				God8.Alignment = text.Start
				God8.TextSize = unit.Dp(14)
				//	c := color.NRGBA{80, 144, 255, 255}
				God8.Color = baShenC
				God8.Font = text.Font{Typeface: "Noto"}
				return God8.Layout(gtx)
			}),
			//天盘三奇六仪
			//layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			//	sky8 := material.H6(th, sqy8)
			//	sky8.Alignment = text.Start
			//	sky8.TextSize = unit.Dp(14)
			//	c := color.NRGBA{80, 128, 255, 255}
			//	sky8.Color = c
			//	sky8.Font = text.Font{Typeface: "Noto"}
			//	return sky8.Layout(gtx)
			//}),
			//地盘三奇六仪
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				e8 := material.H6(th, eqy8)
				e8.Alignment = text.Start
				e8.TextSize = unit.Dp(14)
				//	c := color.NRGBA{80, 128, 255, 255}
				e8.Color = diQiYiC
				e8.Font = text.Font{Typeface: "Noto"}
				return e8.Layout(gtx)
			}),
			//暗干支
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				gz8 := material.H6(th, gzs8)
				gz8.Alignment = text.Start
				gz8.TextSize = unit.Dp(14)
				//	c := color.NRGBA{80, 128, 255, 255}
				gz8.Color = gzC
				gz8.Font = text.Font{Typeface: "Noto"}
				return gz8.Layout(gtx)
			}),
			//八卦名称
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				bg8 := material.H6(th, "艮")
				bg8.Alignment = text.Start
				c := color.NRGBA{30, 144, 255, 255}
				bg8.Color = c
				bg8.Font = text.Font{Typeface: "Noto"}
				return bg8.Layout(gtx)
			}),
		)
	})
}

//1宫
func show1(th *material.Theme, gtx layout.Context) layout.Dimensions {
	//s1 = `九星1`
	//m1 = `八门1`
	//god1 = `八神1`
	//sqy1 = `s奇仪1`
	//eqy1 = `e奇异1`
	//gzs1 = `暗干支1`
	inset1 := layout.Inset{Top: unit.Dp(360), Left: unit.Dp(150)}

	return inset1.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
			//九星
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				star1 := material.H6(th, s1+" "+sqy1)
				star1.TextSize = unit.Dp(14)
				//c := color.NRGBA{10, 144, 255, 255}
				star1.Color = jiuxingC
				star1.Font = text.Font{Typeface: "Noto"}
				return star1.Layout(gtx)
			}),
			//八门
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				door1 := material.H6(th, m1)
				door1.TextSize = unit.Dp(14)
				//	c := color.NRGBA{10, 144, 255, 255}
				door1.Color = baMenC
				door1.Font = text.Font{Typeface: "Noto"}
				return door1.Layout(gtx)
			}),
			//八神
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				God1 := material.H6(th, god1)
				God1.Alignment = text.Start
				God1.TextSize = unit.Dp(14)
				//c := color.NRGBA{10, 144, 255, 255}
				God1.Color = baShenC
				God1.Font = text.Font{Typeface: "Noto"}
				return God1.Layout(gtx)
			}),
			//天盘三奇六仪
			//layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			//	sky1 := material.H6(th, sqy1)
			//	sky1.Alignment = text.Start
			//	sky1.TextSize = unit.Dp(14)
			//	c := color.NRGBA{10, 121, 255, 255}
			//	sky1.Color = c
			//	sky1.Font = text.Font{Typeface: "Noto"}
			//	return sky1.Layout(gtx)
			//}),
			//地盘三奇六仪
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				e1 := material.H6(th, eqy1)
				e1.Alignment = text.Start
				e1.TextSize = unit.Dp(14)
				//c := color.NRGBA{10, 121, 255, 255}
				e1.Color = diQiYiC
				e1.Font = text.Font{Typeface: "Noto"}
				return e1.Layout(gtx)
			}),
			//暗干支
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				gz1 := material.H6(th, gzs1)
				gz1.Alignment = text.Start
				gz1.TextSize = unit.Dp(14)
				//c := color.NRGBA{10, 121, 255, 255}
				gz1.Color = gzC
				gz1.Font = text.Font{Typeface: "Noto"}
				return gz1.Layout(gtx)
			}),
			//八卦名称
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				bg1 := material.H6(th, "坎")
				bg1.Alignment = text.Start
				c := color.NRGBA{30, 144, 255, 255}
				bg1.Color = c
				bg1.Font = text.Font{Typeface: "Noto"}
				return bg1.Layout(gtx)
			}),
		)
	})
}

//6宫
func show6(th *material.Theme, gtx layout.Context) layout.Dimensions {
	//s6 = `九星6`
	//m6 = `八门6`
	//god6 = `八神6`
	//sqy6 = `s奇仪6`
	//eqy6 = `e奇异6`
	//gzs6 = `暗干支6`
	inset6 := layout.Inset{Top: unit.Dp(360), Left: unit.Dp(240)}

	return inset6.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
			//九星
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				star6 := material.H6(th, s6+" "+sqy6)
				star6.Alignment = text.Middle
				star6.TextSize = unit.Dp(14)
				//c := color.NRGBA{60, 255, 255, 255}
				star6.Color = jiuxingC
				star6.Font = text.Font{Typeface: "Noto"}
				return star6.Layout(gtx)
			}),
			//八门
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				door6 := material.H6(th, m6)
				door6.Alignment = text.Middle
				door6.TextSize = unit.Dp(14)
				//c := color.NRGBA{60, 255, 255, 255}
				door6.Color = baMenC
				door6.Font = text.Font{Typeface: "Noto"}
				return door6.Layout(gtx)
			}),
			//八神
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				God6 := material.H6(th, god6)
				God6.Alignment = text.Middle
				God6.TextSize = unit.Dp(14)
				//c := color.NRGBA{60, 255, 255, 255}
				God6.Color = baShenC
				God6.Font = text.Font{Typeface: "Noto"}
				return God6.Layout(gtx)
			}),
			//天盘三奇六仪
			//layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			//	sky6 := material.H6(th, sqy6)
			//	sky6.Alignment = text.Middle
			//	sky6.TextSize = unit.Dp(14)
			//	c := color.NRGBA{60, 255, 255, 255}
			//	sky6.Color = c
			//	sky6.Font = text.Font{Typeface: "Noto"}
			//	return sky6.Layout(gtx)
			//}),
			//地盘三奇六仪
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				e6 := material.H6(th, eqy6)
				e6.Alignment = text.Middle
				e6.TextSize = unit.Dp(14)
				//c := color.NRGBA{60, 255, 255, 255}
				e6.Color = diQiYiC
				e6.Font = text.Font{Typeface: "Noto"}
				return e6.Layout(gtx)
			}),
			//暗干支
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				gz6 := material.H6(th, gzs6)
				gz6.Alignment = text.Middle
				gz6.TextSize = unit.Dp(14)
				//c := color.NRGBA{60, 255, 255, 255}
				gz6.Color = gzC
				gz6.Font = text.Font{Typeface: "Noto"}
				return gz6.Layout(gtx)
			}),
			//八卦名称
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				bg6 := material.H6(th, "乾")
				bg6.Alignment = text.Middle
				c := color.NRGBA{30, 144, 255, 255}
				bg6.Color = c
				bg6.Font = text.Font{Typeface: "Noto"}
				return bg6.Layout(gtx)
			}),
		)
	})
}
