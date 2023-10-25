package main

import (
	"fmt"
	"strings"
)

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

	for {
		firstName, lastName, email, userTickets := userDataRetriever()
		var isValidTicketAmount bool = ticketValidator(userTickets, remainingTickets)
		if !isValidTicketAmount{
			fmt.Printf("Remaining tickets (%v) is less than %v! So, go fuck yourself and try again :)\n", remainingTickets, userTickets)
			break
		} else{
			remainingTickets = remainingTickets - userTickets
		}

		bookings = processBooking(bookings, firstName, lastName, email, userTickets, remainingTickets, conferenceName)
		reprBookings(bookings)
	}


}

func userDataRetriever() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint
	fmt.Println("This comes from userDataRetriever method!")
	fmt.Print("Enter your first name: ")
	fmt.Scan(&firstName)
	fmt.Print("Enter your last name: ")
	fmt.Scan(&lastName)
	fmt.Print("Enter your email: ")
	fmt.Scan(&email)
	fmt.Print("Enter number of tickets: ")
	fmt.Scan(&userTickets)	
	return firstName, lastName, email, userTickets
}

func ticketValidator(userTickets uint, remainingTickets uint) bool {
	if userTickets > remainingTickets || remainingTickets == 0{
		return false
	} else {
		return true
	}
}

func processBooking(bookings []string, firstName string, lastName string, email string, userTickets uint, remainingTickets uint, conferenceName string) []string{
	// bookings[0] = firstName + " " + lastName
	bookings = append(bookings, firstName + " " + lastName)
	fmt.Printf("Thank You %v %v for booking %v tickets. You will receive a confirmation email at %v with more details!\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
	return bookings
}

func reprBookings(bookings []string){
	firstNames := []string{}

	for _, booking := range bookings {
		// It is like .split() from Python
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	fmt.Printf("Bookings: %v\n", firstNames)
}