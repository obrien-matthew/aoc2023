package main

import "fmt"

func partoneIsValidSet(set Set) bool {
	maxValues := Set{
		"red":   12,
		"blue":  14,
		"green": 13,
	}
	for _, color := range []string{"red", "blue", "green"} {
		if set[color] > maxValues[color] {
			return false
		}
	}
	return true
}

func partoneIsValidGame(game Game) bool {
	for _, set := range game {
		if !partoneIsValidSet(set) {
			return false
		}
	}
	return true
}

func partone(games []Game) {
	fmt.Print("Starting solution part one for day 2...\n")
	sum := 0
	for i, game := range games {
		if partoneIsValidGame(game) {
			sum += (i + 1)
		}
	}
	fmt.Printf("Part one result: [%d]\n", sum)
}
