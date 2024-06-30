package initialize

import (
	"fmt"
	"github.com/spf13/viper"
	"go-ecommerce-backend-api/global"
)

func LoadConfig()  {
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
	if err := v.Unmarshal(&global.Config); err != nil {
		fmt.Printf("Unable to decode configuration, %v\n", err)
	}
}
