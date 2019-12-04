package main

import "strconv"

import "log"

func main() {
	from := 264360
	to := 746325

	inRange := 0
	correctNumbers := []string{}

	for i := from; i <= to; i++ {

		current := strconv.Itoa(i)

		hasSameDigits := false
		hasMixedNumbers := false
		var lastDigit rune

		for _, c := range current {
			digitInt, _ := strconv.Atoi(string(c))
			lastInt, _ := strconv.Atoi(string(lastDigit))

			if digitInt == lastInt {
				hasSameDigits = true

				//log.Print("hasSameDigits")

			}

			//log.Printf("Digit: %d , Last %d", digitInt, lastInt)
			if digitInt < lastInt {
				hasMixedNumbers = true
			}

			lastDigit = c
		}
		// log.Printf("hasSameDigits %v hasMixedNumbers %v", hasSameDigits, hasMixedNumbers)
		if hasSameDigits && !hasMixedNumbers {
			inRange++
			correctNumbers = append(correctNumbers, current)
		}
	}

	wrongNumbers := []string{}

	for _, cn := range correctNumbers {
		incorrect := getIncorrectList(cn)

		isValid := false
		for _, v := range incorrect {

			if v == 2 {
				isValid = true
			}
		}
		if isValid {
			wrongNumbers = append(wrongNumbers, cn)
		}
		// if isInvalid {
		// 	inRange--
		// }
	}

	log.Printf("In range %d in list %d , in wrong: %d , left: %d", inRange, len(correctNumbers), len(wrongNumbers), len(correctNumbers)-len(wrongNumbers))
	// for _, c := range correctNumbers {
	// 	existIn := false
	// 	for _, w := range wrongNumbers {
	// 		if c == w {
	// 			existIn = true
	// 		}
	// 	}
	// 	if existIn {
	// 		log.Printf("incorrect2:%s", c)
	// 	} else {
	// 		log.Printf("correct2:%s", c)
	// 	}
	// }

}

func getIncorrectList(cn string) map[rune]int {
	incorrect := map[rune]int{}

	for i, r := range cn {
		if i == 0 {
			continue
		}

		prev := rune(cn[i-1])

		var prev2 rune

		if i-2 > -1 {
			prev2 = rune(cn[i-2])
		}

		if prev == r {
			if prev2 == r {
				incorrect[r]++
			} else {
				if incorrect[r] != 2 {
					incorrect[r] += 2
				}
			}

		}
	}
	return incorrect
}
