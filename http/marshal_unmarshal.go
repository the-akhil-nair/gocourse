package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	user := User{
		Name:  "John Doe",
		Email: "john.doe@example.com",
	}

	marshal, err := json.MarshalIndent(user, "", "\t")

	if err != nil {
		fmt.Println("Unable to Marshal")
	}

	fmt.Println(string(marshal))

	err = json.Unmarshal(marshal, &user)

	if err != nil {
		fmt.Println("Unable to Unmarshal")
	}

	fmt.Println(user)

	jsonData := `{"name": "Jack Daniels", "email": "jack.daniels@test.com"}`
	reader := strings.NewReader(jsonData)

	// Decode decodes json to struct.
	jsonDecoder := json.NewDecoder(reader)

	var user1 User
	err = jsonDecoder.Decode(&user1)
	if err != nil {
		fmt.Println("Unable to decode.")
	}

	fmt.Println(user1)

	var buff bytes.Buffer
	enconder := json.NewEncoder(&buff)

	err = enconder.Encode(User{
		Name:  "Alister Mac",
		Email: "alister.mac@test.com",
	})

	if err != nil {
		fmt.Println("Unable to Encode.")
	}

	fmt.Println(buff.String())

}
