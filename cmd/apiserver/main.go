package main

import (
	"log"

	"github.com/ramizkalabayov/goproject/internal/app/apiserver"
)

func main() {
	s := apiserver.New()
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
