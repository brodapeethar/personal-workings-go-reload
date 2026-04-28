package tree

import "strings"

func IsVowel(nextword string) string {
	for i := 0; i < len(nextword) && strings.ContainsAny(string(nextword[0]), "aeiouhAEIOUH"); i++ {

		return "an"
	}
	return "a"
}
