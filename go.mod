module github.com/Aquarian-Age/sjqm

go 1.15

require liangzi.local/cal v0.0.0

replace (
liangzi.local/cal => ./cal/
github.com/Aquarian-Age/sjqm => ./
)
