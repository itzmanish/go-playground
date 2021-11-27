package playground

import "fmt"

func IsBalanced(input string) bool {

	stack := []byte{}
	mapChar := map[byte]byte{
		byte(')'): byte('('),
		byte('}'): byte('{'),
		byte(']'): byte('['),
	}
	for i := 0; i < len(input); i++ {
		fmt.Println(stack, input[i])
		vv, ok := mapChar[input[i]]
		if ok {
			if vv == stack[len(stack)-1] {
				stack = stack[:len(stack)-1]
			}
		} else {
			stack = append(stack, input[i])
		}
	}

	return len(stack) == 0

}
