package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/clfs/aloe/engine"
	"github.com/clfs/aloe/uci"
)

func main() {
	eng := engine.New()

	adapter := uci.NewAdapter(eng)

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()

		messages, err := adapter.SendLine(line)
		if err != nil {
			log.Fatal(err)
		}

		for _, m := range messages {
			fmt.Println(m)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
