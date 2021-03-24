package main

import (
	"GraphBlender"
	"fmt"
	_ "github.com/joho/godotenv/autoload"
	"log"
	"os"
)

var autoLoader = GraphBlender.AutoLoader{}

func main() {
	autoLoadDir := os.Getenv("GB_AUTOLOAD")
	if err := autoLoader.Run(autoLoadDir); err != nil {
		log.Fatal("error loading datasets", err)
	}
	fmt.Println("dataframes available: ", len(autoLoader.DataFrames))
	fmt.Println(autoLoader.DataFrames.Find("health_and_population").DataFrame.String())
}
