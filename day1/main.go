package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
)

func main() {
	var (
		resWFuel  = 0.
		resWOFuel = 0.
	)

	stdin := os.Stdin
	r := bufio.NewReader(stdin)
	for {
		bytes, _, err := r.ReadLine()
		if err == io.EOF || len(bytes) == 0 {
			break
		}
		t, _ := strconv.ParseFloat(string(bytes), 10)
		resWOFuel += calcForMass(t)
		resWFuel += calcForMassWithFuel(t)
		// fmt.Printf("%s = %v\n", str, res)
	}

	fmt.Printf("For mass only = %.f\n", resWOFuel)
	fmt.Printf("For mass with fuel = %.f\n", resWFuel)
}

func calcForMass(mass float64) float64 {
	return math.Floor(mass/3) - 2
}

func calcForMassWithFuel(mass float64) float64 {
	res := math.Floor(mass/3) - 2
	// fmt.Printf("[DEBUG] For mass %f fuel %.f\n", mass, res)
	if res <= 0 {
		return 0
	}
	return res + calcForMassWithFuel(res)
}
