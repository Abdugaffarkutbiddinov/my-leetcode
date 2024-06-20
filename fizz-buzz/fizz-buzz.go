package fizz_buzz

import "strconv"

func fizzBuzz(n int) []string {
	ans := make([]string, n)
	for i := 0; i < n; i++ {
		switch {
		case (i+1)%3 == 0 && (i+1)%5 == 0:
			ans[i] += "FizzBuzz"
		case (i+1)%3 == 0:
			ans[i] += "Fizz"
		case (i+1)%5 == 0:
			ans[i] += "Buzz"
		default:
			ans[i] = strconv.Itoa(i + 1)
		}
	}
	return ans
}
