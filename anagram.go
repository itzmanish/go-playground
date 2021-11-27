package playground

import (
	"strings"
)

// func main() {
// given 2 strings as input, string1 and string2. Find anagram of string2 within string1
// if there's no possible anagram, print "no possible anagram"
// if there's at least one possibility, print array of index of integer
// The integers within the array are the index of string1 which is the starting position of the anagram of string2

// example:
// string1 := "abcdddd"
// string2 := "abcd"
// output: [0]

// string1 := "aba"
// string2 := "ab"
// output: [0,1]

//string1 := "abcdddd"
//string2 := "abcd"
// 	string1 := "dabcddd"
// 	string2 := "abcd"
// 	out := FindAnagram(string1, string2)
// 	fmt.Println(out)
// }

func FindAnagram(s1 string, s2 string) []int {
	// code here
	// First I will check the exact string2 within string. Because that's the one anagram I will get.

	// we need to also check for the first rune of s1 from where first rune of s2 starts matching
	//func checkAnagram(s1 string, s2 string){
	//if len(s2)=0{
	//return
	//}
	// Need to cover string2 with different order also.
	// For that one possible solution can be generating a list of possible combination from letters of s2
	// and then try to match them by below for loop.
	// If I do that then I won't need to recursively call this method.

	// var output []int
	// var matchedCount, startMatchedIndex int
	// var startedMatching bool

	// for i := 0; i < len(s2); i++ {
	// 	for j := i; j < len(s1); j++ {
	// 		// Trigger comparing s1[j]==s2[i]
	// 		// fmt.Println(s1[j], s2[i])
	// 		if s1[j] == s2[i] {
	// 			matchedCount = matchedCount + 1

	// 			if !startedMatching {
	// 				startMatchedIndex = j
	// 				startedMatching = true
	// 			}
	// 			break
	// 		}

	// if we got match we iterate next word of s1 to match next word of s2
	// i = 0 ; j=0; s2[0] = a ; s1[0]= a; matchedCount = 1, startMatchedIndex=j
	// return here that will increament i to 1
	// i = 1; j=1; s2[1] = b; s1[1] = b; matchedCount =2
	// ....
	// i = 3; j = 3; still match; matchedCount = 4
	// It will exit here
	// if matchedCount=len(s2) then return the first index from where it started matching

	// 	}
	// }
	// if matchedCount == len(s2) {
	// 	output = append(output, startMatchedIndex)
	// }

	// fmt.Println(output)
	// Let's check if recursive calling checkAnagram with len(s2)-1 covers some cases or not.
	//checkAnagram(s1, s2[:len(s2)-1])

	//}

	//func generatePossibleCombinationFromString(s string) []string{
	//	 ss := []string{}

	//
	//	}

	//---------------------------------------------------------------------------------------------------------------------------
	// The above code is what I did on interview. The below one is after the interview
	//---------------------------------------------------------------------------------------------------------------------------
	// Start Time - 2:12 PM
	// End Time - 1st success result - 2:56 PM
	// Logic
	// 1. generate all possible combination from string2
	// 2. check if string 1 contains substring from combination of string 2

	generateCombination := func(s string) []string {
		var result []string
		Perm([]rune(s2), func(r []rune) {
			result = append(result, string(r))
		})
		return result
	}

	getIndexOfSubstring := func(s1 string, values []string) []int {
		var output []int
		for _, v := range values {
			vv := strings.Index(s1, v)
			if vv != -1 {
				output = append(output, vv)
			}
		}
		return output
	}

	possibleStrings2 := generateCombination(s2)

	return getIndexOfSubstring(s1, possibleStrings2)

}

//  Reference := https://yourbasic.org/golang/generate-permutation-slice-string/
// Perm calls f with each permutation of a.
func Perm(a []rune, f func([]rune)) {
	perm(a, f, 0)
}

// Permute the values at index i to len(a)-1.
func perm(a []rune, f func([]rune), i int) {
	if i > len(a) {
		f(a)
		return
	}
	perm(a, f, i+1)
	for j := i + 1; j < len(a); j++ {
		a[i], a[j] = a[j], a[i]
		perm(a, f, i+1)
		a[i], a[j] = a[j], a[i]
	}
}
