module liangzi.local/sjqmui

go 1.15

require (
	gioui.org v0.0.0-20210225120118-f6fba7388544
	gioui.org/x v0.0.0-20210120222453-b55819bc712b
	github.com/gonoto/notosans v0.0.0-20200703162533-d78fef05ce80
	github.com/wunderkind2k1/gorcle v0.0.0-20190729115232-76a13e23937f
	golang.org/x/sys v0.0.0-20210113181707-4bcb84eeeb78 // indirect
	liangzi.local/cal v0.0.0
	liangzi.local/sjqm v0.0.0-00010101000000-000000000000
)

replace (
	liangzi.local/cal => ../src
	liangzi.local/sjqm => ../
)
