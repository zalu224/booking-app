package main

import "fmt"

func main() {
	// variable needs to be used or error, same for packages
	var conferenceName = "Go Conference"
	// conferenceName := "Go Conference"
	// can also do, but doesnt work for constant

	// constants are like variables but cannot be changed
	const conferenceTickets = 50
	// needs for the remaning tickets
	var remainingTickets uint = 50

	// %T gives the type for each variable
	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	// var bookings []string
	bookings := []string{}

	for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint //uint unsigned integer, only positive integers
		// ask user for their name, needs pointer &
		// values stored in memory need a pointer to point to the memory value
		fmt.Println("Enter your first name: ")
		fmt.Scan(&firstName)

		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email: ")
		fmt.Scan(&email)

		fmt.Println("Enter your number of tickets: ")
		fmt.Scan(&userTickets)

		remainingTickets = remainingTickets - userTickets
		//appends right to left and assigns it back to the original bookings
		bookings = append(bookings, firstName+" "+lastName)

		fmt.Printf("The whole slice: %v\n", bookings)
		fmt.Printf("The whole slice: %v\n", bookings[0])
		fmt.Printf("Slice type: %T\n", bookings)
		fmt.Printf("Slice length: %v\n", len(bookings))

		// fmt.Println(remainingTickets) //the value of remainingTickets printed out
		// fmt.Println(&remainingTickets) //will print out memory address of remainingTickets
		fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v.\n", firstName, lastName, userTickets, email)
		fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

		firstNames := []string{}
		for index, booking := bookings 

		fmt.Printf("These are all the bookings: %v\n", bookings)
	}

}
