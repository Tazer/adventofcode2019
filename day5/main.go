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

	startupIntCode(input, 1)

	// _ := strings.Split(res, ",")

	// log.Printf("Output from IntCode computer 🖥 the value of the first output is : %s", result[0])

}

func startupIntCode(input string, pInput int) string {
	intCodes := strings.Split(input, ",")
	return processIntCode(intCodes, pInput)
}

func processIntCode(intCodes []string, pInput int) string {

	for i := 0; i < len(intCodes); {
		processCode := intCodes[i]
		v := intCodes[i]
		if len(processCode) > 2 {
			if v != "99" {
				v = processCode[len(processCode)-1:]
			}

		}

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
		pos3, err := strconv.Atoi(intCodes[i+3])
		if err != nil {
			log.Fatal(err)
		}

		// posRes, err := strconv.Atoi(intCodes[i+4])
		// if err != nil {
		// 	log.Fatal(err)
		// }

		// log.Printf("v %s , pos1 %d pos2 %d pos3 %d code %s", v, pos1, pos2, pos3, processCode)

		v1code := "0"
		v2code := "0"
		//v3code := "0"

		if len(processCode) > 2 {
			v1code = string(processCode[0])
			if len(processCode) > 3 {
				v1code = string(processCode[1])
				v2code = string(processCode[0])
			}
			// if len(processCode) > 4 {
			// 	v3code = processCode[:5]
			// }
		}

		// log.Printf("v1code: %s v2code: %s , processcode: %s", v1code, v2code, processCode)
		v1 := 0
		v2 := 0
		if v != "4" && v != "3" {
			if v1code != "1" {
				v1, err = strconv.Atoi(intCodes[pos1])
				if err != nil {
					log.Fatal(err)
				}
			}
			if v2code != "1" {
				v2, err = strconv.Atoi(intCodes[pos2])
				if err != nil {
					log.Fatal(err)
				}
			}

		}

		// v3, err := strconv.Atoi(intCodes[pos3])
		// if err != nil {
		// 	log.Fatal(err)
		// }

		if v1code == "1" {
			v1 = pos1
		}

		if v2code == "1" {
			v2 = pos2
		}

		// if v3code == "1" {
		// 	v3 = pos3
		// }

		if v == "1" {
			// log.Printf("v1code %s values %d , %d", v1code, v1, v2)
			intCodes[pos3] = strconv.Itoa(v1 + v2)
			i += 4
		}

		if v == "2" {
			intCodes[pos3] = strconv.Itoa(v1 * v2)
			i += 4
		}

		if v == "3" {
			intCodes[pos1] = strconv.Itoa(pInput)
			i += 2
		}

		if v == "4" {
			// log.Printf("Hmm v1code %s , pos1 %d", v1code, pos1)
			i += 2
			if v1code == "1" {
				log.Printf("Output: %s", pos1)
			} else {
				log.Printf("Output: %s", intCodes[pos1])
			}

		}
	}

	res := strings.Join(intCodes, ",")

	return res
}
