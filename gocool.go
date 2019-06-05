package main

import (
	"gocool/config"
	"fmt"
)

func main()  {
	fmt.Println("today is good")
	config.GetConfig()
}