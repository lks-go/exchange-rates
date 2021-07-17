package app

import (
	"fmt"

	"github.com/lks-go/exchange-rates/pkg/exchangeratesapi"
	"github.com/spf13/viper"
)

type Config struct {
	API exchangeratesapi.Config `mapstructure:"exchangeratesapi"`
}

func Configure(cfg *Config) error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	err = viper.Unmarshal(cfg)
	if err != nil {
		fmt.Errorf("unable to decode into struct, %v", err)
	}

	return nil
}
