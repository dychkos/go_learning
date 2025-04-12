package main

import "fmt"

type User struct {
	name  string
	age   int
	city  string
	phone string
}

func NewUser(name string, age int, city string, phone string) User {
	return User{
		name:  name,
		age:   age,
		city:  city,
		phone: phone,
	}
}

func (u User) Introduce() {
	fmt.Printf(
		"Привіт, мене звати %s. Мені %d років, я з %s. Мій номер для звяку: %s. \n",
		u.name, u.age, u.city, u.phone,
	)
}

func main() {
	var user1 User

	user1 = User{
		name:  "Serhii",
		age:   22,
		city:  "Kiyv",
		phone: "4042536",
	}

	user1.Introduce()

	var user2 User

	user2 = NewUser("Diana", 21, "Snihuribka", "32423")

	user2.Introduce()
}
