package controllers

import (
	"sample-memory/models"

	"fmt"
	"strconv"
	"strings"

	"syreclabs.com/go/faker"
)

var users []models.User

func listUsers() {
	for _, item := range users {
		fmt.Println(item.ID, item.Name, item.LastName)
	}
}

func getUserByIndex() (models.User, int) {
	var index string
	_, err := fmt.Scanln(&index)
	if err != nil {
		return models.User{}, 0
	}
	i, _ := strconv.Atoi(index)
	if i > len(users) {
		fmt.Println("Username does not exist")
		return models.User{}, i
	}
	user := users[i]

	decorator := "*"
	title := "* Index user " + index + " *"
	fmt.Println(strings.Repeat(decorator, len(title)))
	fmt.Println(strings.ToUpper(title))
	fmt.Println(strings.Repeat(decorator, len(title)))
	fmt.Println(user.ID, user.Name, user.LastName)
	fmt.Println()
	return user, i

}
func removeUser() {
	var index string
	_, err := fmt.Scanln(&index)
	if err != nil {
		return
	}

	i, _ := strconv.Atoi(index)
	if i > len(users) {
		fmt.Println("Username does not exist")
		return
	}
	users = append(users[:i], users[i+1:]...)
	fmt.Println("user remove successfully")
}

func createUser() {
	var name, lastname string
	fmt.Println("Write the name:")
	_, err := fmt.Scanln(&name)
	if err != nil {
		return
	}

	fmt.Println("Write the lastname:")
	_, err = fmt.Scanln(&lastname)
	if err != nil {
		return
	}

	user := models.User{ID: faker.Number().NumberInt(5), Name: name, LastName: lastname}
	users = append(users, user)
	fmt.Println("user created successfully")

}

func updateUser() {
	_, index := getUserByIndex()
	var name, lastname string
	fmt.Println("Write the name:")
	_, err := fmt.Scanln(&name)
	if err != nil {
		return
	}

	fmt.Println("Write the lastname:")
	_, err = fmt.Scanln(&lastname)
	if err != nil {
		return
	}

	users[index].Name = name
	users[index].LastName = lastname

	fmt.Println("User updated successfully")

}

func FillInitialData(count uint) {
	for i := 0; i <= int(count); i++ {
		user := models.User{ID: faker.Number().NumberInt(5), Name: faker.Name().FirstName(), LastName: faker.Name().LastName()}
		users = append(users, user)
	}
}
