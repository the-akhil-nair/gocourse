package intermediate

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {

	str := "Hello Go!"
	// str10 := `Hello Go. Go World!`

	fmt.Println(len(str))

	// str1 := "Hello"
	// str2 := "World"
	// result := str1 + " " + str2
	// fmt.Println(result)

	// // Printing character will result in rune.
	// // Printing the slice of characters will result in string.
	// fmt.Println(str[0])
	// fmt.Println(str[1:7])

	// String Conversion
	// num := 18
	// str3 := strconv.Itoa(num)
	// fmt.Println(len(str3))

	// // strings splitting
	// fruits := "apple, orange, banana"
	// fruits1 := "apple-orange-banana"
	// parts := strings.Split(fruits, ",")
	// parts1 := strings.Split(fruits1, "-")
	// fmt.Println(parts)
	// fmt.Println(parts1)

	// countries := []string{"Germany", "France", "Italy"}
	// joined := strings.Join(countries, ", ")
	// fmt.Println(joined)

	// fmt.Println(strings.Contains(str, "Go?"))

	// replaced := strings.Replace(str, "Go", "Universe", 1)
	// replaced1 := strings.ReplaceAll(str10, "Go", "Universe")
	// fmt.Println(replaced)
	// fmt.Println(replaced1)

	// strwspace := " Hello Everyone! "
	// fmt.Println(strwspace)
	// fmt.Println(strings.TrimSpace(strwspace))

	// fmt.Println(strings.ToLower(strwspace))
	// fmt.Println(strings.ToUpper(strwspace))

	// fmt.Println(strings.Repeat("foo ", 3))

	// fmt.Println(strings.Count("Hello", "l"))
	// fmt.Println(strings.Count("Hello", "ll"))
	// fmt.Println(strings.HasPrefix("Hello", "he"))
	// fmt.Println(strings.HasSuffix("Hello", "lo"))
	// fmt.Println(strings.HasSuffix("Hello", "la"))

	// str5 := "Hel1lo, 123 Go 11!"
	// re := regexp.MustCompile(`\d+`)
	// // -1 means find all matches
	// matches := re.FindAllString(str5, -1)
	// fmt.Printf("Found %d matches: ", len(matches))
	// fmt.Printf("%v\n", matches)

	str6 := "Hello, 世界"
	// Able to count number of runes in a string which may contain multi-byte characters.
	fmt.Println(utf8.RuneCountInString(str6))

	// STRING BUILDER
	// More efficient for concatenating multiple strings and have better performance in terms of memory allocation.
	// Also prevents memory leaks and also memory efficient.
	var builder strings.Builder

	// Write some strings
	builder.WriteString("Hello")
	builder.WriteString(", ")
	builder.WriteString("world!")

	// Convert builder to a string
	result := builder.String()
	fmt.Println(result)

	// Using Writerune to add a character
	builder.WriteRune(' ')
	builder.WriteString("How are you")

	result = builder.String()
	fmt.Println(result)

	// Reset the builder
	builder.Reset()
	builder.WriteString("Starting fresh!")
	result = builder.String()
	fmt.Println(result)
}
