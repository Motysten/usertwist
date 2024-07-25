package cesar

import "unicode"

const ALPHA_LEN = 26

type Direction int

const (
	Left  Direction = -1
	Right           = 1
)

func Rotate(input string, dir Direction, key int) string {

	runes := []rune(input)

	normalizeKey := rune(key % 26)

	for index, current := range runes {
		if unicode.IsLetter(current) {
			var base rune
			if unicode.IsUpper(current) {
				base = 'A'
			} else {
				base = 'a'
			}

			switch dir {
			case Left:
				current = base + (current-base-normalizeKey+ALPHA_LEN)%ALPHA_LEN
			case Right:
				current = base + (current-base+normalizeKey)%ALPHA_LEN
			}
		}

		runes[index] = current
	}

	return string(runes)

}
