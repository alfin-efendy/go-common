package configs

type Log struct {
	Level      string  `mapstructure:"level"`
	Location   *string `mapstructure:"location"`
	MaxSize    int     `mapstructure:"maxSize"`
	MaxAge     int     `mapstructure:"maxAge"`
	MaxBackups int     `mapstructure:"maxBackups"`
	TimeZone   string  `mapstructure:"timeZone"`
	Compress   bool    `mapstructure:"compress"`
}
