package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Debug        	bool
	Port         	int    `required:"true"`
	ChannelAccessToken		string `envconfig:"channel_access_token" require:"true"`
	ChannelID	 	string `envconfig:"channel_id" require:"true"`
	ChannelSecret 	string	`envconfig:"channel_secret" require:"true"`
	LineTokenAPI 	string `envconfig:"line_token_api" require:"true"`
	DBName			string	`envconfig:"db_name" require:"true"`
	DBUserName		string	`envconfig:"db_user_name" require:"true"`
	DBPassword		string	`envconfig:"db_password" require:"true"`
}

func Read() (Config, error) {
	var cfg Config
	err := envconfig.Process("JENY", &cfg)
	return cfg, err
}
