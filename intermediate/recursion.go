package intermediate

import "fmt"

var FibMap = make(map[int]uint64)

func main() {

	fmt.Println(factorial(5))
	fmt.Println(factorial(10))

	fmt.Println(sumOfDigits(9))
	fmt.Println(sumOfDigits(12))
	fmt.Println(sumOfDigits(12345))

	fmt.Println(fibonacii(11))
	fmt.Println(fibonacii(7))
	fmt.Println(fibonacii(20))

}

func factorial(n int) int {
	// Base case: factorial of 0 is 1
	if n == 0 {
		return 1
	}
	// Recursive case: factorial of n is n * factorial(n - 1)
	return n * factorial(n-1)
	// n * (n - 1) * (n-2) * factorial (n-3)..... factorial(0)

}

func sumOfDigits(n int) int {
	// Base case i.e. 1..10 is 1..10 only as it reached last digit.
	if n < 10 {
		return n
	}
	// 2000 then 2000%10=0 + sumOfDigits(200)
	// 200 then 200%10=0 + sumOfDigits(20)
	// 20 then 20%10=0 + sumOfDigits(2)
	// return as its base case
	return n%10 + sumOfDigits(n/10)
}

func fibonacii(i int) uint64 {
	v, ok := FibMap[i]
	if ok {
		fmt.Println(i, v)
		return v
	}
	// Base case i.e 1
	if i == 1 {
		return 1
	} else if i <= 0 {
		return 0
	}

	// If you pass i instead of i-1 then base condition will never meet and recursion will go on.
	fib := fibonacii(i-1) + fibonacii(i-2)

	v, ok = FibMap[i]
	if !ok {
		FibMap[i] = fib
	}
	return fib
}
