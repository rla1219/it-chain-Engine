package common

import   (
	"github.com/spf13/viper"
	"fmt"
)

func init(){
	initConfig()
}

func initConfig(){

	viper.SetConfigType("yaml")
	viper.SetConfigName("config")
	viper.AddConfigPath("$GOPATH/src/github.com/it-chain/it-chain-Engine/conf")

	//todo 데모용
	//viper.AddConfigPath("./conf")

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}