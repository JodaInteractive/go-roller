package main

import (
	"math/rand"
	"os"
	"regexp"
	"strconv"
	"time"
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

	for i := 0; i < quantity; i++ {
		println(strconv.Itoa(roll_die(sides)))
	}
}

func roll_die(sides int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(sides) + 1
}
