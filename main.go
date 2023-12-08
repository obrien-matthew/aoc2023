package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func makeGet(requestURL string, filename string, sessionCookie http.Cookie) {
	fmt.Printf("Request URL: [%s]\nFilename: [%s]\n", requestURL, filename)
	req, err := http.NewRequest(http.MethodGet, requestURL, nil)
	check(err)
	req.AddCookie(&sessionCookie)
	client := http.Client{Timeout: 30 * time.Second}
	resp, err := client.Do(req)
	check(err)
	defer resp.Body.Close()

	fmt.Printf("HTTP status code: %d\n", resp.StatusCode)

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(resp.Body)
		check(err)
		fo, err := os.Create(filename)
		check(err)

		// close fo when exiting
		defer func() {
			if err := fo.Close(); err != nil {
				panic(err)
			}
		}()

		w := bufio.NewWriter(fo)
		nn, err := w.Write(bodyBytes)
		check(err)
		fmt.Printf("Bytes written: %d\n", nn)
	}
}

func getInput(day int, sessionCookie http.Cookie) {
	requestURL := fmt.Sprintf("https://adventofcode.com/2023/day/%d/input", day)
	filename := fmt.Sprintf("./day%02d/input.txt", day)
	makeGet(requestURL, filename, sessionCookie)
}

func getProblem(day int, sessionCookie http.Cookie) {
	requestURL := fmt.Sprintf("https://adventofcode.com/2023/day/%d", day)
	filename := fmt.Sprintf("./day%02d/problem.html", day)
	makeGet(requestURL, filename, sessionCookie)
}

func main() {
	sessionCookie := http.Cookie{Name: "session", Value: os.Getenv("AOC_SESSION_TOKEN")}
	day, err := strconv.Atoi(os.Args[1])
	check(err)
	fmt.Printf("Making requests for day: %d\n", day)
	newpath := fmt.Sprintf("./day%02d", day)
	err = os.MkdirAll(newpath, os.ModePerm)
	check(err)
	getInput(day, sessionCookie)
	getProblem(day, sessionCookie)
	fmt.Printf("Done!\n")
}
