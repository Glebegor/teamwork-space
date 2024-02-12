package bootstrap

import (
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"os"
)

type Env struct {
	SERVERenv    string
	SERVERport   string
	SERVERsecret string
	DBport       string
	DBhost       string
	DBname       string
	DBsslmode    string
	DBusername   string
	DBpassword   string
}

func NewEnv(EnvName string) (*Env, error) {
	env := &Env{}
	viper.SetConfigName(EnvName)
	viper.AddConfigPath("config")
	if err := viper.ReadInConfig(); err != nil {
		return env, err
	}

	if err := godotenv.Load(); err != nil {
		return env, err
	}
	env.SERVERenv = viper.GetString("server.ENV")
	env.SERVERport = viper.GetString("server.PORT")
	env.DBport = viper.GetString("db.PORT")
	env.DBhost = viper.GetString("db.HOST")
	env.DBname = viper.GetString("db.NAME")
	env.DBsslmode = viper.GetString("db.SSLMODE")
	env.DBusername = viper.GetString("db.USERNAME")

	if env.SERVERenv == "tests" {
		env.DBpassword = viper.GetString("db.TESTDBPASSWORD")
		env.SERVERsecret = viper.GetString("server.TESTSECRET")
	} else if env.SERVERenv == "product" {
		env.DBpassword = viper.GetString("db.TESTDBPASSWORD")
		env.SERVERsecret = viper.GetString("server.TESTSECRET")
	} else {
		env.DBpassword = os.Getenv("DB_PASSWORD")
		env.SERVERsecret = os.Getenv("SECRET")
	}

	return env, nil
}
