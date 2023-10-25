package helper

import (
	"fmt"
	"strings"
)

func UserDataRetriever() (string, string, string, uint, string) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint
	var errorMessage string	
	fmt.Print("Enter your first name: ")
	fmt.Scan(&firstName)
	fmt.Print("Enter your last name: ")
	fmt.Scan(&lastName)
	fmt.Print("Enter your email: ")
	fmt.Scan(&email)
	fmt.Print("Enter number of tickets: ")
	fmt.Scan(&userTickets)	

	if (len(firstName) < 2 || len(lastName) < 2){
		errorMessage = "First and Last names must not be less than 2 chars"
	} else if !strings.Contains(email, "@"){
		errorMessage = "Invalid Email"
	}
	return firstName, lastName, email, userTickets, errorMessage
}

func TicketValidator(userTickets uint, remainingTickets uint) bool {
	if userTickets > remainingTickets || remainingTickets == 0{
		return false
	} else {
		return true
	}
}

func ProcessBooking(bookings []string, firstName string, lastName string, email string, userTickets uint, remainingTickets uint, conferenceName string) []string{
	// bookings[0] = firstName + " " + lastName
	bookings = append(bookings, firstName + " " + lastName)
	fmt.Printf("Thank You %v %v for booking %v tickets. You will receive a confirmation email at %v with more details!\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
	return bookings
}

func ReprBookings(bookings []string){
	firstNames := []string{}

	for _, booking := range bookings {
		// It is like .split() from Python
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	fmt.Printf("Bookings: %v\n", firstNames)
}