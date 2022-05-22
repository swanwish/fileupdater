package settings

import (
	"github.com/spf13/viper"
	"github.com/swanwish/go-common/logs"
)

var (
	ConfigFilePath = "conf/app.ini"
)

func LoadAppSetting() error {
	viper.SetConfigName("app")
	viper.SetConfigType("toml")
	viper.AddConfigPath("./conf/")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			logs.Infof("The configuration file does not found")
		} else {
			// Config file was found but another error was produced
			logs.Errorf("Failed to read config file, the error is %v", err)
			return err
		}
	}

	logLevel := viper.GetString("log.log_level")
	if logLevel != "" {
		logs.SetLogLevel(logLevel)
	}
	return nil
}
