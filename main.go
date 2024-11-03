package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50 // in our case remaining tickets value never be negative, so assigned uint
	var booking []string

	fmt.Printf("Welcome to %v booking application.\n", conferenceName)
	fmt.Printf("We have total tickets %v and remaining tickets available %v.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your Tickets here to attend")

	for {
		var firstName string
		var lastName string
		var email string
		var userTickets int

		fmt.Println("Enter your first name: ")
		fmt.Scan(&firstName)

		fmt.Println("Enter your Last name: ")
		fmt.Scan(&lastName)

		fmt.Println("Enter your Email address: ")
		fmt.Scan(&email)

		fmt.Println("Enter No. of Tickets: ")
		fmt.Scan(&userTickets)
		//fmt.Println(&userTickets)  // this will print the memory address of userTickets where it going to store value of variable

		if userTickets <= int(remainingTickets) {
			remainingTickets = remainingTickets - uint(userTickets)

			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive confirmation Email at %v.\n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

			booking = append(booking, firstName+" "+lastName)

			var firstNames = []string{}
			for _, bookingEntry := range booking {
				var splitName = strings.Fields(bookingEntry)
				var splitted_first_name = splitName[0]
				firstNames = append(firstNames, splitted_first_name)
			}

			fmt.Printf("List of Bookings are: %v\n", firstNames)
		} else {
			fmt.Println("Sorry, we don't have enough tickets for you.")
		}
	}
}
