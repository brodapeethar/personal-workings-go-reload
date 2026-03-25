package tree

import (
	"strconv"
	"strings"
)

func Capitalize(text string) string {
	input := strings.Split(text, " ")
	for i := 0; i < len(input); i++ {
		if input[i] == "(cap)" {
			input[i-1] = TittleCase(input[i-1])
			input = append(input[:i], input[i+1:]...)

		}

		if input[i] == "(cap," {
			word := strings.TrimRight(input[i+1], ")")
			value, _ := strconv.Atoi(word)
			input = append(input[:i], input[i+2:]...)

			for j := 0; j <= value; j++ {
				input[i-j] = TittleCase(input[i-j])
			}
		}
	}
	result := strings.Join(input, " ")
	return result
}

func TittleCase(text string) string {
	return strings.ToUpper(text[:1]) + text[1:]
}
