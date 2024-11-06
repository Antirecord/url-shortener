package main

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func main() {
	client := resty.New()
	users := []User{}
	_, err := client.R().
		SetResult(&users).
		Get("https://jsonplaceholder.typicode.com/users")

	if err != nil {
		fmt.Println("Error: ", err.Error())
	}
	for _, u := range users {
		fmt.Print(u.Username, " ")
	}
}
