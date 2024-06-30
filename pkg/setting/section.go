package setting

type Config struct {
	Server struct {
		Port int    `mapstructure:"port"`
		Host string `mapstructure:"host"`
	} `mapstructure:"server"`
	MySQLSetting MySQLSetting `mapstructure:"mysql"`
	LogSetting   LogSetting   `mapstructure:"logger"`
}

type MySQLSetting struct {
	Host            string `mapstructure:"host"`
	Port            int    `mapstructure:"port"`
	User            string `mapstructure:"user"`
	Password        string `mapstructure:"password"`
	Dbname          string `mapstructure:"dbname"`
	MaxIdleConns    int    `mapstructure:"maxIdleConns"`
	MaxOpenConns    int    `mapstructure:"maxOpenConns"`
	ConnMaxLifeTime int    `mapstructure:"connMaxLifeTime"`
}

type LogSetting struct {
	LogLevel    string `mapstructure:"level"`
	LogFileName string `mapstructure:"file_log_name"`
	MaxSize     int    `mapstructure:"max_size"`
	MaxBackups  int    `mapstructure:"max_backups"`
	MaxAge      int    `mapstructure:"max_age"`
	Compress    bool   `mapstructure:"compress"`
}
