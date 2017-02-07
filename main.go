package main

import "fmt"
import "./server"

func main()  {
	fmt.Println("Airdisk server started")
	airdisk.Run()
	fmt.Println("Airdisk exit")
}
