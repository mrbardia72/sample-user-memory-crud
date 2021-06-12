package controllers

import (
	"fmt"
	"strings"
)

func PrintMenu() {
	fmt.Println(strings.ToUpper("0- remove user"))
	fmt.Println(strings.ToUpper("1- listUsers"))
	fmt.Println(strings.ToUpper("2- getUserByIndex"))
	fmt.Println(strings.ToUpper("3- createUser"))
	fmt.Println(strings.ToUpper("4- updateUser"))
	fmt.Println(strings.ToUpper("5- exit"))
	doTask()
}

func doTask() {
	var option string
	_, err := fmt.Scanln(&option)
	if err != nil {
		return
	}

	switch option {
	case "0":
		removeUser()
		PrintMenu()
	case "1":
		listUsers()
		PrintMenu()
	case "2":
		getUserByIndex()
		PrintMenu()

	case "3":
		createUser()
		PrintMenu()
	case "4":
		updateUser()
		PrintMenu()
	case "5":
		break
	default:
		fmt.Println("Select an option from the list")
		PrintMenu()
	}
}

func PrintHeader() {
	decorator := "*"
	title := "* user memory crud *"
	fmt.Println(strings.Repeat(decorator, len(title)))
	fmt.Println(strings.ToUpper(title))
	fmt.Println(strings.Repeat(decorator, len(title)))
}
