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

	res := startupIntCode(input, 1)

	result := strings.Split(res, ",")

	log.Printf("Output from IntCode computer 🖥 the value of the first output is : %s", result[0])

}

func startupIntCode(input string, pInput int) string {
	intCodes := strings.Split(input, ",")
	return processIntCode(intCodes, pInput)
}

func processIntCode(intCodes []string, pInput int) string {

	for i := 0; i < len(intCodes); {
		v := intCodes[i]

		if v == "99" {
			break
		}

		vI, err := strconv.Atoi(v)
		if err != nil {
			log.Fatal(err)
		}

		pos1, err := strconv.Atoi(intCodes[i+1])
		if err != nil {
			log.Fatal(err)
		}
		pos2, err := strconv.Atoi(intCodes[i+2])
		if err != nil {
			log.Fatal(err)
		}
		pos3, err := strconv.Atoi(intCodes[i+3])
		if err != nil {
			log.Fatal(err)
		}

		posRes, err := strconv.Atoi(intCodes[i+4])
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

		log.Printf("v %s , pos1 %d pos2 %d pos3 %d posRes %d", v, pos1, pos2, pos3, posRes)

		if v == "1" {
			intCodes[posRes] = strconv.Itoa(v1 + v2)
			i += 4
		}

		if v == "2" {
			intCodes[posRes] = strconv.Itoa(v1 * v2)
			i += 4
		}

		if v == "3" {
			intCodes[pos1] = strconv.Itoa(pInput)
			i += 2
		}

		if v == "4" {
			i += 2
			log.Printf("Output: %s", intCodes[pos1])
		}
	}

	res := strings.Join(intCodes, ",")

	return res
}
