package intermediate

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	user := os.Getenv("USER")
	home := os.Getenv("HOME")

	fmt.Println("User env var:", user)
	fmt.Println("Home env var:", home)

	err := os.Setenv("FRUIT", "APPLE")
	if err != nil {
		fmt.Println("Error setting environment variable:", err)
	}
	fmt.Println("Environment variable FRUIT set successfully.")

	fmt.Println("FRUIT env var:", os.Getenv("FRUIT"))

	for _, e := range os.Environ() {
		// int is how mant fragments we want after splitting
		// SplitN is useful here because we do not want to split further on "=" in case the value contains "="
		kvpair := strings.SplitN(e, "=", 2)
		fmt.Println(kvpair)
	}

	err = os.Unsetenv("FRUIT")
	if err != nil {
		fmt.Println("Error unsetting environment variable:", err)
		return
	}
	fmt.Println("Unset env var done on key FRUIT")

	fmt.Println("FRUIT env var:", os.Getenv("FRUIT"))
	fmt.Println("---------------------------------------------------")

	str := "a=b=c=d=e"
	// n = 1 returns "a=b=c=d"
	// n = 2 returns "a" and "b=c=d"
	// n = 3 returns "a" and "b" and "c=d"
	// n = 4 returns "a" and "b" and "c" and "d"
	fmt.Println(strings.SplitN(str, "=", -1))
	fmt.Println(strings.SplitN(str, "=", 0))
	fmt.Println(strings.SplitN(str, "=", 1))
	fmt.Println(strings.SplitN(str, "=", 2))
	fmt.Println(strings.SplitN(str, "=", 3))
	fmt.Println(strings.SplitN(str, "=", 4))
	fmt.Println(strings.SplitN(str, "=", 5))

}
