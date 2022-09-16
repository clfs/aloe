package main

import (
	"log"
	"os"

	"github.com/clfs/aloe/engine"
	"github.com/clfs/aloe/uci"
)

func main() {
	eng := engine.New()

	client := uci.NewClient(eng, os.Stdout)

	if err := client.Run(os.Stdin); err != nil {
		log.Fatal(err)
	}
}
