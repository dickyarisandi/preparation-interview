package main

import (
	"fmt"
	"math"
	"strings"
)

// Given variable string contains of all alphabet from A to Z
// ABCDEFGHIJKLMNOPQRSTUVWXYZ

// Build a function that receive two paramters: first letter and last letter.
// That will do to find middle letter between of the specified letter.
// Example:

// The middle between Q and U is S
// The middle between R and U is ST
// The middle between T and Z is W
// The middle between Q and Z is UV

func main() {
	alphabets := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	firstAlphabet := "Y"
	lastAlphabet := "Z"

	fmt.Println(middleAlphabet(alphabets, firstAlphabet, lastAlphabet))
}

func middleAlphabet(alphabets, firstAlphabet, lastAlphabet string) string {
	first := float64(strings.Index(alphabets, firstAlphabet))
	last := float64(strings.Index(alphabets, lastAlphabet))

	max := math.Max(first, last)
	min := math.Min(first, last)
	center := max - min
	alpha := func(pos float64) string {
		return string(alphabets[int(pos)])
	}

	if center == 0 {
		return alpha(max)
	}

	modulo := int(center) % 2
	devide := math.Floor(center / 2)
	if modulo == 0 {
		return alpha(min + devide)
	}

	return alpha(min+devide) + alpha(min+devide+1)
}
