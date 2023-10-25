package models

type UserData struct{
	FirstName string
	LastName string
	Email string
	NumberOfTickets uint
}

func FullName(user UserData) string{
	return user.FirstName + " " + user.LastName
}