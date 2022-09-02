package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	count := -1               // slate at -1 to count for the intital depth being > 0
	var current, previous int // current will be 0 so first depth will be greater
	for scanner.Scan() {
		strValue := scanner.Text()
		current, _ = strconv.Atoi(strValue)
		if current > previous {
			count += 1
		}
		previous = current
	}
	fmt.Printf("Total number of ints = %d", count)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
