package main

import (
	"log"

	"github.com/hyperyuri/test-go/services"
)

func main() {
	s := services.Sum(2, 2)
	m := services.Multiply(2, 2)

	log.Print(m, s)
}
