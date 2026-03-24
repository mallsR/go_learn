module go_learn

go 1.26.1

require (
	go_learn/greetings v0.0.0-00010101000000-000000000000
	rsc.io/quote v1.5.2
)

require (
	github.com/duke-git/lancet/v2 v2.3.9
	golang.org/x/text v0.9.0 // indirect
	rsc.io/sampler v1.3.0 // indirect
)

replace go_learn/greetings => ./greetings // 由于go_learn/greetings模块尚
