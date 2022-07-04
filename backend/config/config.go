package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct{
	Server 				ServerConfig
	Logger 				LoggerConfig
	LumberjackLogger 	LumberjackConfig
	DB 					DBConfig
	JWT 				JWTConfig
}

type ServerConfig struct{
	AppName		string
	AppPort		string
	Mode 		string
}

type LoggerConfig struct {
	Development 	bool
	Encoding 		string
	Level			string
}

type LumberjackConfig struct {
	MaxSize		uint32
    MaxBackups	uint16
    MaxAge		uint16
    Compress	bool
}

type JWTConfig struct {
	SecretKey	 		string
	RefreshSecretKey 	string
	Duration		 	uint32
	RefreshDuration		uint32
}

type DBConfig struct {
	Host string
	Port string
	User string
	Pass string
	Name string

}

func (c *Config) LoadConfig(configName string){
	v := viper.New()
	v.SetConfigName(configName)
	v.AddConfigPath("./config")
	v.AutomaticEnv()
	err := v.ReadInConfig()
	if err != nil {
		log.Fatalln("[Viper Config]", "Failed to load config.", err)
	}
	c.parseConfig(v)
}

func (c *Config) parseConfig(v *viper.Viper){
	err := v.Unmarshal(c)
	if err != nil {
		log.Fatalln("[Viper Config]", "Failed to load config", err)
	}
}