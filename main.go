package main

import (
	"log"

	"github.com/phrabec/projector-controller/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
