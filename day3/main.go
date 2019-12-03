package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	line1 := ""
	line2 := ""

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if line1 == "" {
			line1 = scanner.Text()
			continue
		}
		line2 = scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	log.Printf("Line1: %s Line2: %s", line1, line2)

}

func calculateManhattanDistance(line1,line2 string) int {
	return 0
}
