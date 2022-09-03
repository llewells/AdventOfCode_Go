package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func readFile(path string) []int {
	f, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var lines []int
	for scanner.Scan() {
		strValue := scanner.Text()
		intValue, _ := strconv.Atoi(strValue)
		lines = append(lines, intValue)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return lines
}

func main() {
	data := readFile("input.txt")
	var current, previous, count int
	for i := 0; i <= len(data); i++ {
		if (i + 2) >= len(data) {
			break
		}
		current = data[i] + data[i+1] + data[i+2]

		if i > 0 {
			if current > previous {
				count += 1
			}
		}
		previous = current
	}
	fmt.Printf("Total number of incs: %d", count)

}
