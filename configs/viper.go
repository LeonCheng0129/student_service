package configs

import (
	"strings"
	"sync"

	"github.com/spf13/viper"
)

var once sync.Once

func NewViperConfig() (err error) {
	once.Do(func() {
		err = newViperConfig()
	})
	return
}

func newViperConfig() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	viper.AddConfigPath(".")
	viper.AddConfigPath("./configs")

	viper.EnvKeyReplacer(strings.NewReplacer(".", "_")) // 允许用 MYSQL_USER 覆盖 mysql.user
	viper.AutomaticEnv()

	return viper.ReadInConfig()
}
