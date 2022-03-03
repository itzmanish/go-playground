// Fredo and Array Update
// Max. score: 100
// Fredo is assigned a new task today. He is given an array A containing N integers.
// His task is to update all elements of array to some minimum value x , that is,  ;
// such that sum of this new array is strictly greater than the sum of the initial array.
// Note that x should be as minimum as possible such that sum of the new array is greater than
// the sum of the initial array.

// Input Format:
// First line of input consists of an integer N denoting the number of elements in the array A.
// Second line consists of N space separated integers denoting the array elements.

// Output Format:
// The only line of output consists of the value of x.

// Input Constraints:

// SAMPLE INPUT
// 5
// 1 2 3 4 5
// SAMPLE OUTPUT
// 4
// Explanation
// Initial sum of array
// When we update all elements to 4, sum of array  which is greater than .
// Note that if we had updated the array elements to 3,  which is not greater than .
//  So, 4 is the minimum value to which array elements need to be updated.

// Time Limit:	1.0 sec(s) for each input file.
// Memory Limit:	256 MB
// Source Limit:	1024 KB
// Marking Scheme:	Score is assigned if any testcase passes.
// Allowed Languages:	C, C++, Clojure, C#, D, Erlang, F#, Go, Groovy, Haskell, Java, Java 8, JavaScript(Rhino),
// JavaScript(Node.js), Lisp, Lisp (SBCL), Lua, Objective-C, OCaml, Octave, Pascal, Perl, PHP, Python, Python 3,
//  R(RScript), Racket, Ruby, Rust, Scala, Swift, Visual Basic

package playground

import (
	"sort"
)

func CalculateFredoArrayUpdate(array []int) int {
	sort.Ints(array)
	mid := 0
	if len(array)%2 == 0 {
		mid = len(array) / 2
	} else {
		mid = (len(array) + 1) / 2
	}
	sum := sumArr(array)
	for i := mid - 1; i < len(array); i++ {
		replacableValue := array[i]
		newSum := len(array) * replacableValue
		if newSum > sum {
			return replacableValue
		}
	}
	return array[mid-1]
}

func sumArr(arr []int) int {
	sumArr := 0
	for j := 0; j < len(arr); j++ {
		sumArr += arr[j]
	}
	return sumArr
}

func replaceArrWithX(arr []int, x int) []int {
	copyArr := make([]int, len(arr))
	copy(copyArr, arr)
	for i := 0; i < len(arr); i++ {
		copyArr[i] = x
	}
	return copyArr
}
