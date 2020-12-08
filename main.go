package main

import (
	"fmt"
	"game/game"
	"math/rand"
	"time"
)

func main() {
	max := 100
	fmt.Println("Lets the game begin")
	rand.Seed(time.Now().UnixNano())
	game.Game(rand.Intn(max) + 1)
}
