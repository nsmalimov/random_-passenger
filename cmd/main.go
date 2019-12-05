package main

import (
	"fmt"
	"math/rand"
	"random_passenger/internal/addresses"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	adrServ := addresses.New()

	c1, c2 := adrServ.GenCoordinates(55.752818, 37.621753, 20000)

	fmt.Println(c1, c2)
}
