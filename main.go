package main

import (
	"fmt"
	"log"

	"github.com/WatchJani/env/env"
)

func main() {
	env, err := env.Load()

	if err != nil {
		log.Println(err)
	}

	fmt.Println(env["DRVO"])
}
