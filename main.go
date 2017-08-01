package main

import (
	"log"

	"github.com/gobuffalo/envy"
	"github.com/konart/tft/actions"
)

func main() {
	port := envy.Get("PORT", "3000")
	app := actions.App()
	log.Fatal(app.Start(port))
}
