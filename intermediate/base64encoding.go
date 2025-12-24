package intermediate

import (
	"encoding/base64"
	"fmt"
)

func main() {

	data := []byte("He~lo, Base64 Encoding")

	// Encode to Base64
	// Binary to text encoding scheme
	// It represents binary data in an ASCII string format by translating it into a radix-64 representation.
	// Binary data is divided into chunks of 6 bits, each mapped to a specific character in the Base64 alphabet.
	// Used in email via MIME, storing complex data in XML or JSON, embedding media in web pages, etc.
	// Base64 encoding is not cryptographically secure, as it can be easily decoded back to its original form.
	encoded := base64.StdEncoding.EncodeToString(data)
	// Same output for same input.
	fmt.Println(encoded)

	// Decode from Base64
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		fmt.Println("Error in decoding:", err)
		return
	}

	// Convert byte slice to string for display as DecodeString returns byte slice.
	fmt.Println("Decoded:", string(decoded))

	// URL safe, avoid '/' and '+'

	urlSafeEncoded := base64.URLEncoding.EncodeToString(data)
	fmt.Println("URL Safe encoded:", urlSafeEncoded)

}
