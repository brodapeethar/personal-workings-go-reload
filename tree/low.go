package tree

import "strings"

func LOW(s string) string {
	input := strings.Split(s, " ")
	for i := 0; i < len(input); i++ {

		if input[i] == "(low)" && i > 0 {
			input = append(input[:i], input[i+1:]...)
			i--
		}
		input[i] = strings.ToLower(input[i])
	}
	result := strings.Join(input, " ")
	return result
}