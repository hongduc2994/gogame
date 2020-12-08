package game

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Game is the game
func Game(number int) {
	tries := 1
	random := number
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
			fmt.Println("Wrong!")
			if tries > 10 {
				fmt.Println("The number was " + fmt.Sprint(random) + " you Pepeg")
			}
		}
	}
}

func inputNumber() float64 {
	fmt.Println("Please input a number")
	var output float64
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	} else {
		output, err = strconv.ParseFloat(strings.TrimSpace(input), 64)
		if err != nil {
			output = -1
		}
	}
	return output
}
