package main

import (
	"github.com/pocketbase/pocketbase"

	_ "github.com/zcharym/pocketbase-client/migrations"
)

func main() {
	app := pocketbase.New()

	if err := app.Start(); err != nil {
		panic(err)
	}
}
