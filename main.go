package main

import "fmt"

func main(){
	conferenceName := "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50
	fmt.Printf("confrence ticket is %T, remainig Tickets is %T, conferenceName is %T\n",conferenceTickets,remainingTickets,conferenceName)
	fmt.Printf("Welcome to %v booking application\n",conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n",conferenceTickets,remainingTickets)
	fmt.Println("Get your tickets here to attend")
    
}

