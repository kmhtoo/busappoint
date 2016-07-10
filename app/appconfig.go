package app

import (
	"bytes"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"io/ioutil"
	"log"
	"os"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	section := os.Getenv("ENV_MODE")
	configPath := "./conf/config.json"
	if section != "prod" {
		configPath = "./conf/local/config.json"
	}

	viper.SetConfigType("json")
	b, err := ioutil.ReadFile(configPath)
	if err != nil {
		panic(err)
	}
	err = viper.ReadConfig(bytes.NewBuffer(b))
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}

func ParamString(param string) string {
	return viper.GetString(param)
}

func ParamInt(param string) int {
	return viper.GetInt(param)
}

func ParamBool(param string) bool {
	return viper.GetBool(param)
}

func ParamDouelb(param string) float64 {
	return viper.GetFloat64(param)
}

func Params(param string) map[string]interface{} {
	return viper.GetStringMap(param)
}
