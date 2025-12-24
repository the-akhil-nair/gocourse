package interrmediate

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
)

func main() {

	// Hashing is the transformation of data into a fixed-size string of characters, which is typically a sequence of numbers and letters.
	// It is a one-way function, meaning that once data is hashed, it cannot be easily reversed back to its original form.
	// Common hashing algorithms include MD5, SHA-1, SHA-256, and SHA-512.
	// Hashing is commonly used for data integrity verification, password storage, digital signatures, and more.
	// Same input will always produce the same output. Small change in input will produce a significantly different output and called avalanche effect.
	// Hashing is done for irreversible input i.e. one way function.
	// Hashing is fast and efficient.

	// Example: Hashing a password before storing it in a database.
	// When a user creates an account, their password is hashed and the hash is stored in the database.
	// When the user logs in, the entered password is hashed again and compared to the stored hash.
	// If they match, the user is authenticated.

	// Hash also tell if the file is altered or not and used for data integrity verification.
	// Hash functions generate fixed-size output (hash) for any input data. Since the information is lost during hashing, it is not possible to retrieve the original data from the hash.
	// Good hash functions minimize collisions, where different inputs produce the same hash output.

	// Sha256 and Sha512 are cryptographic hash functions that produce a fixed-size output (hash) for any input data.
	// That is 2^256 and 2^512 possible hash values respectively.
	// They are designed to be secure against various types of attacks, including collision attacks, pre-image attacks, and second pre-image attacks.
	// They are widely used in various security applications and protocols, including SSL/TLS, digital signatures, and password hashing.
	password := "password123"

	// hash := sha256.Sum256([]byte(password))
	// hash512 := sha512.Sum512([]byte(password))

	// fmt.Println(password)
	// fmt.Println(hash)
	// fmt.Println(hash512)
	// fmt.Printf("SHA-256 Hash hex val: %x\n", hash)
	// fmt.Printf("SHA-512 Hash hex val: %x\n", hash512)

	// Salting: Adding random data to the input of a hash function to ensure that the output (hash) is unique even for identical inputs.
	// It protects against rainbow table attacks, where attackers use precomputed tables of hash values to crack passwords.
	// Salting ensures that even if two users have the same password, their hashes will be different due to the unique salt added to each password before hashing.
	salt, err := generateSalt()
	fmt.Println("Original Salt:", salt)
	fmt.Printf("Original Salt: %x\n", salt)
	if err != nil {
		fmt.Println("Error generating salt:", err)
		return
	}

	// Hash the password with salt
	signUpHash := hashPassword(password, salt)

	// Store the salt and password in database, just printing as of now
	saltStr := base64.StdEncoding.EncodeToString(salt)
	fmt.Println("Salt:", saltStr)    // simulate as storing in database
	fmt.Println("Hash:", signUpHash) // simulate as storing in database
	hashOriginalPassword := sha256.Sum256([]byte(password))
	fmt.Println("Hash of just the password string without salt:", base64.StdEncoding.EncodeToString(hashOriginalPassword[:]))

	// verify
	// retrieve the saltStr and decode it
	decodedSalt, err := base64.StdEncoding.DecodeString(saltStr)
	if err != nil {
		fmt.Println("Unable to decode salt:", err)
		return
	}
	//loginHash := hashPassword("password124", decodedSalt)
	loginHash := hashPassword(password, decodedSalt)

	// Compare the stored signUpHash with loginHash
	// Each time the generated hash will be different due to different salt.
	if signUpHash == loginHash {
		fmt.Println("Password is correct. You are logged in.")
	} else {
		fmt.Println("Login failed. Please check user credentials.")
	}

}

func generateSalt() ([]byte, error) {
	salt := make([]byte, 32)
	// rand.Reader generates cryptographically secure random numbers.
	// ReadFull reads exactly len(salt) bytes from rand.Reader into salt and stop feeding when slice is full.
	// Cryptographihcally secure random numbers are important in security applications to ensure unpredictability and resistance to attacks.
	// Seeding is not required as they are automatically seeded from the operating system's source of randomness.	
	n, err := io.ReadFull(rand.Reader, salt)
	fmt.Println("Bytes read for salt:", n)
	if err != nil {
		return nil, err
	}
	return salt, nil
}

// Function to hash password
func hashPassword(password string, salt []byte) string {
	// var a []int = []int{1, 2, 3}
	// var b []int = []int{4, 5, 6}
	// a = append(a, b...)
	// fmt.Println(a)
	// Combine salt and password
	saltedPassword := append(salt, []byte(password)...)
	hash := sha256.Sum256(saltedPassword)
	return base64.StdEncoding.EncodeToString(hash[:])
}
