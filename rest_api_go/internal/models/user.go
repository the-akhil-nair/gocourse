package models

type User struct {
	Name string `json:"name"`
	Age  uint16 `json:"age"`
	City string `json:"city"`
}
