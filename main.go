package main

import (
	"fmt"
	"main/config"
	"main/discord"
	"main/logger"
)

func main() {
	config, err := config.GetConfig()
	if err != nil {
		fmt.Println(err.Error())
	}

	logger := logger.Logger(config.LogLevel)

	go discord.Run(config, logger)

}
