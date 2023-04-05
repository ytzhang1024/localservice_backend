package conf

import (
	"fmt"
	"github.com/spf13/viper"
)

// InitConfig
//
//	@Description: Handle the system configuration by using viper
func InitConfig() {
	viper.SetConfigName("settings")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./conf/")
	err := viper.ReadInConfig()

	if err != nil {
		panic(fmt.Sprintf("Load Config Error: %s", err.Error()))
	}

	fmt.Println(viper.GetString("server.port"))
}
