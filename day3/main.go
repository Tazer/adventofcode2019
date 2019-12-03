package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
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

	res, steps, err := calculateManhattanDistance(line1, line2)

	if err != nil {
		log.Fatalf("Problemo err:%v", err)
	}

	log.Printf("The result is %d and steps %d 🥳", res, steps)
}

type Point struct {
	X     int
	Y     int
	Steps int
}

func calculateManhattanDistance(line1, line2 string) (int, int, error) {
	l1i := strings.Split(line1, ",")
	l2i := strings.Split(line2, ",")

	line1Points := []Point{Point{
		X:     0,
		Y:     0,
		Steps: 0,
	}}

	line2Points := []Point{Point{
		X:     0,
		Y:     0,
		Steps: 0,
	}}

	for _, l := range l1i {
		direction := string(l[0])
		value, err := strconv.Atoi(l[1:])

		if err != nil {
			return 0, 0, err
		}

		var currentPoint Point

		if len(line1Points) > 0 {
			currentPoint = line1Points[len(line1Points)-1]
		}

		newPoint := processMovement(direction, value, currentPoint)

		for _, p := range newPoint {
			line1Points = append(line1Points, p)
		}
	}

	for _, l := range l2i {
		direction := string(l[0])
		value, err := strconv.Atoi(l[1:])

		if err != nil {
			return 0, 0, err
		}

		var currentPoint Point

		if len(line2Points) > 0 {
			currentPoint = line2Points[len(line2Points)-1]
		}

		newPoint := processMovement(direction, value, currentPoint)

		for _, p := range newPoint {
			line2Points = append(line2Points, p)
		}

	}

	matches := []map[string]Point{}

	var shortest []Point
	var longest []Point

	if len(line1Points) > len(line2Points) {
		longest = line1Points
		shortest = line2Points
	} else {
		shortest = line1Points
		longest = line2Points
	}

	for _, l1 := range longest {
		for _, l2 := range shortest {
			if l1.X == l2.X && l1.Y == l2.Y {
				m := map[string]Point{}
				m["l1"] = l1
				m["l2"] = l2
				matches = append(matches, m)
			}
		}
	}

	log.Printf("matches points %d  , %+v", len(matches), matches)

	distance := 0.0
	steps := 0

	for _, m := range matches {

		d := math.Abs(float64(m["l1"].X-0)) + math.Abs(float64(m["l1"].Y-0))

		if distance > d || distance == 0 {
			distance = d
			log.Printf("steps %d", steps)

		}

		cSteps := 0

		for k, v := range m {
			log.Printf("k: %s , steps: %d", k, v.Steps)
			cSteps += v.Steps
		}

		if steps > cSteps || steps == 0 {
			steps = cSteps
		}

	}

	return int(distance), steps, nil
}

func processMovement(direction string, value int, current Point) []Point {
	processPoints := []Point{}

	if direction == "U" {
		for i := 1; i <= value; i++ {
			current.Steps++
			processPoints = append(processPoints, Point{X: current.X, Y: current.Y + i, Steps: current.Steps})
		}
	}

	if direction == "D" {
		for i := 1; i <= value; i++ {
			current.Steps++
			processPoints = append(processPoints, Point{X: current.X, Y: current.Y - i, Steps: current.Steps})
		}
	}

	if direction == "L" {
		for i := 1; i <= value; i++ {
			current.Steps++
			processPoints = append(processPoints, Point{X: current.X - i, Y: current.Y, Steps: current.Steps})
		}
	}

	if direction == "R" {
		for i := 1; i <= value; i++ {
			current.Steps++
			processPoints = append(processPoints, Point{X: current.X + i, Y: current.Y, Steps: current.Steps})
		}
	}

	return processPoints
}
