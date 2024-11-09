package main

import (
	"fmt"
	"strings"
)

var conferenceName string = "Go Conference"

const conferenceTickets = 50

var remainingTickets uint = 50 // in our case remaining tickets value never be negative, so assigned uint
var booking []string

var firstName string
var lastName string
var email string
var userTickets int

func main() {
	getGreeting()

	for {

		firstName, lastName, email, userTickets := getUserInput()
		//fmt.Println(&userTickets)  // this will print the memory address of userTickets where it going to store value of variable

		// Assign the returned values to variables when calling validateUserinput.
		isValidName, isValidEmail, isUserTickets := validateUserinput()

		if isUserTickets && isValidEmail && isValidName {
			remainingTickets = remainingTickets - uint(userTickets)

			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive confirmation Email at %v.\n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

			booking = append(booking, firstName+" "+lastName)

			firstNames := getFirstNames()
			fmt.Printf("List of Bookings are: %v\n", firstNames)

		} else if !isUserTickets {
			fmt.Println("Error : Tickets value is wrong or exceed over available tickets")
		} else if !isValidEmail {
			fmt.Println("Error : Invalid Email address")
		} else if !isValidName {
			fmt.Println("Error : Name at least contains 2 characters")
		}
	}
}

func getGreeting() {
	fmt.Printf("Welcome to %v booking application.\n", conferenceName)
	fmt.Printf("We have total tickets %v and from that remaining tickets available %v.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your Tickets here to attend")
}

func validateUserinput() (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isUserTickets := userTickets > 0 && userTickets <= int(remainingTickets)

	return isValidName, isValidEmail, isUserTickets
}

func getFirstNames() []string {
	var firstNames = []string{}
	for _, bookingEntry := range booking {
		var splitName = strings.Fields(bookingEntry)
		var splitted_first_name = splitName[0]
		firstNames = append(firstNames, splitted_first_name)
	}

	return firstNames
}

func getUserInput() (string, string, string, int) {

	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your Last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your Email address: ")
	fmt.Scan(&email)

	fmt.Println("Enter No. of Tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}
