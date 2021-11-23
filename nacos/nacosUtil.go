package nacos

import(
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

type ServerCfg struct {
	IpAddr string
	Port uint64
	Group string
	DataId string
}

type ClientCfg struct {
	NamespaceId string
	TimeoutMs uint64
	NotLoadCacheAtStart bool
	LogDir string
	CacheDir string
	RotateTime string
	MaxAge  int64
	LogLevel string
}


func InitConfig(scfg ServerCfg,ccfg ClientCfg,listener func(namespace, group, dataId, data string)) string{
	sc := []constant.ServerConfig{
		{
			IpAddr: scfg.IpAddr,
			Port:   scfg.Port,
		},
	}

	cc := constant.ClientConfig{
		NamespaceId:         ccfg.NamespaceId,
		TimeoutMs:           ccfg.TimeoutMs,
		NotLoadCacheAtStart: ccfg.NotLoadCacheAtStart,
		LogDir:              ccfg.LogDir,
		CacheDir:            ccfg.CacheDir,
		RotateTime:          ccfg.RotateTime,
		MaxAge:              ccfg.MaxAge,
		LogLevel:            ccfg.LogLevel,
	}

	configClient, err := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": sc,
		"clientConfig":  cc,
	})
	if err != nil {
		panic(err)
	}

	content, err := configClient.GetConfig(vo.ConfigParam{DataId: scfg.DataId, Group:  scfg.Group})

	if err != nil {
		panic(err)
	}

	err = configClient.ListenConfig(vo.ConfigParam{
		DataId: scfg.DataId,
		Group:  scfg.Group,
		OnChange: listener,
	})

	return content
}