package utils

import "github.com/spf13/viper"

// All app config is stored in this structure
// The values are read by viper from a config file or enviroment variables

type Config struct {
	// General Config
	APPName string `mapstructure:"APP_NAME"`
	APPPort string `mapstructure:"APP_PORT"`
	// Database Config
	DBLog          bool   `mapstructure:"DB_LOG"`
	DBHost         string `mapstructure:"DB_HOST"`
	DBPort         string `mapstructure:"DB_PORT"`
	DBName         string `mapstructure:"DB_NAME"`
	DBUser         string `mapstructure:"DB_USER"`
	DBPass         string `mapstructure:"DB_PASS"`
	DBChar         string `mapstructure:"DB_CHAR"`
	DBMaxIdleConns int64  `mapstructure:"DB_MAXIDLECONNS"`
	DBMaxOpenConns int64  `mapstructure:"DB_MAXOPENCONNS"`

	// Mail Config
	MAILMailer string `mapstructure:"MAIL_MAILER"`
	MAILHost   string `mapstructure:"MAIL_HOST"`
	MAILPort   string `mapstructure:"MAIL_PORT"`
	MAILUser   string `mapstructure:"MAIL_USERNAME"`
	MAILPass   string `mapstructure:"MAIL_PASSWORD"`
}

// LoadConfig reads configuration from file or enviroment variables
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
