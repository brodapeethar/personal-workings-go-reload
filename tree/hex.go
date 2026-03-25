package tree

import (
	"fmt"
	"strconv"
	"strings"
)

func Hexadecimal(s string) string {
	input := strings.Split(s, " ")

	for p := 0; p < len(input); p++ {

		if input[p] == "(hex)" {
			value, _ := strconv.ParseInt(input[p-1], 16, 64)

			input[p-1] = fmt.Sprintf("%d", value)
			input = append(input[:p], input[p+1:]...)
		}
	}
	result := strings.Join(input, " ")
	return result
}
