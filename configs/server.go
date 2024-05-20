package configs

type Server struct {
	RestAPI *RestAPI `mapstructure:"restAPI"`
}

type RestAPI struct {
	Port   int   `mapstructure:"port"`
	Stdout bool  `mapstructure:"stdout"`
	Cors   *Cors `mapstructure:"cors"`
	Enable bool  `mapstructure:"enable"`
}

type Cors struct {
	AllowOrigins     []string `mapstructure:"allowOrigins"`
	AllowMethods     []string `mapstructure:"allowMethods"`
	AllowHeaders     []string `mapstructure:"allowHeaders"`
	AllowCredentials bool     `mapstructure:"allowCredentials"`
	ExposeHeaders    []string `mapstructure:"exposeHeaders"`
	MaxAge           int      `mapstructure:"maxAge"`
}
