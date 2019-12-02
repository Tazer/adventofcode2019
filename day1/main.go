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
	totalFuelWithFuelMass := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		i1, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}

		totalFuel += calculateFuelForModule(i1)
		totalFuelWithFuelMass += calculateFuelForModuleWithFuelMass(i1)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	log.Printf("The total fuel needed is: %d  ✅", totalFuel)
	log.Printf("The total fuel with fuel mass needed is: %d  ✅", totalFuelWithFuelMass)
}

func calculateFuelForModule(input int) int {
	d := input / 3

	s := d - 2

	return s
}

func calculateFuelForModuleWithFuelMass(input int) int {
	fuel := calculateFuelForModule(input)

	neededfuel := fuel

	for {
		fuel = calculateFuelForModule(fuel)
		neededfuel += fuel
		if fuel < 1 {
			break
		}
	}

	return neededfuel
}
