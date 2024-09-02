package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Directory struct {
		Root         string
		Kernelsource string
		Kernelbuild  string
		Clang        string
	}
	Url struct {
		Source     string
		Toolchain  string
		Anykernel3 string
	}
	Buildenv string
}

// Read reads the config file.
func (c Config) Read() error {
	cfgFile := "config.yaml"
	viper.SetConfigFile(cfgFile)
	return nil
}

// Validate checks that read config is valid.
func (c Config) Validate() bool {
	return true
}
