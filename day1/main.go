package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	totalFuel := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		i1, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}

		totalFuel += calculateFuelForModule(i1)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	log.Printf("The total fuel needed is: %d  ✅", totalFuel)
}

func calculateFuelForModule(input int) int {
	d := input / 3

	s := d - 2

	return s
}
