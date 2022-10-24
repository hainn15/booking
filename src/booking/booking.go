package main

import (
	"fmt"
	"strings"
)

func main() {

	var conferenceName1 = "Go Conference"
	const conferenceTicket = 50 //const Constants are like variables, exept that their value cannot be changed

	//variable for ticket remaining
	var remainingTickets = 50

	//input welcome to function

	greetUser(conferenceName1, conferenceTicket, remainingTickets)

	//print formatted data
	//%v the value in a defualt format

	//BEGINING OF LOGIC WE WRITE FOR LOOP.

	for { //infinite loop - keep asking after 1 booking

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

		//int8, int16 (-32768-32768), int32, int64 - uint8, uint16 (0 - 65535), uint32, uint64

		//CHECK USER INPUT VALIDATION

		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isValidTickets := userticket > 0 && userticket < 5 && userticket <= remainingTickets

		//collect user input, write to a file.

		//pointer iss variable that points to the memory address of another variable in Golang pointer also call special variable
		fmt.Println(remainingTickets)  //print value remaining ticket
		fmt.Println(&remainingTickets) //print hash memory access of value variable remainingTicket -> that basicly point

		//CHECK USER BOOK NO MORE THAN TICKER AVAILABLE or LESSER TICKET AVAIABLE.
		if isValidName && isValidEmail && isValidTickets {

			remainingTickets = remainingTickets - userticket
			fmt.Printf("Remaining Tickets: %v \n", remainingTickets)

			//array and Slices
			//arry in Go fixed size (how many elements the array can hold)
			//array Only the same data type can be stored

			//var bookings = [50]string{"a","b","c"} //array size 50 with 3 elements

			var bookings [50]string

			//adding new elements accessing by thir index(position)
			bookings[0] = firstName
			bookings[1] = lastName
			bookings[2] = email

			fmt.Println(bookings)

			//PRINT OUT SUMMARY BOOKING.

			// a list that is more dynamic in size, automatic expands when new elemants are added.
			//Slices are more felixible and powerfull slices are index-based and have a size, but is resized when needed
			var booking1 []string
			var numTickets []int
			//append - adds the element at the end of the slice, grow the slice if a grater capacity is needed
			booking1 = append(booking1, firstName+" "+lastName, email)
			numTickets = append(numTickets, userticket)

			//print out first name each.
			//for each loop slices or array
			//Call Print first name function

			printFirstName(booking1) //call funtion and parameter

			fmt.Printf("This is all our booking list %v \n", firstNames)
			fmt.Printf("Thank Mr. %v %v booked %v tickets and please check email address: %v confirm your booking schedule.\n", lastName, firstName, userticket, email)

		} else {

			//Remind User input invalid data

			if !isValidName {
				fmt.Println("First Name or Last name is too short")
			}
			if !isValidEmail {
				fmt.Println("Email Address is  invalid!")
			}
			if !isValidTickets {
				fmt.Println("Your number ticket is invalid")
			}

			fmt.Printf("We only have %v ticket remaining, so you can book %v tickets \n", remainingTickets, userticket)

		} //end of if

		// MAKE THE CHECK OF SOLD OUT TICKET

		var noTicketremaining bool = remainingTickets == 0 //boolean datatype

		if noTicketremaining {
			// break the loop
			fmt.Println("Our Conference is booked out.")
			break

		}

	} //end of for loop

} // end of main funtion

// funtcion additional can execute when it called in main function

func greetUser(conferenceName1 string, conferenceTicket int, remainingTickets int) {

	fmt.Printf("Welcome to Our %v! \n", conferenceName1)
	fmt.Printf("We have total %v of tickets and %v are still avaiable \n", conferenceTicket, remainingTickets)
	fmt.Println("Get your Ticket here to attend")

} //end of func greetUser

func printFirstName(booking1 []string) {
	firstNames := []string{}
	for _, booking := range booking1 { //index replace by _ ignore a variable you don't want to use.
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
}
