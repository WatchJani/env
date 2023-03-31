package main

import (
	"fmt"
	"log"
	"root/env"
)

func main() {
	env, err := env.Load()

	if err != nil {
		log.Println(err)
	}

	fmt.Println(env["JANKO"])
}
