package main

import (
	"fmt"
	"bufio"
	"os"
	"unicode"
	"strconv"
)


// CODE FOR DAY 1


func digitinstring(line string, this map[string]string)(int){
	var sliceofletters []string
	for i,e  := range line {
		if unicode.IsDigit(e) {
			sliceofletters = append(sliceofletters, string(e))
		 }else {
			for key,value := range this {
				length := len(key)
				if i + length <= len(line) {
					slice := line[i : i + length]
					if slice == key {
						sliceofletters = append(sliceofletters, value)
					}
				}
			}
		}
		}
	concat,err := strconv.Atoi(sliceofletters[0] + sliceofletters[len(sliceofletters) - 1])
	if err != nil {
		fmt.Println("error",err)
		os.Exit(1)
	}
	return concat
}


func main() {
	myMap := make(map[string]string)
	myMap["one"] = "1"
	myMap["two"] = "2"
	myMap["three"] = "3"
	myMap["four"] = "4"
	myMap["five"] = "5"
	myMap["six"] = "6"
	myMap["seven"] = "7"
	myMap["eight"] = "8"
	myMap["nine"] = "9"
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
		sum += digitinstring(line,myMap)
		
	}
	fmt.Println(sum)
}









