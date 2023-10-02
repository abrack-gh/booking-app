package main

import (
	"fmt"
	"strings"
)

func main() {

	var conferenceName = "Learn Go Conference"
	const conferenceTickets = 50

	// fmt.Printf("conferenceTickets is %T, remaining tickets is %T, conference name is %T \n", conferenceTickets, remainingTickets, conferenceName)

	greetUser(conferenceName)

	for {

		var firstName string
		var lastName string
		var userTickets uint
		var userEmail string

		fmt.Println("Enter your first name, or type 'q' to exit.")
		fmt.Scan(&firstName)

		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email for confirmation: ")
		fmt.Scan(&userEmail)

		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(userEmail, "@")

		if !isValidEmail || !isValidName {

			fmt.Println("Invalid name or email detected, please try again.")
			continue

		}

		fmt.Println("Enter your tickets required: ")
		fmt.Scan(&userTickets)

		if userTickets > conferenceTickets || userTickets <= 0 {

			fmt.Println("Invalid ticket order")
			continue

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

		var firstNames = getNames(bookings)
		fmt.Printf("The names booked so far are %v\n", firstNames)

		if ticketsLeft == 0 {

			//If tickets = 0, end application
			fmt.Println("Tickets sold out!")
			break
		}

	}

	city := "Berlin"

	switch city {
	case "New York", "Mexico City":
		//Same code block for NY and Mexico City

	case "Singapore", "Hong Kong":
	//Same code block for Singapore/HK

	case "Berlin", "London":
	//Same code block for Berlin and London

	default:
		fmt.Println("Invalid city selected")
	}

}

func greetUser(confName string) {

	fmt.Printf("Hello, welcome to %v booking form.\n", confName)

}

func getNames(bookings []string) []string {

	firstNames := []string{}
	for _, bookings := range bookings {
		var names = strings.Fields(bookings)
		firstNames = append(firstNames, names[0])

	}

	return firstNames

}
