package nacos

import (
	"fmt"
	"testing"
)


func TestGetConfig(t *testing.T) {
	sc := ServerCfg{
			IpAddr: "http://xxx.com",
			Port:   80,
			Group: "tsp-mock",
			DataId: "main-config.yaml",
	}

	cc := ClientCfg{
		NamespaceId:         "tsp-mock",
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "tmp/nacos/log",
		CacheDir:            "tmp/nacos/cache",
		RotateTime:          "2h",
		MaxAge:              3,
		LogLevel:            "info",
	}

	content :=InitConfig(sc,cc,func(namespace, group, dataId, data string) {
		fmt.Println("receive config change")
		fmt.Println("group:" + group + ", dataId:" + dataId + ", data:" + data)
	})

	fmt.Println(content)
}
