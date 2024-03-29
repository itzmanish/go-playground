package playground

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindAnagram(t *testing.T) {
	testcases := []struct {
		string1 string
		string2 string
		output  []int
	}{
		{
			string1: "abcdddd",
			string2: "abcd",
			output:  []int{0},
		},
		{
			string1: "aba",
			string2: "ab",
			output:  []int{0, 1},
		},
		{
			string1: "dabcddd",
			string2: "abcd",
			output:  []int{1, 0},
		},
	}
	for i, test := range testcases {
		t.Run(fmt.Sprintf("Testcase: %d", i), func(t *testing.T) {
			out := FindAnagram(test.string1, test.string2)
			assert.Equal(t, out, test.output)
		})
	}
}

func TestIsBalancedBrackets(t *testing.T) {
	input := "({})[()]"
	assert.True(t, IsBalanced(input))
}

func TestCalculateFredoArrayUpdate(t *testing.T) {
	arr := []int{65, 83, 57, 4, 26, 74, 99, 28, 41, 007, 578}
	expected := 4
	result := CalculateFredoArrayUpdate(arr)
	assert.Equal(t, expected, result)
}
