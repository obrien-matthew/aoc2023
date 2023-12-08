package main

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Set map[string]int
type Game []Set

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func parseSet(setString string) Set {
	resultRegexp := regexp.MustCompile(`(?P<count>\d+) (?P<color>red|green|blue)`)
	setMap := Set{
		"red":   0,
		"blue":  0,
		"green": 0,
	}
	for _, result := range strings.Split(setString, ", ") {
		// fmt.Printf("\tresult: [%s]\n", result)
		matches := resultRegexp.FindStringSubmatch(result)
		count, err := strconv.Atoi(matches[1])
		check(err)
		color := matches[2]
		// fmt.Printf("\t\tcount: %d\n\t\tcolor: %s\n", count, color)
		setMap[color] += count
	}
	// fmt.Println(setMap)
	return setMap
}

func parseGame(setsString string) Game {
	// 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
	var resultsArr Game
	for _, set := range strings.Split(setsString, "; ") {
		// fmt.Printf("set: [%s]\n", set)
		resultsArr = append(resultsArr, parseSet(set))
	}
	return resultsArr

}

func parseLine(line string) (int, Game) {
	// Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
	splitStr := strings.Split(line, ": ")
	if len(splitStr) < 1 {
		return 0, nil
	}
	gameStr := splitStr[0]
	r, err := regexp.Compile(`Game (\d+)`)
	check(err)
	// fmt.Printf("gameStr: [%s]\n", gameStr)
	matches := r.FindStringSubmatch(gameStr)
	// for i, match := range matches {
	// 	fmt.Printf("\t[%d]: [%s]\n", i, match)
	// }
	gameId, err := strconv.Atoi(matches[1])
	check(err)
	// fmt.Printf("Game ID: [%d]\n", gameId)
	return gameId, parseGame(splitStr[1])
}

func main() {
	file, err := os.Open("./input.txt")
	check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var games []Game
	for scanner.Scan() {
		lineText := scanner.Text()
		_, game := parseLine(lineText)
		games = append(games, game)
	}
	partone(games)
	parttwo(games)
}
