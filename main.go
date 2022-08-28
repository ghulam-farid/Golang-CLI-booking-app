package main

import (
	"fmt"
	"sync"
	"time"
)

var conferenceName = "Go Conference" // when using syntethic sugar (:) not need to write keyword var and var_type. works for var only
const conferenceTickets uint = 50

var remaningTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	greetUsers()

	firstName, lastName, email, userTickets := getUserInput()
	isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets, remaningTickets)

	if isValidName && isValidEmail && isValidTicketNumber {
		bookTicket(userTickets, firstName, lastName, email)
		wg.Add(1)
		go sendTicket(userTickets, firstName, lastName, email)

		firstNames := printFirstNames()

		fmt.Printf("The first names of bookings are :%v\n", firstNames)
		if remaningTickets == 0 {
			fmt.Println("All tickets are sold out. Please come next year")
			// break
		}
	} else {
		if !isValidName {
			fmt.Println("First name or Last name your entered is too short")
		}
		if !isValidEmail {
			fmt.Println("Email address you entered doestn't contain @ symbol")
		}
		if !isValidTicketNumber {
			fmt.Println("Number of tickets you entered is invalid")
		}
	}
	wg.Wait()
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remaningTickets)
	fmt.Println("Get your tickets here to attend the conferrance")
}
func printFirstNames() []string {
	firstNames := []string{}
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

	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}
func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remaningTickets -= userTickets

	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)

	fmt.Printf("The new list of booking is %v\n", bookings)
	fmt.Printf("Thank you %v %v for bookings %v tickets. You will receive confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remaningTickets, conferenceName)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for :%v %v", userTickets, firstName, lastName)
	fmt.Println("===================")
	fmt.Printf("Sending Tickets\n %v \n to email address %v\n", ticket, email)
	fmt.Println("===================")
	wg.Done()
}
