package mmutil

import (
	"unicode/utf8"
)

const asat = "်"

func isSpace(n int) bool {
	return n == 32
}

func isConsonant(n int) bool {
	return (n >= 4096 && n <= 4138) || (n >= 4159 && n <= 4175)
}

func isStacked(n int) bool {
	return n == 4153
}

func isAsat(n int) bool {
	return n == 4154
}

func isVisarga(n int) bool {
	return n == 4152 || n == 4232 || n == 4234 || n == 4239 || n == 4252
}

func isValidCode(n int) bool {
	return (n >= 4096 && n <= 4255) || n == 32
}

// SplitWords break an input string written in myanmar unicode into individual words as an array
// TODO - kinzi words
// TODO - handle english characters
func SplitWords(s string) []string {
	var xs []string
	var cur string
	b := []byte(s)
	for len(b) > 0 {
		r, size := utf8.DecodeRune(b)
		if isValidCode(int(r)) {
			switch {
			// If the word is consonant (က, ခ, ဂ, etc) or a space, append cached string to the result,
			// and replace the cache with current word
			case isConsonant(int(r)):
				xs = append(xs, cur)
				cur = string(r)
			// If the word is an asat ( ် ), append asat to cached string with last element of the result
			// as an prefix and remove the last element from the result
			case isAsat(int(r)):
				cur = xs[len(xs)-1] + cur + string(r)
				xs = xs[:len(xs)-1]
				next, nextSize := utf8.DecodeRune(b[size:])
				if isVisarga(int(next)) {
					cur += string(next)
					b = b[nextSize:]
				}
			case isStacked(int(r)):
				cur = xs[len(xs)-1] + cur + asat
				xs = xs[:len(xs)-1]
			// Do nothing and append the word to the cached string
			default:
				cur += string(r)
			}
		}
		b = b[size:]
	}
	xs = append(xs, cur)
	return xs
}
