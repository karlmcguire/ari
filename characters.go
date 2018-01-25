package ari

var characters map[byte]struct{} = map[byte]struct{}{
	'a': struct{}{}, 'A': struct{}{},
	'b': struct{}{}, 'B': struct{}{},
	'c': struct{}{}, 'C': struct{}{},
	'd': struct{}{}, 'D': struct{}{},
	'e': struct{}{}, 'E': struct{}{},
	'f': struct{}{}, 'F': struct{}{},
	'g': struct{}{}, 'G': struct{}{},
	'h': struct{}{}, 'H': struct{}{},
	'i': struct{}{}, 'I': struct{}{},
	'j': struct{}{}, 'J': struct{}{},
	'k': struct{}{}, 'K': struct{}{},
	'l': struct{}{}, 'L': struct{}{},
	'm': struct{}{}, 'M': struct{}{},
	'n': struct{}{}, 'N': struct{}{},
	'o': struct{}{}, 'O': struct{}{},
	'p': struct{}{}, 'P': struct{}{},
	'q': struct{}{}, 'Q': struct{}{},
	'r': struct{}{}, 'R': struct{}{},
	's': struct{}{}, 'S': struct{}{},
	't': struct{}{}, 'T': struct{}{},
	'u': struct{}{}, 'U': struct{}{},
	'v': struct{}{}, 'V': struct{}{},
	'w': struct{}{}, 'W': struct{}{},
	'x': struct{}{}, 'X': struct{}{},
	'y': struct{}{}, 'Y': struct{}{},
	'z': struct{}{}, 'Z': struct{}{},
	'0': struct{}{},
	'1': struct{}{},
	'2': struct{}{},
	'3': struct{}{},
	'4': struct{}{},
	'5': struct{}{},
	'6': struct{}{},
	'7': struct{}{},
	'8': struct{}{},
	'9': struct{}{},
}

func charCount(text string) int {
	var count int

	for _, b := range []byte(text) {
		if _, valid := characters[b]; valid {
			count++
		}
	}

	return count
}
