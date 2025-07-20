package main

import (
	"log"

	"github.com/opoccomaxao/myownranking/pkg/app"
)

func main() {
	err := app.Run()
	if err != nil {
		log.Fatalf("%+v", err)
	}
}
