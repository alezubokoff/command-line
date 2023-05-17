package main

import (
	"fmt"
	"log"
	"os"

	"github.com/alezubokoff/command-line/app"
)

func main() {
	fmt.Println("Sistema de linha de comando")

	app := app.Generate()
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}

}
