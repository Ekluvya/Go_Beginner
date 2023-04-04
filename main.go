package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50
	fmt.Printf("confrence ticket is %T, remainig Tickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
	var bookings []string

	for {

		var firstNames []string
		var firstName string
		var lastName string
		var email string
		var userTickets uint
		fmt.Println("Enter your first name:")
		fmt.Scan(&firstName)
		fmt.Println("Enter your last name:")
		fmt.Scan(&lastName)
		fmt.Println("Enter your email address:")
		fmt.Scan(&email)
		fmt.Println("Enter numbers of tickets:")
		fmt.Scan(&userTickets)
        
		isValidName := len(firstName) >=2 && len(lastName)>=2
		isValidEmail := strings.Contains(email,"@")
	    isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
		if isValidTicketNumber && isValidEmail && isValidName {
			remainingTickets = remainingTickets - uint(userTickets)
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve a confirmation email at %v\n", firstName, lastName, userTickets, email)
			fmt.Printf("%v ticket remaining for %v\n", remainingTickets, conferenceName)

			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}

			fmt.Printf("These are all our bookings: %v\n\n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("Our conference is booked out, Come back next year!")
				break
			}
		} else {
			if !isValidName{
				fmt.Println("First or Last name you entered is too short.")
			} 
			if !isValidEmail{
				fmt.Println("Email you entered is wrong.")
			}
			if !isValidTicketNumber{
				fmt.Println("Number of tickets you entered is invalid.")
			}
			fmt.Println("Please check your input.\n")
			continue
		}
	}
}
