package tree

import (
	"fmt"
	"strconv"
	"strings"
)

func Binary(s string) string {
	input := strings.Split(s, " ")

	for w := 0; w < len(input); w++ {

		if input[w] == "(bin)" {
			value, _ := strconv.ParseInt(input[w-1], 2, 64)

			input[w-1] = fmt.Sprintf("%d", value)
			input = append(input[:w], input[w+1:]...)

		}
	}
	result := strings.Join(input, " ")
	return result

}
