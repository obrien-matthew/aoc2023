package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func parttwo() {
	numMap := map[string]int{
		"zero":  0,
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}
	fmt.Print("Starting solution part 2 for day 1...\n")
	file, err := os.Open("./input.txt")
	check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan() {
		setFirst := false
		first, last := 0, 0
		lineText := scanner.Text()
		// fmt.Printf("Line: [%s]\n", lineText)
		for i, c := range lineText {
			if unicode.IsDigit(c) {
				val := int(c - '0')
				// fmt.Printf("\tval: [%d]\n", val)
				if !setFirst {
					setFirst = true
					first = val
				}
				last = val
			} else {
				lastIndex := len(lineText)
				if val, ok := numMap[string(lineText[i:min(i+3, lastIndex)])]; ok {
					// fmt.Printf("\tval: [%d]\n", val)
					if !setFirst {
						setFirst = true
						first = val
					}
					last = val
				} else if val, ok := numMap[string(lineText[i:min(i+4, lastIndex)])]; ok {
					// fmt.Printf("\tval: [%d]\n", val)
					if !setFirst {
						setFirst = true
						first = val
					}
					last = val
				} else if val, ok := numMap[string(lineText[i:min(i+5, lastIndex)])]; ok {
					// fmt.Printf("\tval: [%d]\n", val)
					if !setFirst {
						setFirst = true
						first = val
					}
					last = val
				}
			}
		}
		val := (first * 10) + last
		// fmt.Printf("Val:  [%d]\n", val)
		sum += val
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
	fmt.Printf("Done! Part two solution: [%d]\n", sum)
}

func partone() {
	fmt.Print("Starting solution part one for day 1...\n")
	file, err := os.Open("./input.txt")
	check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan() {
		setFirst := false
		first, last := 0, 0
		lineText := scanner.Text()
		// fmt.Printf("Line: [%s]\n", lineText)
		for _, c := range lineText {
			if unicode.IsDigit(c) {
				if !setFirst {
					setFirst = true
					first = int(c - '0')
					check(err)
				}
				last = int(c - '0')
				check(err)
			}
		}
		val := (first * 10) + last
		// fmt.Printf("Val:  [%d]\n", val)
		sum += val
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
	fmt.Printf("Done! Part one solution: [%d]\n", sum)
}

func main() {
	partone()
	parttwo()
}
