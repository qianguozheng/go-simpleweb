package main

import "fmt"
import "./server"
import "./models"

func main()  {
	airdisk.Opts = airdisk.ParseArgs()

	fmt.Println("Airdisk server started:", airdisk.Opts.Database, airdisk.Opts.Port, airdisk.Opts.Logto, airdisk.Opts.Loglevel)
	models.InitDB(airdisk.Opts.Database)
	airdisk.Run()
	fmt.Println("Airdisk exit")
}
