package main

import (
	"github.com/KadyrPoyraz/httplayout/config"
	"github.com/KadyrPoyraz/httplayout/internal/app"
)

func main() {
	cnf, err := config.New()
    if err != nil {
        panic(err)
    }

	application := app.New(cnf)
    err = application.Run()
    if err != nil {
        panic(err)
    }
}
