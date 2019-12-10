package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	inputs := []string{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		inputs = append(inputs, scanner.Text())
	}

	res := calculateOrbitConnections(inputs)
	log.Printf("the connetion was %d 🌚", res)
}

func calculateOrbitConnections(input []string) int {

	connections := 0

	for i := len(input) - 1; i >= 0; i-- {
		r := strings.Split(strings.TrimSpace(input[i]), ")")
		//a := r[0]
		b := strings.TrimSpace(r[0])

		added := 0

		usedChars := map[string]bool{}

		// connections++

		// log.Printf("direct %v", r)

		for ii := len(input) - 1; ii >= 0; ii-- {

			r := strings.Split(input[ii], ")")

			for iii := len(input) - 1; iii >= 0; iii-- {

				r := strings.Split(input[iii], ")")

				// log.Printf("checking if %s is %s", r[1], b)

				if ok, _ := usedChars[b]; ok {
					// log.Printf("skipping cause %s is already done", b)
					continue
				}

				if input[i] == input[iii] || input[ii] == input[iii] {
					// log.Print("Skipping same as input")
					continue
				}

				if strings.TrimSpace(r[1]) == b {
					usedChars[strings.TrimSpace(b)] = true
					log.Printf("indirect %v", r)
					added++
				}
			}

			b = strings.TrimSpace(r[0])
		}
		// log.Printf("Added %d for %v", added, r)
		connections += added
	}

	log.Printf("connetion %d", connections)

	return connections
}
