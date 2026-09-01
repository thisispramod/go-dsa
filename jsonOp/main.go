package main

import (
	"encoding/json"
	"fmt"
)

// {
// 	"user_id" : 1,
// 	"name" : "Pramod",
// 	"age" : 26,
// 	"address" : "Aama Mafee Post Pachhar Bazar ",
// 	"password" : "password",
// 	"role" :["admin","author","moderator"]
// }
/*
type User struct {
	ID          int      `json:"user_id"`
	Name        string   `json:"-"`
	Age         int      `json: "age"`
	Address     string   `json: "address"`
	Password    string   `json:"-"`
	Permissions []string `json: "role"`
}
*/

type Person struct {
	Name       string `json:"full_name"`
	Age        int    `json:"age"`
	Occupation string `json:"-"`
	Language   string `json:"language"`
}

func main() {
	/*
		u := User{
			ID:          1,
			Name:        "Pramod",
			Age:         20,
			Address:     "aama mafe post pachhar bazar",
			Password:    "my-password",
			Permissions: []string{"admin", "group-member"},
		}

		users, err := json.Marshal(u)

		if err != nil {
			fmt.Println("erro marshalling json:", err)
			panic(err)
		}

		fmt.Println(string(users))
	*/
	// now unmarshalling how its work

	jsonData := `{"full_name":"Pramod", "age":27, "language":"student"}`

	var person Person
	err := json.Unmarshal([]byte(jsonData), &person)

	if err != nil {
		fmt.Println("something is wrong")
		panic(err)
	}

	fmt.Println(person.Name)
	fmt.Println(person.Language)
	fmt.Println(person.Age)

	fmt.Println(person.Occupation)

}
