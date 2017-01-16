package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

const (
	BOILING   float64 = 5         // 5
	HOT       float64 = 10 * iota // 10
	WARM                          // 20
	LUKEWARM                      // 30
	COLD                          // 40
	MAXNUMBER int     = 100
	MAXTURNS  int     = 7 // Using an ideal strategy, the most turns it could take to win is 8.
)

func init() {
	rand.Seed(time.Now().UnixNano()) // Initialize random seed
}

func main() {
	fmt.Println("Welcome to the Guessing Game!")
	fmt.Printf("We haven chosen a random number from 1 to %d; your job is to guess it.\n", MAXNUMBER)
	fmt.Println("If you guess wrong we will tell you how close you were using a hot/cold system.")
	fmt.Printf("You will only have %d turns to guess right.\n", MAXTURNS)
	number := float64(rand.Intn(MAXNUMBER) + 1) // Normally it ranges from 0, MAXNUMBER -1. This way it ranges from 1, MAXNUMBER.
	var guess float64
	var turns int
Loop:
	for turns = 0; turns < MAXTURNS; turns++ {
		fmt.Printf("You have %d turns left. Take your guess: ", MAXTURNS - turns)
		fmt.Scanln(&guess)
		switch {
		case guess == number:
			break Loop
		case math.Abs(guess-number) <= BOILING:
			fmt.Printf("Boiling. (%.0f or closer.)\n", BOILING)
		case math.Abs(guess-number) <= HOT:
			fmt.Printf("Hot. (%.0f or closer.)\n", HOT)
		case math.Abs(guess-number) <= WARM:
			fmt.Printf("Warm. (%.0f or closer.)\n", WARM)
		case math.Abs(guess-number) <= LUKEWARM:
			fmt.Printf("Lukewarm. (%.0f or closer.)\n", LUKEWARM)
		case math.Abs(guess-number) <= COLD:
			fmt.Printf("Cold. (%.0f or closer.)\n", COLD)
		default:
			fmt.Printf("Freezing. (further than %.0f.)\n", COLD)
		}
	}
	if turns == MAXTURNS {
		fmt.Println("You have lost. Better luck next time!")
	} else {
		fmt.Printf("You have won with %d turn(s) left. Good thing there isn't a prize.\n", MAXTURNS-turns)
	}
}
