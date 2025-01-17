package config

import (
	"errors"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"bitbucket.org/hasaki-tech/zeus/config"
	"github.com/spf13/viper"
)

type Config struct {
	DiscordBot  DiscordBot
	Mysql       Mysql
	Makersuite  MakersuiteConfig
	Redis       RedisConfig
	KafkaQc     config.KafkaConfig
	KafkaProd   config.KafkaConfig
	ApiExternal ApiExternalConfig
}

type Butler struct {
	Token string
}
type DiscordBot struct {
	Butler Butler
}

type Mysql struct {
	Username string
	Password string
	Host     string
	Port     int64
	DBName   string
}

type RedisConfig struct {
	RedisAddr      string
	RedisPassword  string
	RedisDB        string
	RedisDefaultdb string
	MinIdleConns   int
	PoolSize       int
	PoolTimeout    int
	Password       string
	DB             int
}

type MakersuiteConfig struct {
	Endpoint string
	ApiKey   string
	Model    string
}

type ApiExternalConfig struct {
	Wms     WmsConfig
	Discord DiscordConfig
}

type WmsConfig struct {
	Url      string
	Email    string
	Password string
}

type DiscordConfig struct {
	Url      string
	Login    string
	Password string
	Undelete bool
}

func GetConfig() (*Config, error) {
	configPath := getConfigPath(os.Getenv("ENVIRONMENT"))
	cfgFile, err := LoadConfig(configPath)
	if err != nil {
		return nil, err
	}

	cfg, err := ParseConfig(cfgFile)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}

// Load config file from given path
func LoadConfig(filename string) (*viper.Viper, error) {
	v := viper.New()
	v.SetConfigName(filename)
	_, b, _, _ := runtime.Caller(0)
	path_config := filepath.Join(filepath.Dir(b), "..")
	v.AddConfigPath(path_config)
	v.AutomaticEnv()
	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, errors.New("config file not found")
		}
		return nil, err
	}
	return v, nil
}

// Parse config file
func ParseConfig(v *viper.Viper) (*Config, error) {
	var c Config
	err := v.Unmarshal(&c)
	if err != nil {
		log.Printf("unable to decode into struct, %v", err)
		return nil, err
	}
	return &c, nil
}

func getConfigPath(env string) string {
	if env == "qc" {
		return "./config/config-qc"
	}
	if env == "staging" {
		return "./config/config-staging"
	}
	if env == "prod" {
		return "./config/config-prod"
	}
	return "./config/config"
}
