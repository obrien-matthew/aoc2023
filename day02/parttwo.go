package main

import "fmt"

func powerOfGameMins(game Game) int {
	mins := Set{
		"red":   0,
		"blue":  0,
		"green": 0,
	}
	for _, set := range game {
		for _, color := range []string{"red", "blue", "green"} {
			if mins[color] < set[color] {
				mins[color] = set[color]
			}
		}
	}
	// fmt.Println(mins)
	return mins["red"] * mins["blue"] * mins["green"]
}

func parttwo(games []Game) {
	fmt.Print("Starting solution part two for day 2...\n")
	sum := 0
	for _, game := range games {
		sum += powerOfGameMins(game)
	}
	fmt.Printf("Part two result: [%d]\n", sum)
}
