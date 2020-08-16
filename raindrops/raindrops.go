package raindrops

import (
	"strconv"
)

func dividesEvenly(dividend int, divisor int) bool {
	return dividend%divisor == 0
}

// Convert generates realistic raindrop sounds for a given int
func Convert(raindrops int) string {
	var sounds = ""
	if dividesEvenly(raindrops, 3) {
		sounds += "Pling"
	}
	if dividesEvenly(raindrops, 5) {
		sounds += "Plang"
	}
	if dividesEvenly(raindrops, 7) {
		sounds += "Plong"
	}
	if sounds == "" {
		sounds = strconv.Itoa(raindrops)
	}
	return sounds
}
