package main

import (
	"strconv"
	"time"

	"github.com/Aquarian-Age/sjqm/v2"
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types/colors"
)

type TMainForm struct {
	*vcl.TForm
	year                                    *vcl.TEdit
	month, day, hour                        *vcl.TSpinEdit
	ok                                      *vcl.TButton
	ymds                                    *vcl.TLabel //基本纪年信息
	label1                                  *vcl.TLabel //排盘信息
	kan, gen, zhen, xun, li, kun, dui, qian *vcl.TLabel //12宫卦名
	buffBmp                                 *vcl.TBitmap
	qm                                      *sjqm.SJQM
	g1, g8, g3, g4, g9, g2, g7, g6, g5      *vcl.TLabel //九宫信息
}

var (
	mainForm *TMainForm
	t        = time.Now().Local()
)

func main() {
	vcl.RunApp(&mainForm)
}
func (f *TMainForm) OnFormCreate(sender vcl.IObject) {
	const (
		width  = 500
		height = 800
	)
	f.SetWidth(width)
	f.SetHeight(height)
	f.SetCaption("时家奇门")
	f.Font().SetSize(13)
	//
	f.draw()
	f.OnFormCreatG12()
	//
	f.OnFormGet(sender)
	//
	f.OnFormDefault(sender)
}

//获取输入信息
func (f *TMainForm) OnFormGet(sender vcl.IObject) {
	const height = 800
	const step = 33
	f.year = vcl.NewEdit(f)
	f.year.SetParent(f)
	f.year.SetTop(height - step)
	f.year.SetWidth(2 * step)
	f.year.Font().SetSize(11)
	f.year.SetTextBuf(strconv.Itoa(t.Year()))

	f.month = vcl.NewSpinEdit(f)
	f.month.SetParent(f)
	f.month.SetTop(height - step)
	f.month.SetLeft(2 * step)
	f.month.SetWidth(2 * step)
	f.month.Font().SetSize(11)
	f.month.SetMinValue(1)
	f.month.SetMaxValue(12)

	f.day = vcl.NewSpinEdit(f)
	f.day.SetParent(f)
	f.day.SetTop(height - step)
	f.day.SetLeft(4 * step)
	f.day.SetWidth(2 * step)
	f.day.Font().SetSize(11)
	f.day.SetMinValue(1)
	f.day.SetMaxValue(31)

	f.hour = vcl.NewSpinEdit(f)
	f.hour.SetParent(f)
	f.hour.SetTop(height - step)
	f.hour.SetLeft(6 * step)
	f.hour.SetWidth(2 * step)
	f.hour.Font().SetSize(11)
	f.hour.SetMinValue(0)
	f.hour.SetMaxValue(24)

	f.ok = vcl.NewButton(f)
	f.ok.SetParent(f)
	f.ok.SetTop(height - step)
	f.ok.SetLeft(8 * step)
	f.ok.SetWidth(step)
	f.ok.SetHeight(step)
	f.ok.Font().SetSize(13)
	f.ok.Font().SetColor(colors.ClBlack)
	f.ok.SetCaption("ok")
	f.ok.SetOnClick(f.OnFormBtnOK)
	//

}

func (f *TMainForm) OnFormBtnOK(sender vcl.IObject) {
	f.OnFormSpecifiedTime(sender)
}
