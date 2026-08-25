package main

import (
	"log"

	"github.com/EgoEquusFebrianto/LinguaLearn/internal/app"
)

func main() {
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}