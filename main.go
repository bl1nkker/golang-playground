package main

import "fmt"

// Entry point for the application
func main() {
	// Declaring a variable
	var conferenceName string = "Hello Golang Conference"
	var conferenceTickets int = 50
	var remainingTickets int = 10
	fmt.Println("Welcome to", conferenceName)
	fmt.Printf("We have total %v tickets\n", conferenceTickets)
	fmt.Printf("We have %d tickets remaining\n", remainingTickets)

	
}