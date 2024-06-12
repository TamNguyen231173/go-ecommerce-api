package main

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port int `mapstructure:"port"`
		Host string `mapstructure:"host"`
	} `mapstructure:"server"`
	Databases []struct {
		User  string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		Host string `mapstructure:"host"`
	} `mapstructure:"databases"`
}

func main()  {
	v := viper.New()
	v.AddConfigPath("config/")
	v.SetConfigName("local")
	v.SetConfigType("yaml")

	// read configuration
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	// read server configuration
	fmt.Println("Server Port::", v.GetInt("server.port"))
	fmt.Println("Server Host::", v.GetString("server.host"))
	fmt.Println("Server Security::", v.GetString("security.jwt.key"))

	// configure struct
	var config Config
	if err := v.Unmarshal(&config); err != nil {
		fmt.Printf("Unable to decode configuration, %v\n", err)
	}

	fmt.Println("Server Port::", config.Server.Port)

	for _, db := range config.Databases {
		fmt.Println("Database User::", db.User)
		fmt.Println("Database Password::", db.Password)
		fmt.Println("Database Host::", db.Host)
	}
}
