package config

import "github.com/spf13/viper"

type Config struct {
	Bn            *BN            `json:"bn" mapstructure:"bn"`
	BnCredentials *BNCredentials `json:"bn_credentials" mapstructure:"bn_credentials"`
	Dynamodb      *Dynamodb      `json:"dynamodb" mapstructure:"dynamodb"`
}

type BN struct {
	BaseURL  string   `json:"base_url" mapstructure:"base_url"`
	EndPoint EndPoint `json:"end_point" mapstructure:"end_point"`
}

type EndPoint struct {
	PositionInformation string `json:"position_information" mapstructure:"position_information"`
}

type BNCredentials struct {
	APIKey    string `json:"api_key" mapstructure:"api_key"`
	SecretKey string `json:"secret_key" mapstructure:"secret_key"`
}

type Dynamodb struct {
	Region   string `mapstructure:"region" json:"region"`
	Ak       string `mapstructure:"ak" json:"ak"`
	Sk       string `mapstructure:"sk" json:"sk"`
	Endpoint string `mapstructure:"endpoint" json:"endpoint"`
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
