package main

import "fmt"

func main() {

	var conferenceName = "Learn Go Conference"
	const conferenceTickets = 50
	var remainingTickets = 50

	fmt.Printf("Welcome to our %v lesson booking application!\n", conferenceName)
	fmt.Printf("There are currently %v remaining tickets out of a total of %v. \n", remainingTickets, conferenceTickets)

}
