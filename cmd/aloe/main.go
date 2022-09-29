package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/clfs/aloe/engine"
	"github.com/clfs/aloe/uci"
)

func main() {
	eng := engine.New()

	go func() {
		for {
			resp, err := eng.Respond()
			if errors.Is(err, uci.ErrEngineClosed) {
				return
			} else if err != nil {
				log.Fatal(err)
			}

			text, err := resp.MarshalText()
			if err != nil {
				log.Fatal(err)
			}

			fmt.Printf("%s\n", text)
		}
	}()

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()

		req, err := uci.Parse(line)
		if err != nil {
			log.Fatal(err)
		}

		err = eng.Do(req)
		if errors.Is(err, uci.ErrEngineClosed) {
			return
		} else if err != nil {
			log.Fatal(err)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
