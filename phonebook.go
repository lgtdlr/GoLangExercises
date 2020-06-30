package main

import (
	"fmt"
)

type Contact struct {
	name        string
	surname     string
	phoneNumber string
}

func main() {

	larry := Contact{"Larry", "Fisher", "2049485772"}
	billy := Contact{"Billy", "Raider", "2015439401"}
	william := Contact{"William", "Ar", "7870924441"}

	//phoneBook := []Contact{larry, billy, william}

	phoneBookMap := map[string]Contact{larry.name: larry, billy.name: billy, william.name: william}

	phoneBookMap2 := map[string]string{larry.name: larry.phoneNumber, billy.name: billy.phoneNumber, william.name: william.phoneNumber}
	for {
		fmt.Println("Enter 1 to add a new contact, 2 to remove a contact, 3 to search for a contact's number by name")
		//reader := bufio.NewReader(os.Stdin)
		var input string
		var name string
		var surname string
		var phoneNumber string

		fmt.Scanln(&input)

		switch input {
		case "1":
			fmt.Println("Adding\n ")
			fmt.Println("Enter name: ")
			fmt.Scanln(&name)
			fmt.Println("Enter surname: ")
			fmt.Scanln(&surname)
			fmt.Println("Enter phone number: ")
			fmt.Scanln(&phoneNumber)
			newContact := Contact{name, surname, phoneNumber}
			phoneBookMap[name] = newContact
			phoneBookMap2[name] = phoneNumber

		case "2":
			fmt.Println(phoneBookMap2)
			fmt.Println("Enter name of person to be deleted ")
			fmt.Scanln(&name)
			delete(phoneBookMap2, name)
			fmt.Println(phoneBookMap2)

		case "3":
			fmt.Println("Searching\n ")
			fmt.Println("Enter name: ")
			fmt.Scanln(&name)
			fmt.Println(phoneBookMap2[name])
		}
	}

}
