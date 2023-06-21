package config

import (
	"strings"

	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

// Load default config file from "default.yaml" and override it with environment variables
func Load(cfgFile string) (config Config, err error) {
	v := viper.New()
	v.SetConfigFile(cfgFile)
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()
	if err = v.ReadInConfig(); err != nil {
		err = errors.Wrapf(err, "failed to read config file %q", cfgFile)
		return
	}
	if err = v.Unmarshal(&config); err != nil {
		err = errors.Wrapf(err, "failed to unmarshal config file")
		return
	}
	return
}
