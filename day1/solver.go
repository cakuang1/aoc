package main

import (
	"fmt"
	"bufio"
	"os"
	"unicode"
	"strconv"
)


// CODE FOR DAY 1

func digitinstring(line string)(int){
	foundfirst := false
	var first,last rune
	for _,e  := range line {
		if unicode.IsDigit(e) {
			if !foundfirst {
				foundfirst = true
				first = e
			}
			last = e
		}
	}
	
	concat,err := strconv.Atoi(string(first) + string(last))
	if err != nil {
		fmt.Println("error",err)
		os.Exit(1)
	}
	return concat
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
		sum += digitinstring(line)

	}
	fmt.Println(sum)
}









