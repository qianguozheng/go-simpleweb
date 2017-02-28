package main

import "fmt"
import "./server"
import "./models"
import (
	"./regserver"
	"github.com/goless/config"
	"./binhtml"

)

func main()  {
	airdisk.Opts = airdisk.ParseArgs()

	if airdisk.Exist(airdisk.Opts.Config) {
		config := config.New(airdisk.Opts.Config)
		regserver.Token = config.Get("token").(string)
		database := config.Get("db")
		models.InitDB(database.(string))
		airdisk.AppId = config.Get("appid").(string)
		airdisk.AppSecret = config.Get("appsecret").(string)
		airdisk.SecretKey = config.Get("secretkey").(string) //TODO: remove from parameter
		airdisk.Opts.Port = config.Get("port").(string)
		airdisk.Opts.Logto = config.Get("log").(string)
		airdisk.Opts.Loglevel = config.Get("log-level").(string)
	} else {
		regserver.Token = airdisk.Opts.Token
		models.InitDB(airdisk.Opts.Database)
		airdisk.AppId = airdisk.Opts.AppId
		airdisk.AppSecret = airdisk.Opts.AppSecret
		airdisk.SecretKey = airdisk.Opts.SecretKey
	}

	fmt.Println("Airdisk server started:", airdisk.Opts.Database, airdisk.Opts.Port,
			airdisk.Opts.Logto, airdisk.Opts.Loglevel)

	bt := binhtml.New(binhtml.Asset, binhtml.AssetDir)
	go regserver.Run()
	airdisk.Run(bt)
	fmt.Println("Airdisk exit")
}
