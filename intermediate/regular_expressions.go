package intermediate

import (
	"fmt"
	"regexp"
)

func main() {

	fmt.Println("He said, \"I am great\"")
	fmt.Println(`He said, "I am great"`)

	// Compile a regex pattern to match email address
	// re := regexp.MustCompile(`[a-zA-Z0-9._+%-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}`)
	re := regexp.MustCompile(`^\w+@\w+\.[a-zA-Z]{2,}$`)

	// Test strings
	email1 := "user@email.com"
	email2 := "invalid_email23@g.com"

	// Match
	fmt.Println("Email1:", re.MatchString(email1))
	fmt.Println("Email2:", re.MatchString(email2))

	// Capturing Groups
	// Compile a regex pattern to capture date components
	re = regexp.MustCompile(`(\d{4})-(\d{2})-(\d{2})`)

	// Test string
	date := "2024-07-30"

	// Find all submatches
	submatches := re.FindStringSubmatch(date)
	// Will print the full match and each capturing group, simlar to groups in python
	fmt.Println(submatches)
	// Accessing individual groups, similar to group(0), group(1) etc in python
	fmt.Println(submatches[0])
	fmt.Println(submatches[1])
	fmt.Println(submatches[2])
	fmt.Println(submatches[3])

	// Source string
	str := "Hello World"

	re = regexp.MustCompile(`[aeiou]`)

	result := re.ReplaceAllString(str, "*")
	fmt.Println(result)

	// i - case insensitive
	// m - multi line model
	// s - dot matches all

	re = regexp.MustCompile(`(?i)go`)
	// re = regexp.MustCompile(`go`)

	// Test string
	text := "Golang is great"

	// Match
	fmt.Println("Match:", re.MatchString(text))
}
