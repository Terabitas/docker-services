package config

import (
	"strings"

	"bufio"
	"os"

	"github.com/spf13/viper"
)

type (
	// Config type
	Config struct {
		Verbosity int      `yaml:"verbosity"`
		Dir       string   `yaml:"dir"`
		Host      string   `yaml:"host"`
		Services  Services `yaml:"services"`
	}

	// Service ...
	Service struct {
		Branch          string  `yaml:"branch"`
		Host            *string `yaml:"host,omitempty"`
		Env             *string `yaml:"env,omitempty"`
		ComposeFilePath *string `yaml:"compose_path,omitempty"`
	}

	// Name ...
	Name string

	// Services ...
	Services map[Name]Service
)

const (
	DefaultConfigFileName = "docker-services"
)

// LoadConfig ...
func LoadConfig(pathToCfg string) error {
	cfg := viper.New()
	if pathToCfg == "" {
		cfg.SetConfigType("yaml")
		cfg.SetConfigName(DefaultConfigFileName)
		cfg.AddConfigPath(".")
		err := cfg.ReadInConfig()
		if err != nil {
			return err
		}
	} else {
		f, err := os.Open(pathToCfg)
		if err != nil {
			return err
		}

		err = cfg.ReadConfig(bufio.NewReader(f))
		if err != nil {
			return err
		}
	}

	var c Config
	cfg.Unmarshal(&c)

	return nil
}

// StringToSlice return slice from "x,y,z"
func StringToSlice(in string) []string {
	rez := []string{}
	for _, val := range strings.Split(in, ",") {
		rez = append(rez, val)
	}

	return rez
}
