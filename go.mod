module github.com/jerryzhengj/utils

go 1.12

require (
	github.com/BurntSushi/toml v0.3.1
	github.com/fastly/go-utils v0.0.0-20180712184237-d95a45783239 // indirect
	github.com/jehiah/go-strftime v0.0.0-20171201141054-1d33003b3869 // indirect
	github.com/jonboulle/clockwork v0.1.0 // indirect
	github.com/lestrrat/go-envload v0.0.0-20180220120943-6ed08b54a570 // indirect
	github.com/lestrrat/go-file-rotatelogs v0.0.0-20180223000712-d3151e2a480f
	github.com/lestrrat/go-strftime v0.0.0-20180220042222-ba3bf9c1d042 // indirect
	github.com/natefinch/lumberjack v2.0.0+incompatible
	github.com/tarm/serial v0.0.0-20180830185346-98f6abe2eb07
	github.com/tebeka/strftime v0.1.3 // indirect
	go.uber.org/zap v1.13.0
	golang.org/x/sys v0.0.0-20190606165138-5da285871e9c // indirect
	gopkg.in/check.v1 v1.0.0-20180628173108-788fd7840127
	gopkg.in/natefinch/lumberjack.v2 v2.0.0 // indirect
)

replace (
	golang.org/x/lint => github.com/golang/lint v0.0.0-20190930215403-16217165b5de
	golang.org/x/sys => github.com/golang/sys v0.0.0-20190209173611-3b5209105503
	golang.org/x/tools => github.com/golang/tools v0.0.0-20191029190741-b9c20aec41a5
)
