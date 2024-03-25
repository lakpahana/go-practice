module example/hello

go 1.22.1

require rsc.io/quote v1.5.2

require example.com/greetings v1.1.0

require (
	golang.org/x/text v0.14.0 // indirect
	rsc.io/sampler v1.99.99 // indirect
)

replace example.com/greetings => ../greetings
