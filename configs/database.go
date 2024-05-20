package configs

type Database struct {
	Driver            string            `mapstructure:"driver"`
	Host              string            `mapstructure:"host"`
	Port              int               `mapstructure:"port"`
	Name              string            `mapstructure:"name"`
	Username          string            `mapstructure:"username"`
	Password          string            `mapstructure:"password"`
	PoolingConnection PoolingConnection `mapstructure:"poolingConnection"`
}

type PoolingConnection struct {
	MaxIdle     int   `mapstructure:"maxIdle"`
	MaxOpen     int   `mapstructure:"maxOpen"`
	MaxLifetime int64 `mapstructure:"maxLifetime"`
}
