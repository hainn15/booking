package main

import "fmt"

func main() {

	var conferenceName1 = "Go Conference"
	const conferenceTicket = 50 //const Constants are like variables, exept that their value cannot be changed
	//variable for ticket remaining
	var remainingTickets = 50

	fmt.Println("Welcome to ", conferenceName1, " Booking application")
	fmt.Println("We have total of", conferenceTicket, "tickets and ", remainingTickets, "are still avaiable.")
	fmt.Println("Get your ticket here to attend")

	//print formatted data
	//%v the value in a defualt format
	fmt.Printf("Welcome to %v Booking application \n", conferenceName1)
	fmt.Printf("We have total %v of tickets and %v are still avaiable \n", conferenceTicket, remainingTickets)
	fmt.Println("Get your Ticket here to attend")

	//Data type
	var firstName string
	var lastName string
	var email string
	var userticket int

	// ask their name
	fmt.Println("Enter Your Last Name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter Your First Name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter Your Email Address: ")
	fmt.Scan(&email)

	fmt.Println("How many tickets you want to book?")
	fmt.Scan(&userticket) //Scan funtion read the value user enter in the memory and assign that value to userName Variable
	fmt.Printf("Thank Mr. %v %v booked %v tickets and please check email address: %v confirm your booking schedule.\n", lastName, firstName, userticket, email)

	//int8, int16 (-32768-32768), int32, int64 - uint8, uint16 (0 - 65535), uint32, uint64

	//collect user input, write to a file.

	//pointer iss variable that points to the memory address of another variable in Golang pointer also call special variable
	fmt.Println(remainingTickets)  //print value remaining ticket
	fmt.Println(&remainingTickets) //print hash memory access of value variable remainingTicket -> that basicly point

	remainingTickets = remainingTickets - userticket
	fmt.Printf("Remaining Tickets: %v", remainingTickets)

}
