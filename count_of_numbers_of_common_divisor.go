// Given two number find the number of common divisor of both

package playground

import (
	"fmt"
	"math"
)

func CountNoOfCommonDivisor(a, b int) {
	n := gcd(a, b)
	count := 0
	for i := 1; i <= (int(math.Sqrt(float64(n))) + 1); i++ {
		if n%i == 0 {
			if n/i == i {
				count += 1
			} else {
				count += 2
			}
		}
	}
	fmt.Println(count)
}

func gcd(a, b int) int {
	if a < b {
		a, b = b, a
	}
	if a == 0 {
		return b
	}
	return gcd(b%a, a)
}
