package main

// Example line
//Game 1: 4 red, 18 green, 15 blue; 17 green, 18 blue, 9 red; 8 red, 14 green, 6 blue; 14 green, 12 blue, 2 red
import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

func helper(line string) (int) {
	current := true
	parts := strings.Split(line, ":")
	gameid := strings.Split(parts[0]," ")[1]
	num,err := strconv.Atoi(gameid)
	if err != nil {
		// Handle the error if the conversion fails
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	part2 := strings.Split(parts[1],";")
	
	for _,e := range(part2) {
		parts := strings.Split(e, ", ")

		for _, part := range parts {
			// Split each part by space to separate count and color
			part = strings.TrimSpace(part)
			subparts := strings.Split(part, " ")
			fmt.Println(subparts)
			if len(subparts) >= 2 {
				// Extract count and color
				countStr := subparts[0]
				color := subparts[1]
				// Convert count from string to integer
				count, err := strconv.Atoi(countStr)
				if err != nil {
					fmt.Println("Error:", err)
					os.Exit(1)
				}
				if color == "red" && count > 12 {
					current = false
				}
				if color == "green" && count > 13 {
					current = false
				}
				if color == "blue" && count > 14 {
					current = false
				}
			}
		}
	}
		
	if current == false{
		return 0
	}
	return num
}

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
		sum += helper(line)
	}
	
	fmt.Print(sum)

}