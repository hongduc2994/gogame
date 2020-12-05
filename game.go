package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

// Game is the game
func Game() {
	max := 100
	tries := 1
	random := rand.Intn(max) + 1
	fmt.Print(random)
	for tries <= 10 {
		input := inputNumber()
		if input == float64(random) {
			fmt.Println("You are correct!")
			break
		} else if input == -1 {
			tries++
			fmt.Println("You did not enter a number!")
		} else {
			tries++
			fmt.Println(tries)
			fmt.Println("Wrong!")
		}
	}
}

func inputNumber() float64 {
	fmt.Println("Please input a number")
	var output float64
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		output = -1
	} else {
		output, err = strconv.ParseFloat(strings.TrimSpace(input), 64)
		if err != nil {
			log.Fatal(err)
		}
	}
	return output
}
