package main

import (
	"fmt"
)

func main() {

	var conferenceName = "Learn Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = conferenceTickets

	// fmt.Printf("conferenceTickets is %T, remaining tickets is %T, conference name is %T \n", conferenceTickets, remainingTickets, conferenceName)

	fmt.Printf("Welcome to our %v lesson booking application!\n", conferenceName)
	fmt.Printf("There are currently %v remaining tickets out of a total of %v. \n", remainingTickets, conferenceTickets)

	for {

		var firstName string
		var lastName string
		var userTickets uint
		var userEmail string

		fmt.Println("Enter your first name: ")
		fmt.Scan(&firstName)
		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)
		fmt.Println("Enter your tickets required: ")
		fmt.Scan(&userTickets)

		fmt.Println("Enter your email for confirmation: ")
		fmt.Scan(&userEmail)

		fmt.Printf("Hi, %v %v. You have requested %v tickets. \n", firstName, lastName, userTickets)
		fmt.Printf("A confirmation email has been sent to %v. Thanks for your booking!\n", userEmail)

		var bookings = []string{}
		bookings = append(bookings, firstName+" "+lastName)

		fmt.Printf("The whole slice: %v \n", bookings)
		fmt.Printf("First index: %v \n", bookings[0])
		fmt.Printf("Length of slice: %v \n", len(bookings))

		var ticketsLeft uint = remainingTickets - userTickets
		fmt.Printf("%v tickets remaining. \n", ticketsLeft)

		fmt.Printf("All of our bookings: %v\n", bookings)

	}

}
