package main

import "fmt"

// Entry point for the application
func main() {
	// Declaring a variable
	var conferenceName string = "Hello Golang Conference"
	var conferenceTickets int = 50
	var remainingTickets int = 10
	// Types
	fmt.Printf("conferenceName is %T, and conferenceTickets is %T, and remainingTickets is %T \n", conferenceName, conferenceTickets, remainingTickets)
	fmt.Println("Welcome to", conferenceName)
	fmt.Printf("We have total %v tickets\n", conferenceTickets)
	fmt.Printf("We have %d tickets remaining\n", remainingTickets)

	var firstName string
	var lastName string
	var email string
	var userTickets int
	fmt.Print("Enter your first name: ")
	fmt.Scan(&firstName)
	fmt.Print("Enter your last name: ")
	fmt.Scan(&lastName)
	fmt.Print("Enter your email: ")
	fmt.Scan(&email)
	fmt.Print("Enter number of tickets: ")
	fmt.Scan(&userTickets)
	fmt.Printf("Thank You %v %v for booking %v tickets. You will receive a confirmation email at %v with more details!", firstName, lastName, userTickets, email)
}