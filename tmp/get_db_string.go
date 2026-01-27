package main

import (
	"Api-Web/app/config"
	"fmt"
)

func main() {
	config.LoadConfig()
	fmt.Println(config.StringConnectDatabase)
}
