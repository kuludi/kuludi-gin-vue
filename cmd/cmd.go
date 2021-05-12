package cmd

import (
	"github.com/kuludi/kuludi-gin-vue/config"
	"github.com/kuludi/kuludi-gin-vue/db"
	"github.com/kuludi/kuludi-gin-vue/model"
	"github.com/kuludi/kuludi-gin-vue/router"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"os"
)

var (
	cfgFile string
	logger  = &logrus.Logger{}
	rootCmd = &cobra.Command{}
)

func initConfig() {
	config.MustInit(os.Stdout, cfgFile)
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "config/dev.yaml", "config file (default is $HOME/.cobra.yaml)")
	rootCmd.PersistentFlags().Bool("debug", true, "开启debug")
	viper.SetDefault("gin.mode", rootCmd.PersistentFlags().Lookup("debug"))
}

func Execute() error {
	rootCmd.RunE = func(cmd *cobra.Command, args []string) error {

		_, err := db.Mysql(
			viper.GetString("db.hostname"),
			viper.GetString("db.port"),
			viper.GetString("db.username"),
			viper.GetString("db.password"),
			viper.GetString("db.dbname"),
		)
		if err != nil {
			return err
		}

		r := router.SetupRouter()

		r.Run()

		db.DB.AutoMigrate(&model.User{})

		port := viper.GetString("server.port")

		log.Println("port = *** =", port)

		return http.ListenAndServe(port, nil)

	}

	return rootCmd.Execute()
}
