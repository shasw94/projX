package config

import (
	"fmt"
	"github.com/spf13/viper"
	"strings"
)

type Schema struct {
	Env          string `mapstructure:"env"`
	DefaultLimit int    `mapstructure:"defailt_limit"`
	MaxLimit     int    `mapstructure:"max_limit"`
	Database     struct {
		Host     string `mapstructure:"host"`
		Port     int    `mapstructure:"port"`
		Name     string `mapstructure:"name"`
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		Env      string `mapstructure:"env"`
		SSLMode  string `mapstructure:"sslmode"`
	} `mapstructure:"database"`

	Redis struct {
		Host     string `mapstructure:"host"`
		Port     int    `mapstructure:"port"`
		Password string `mapstructure:"password"`
		Database int    `mapstructure:"database"`
	} `mapstructure:"redis"`

	Cache struct {
		Enable     bool `mapstructure:"enable"`
		ExpiryTime int  `mapstructure:"expiry_time"`
	} `mapstructure:"cache"`

	JWTAuth struct {
		SigningKey          string `mapstructure:"signing_key"`
		Expired             int    `mapstructure:"expired"`
		SigningRefreshKey   string `mapstructure:"signing_refresh_key"`
		ExpiredRefreshToken int    `mapstructure:"expired_refresh_token"`
	} `mapstructure:"jwt_auth"`

	Casbin struct {
		Enable           bool   `mapstructure:"enable""`
		Debug            bool   `mapstructure:"debug"`
		Model            string `mapstructure:"model"`
		AutoLoad         bool   `mapstructure:"auto_load"`
		AutoLoadInternal int    `mapstructure:"auto_load_internal"`
	} `mapstructure:"casbin"`

	CORS struct {
		Enable           bool     `mapstructure:"enable"`
		AllowOrigins     []string `mapstructure:"allow_origins"`
		AllowMethods     []string `mapstructure:"allow_methods"`
		AllowHeaders     []string `mapstructure:"allow_headers"`
		AllowCredentials bool     `mapstructure:"allow_credentials"`
		MaxAge           int      `mapstructure:"max_age"`
	} `mapstructure:"cors"`
}

// Config global parameter config
var Config Schema

func init() {
	config := viper.New()
	config.SetConfigName("config")
	config.AddConfigPath(".")
	config.AddConfigPath("config/")
	config.AddConfigPath("../config/")
	config.AddConfigPath("../")
	config.AddConfigPath("../../config/")
	config.AddConfigPath("../../")

	config.SetEnvKeyReplacer(strings.NewReplacer(".", "__"))
	config.AutomaticEnv()

	err := config.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}

	err = config.Unmarshal(&Config)
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}
}
