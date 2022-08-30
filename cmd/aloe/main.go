package main

import (
	"log"
	"os"

	"github.com/clfs/aloe/engine"
	"github.com/clfs/aloe/uci"
)

func main() {
	eng := engine.New()

	adapter := uci.NewAdapter(eng, os.Stdin, os.Stdout)

	if err := adapter.Run(); err != nil {
		log.Fatal(err)
	}
}
