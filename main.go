package main

import (
	"fmt"
	"strings"
)

func main() {

	var conferenceName = "Learn Go Conference"
	const conferenceTickets = 50

	// fmt.Printf("conferenceTickets is %T, remaining tickets is %T, conference name is %T \n", conferenceTickets, remainingTickets, conferenceName)

	fmt.Printf("Welcome to our %v lesson booking application!\n", conferenceName)

	for {

		var firstName string
		var lastName string
		var userTickets uint
		var userEmail string

		fmt.Println("Enter your first name: ")
		fmt.Scan(&firstName)

		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email for confirmation: ")
		fmt.Scan(&userEmail)

		fmt.Println("Enter your tickets required: ")
		fmt.Scan(&userTickets)

		if userTickets > conferenceTickets || userTickets <= 0 {

			fmt.Println("Invalid ticket order")
			return

		}

		fmt.Printf("Hi, %v %v. You have requested %v tickets. \n", firstName, lastName, userTickets)
		fmt.Printf("A confirmation email has been sent to %v. Thanks for your booking!\n", userEmail)

		var bookings = []string{}
		bookings = append(bookings, firstName+" "+lastName)

		// fmt.Printf("The whole slice: %v \n", bookings)
		// fmt.Printf("First index: %v \n", bookings[0])
		// fmt.Printf("Length of slice: %v \n", len(bookings))

		var ticketsLeft uint = conferenceTickets - userTickets
		fmt.Printf("%v tickets remaining. \n", ticketsLeft)

		firstNames := []string{}
		for _, bookings := range bookings {
			var names = strings.Fields(bookings)
			firstNames = append(firstNames, names[0])

		}

		fmt.Printf("All of our bookings: %v\n", firstNames)

		if ticketsLeft == 0 {

			//If tickets = 0, end application
			fmt.Println("Tickets sold out!")
			break
		}

	}

}
