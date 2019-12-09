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

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		log.Print(scanner.Text())
	}
}

func calculateOrbitConnections(input []string) int {

	connections := 0

	for i := len(input) - 1; i >= 0; i-- {
		r := strings.Split(strings.TrimSpace(input[i]), ")")
		//a := r[0]
		b := strings.TrimSpace(r[0])

		connections++

		// log.Printf("direct %v", r)

		offset := 0

		next := i - 1

		for ii := next; ii >= 0; ii-- {

			r := strings.Split(input[ii], ")")

			for iii := next - offset; iii >= 0; iii-- {

				r := strings.Split(input[iii], ")")

				log.Printf("checking if %s is %s", r[1], b)

				if strings.TrimSpace(r[1]) == b {
					log.Printf("indirect %v", r)
					connections++
				}
			}

			b = strings.TrimSpace(r[0])
			offset = 1
		}

	}

	return connections
}
