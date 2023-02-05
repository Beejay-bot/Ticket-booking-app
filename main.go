package main

import (
	"fmt"
	"strconv"
	"strings"
)


var conferenceName = "Go conference 2023"
const conferenceTickets=50
var remainingTickets uint= 50
// declaring or creating a slice here
var bookings = make([]map[string]string, 0)

func main()  {
	greetUsers()
	
	for remainingTickets > 0  && len(bookings) < 50 {
		firstName, lastName, email, userTickets := getUserInputs()
		isValidName, isEmailValid, isValidTicketNumber := validateUserInputs(firstName, lastName, email, userTickets)

		if !isValidName || !isEmailValid || !isValidTicketNumber {
			if !isValidName{
				fmt.Println("First name or last name is too short")
			}
			if !isEmailValid {
				fmt.Println("Email does not contain '@' ")
			}
			if !isValidTicketNumber {
				fmt.Printf("We only have %v tickets remaining, so you can't book %v tickets.\n", remainingTickets, userTickets)
			}
			// The "continue" keyword will cause the loop to start all over again skipping the remianing code && instructions. 
			continue
		}
		bookTicket(userTickets, firstName, lastName, email)
		sendTicket(userTickets, firstName, lastName, email)
		firstNames:= getFirstNames()
		fmt.Printf("The first names of all bookings are: %v \n", firstNames)

		fmt.Println("=================================")
		fmt.Println("=================================")
		
		if remainingTickets ==0{
			fmt.Printf("The %v tickets has been sold out. Thank you", conferenceName)
			break
		}
	}	
}

func greetUsers()  {
	// The Printf is used for formatting values or placeholders in a print statement
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v  are still available \n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend now!!!!!")
}

func getFirstNames() []string {
	firstNames := []string{}
		// In go _ are used to ignore variables you don't want use. 
		// In go, you need to make unused variables explicit. 
		for _, booking := range bookings{
			firstNames = append(firstNames,	booking["firstName"])
		}
		return firstNames
}

func validateUserInputs(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool)  {
	isValidName:= len(firstName) >=2 && len(lastName) >= 2
	isEmailValid := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isEmailValid, isValidTicketNumber
}

func getUserInputs() (string, string, string, uint)  {
	    var firstName string
		var lastName string
		var email string
		var userTickets uint

		//ask for firstName
		fmt.Println("Please enter your firstName:")
		// A pointer is a variable that points to the memory address of another variable
		fmt.Scan(&firstName)

		fmt.Println("Please enter your lastName:")
		fmt.Scan(&lastName)


		fmt.Println("Please enter your email:")
		fmt.Scan(&email)

		fmt.Println("Enter number of tickets: ")
		fmt.Scan(&userTickets)

		return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string)  {
		remainingTickets -= userTickets

		// create a map for a user
		var userData =make(map[string]string)
		userData["firstName"] = firstName
		userData["lastName"] = lastName
		userData["email"] = email
		userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)
		bookings = append(bookings, userData)
		fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve a confirmation at %v\n", firstName, lastName, userTickets, email)
		fmt.Printf("%v tickets remaining for %v \n", remainingTickets, conferenceName)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string){
	// Sprintf helps put together a formatted but instead of printing them, you store them in a variable.
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("###################")
	fmt.Printf("Sending ticket: \n %v \nto email address %v\n", ticket, email)
	fmt.Println("###################")
}