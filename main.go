package main

import (
	"fmt"
	"golang-playground/helper"
	"golang-playground/models"
	"sync"
)

var wg = sync.WaitGroup{}

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
	bookings := make([]models.UserData, 0)
	// var bookings = []string{"Josh", "Mike"}

	firstName, lastName, email, userTickets, errorMessage := helper.UserDataRetriever()
	if len(errorMessage) != 0{
		fmt.Printf("Validation Error: %v. Try again!\n", errorMessage)
		// continue
	}
	var isValidTicketAmount bool = helper.TicketValidator(userTickets, remainingTickets)
	if !isValidTicketAmount{
		fmt.Printf("Remaining tickets (%v) is less than %v! So, go fuck yourself and try again :)\n", remainingTickets, userTickets)
		// break
	} else{
		remainingTickets = remainingTickets - userTickets
	}

	bookings = helper.ProcessBooking(bookings, firstName, lastName, email, userTickets, remainingTickets, conferenceName)
	wg.Add(1)
	go helper.SendEmail(&wg, firstName, lastName, email, userTickets)
	helper.ReprBookings(bookings)

	wg.Wait()
}