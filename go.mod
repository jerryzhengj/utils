module github.com/jerryzhengj/utils

go 1.12

require (
	github.com/BurntSushi/toml v0.3.1
	github.com/lestrrat/go-file-rotatelogs v0.0.0-20180223000712-d3151e2a480f
	github.com/nacos-group/nacos-sdk-go v1.0.9
	github.com/natefinch/lumberjack v2.0.0+incompatible
	github.com/tarm/serial v0.0.0-20180830185346-98f6abe2eb07
	go.uber.org/zap v1.15.0
	golang.org/x/sys v0.0.0-20190606165138-5da285871e9c // indirect
	gopkg.in/check.v1 v1.0.0-20180628173108-788fd7840127
	gopkg.in/natefinch/lumberjack.v2 v2.0.0 // indirect
)

replace (
	golang.org/x/lint => github.com/golang/lint v0.0.0-20190930215403-16217165b5de
	golang.org/x/sys => github.com/golang/sys v0.0.0-20190209173611-3b5209105503
	golang.org/x/tools => github.com/golang/tools v0.0.0-20191029190741-b9c20aec41a5
)
