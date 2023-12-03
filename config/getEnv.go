package config

import (
	"log"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	AppName string `mapstructure:"APP_NAME"`
	Env     string `mapstructure:"ENVIROTMENT"`
	Host    string `mapstructure:"HOST"`
	Port    string `mapstructure:"PORT"`
	Timeout int64  `mapstructure:"TIMEOUT"`

	DBHost         string `mapstructure:"POSTGRES_HOST"`
	DBUserName     string `mapstructure:"POSTGRES_NAME"`
	DBUserPassword string `mapstructure:"POSTGRES_PASS"`
	DBName         string `mapstructure:"POSTGRES_DB"`
	DBPort         string `mapstructure:"POSTGRES_PORT"`

	RedisAddress  string `mapstructure:"REDIS_ADDRESS"`
	RedisPassword string `mapstructure:"REDIS_PASSWORD"`
	RedisDb       int    `mapstructure:"REDIS_DB"`
	RedisExpired  int    `mapstructure:"REDIS_EXPIRED"`

	AccTokenPrivateKey string        `mapstructure:"ACCESS_TOKEN_PRIVATE_KEY"`
	AccTokenPublicKey  string        `mapstructure:"ACCESS_TOKEN_PUBLIC_KEY"`
	AccTokenExpireIn   time.Duration `mapstructure:"ACCESS_TOKEN_EXPIRED_IN"`
	AccTokenMaxEge     int           `mapstructure:"ACCESS_TOKEN_MAXAGE"`

	EmailFrom string `mapstructure:"EMAIL_FROM"`
	SMTPHost  string `mapstructure:"SMTP_HOST"`
	SMTPUser  string `mapstructure:"SMTP_USER"`
	SMTPPass  string `mapstructure:"SMTP_PASS"`
	SMTPPort  int    `mapstructure:"SMTP_PORT"`
}

var AppEnv Config

func init() {
	log.Println("Start Load Config")

	var err error
	if AppEnv, err = LoaAppEnv("."); err != nil {
		log.Println("Load Config Failed", err.Error())
		panic(err)
	}
}

func LoaAppEnv(path string) (conf Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName("app")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&conf)
	return
}
