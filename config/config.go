package config

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

func Initialize(filename string) error {
	splits := strings.Split(filepath.Base(filename), ".")

	viper.SetConfigName(filepath.Base(splits[0]))
	viper.AddConfigPath(filepath.Dir(filename))
	viper.SetConfigType(splits[1])

	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	return nil
}

func isSet(key string) {
	if !viper.IsSet(key) {
		log.Fatalf("Configuration key %s not found", key)
		os.Exit(1)
	}
}

func GetString(key string) string {
	isSet(key)
	return viper.GetString(key)
}

func GetInt(key string) int {
	isSet(key)
	return viper.GetInt(key)
}

func GetBool(key string) bool {
	isSet(key)
	return viper.GetBool(key)
}

func ServerMode() string {
	return GetString("server.mode")
}
