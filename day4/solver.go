package main

// Example line
// Card   8: 70 64 82  4 16  6 19 13  9 29 | 21 93 37 69 24 62 60  3 90 83  8 66 20 34 55 22  6 84 99 50 33 26 65 98 86

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func main() {	
	file, err := os.Open("puzzle.txt")
	if err != nil {
		fmt.Println("Error occured", err)
		os.Exit(1)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		firstsplit := strings.Split(line, ":")[1]
		secondsplit := strings.Split(firstsplit, "|")
		cardsgiven := secondsplit[0]
		matchingcards := secondsplit[1]
		fmt.Println(cardsgiven)
		fmt.Print(matchingcards)
	}
	
	fmt.Print(sum)

}



func processline(line string) {


}