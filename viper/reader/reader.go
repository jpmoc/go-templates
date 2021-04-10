package main

import (
	"fmt"

	"github.com/spf13/viper"
)

func main() {
	fmt.Println("vim-go")

	viper.SetConfigName("config")         // name of config file (without extension)
	viper.SetConfigType("yaml")           // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("/etc/appname/")  // path to look for the config file in
	viper.AddConfigPath("$HOME/.appname") // call multiple times to add many search paths
	viper.AddConfigPath(".")              // optionally look for config in the working directory
	err := viper.ReadInConfig()           // Find and read the config file
	if err != nil {                       // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	fmt.Println(viper.Get("app"))
	fmt.Println(viper.Get("app.function.name"))
	fmt.Println(viper.Get("app.function.build"))
	fmt.Println(viper.GetString("app.function.build"))
	// fmt.Println(viper.GetMap("app.function.build"))
}
