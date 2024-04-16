package config

import (
	"fmt"
	"navidmafi/walli/logger"
	"navidmafi/walli/providers"
	"os"

	"github.com/spf13/viper"
)

func Init() {

	viper.SetConfigType("yaml")
	// viper.SetDefault("backend", backends.Swww)

	viper.SetDefault("provider", providers.Unsplash)
	viper.AddConfigPath("$HOME/.config/walli")
	viper.SetConfigName("walli")

	err := viper.ReadInConfig()
	if err != nil {
		logger.Logger.Debug(err)
		mkdirErr := os.Mkdir(fmt.Sprintf("%s/.config/walli", os.Getenv("HOME")), os.ModePerm)
		logger.Logger.Debug(mkdirErr)
		viper.SafeWriteConfig()
	}
}

func SetDefaultProvider(provider string) {
	viper.Set("provider", provider)
	viper.WriteConfig()
}

func GetDefaultProvider() string {
	return viper.GetString("provider")
}
func GetDefaultBackend() string {
	return viper.GetString("backend")
}

func SetDefaultBackend(backend string) {
	viper.Set("backend", backend)
	err := viper.WriteConfig()
	if err != nil {
		logger.Logger.Error(err)
	}
}

func SetSecret(provider string, secret string) {
	viper.Set(fmt.Sprintf("%s.secret", provider), secret)
	viper.WriteConfig()
}

func GetSecret(provider string) string {
	secret := viper.GetString(fmt.Sprintf("%s.secret", provider))

	return secret
}

// func setAPIKey()
