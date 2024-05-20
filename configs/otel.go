package configs

type Otel struct {
	Trace  *OtelTrace  `mapstructure:"trace"`
	Metric *OtelMetric `mapstructure:"metric"`
	Enable bool        `mapstructure:"enable"`
}

type OtelTrace struct {
	Exporters *OtelExporters `mapstructure:"exporters"`
}

type OtelMetric struct {
	InstrumentationName string         `mapstructure:"instrumentationName"`
	Exporters           *OtelExporters `mapstructure:"exporters"`
}

type OtelExporters struct {
	Otlp   *OtelExportersOtlp `mapstructure:"otlp"`
	Enable bool               `mapstructure:"enable"`
}

type OtelExportersOtlp struct {
	Address                     string `mapstructure:"address"`
	Timeout                     int    `mapstructure:"timeout"`
	ClientMaxReceiveMessageSize string `mapstructure:"clientMaxReceiveMessageSize"`
}
