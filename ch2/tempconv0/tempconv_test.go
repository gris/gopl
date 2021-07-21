package tempconv

import (
	"fmt"
	"testing"
)

func TestExample_one(t *testing.T) {
	{
		fmt.Printf("%g\n", BoilingC-FreezingC) // "100" °C
		if (BoilingC - FreezingC) != 100 {
			t.Error("error")
		}
		boilingF := CToF(BoilingC)
		fmt.Printf("%g\n", boilingF-CToF(FreezingC)) // "180" °F
	}
}
func Example_two() {
	c := FToC(212.0)
	fmt.Println(c.String()) // "100°C"
	fmt.Printf("%v\n", c)   // "100°C"; no need to call String explicitly
	fmt.Printf("%s\n", c)   // "100°C"
	fmt.Println(c)          // "100°C"
	fmt.Printf("%g\n", c)   // "100"; does not call String
	fmt.Println(float64(c)) // "100"; does not call String
}
