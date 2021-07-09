module liangzi.local/sjqmui

go 1.16

require (
	gioui.org v0.0.0-20210518185901-8611894b4bb3
	gioui.org/cmd v0.0.0-20210707135100-0e6093585688 // indirect
	github.com/Aquarian-Age/sjqm v0.0.0
	github.com/gonoto/notosans v0.0.0-20200703162533-d78fef05ce80
	github.com/wunderkind2k1/gorcle v0.0.0-20190729115232-76a13e23937f
	golang.org/x/exp v0.0.0-20201229011636-eab1b5eb1a03 // indirect
	golang.org/x/image v0.0.0-20200927104501-e162460cd6b5 // indirect
	golang.org/x/text v0.3.4 // indirect
	liangzi.local/cal v0.0.0
)

replace (
	github.com/Aquarian-Age/sjqm => ../
	liangzi.local/cal => ../cal/
)
