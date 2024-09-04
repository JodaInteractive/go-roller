package main

import (
	"math"
	"math/rand"
	"os"
	"regexp"
	"strconv"
)

var dieRegex = regexp.MustCompile(`(^\d+|)d(\d+$)`)

func main() {
	roller, quantity, sides := getRoller()
	println("Rolling " + roller + "...")
	println("Results:")
	results := roll(quantity, sides)

	if quantity > 1 {
		printMultiRollStats(results)
	}
}

// getRoller returns the roller, quantity, and sides of the die to roll
// string - the roller
// int - the quantity of dice to roll
// int - the number of sides on the die
func getRoller() (string, int, int) {
	if len(os.Args) < 2 {
		println("You did not provide a something to roll")
		os.Exit(0)
	}

	roller := os.Args[1]
	matches := dieRegex.FindStringSubmatch(roller)
	if len(matches) < 3 || matches[2] == "" || matches[2] == "0" || matches[1] == "0" {
		println("Invalid roller. Must match format '2d6' or simply 'd4' for a single roll.")
		os.Exit(0)
	}

	quantity := 1
	if matches[1] != "" {
		q, err := strconv.Atoi(matches[1])
		if err != nil {
			println("Invalid quantity")
			os.Exit(0)
		}
		quantity = q
	}

	sides, err := strconv.Atoi(matches[2])
	if err != nil {
		println("Invalid sides")
		os.Exit(0)
	}

	return roller, quantity, sides
}

// rolls the dice; prints and returns the result of each roll
func roll(quantity int, sides int) []int {
	results := []int{}
	
	for i := 0; i < quantity; i++ {
		result := rollDie(sides)
		results = append(results, result)
		println(result)
	}

	return results
}

// rolls a single die with the given number of sides and returns the result
func rollDie(sides int) int {
	return rand.Intn(sides) + 1
}

// prints the total, lowest, highest, and average of the results of multiple rolls
func printMultiRollStats(results []int) {
	total := 0
	lowest := math.MaxInt32
	highest := 0
	for _, result := range results {
		total += result
		if result < lowest {
			lowest = result
		}
		if result > highest {
			highest = result
		}
	}
	average := total / len(results)

	println("Total: " + strconv.Itoa(total))
	println("Lowest: " + strconv.Itoa(lowest))
	println("Highest: " + strconv.Itoa(highest))
	println("Average: " + strconv.Itoa(average))
}
