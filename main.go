package main

import (
	"math"
	"math/rand"
	"os"
	"regexp"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		println("You did not provide a something to roll")
		os.Exit(0)
	}

	roller := os.Args[1]

	roll(roller)
}

func roll(roller string) {
	println("Rolling " + roller + "...")

	regexp := regexp.MustCompile(`(^\d+|)d(\d+$)`)
	matches := regexp.FindStringSubmatch(roller)

	quantity := 0
	if matches[1] == "" {
		quantity = 1
	} else {
		q, err := strconv.Atoi(matches[1])
		if err != nil {
			println("Invalid quantity")
			os.Exit(0)
		}
		quantity = q
	}

	sides := 0
	if matches[2] == "" {
		println("Invalid sides")
		os.Exit(0)
	} else {
		s, err := strconv.Atoi(matches[2])
		if err != nil {
			println("Invalid sides")
			os.Exit(0)
		}
		sides = s
	} 
	
	println("Results:")
	results := []int{}
	for i := 0; i < quantity; i++ {
		result := roll_die(sides)
		println(result)
		results = append(results, result)
	}

	if quantity > 1 {
		total := 0
		for _, result := range results {
			total += result
		}
		println("Total: " + strconv.Itoa(total))

		lowest := math.MaxInt32
		for _, result := range results {
			if result < lowest {
				lowest = result
			}
		}
		println("Lowest: " + strconv.Itoa(lowest))

		highest := 0
		for _, result := range results {
			if result > highest {
				highest = result
			}
		}
		println("Highest: " + strconv.Itoa(highest))
	}
}

func roll_die(sides int) int {
	return rand.Intn(sides) + 1
}
