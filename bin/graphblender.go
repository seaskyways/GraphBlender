package main

import (
	"GraphBlender"
	"fmt"
	"github.com/graphql-go/handler"
	_ "github.com/joho/godotenv/autoload"
	"log"
	"net/http"
	"os"
)

var autoLoader = GraphBlender.AutoLoader{}

func main() {
	autoLoadDir := os.Getenv("GB_AUTOLOAD")
	if err := autoLoader.Run(autoLoadDir); err != nil {
		log.Fatal("error loading datasets", err)
	}

	fmt.Println("dataframes available: ", len(autoLoader.DataFrames))

	processor := GraphBlender.NewProcessor(autoLoader.DataFrames[0].DataFrame)
	processor.CleanHeaders()

	blender := GraphBlender.New(autoLoader.DataFrames)
	gql, err := blender.Blend()
	if err != nil {
		log.Fatal("error building GraphQL", err)
	}

	h := handler.New(&handler.Config{
		Schema:   &gql,
		Pretty:   true,
		GraphiQL: true,
	})

	http.Handle("/graphql", h)
	http.ListenAndServe(":8080", nil)
}
