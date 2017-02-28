package airdisk

import (
	"flag"
)


type Options struct {
	Database string
	Port string
	Logto string
	Loglevel string
	AppId string
	AppSecret string
	SecretKey string
	Token string
	Config string
}

func ParseArgs() *Options{
	database := flag.String("db", "./airdisk.db", "Assign db used to connect")
	port := flag.String("port", ":38001", "Assign port used to serve")
	logto := flag.String("log","stdout", "Write log messages to this file.")
	loglevel := flag.String("log-level", "DEBUG", "The level of messages to log. one of: DEBUG, INFO, WARNING, ERROR")
	appId := flag.String("appid", "wx0bc5ea9ffc61b2dg", "The wechat AppId")
	appSecret := flag.String("appsecret","2c0f03863a6eca71c7f9218916bcb238", "The wechat AppSecret")
	secretKey := flag.String("secretkey","685aec96360b737c175b13343cc53388","The wechat secretKey")
	token := flag.String("token", "12345678901234567890qwertyuioqgz", "The wechat server token")
	config := flag.String("config", "config.json", "Config file containing parameters needed")
	flag.Parse()
	return &Options{
		Database: *database,
		Port : *port,
		Logto: *logto,
		Loglevel: *loglevel,
		AppId: *appId,
		AppSecret: *appSecret,
		SecretKey: *secretKey,
		Token: *token,
		Config: *config,
	}
}
