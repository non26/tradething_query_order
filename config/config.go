package config

import "github.com/spf13/viper"

type Config struct {
	Bn            *BN            `json:"bn" mapstructure:"bn"`
	BnCredentials *BNCredentials `json:"bn_credentials" mapstructure:"bn_credentials"`
}

type BN struct {
	WsURL  string `json:"ws_url" mapstructure:"ws_url"`
	Method Method `json:"method" mapstructure:"method"`
}

type Method struct {
	PositionInformation string `json:"position_information" mapstructure:"position_information"`
}

type BNCredentials struct {
	APIKey     string `json:"api_key" mapstructure:"api_key"`
	PrivateKey string `json:"private_key" mapstructure:"private_key"`
}

func ReadConfig() (c *Config, err error) {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath("./config")
	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	viper.Unmarshal(&c)
	return c, nil
}
