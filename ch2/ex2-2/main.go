// Converte seu argumento numérico para diferentes unidades de medida
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"gopl.io/ch2/ex2-2/unitconv"
	"gopl.io/ch2/tempconv"
)

func main() {
	if len(os.Args[1:]) == 0 {
		input := bufio.NewScanner(os.Stdin)
		for input.Scan() {
			unitConv(input.Text())
		}
	} else {
		for _, arg := range os.Args[1:] {
			unitConv(arg)
		}

	}
}

func unitConv(arg string) {
	t, err := strconv.ParseFloat(arg, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "cf: %v\n", err)
		os.Exit(1)
	}
	f := tempconv.Fahrenheit(t)
	c := tempconv.Celsius(t)
	m := unitconv.Meter(t)
	fe := unitconv.Feet(t)
	k := unitconv.Kilogram(t)
	p := unitconv.Pound(t)

	fmt.Printf("%s = %s, %s = %s\n", f, tempconv.FToC(f), c, tempconv.CToF(c))
	fmt.Printf("%s = %s, %s = %s\n", m, unitconv.MeterToFeet(m), fe, unitconv.FeetToMeter(fe))
	fmt.Printf("%s = %s, %s = %s\n", k, unitconv.KilogramToPound(k), p, unitconv.PoundToKilogram(p))
}
