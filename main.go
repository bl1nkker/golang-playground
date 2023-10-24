package main

import "fmt"

// Entry point for the application
func main() {
	// Declaring a variable
	conferenceName := "Hello Golang Conference"
	var conferenceTickets int = 50
	var remainingTickets uint = 10
	// Types
	fmt.Printf("conferenceName is %T, and conferenceTickets is %T, and remainingTickets is %T \n", conferenceName, conferenceTickets, remainingTickets)
	fmt.Println("Welcome to", conferenceName)
	fmt.Printf("We have total %v tickets\n", conferenceTickets)
	fmt.Printf("We have %d tickets remaining\n", remainingTickets)

	// Array
	// var bookings [50]string
	// var bookings = [50]string{"Josh", "Mike"}
	// var bookings = [50]string{"Josh", "Mike"}

	// Slice
	// var bookings []string
	bookings := []string{}
	// var bookings = []string{"Josh", "Mike"}

	var firstName string
	var lastName string
	var email string
	var userTickets uint
	fmt.Print("Enter your first name: ")
	fmt.Scan(&firstName)
	fmt.Print("Enter your last name: ")
	fmt.Scan(&lastName)
	fmt.Print("Enter your email: ")
	fmt.Scan(&email)
	fmt.Print("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	remainingTickets = remainingTickets - userTickets
	// bookings[0] = firstName + " " + lastName
	bookings = append(bookings, firstName + " " + lastName)

	fmt.Printf("Thank You %v %v for booking %v tickets. You will receive a confirmation email at %v with more details!\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
	fmt.Printf("Bookings: %v", bookings)

}