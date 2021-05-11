package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"io"
	"io/ioutil"
	"log"
)

const configType = "yaml"

func Init(output io.Writer, configFile string) error {
	if output == nil {
		output = ioutil.Discard
	}
	viper.SetConfigFile(configFile)
	viper.SetConfigType(configType)
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		_, _ = fmt.Fprintf(output, "Config file changed %s \n", e.Name)
	})
	return nil
}

func MustInit(output io.Writer, conf string) {
	if err := Init(output, conf); err != nil {
		log.Fatal("Fatal error config file: %s \n", err)
	}
}
