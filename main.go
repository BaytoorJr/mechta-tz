package main

import (
	"fmt"
	"log"
	"mechta-tz.github.com/src/config"
	"mechta-tz.github.com/src/service"
	"mechta-tz.github.com/src/utils"
)

func main() {
	err := config.InitConfig()
	if err != nil {
		log.Fatal(err)
	}

	data, err := utils.FetchFileData(config.MainConfig.ExampleFilePath)
	if err != nil {
		log.Fatal(err)
	}

	totalSum, err := service.CalculateJson(data)
	if err != nil {
		log.Fatal(err)
	}

	if totalSum == nil {
		log.Fatal("got empty response from CalculateJson method")
	}

	fmt.Printf("Total sum: %d\n", *totalSum)
}
