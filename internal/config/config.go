package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// Config contains all the configuration parameters
// known by iot-cloud-cli
type Config struct {
	// Client ID of the user
	Client string `yaml:"client"`
	// Secret ID of the user, unique for each Client ID
	Secret string `yaml:"secret"`
}

// Retrieve returns the actual parameters contained in the
// configuration file, if any. Returns error if no config file is found.
func Retrieve() (*Config, error) {
	conf := &Config{}
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath(".")
	err := v.ReadInConfig()
	if err != nil {
		err = fmt.Errorf("%s: %w", "retrieving config file", err)
		return nil, err
	}

	v.Unmarshal(conf)
	return conf, nil
}