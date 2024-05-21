package configs

type Database struct {
	Sql           map[string]*Sql   `mapstructure:"sql"`
	Redis         map[string]*Redis `mapstructure:"redis"`
	ElasticSearch *ElasticSearch    `mapstructure:"elasticSearch"`
	Mongo         *Mongo            `mapstructure:"mongo"`
}

type Sql struct {
	Driver            string             `mapstructure:"driver"`
	Host              string             `mapstructure:"host"`
	Port              int                `mapstructure:"port"`
	Database          string             `mapstructure:"database"`
	Username          string             `mapstructure:"username"`
	Password          string             `mapstructure:"password"`
	PoolingConnection *PoolingConnection `mapstructure:"poolingConnection"`
}

type PoolingConnection struct {
	MaxIdle     int   `mapstructure:"maxIdle"`
	MaxOpen     int   `mapstructure:"maxOpen"`
	MaxLifetime int64 `mapstructure:"maxLifetime"`
}

type Redis struct {
	Mode string `mapstructure:"mode"`
	RedisCluster
}

type RedisSingle struct {
	Address         string  `mapstructure:"address"`
	Username        *string `mapstructure:"username"`
	Password        *string `mapstructure:"password"`
	DB              *int    `mapstructure:"db"`
	Network         *string `mapstructure:"network"`
	MaxRetries      *int    `mapstructure:"maxRetries"`
	MaxRetryBackoff *int    `mapstructure:"maxRetryBackoff"`
	MinRetryBackoff *int    `mapstructure:"minRetryBackoff"`
	DialTimeout     *int    `mapstructure:"dialTimeout"`
	ReadTimeout     *int    `mapstructure:"readTimeout"`
	WriteTimeout    *int    `mapstructure:"writeTimeout"`
	PoolFIFO        *bool   `mapstructure:"poolFIFO"`
	PoolSize        *int    `mapstructure:"poolSize"`
	PoolTimeout     *int    `mapstructure:"poolTimeout"`
	MinIdleConns    *int    `mapstructure:"minIdleConns"`
	MaxIdleConns    *int    `mapstructure:"maxIdleConns"`
}

type RedisCluster struct {
	RedisSingle
	SentinelAddress         []string `mapstructure:"sentinelAddress"`
	MasterName              string   `mapstructure:"masterName"`
	RouteByLatency          *bool    `mapstructure:"routeByLatency"`
	RouteRandomly           *bool    `mapstructure:"routeRandomly"`
	ReplicaOnly             *bool    `mapstructure:"replicaOnly"`
	UseDisconnectedReplicas *bool    `mapstructure:"useDisconnectedReplicas"`
}

type ElasticSearch struct {
	Address  []string `mapstructure:"address"`
	Username string   `mapstructure:"username"`
	Password string   `mapstructure:"password"`
}

type Mongo struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	DataBase string `mapstructure:"database"`
}
