package conf

import (
	"errors"
	"github.com/BurntSushi/toml"
	log "github.com/jerryzhengj/utils/log/zap"
	"os"
)

type configEnv struct{
	configFile string
	configModule interface{}
	metaData toml.MetaData
	isLoad bool
}


func New(configFile string ,configModule interface{}) *configEnv{
	return &configEnv{
		configFile: configFile,
		configModule: configModule,
		isLoad: false,
	}
}

func (conf *configEnv)Load() (interface{},error){

	if m, err := toml.DecodeFile(conf.configFile, conf.configModule); err == nil {
		log.Infof("Load config[%s] success",conf.configFile)
		conf.metaData = m
		conf.isLoad = true

		return conf.configModule,nil
	}else{
		return conf.configModule,err
	}
}

func (conf *configEnv)Refresh() error{
     _ ,err :=conf.Load()
     return err
}

func (conf *configEnv)GetCurrentConfig()(interface{},error){
	if !conf.isLoad{
		return conf.Load()
	}else{
		return conf.configModule,nil
	}
}

func (conf *configEnv)Save()error{
	if !conf.isLoad{
		return errors.New("config isn't loaded")
	}

	f, err := os.OpenFile(conf.configFile, os.O_RDWR | os.O_CREATE|os.O_TRUNC, 0666)
	if err == nil {
		defer f.Close()
		return toml.NewEncoder(f).Encode(conf.configModule)
	}else{
		return errors.New("write file["+conf.configFile+"] failed:"+err.Error())
	}
}