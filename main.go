package main

import (
	"booking-app/helper"
	"fmt"
	"time"
)

var conferenceName = "Go Conference"
const conferenceTickets = 50
var remainingTickets uint = 50
var bookings = make([]UserData,0)

type UserData struct{
	firstName string
	lastName string
	email string
	numberOfTickets uint
}

func main() {
	greetUsers()
	for {
		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)
		if isValidTicketNumber && isValidEmail && isValidName {
			bookTicket(userTickets, firstName, lastName, email)
            sendTicket(userTickets,firstName,lastName,email)
			firstNames := printFirstNames()
			fmt.Printf("These are all our bookings: %v\n\n", firstNames)
			if remainingTickets == 0 {
				fmt.Println("Our conference is booked out, Come back next year!")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("First or Last name you entered is too short.")
			}
			if !isValidEmail {
				fmt.Println("Email you entered is wrong.")
			}
			if !isValidTicketNumber {
				fmt.Println("Number of tickets you entered is invalid.")
			}
			fmt.Println("Please check your input.\n")
			continue
		}
	}
}

func greetUsers() {
	fmt.Printf("confrence ticket is %T, remainig Tickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func printFirstNames() []string {
	var firstNames []string
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}

	return firstNames
}

func getUserInput() (string, string, string, uint) {
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

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - uint(userTickets)
	
	//create a map for user
	var userData = UserData{
		firstName: firstName,
		lastName: lastName,
		email: email,
		numberOfTickets: userTickets,
	}
	
	bookings = append(bookings,userData)
	fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v ticket remaining for %v\n", remainingTickets, conferenceName)
}

func sendTicket(userTickets uint,firstName string,lastName string,email string){
	time.Sleep(10*time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v",userTickets,firstName,lastName)
	fmt.Println("################")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n",ticket,email)
	fmt.Println("################")
}