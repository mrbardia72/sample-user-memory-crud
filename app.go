package main
import (
	"fmt"
	"strings"
)

func init() {
	fmt.Println(strings.ToUpper("0- remove user"))
	fmt.Println(strings.ToUpper("1- listUsers"))
	fmt.Println(strings.ToUpper("2- getUserByIndex"))
	fmt.Println(strings.ToUpper("3- createUser"))
	fmt.Println(strings.ToUpper("4- updateUser"))
	fmt.Println(strings.ToUpper("5- exit"))
	//doTask()
}

func main() {
	var option string
	_, err := fmt.Scanln(&option)
	if err != nil {
		return
	}

	switch option {
	case "0":
		fmt.Println("00000000")
	case "1":
		fmt.Println("1111111111")
	case "2":
		fmt.Println("2222222222")

	case "3":
		break
	default:
		fmt.Println("Select an option from the list")
		//PrintMenu()
	}
}
