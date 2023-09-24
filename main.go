package main

import "fmt"

func main() {

	var conferenceName = "Learn Go Conference"
	const conferenceTickets = 50
	var remainingTickets = 50

	fmt.Println("Welcome to our", conferenceName, "lesson booking application!")
	fmt.Println("There are currently", remainingTickets, "remaining tickets out of a total of", conferenceTickets)

}
