package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	name string
}

func Init(cfg string)error {
	c:= Config{
		name:cfg,
	}

	if err:= c.initConfig();err!=nil {
		return  err
	}
	return nil
}

func (c *Config)initConfig()error  {
	if c.name !=""{
		viper.SetConfigFile(c.name)
	}else {
		viper.AddConfigPath("conf")
		viper.SetConfigName("config")
	}

	viper.SetConfigType("yaml")

	if err:=viper.ReadInConfig(); err!=nil{
		return err
	}

	return nil
}