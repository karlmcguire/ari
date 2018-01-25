package ari

import (
	"math"
	"strings"
)

func ARI(text string) int {
	var (
		chars = float64(charCount(text))
		words = float64(strings.Count(text, " "))
		sente = float64(strings.Count(text, ".") + strings.Count(text, "!") +
			strings.Count(text, "?"),
		)
	)

	return int(
		math.Ceil((4.71 * (chars / words)) + (0.5 * (words / sente)) - (21.43)),
	)
}
