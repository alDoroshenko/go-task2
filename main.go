package main

import (
	"fmt"
	"log"
)

type User struct {
	name string
	age  int
}

func (u *User) Age() int {
	return u.age
}

func NewUser(name string, age int) *User {
	return &User{name: name, age: age}
}

func main() {
	source := inUsers()
	lim := inLimit()
	target := filter(source, lim)
	fmt.Println("Source: ", source)
	fmt.Println("Target: ", target)

}

func inUsers() []User {
	var name string
	var age int
	var users []User
	fmt.Println("Введите имя и возраст пользователя через пробел, по окончании ввода введите exit:  ")
	for {
		_, err := fmt.Scanln(&name, &age)
		if name == "exit" {
			break
		}
		if err != nil {
			log.Fatal("Ошибка при вводе")
		}
		users = append(users, *NewUser(name, age))

	}
	return users
}

func inLimit() int {
	var val int
	fmt.Print("Введите лимит возраста, по которому будет фильтроваться список: ")
	_, err := fmt.Scan(&val)
	if err != nil {
		log.Fatal("Ошибка при вводе")
	}
	if val == 0 {
		log.Fatal("Введено некорректное значение")
	}
	return val
}

func filter(users []User, lim int) []User {
	var res []User
	for _, user := range users {
		if user.Age() > lim {
			res = append(res, user)
		}
	}
	return res
}
