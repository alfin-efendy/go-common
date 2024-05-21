package configs

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/alfin87aa/go-common/constants/char"
	"github.com/alfin87aa/go-common/constants/integer"
	"github.com/spf13/viper"
)

type Config struct {
	App    App      `mapstructure:"app"`
	Log    Log      `mapstructure:"log"`
	Server Server   `mapstructure:"server"`
	DB     Database `mapstructure:"database"`
	Otel   Otel     `mapstructure:"otel"`
}

var (
	Configs *Config
	raw     map[string]interface{}
)

func Load() {
	// config file must be named app.yaml and placed in the root of the project
	viper.SetConfigName("app")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	// read the config file
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	// unmarshal the config file into the Config struct
	config := &Config{}
	if err := viper.Unmarshal(config); err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	Configs = config

	// store the raw config for later use
	raw = viper.AllSettings()
}

func getVal(key string, config map[string]interface{}) interface{} {
	if key == char.Empty {
		return nil
	}

	// split the key by dot
	keys := strings.SplitN(key, char.Dot, 2)

	// if the key is not nested
	if v, ok := config[keys[integer.Zero]]; ok {
		switch v := v.(type) {
		// if the value is a map, then it's nested
		case map[string]interface{}:
			return getVal(keys[1], v)
		default:
			return v
		}
	}
	return nil
}

// GetString use dot to get value from nested key
// ex: sql.host
func GetValue(key string) (string, error) {
	value := getVal(key, raw)
	if value == nil {
		return char.Empty, nil
	}

	switch v := value.(type) {
	case string:
		return v, nil
	case int:
		return strconv.Itoa(v), nil
	case bool:
		return strconv.FormatBool(v), nil
	default:
		return fmt.Sprintf("%v", v), nil
	}
}
