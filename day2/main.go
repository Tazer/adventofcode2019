package main

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	b, err := ioutil.ReadFile("input.txt") // just pass the file name
	if err != nil {
		log.Fatal(err)
	}

	input := string(b)

	intCodes := strings.Split(input, ",")

	intCodes[1] = "12"
	intCodes[2] = "2"

	res := processIntCode(intCodes)

	result := strings.Split(res, ",")

	log.Printf("Output from IntCode computer 🖥 the value of the first output is : %s", result[0])

}

func processIntCode(intCodes []string) string {

	for i := 0; i < len(intCodes); i += 4 {
		v := intCodes[i]

		if v == "99" {
			log.Print("Breaking the code")
			break
		}

		pos1, err := strconv.Atoi(intCodes[i+1])
		if err != nil {
			log.Fatal(err)
		}
		pos2, err := strconv.Atoi(intCodes[i+2])
		if err != nil {
			log.Fatal(err)
		}
		posRes, err := strconv.Atoi(intCodes[i+3])
		if err != nil {
			log.Fatal(err)
		}

		v1, err := strconv.Atoi(intCodes[pos1])
		if err != nil {
			log.Fatal(err)
		}

		v2, err := strconv.Atoi(intCodes[pos2])
		if err != nil {
			log.Fatal(err)
		}

		if v == "1" {
			intCodes[posRes] = strconv.Itoa(v1 + v2)
		}

		if v == "2" {
			intCodes[posRes] = strconv.Itoa(v1 * v2)
		}
	}

	res := strings.Join(intCodes, ",")

	return res
}
