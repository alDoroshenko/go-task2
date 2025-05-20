package main

import (
	"fmt"
	"log"
	"sort"
)

type User struct {
	name string
	age  int
}

func (u *User) Age() int {
	return u.age
}
func (u *User) Name() string {
	return u.name
}

func NewUser(name string, age int) *User {
	return &User{name: name, age: age}
}

func main() {
	source := inUsers()
	lim := inLimit()
	targetBigger, targetLess := filter(source, lim)
	fmt.Println("Source: ", source)
	fmt.Println("TargetBigger: ", targetBigger)
	fmt.Println("targetLess: ", targetLess)

}

func inUsers() []User {
	var name string
	var age int
	users := make([]User, 0, 10)
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

func filter(users []User, lim int) ([]User, []User) {
	resLess := make([]User, 0, len(users))
	resBigger := make([]User, 0, len(users))

	for _, user := range users {
		if user.Age() > lim {
			resBigger = append(resBigger, user)
		} else {
			resLess = append(resLess, user)
		}
	}
	sort.Slice(resBigger, func(i, j int) bool {
		return resBigger[i].Name() < resBigger[j].Name()
	})
	sort.Slice(resLess, func(i, j int) bool {
		return resLess[i].Name() < resLess[j].Name()
	})
	return resBigger, resLess
}
