package main

import (
	"errors"
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

	res := startupIntCode("12", "2", input)

	result := strings.Split(res, ",")

	log.Printf("Output from IntCode computer 🖥 the value of the first output is : %s", result[0])

	log.Print("Gonna find correct input")
	n, v, err := findoutput("19690720", input)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Noun: %d Verb: %d Correct result: %d 🔥", n, v, 100*n+v)

}

func findoutput(expected, input string) (int, int, error) {
	for n := 0; n < 100; n++ {
		for v := 0; v < 100; v++ {

			res := startupIntCode(strconv.Itoa(n), strconv.Itoa(v), input)
			result := strings.Split(res, ",")
			if result[0] == expected {
				return n, v, nil
			}
		}
	}
	return 0, 0, errors.New("didnt find correct output")
}

func startupIntCode(noun, verb, input string) string {
	intCodes := strings.Split(input, ",")

	intCodes[1] = noun
	intCodes[2] = verb

	return processIntCode(intCodes)
}

func processIntCode(intCodes []string) string {

	for i := 0; i < len(intCodes); i += 4 {
		v := intCodes[i]

		if v == "99" {
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
