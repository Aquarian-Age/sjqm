package main

import (
	"gioui.org/app"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"log"
	"os"
	"time"
)

func main() {
	w := app.NewWindow(app.Title("sjqm"),
		app.MaxSize(unit.Dp(360), unit.Dp(640)))
	go func() {
		if err := loop(w); err != nil {
			log.Println("window:", err)
		}
		os.Exit(0)
	}()
	app.Main()
}
func loop(w *app.Window) error {
	ticker := time.NewTicker(time.Second)
	var ops op.Ops
	th := chineseFont()
	for {
		select {
		case e := <-w.Events():
			switch e := e.(type) {
			case system.FrameEvent:
				gtx := layout.NewContext(&ops, e)
				//输入信息布局
				layoutEditAndBtn(th, gtx)
				//输入年月日时之后重置界面 显示九宫排盘
				//画九宫格 ...
				drawLine1(gtx)
				drawLine2(gtx)
				drawLine3(gtx)
				drawLine4(gtx)
				//显示阳历当前时间
				timeS = time.Now().Format("2006-01-02 15:04:05")
				YmdsLayout(gtx, th)
				//显示干支历
				ganZhiS = ymd(t)
				layoutGanZhiS(gtx, th)
				////////////////
				qiMenS(y, m, d, h)
				show4(th, gtx)
				show9(th, gtx)
				show2(th, gtx)
				show3(th, gtx)
				show5(th, gtx)
				show7(th, gtx)
				show8(th, gtx)
				show1(th, gtx)
				show6(th, gtx)
				//画单个八卦图标
				Kan(gtx)
				Qian(gtx)
				Dui(gtx)
				Kun(gtx)
				Li(gtx)
				XunGua(gtx)
				Zhen(gtx)
				Gen(gtx)

				if btn1.Clicked() {
					y = convEdit(ed1.Text())
					m = convEdit(ed2.Text())
					d = convEdit(ed3.Text())
					h = convEdit(ed4.Text())
					t = time.Date(y, time.Month(m), d, h, 0, 0, 0, time.Local)
					ymd(t)
					/////////////
					qiMenS(y, m, d, h)
					//drawCircle(gtx)
				}
				e.Frame(gtx.Ops)
			case system.DestroyEvent:
				return e.Err
			}
		case <-ticker.C:
			w.Invalidate()
		}
	}
}
