package reverse_string

import "fmt"

func ReverseString(input string) (output string) {
	bts := []rune(input)
	fmt.Println(bts)

	for _, v := range bts {
		fmt.Println(v)
		output = string(v) + output
	}

	return output
}
