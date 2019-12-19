package zap


const(
	 defaultEnableConsole = false
	 defaultDailyBackup = false
	 deafultMaxSize = 100
	 defaultMaxBackups = 7
	 defaultMaxAge = 7
)


type ZaplogConf struct{
	EnableConsole bool `toml:"enableConsole"`
	Filename string `toml:"filename"`
	LogLevel string `toml:"logLevel"`
	DailyBackup bool `toml:"dailyBackup"`
	// unit:MB
	MaxSize int `toml:"maxSize"`
	MaxBackups int `toml:"maxBackups"`
	// unit:day
	MaxAge int `toml:"maxAge"`
}
