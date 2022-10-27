package main

// import global package format
import "fmt"

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50

	fmt.Printf("Welcome to %v Booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v tickets are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Book tickets here to attend")

	var userName string
	var userTickets int
	// as user for their name and ticket number

	// a pointer is a variable that points to the memory address of another variable
	fmt.Println("Enter Your Username")
	fmt.Scan(&userName)

	userTickets = 2
	fmt.Printf("user %v booked %v tickets.\n", userName, userTickets)

}
