package config

import "github.com/spf13/viper"

var (
	Host     string
	Port     string
	Username string
	Password string
	Database string
)

func Init() {
	viper.SetConfigName("config")
	viper.AddConfigPath("config/")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	Host = viper.GetString("datasource.host")
	Port = viper.GetString("datasource.port")
	Username = viper.GetString("datasource.username")
	Password = viper.GetString("datasource.password")
	Database = viper.GetString("datasource.database")
}
