package main

import (
	"os"
	"time"

	"github.com/hinha/reavers/config"
	"github.com/hinha/reavers/provider/command"
	"github.com/hinha/reavers/provider/web"
	"github.com/spf13/viper"
)

func init() {
	_ = os.Setenv("TZ", "Asia/Jakarta")
	loc, _ := time.LoadLocation(os.Getenv("TZ"))
	time.Local = loc
}

func main() {
	// Set the file name of the configurations file
	viper.SetConfigName("config")

	// Set the path to look for the configurations file
	viper.AddConfigPath(".")

	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()
	viper.SetConfigType("yml")
	var cfg config.Configurations

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	// Set undefined variables
	viper.SetDefault("database.dbname", "test_db")
	err := viper.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}

	cmd := command.Fabricate()

	webEngine := web.Fabricate(cfg.Server.Port)
	webEngine.FabricateCommand(cmd)

	if err := cmd.Execute(); err != nil {
		panic(err)
	}
}
