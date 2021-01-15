module go-sciter

go 1.15

require (
	github.com/lxn/win v0.0.0-20201111105847-2a20daff6a55 // indirect
	github.com/sciter-sdk/go-sciter v0.5.0
	liangzi.local/auth v0.0.0
	liangzi.local/nongli v0.0.0
	liangzi.local/sjqm v0.0.0
	liangzi.local/ts v0.0.0 // indirect
)

replace (
	liangzi.local/auth => /home/xuan/src/auth/
	liangzi.local/nongli => /home/xuan/src/ccal-cli/
	liangzi.local/sjqm => /home/xuan/src/sjqm/
	liangzi.local/ts => /home/xuan/src/ts/v0.6.9/
)
