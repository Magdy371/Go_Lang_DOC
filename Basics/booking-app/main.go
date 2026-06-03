package main

//fmt stands for format
import (
	"booking-app/helper"
	"fmt"
	"strings"
	"sync"
	"time"
)

const conferenceTickets uint = 50

var confrenceName string = "Go Conference"
var remainingTickets uint = 50

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var booking = make([]UserData, 0)
var wg = sync.WaitGroup{}

func main() {

	for {

		greetUser()

		firstName, lastName, email, userTickets := userInput()

		isValidFullName, isValidEmail, isValidTickets := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidFullName && isValidEmail && isValidTickets {
			bookingTicketa(userTickets, firstName, lastName, email)
			wg.Add(1)
			go sendTicket(userTickets, firstName, lastName, email)

			fmt.Println(`The first names of bookings:`, getFirstNames())
			if remainingTickets == 0 {
				fmt.Printf("Our %v is booked out comeback next year. :(\n", confrenceName)
				break
			}
		} else {
			if !isValidFullName {
				fmt.Println("First or last name is too short")
			}
			if !isValidEmail {
				fmt.Println("Email is not valid")
			}
			if !isValidTickets {
				fmt.Println("Not enough tickets")
			}
		}
	}
	wg.Wait()
}

func greetUser() {
	fmt.Printf("confrenceName type is %T, conferenceTickets type is %T, remainingTickets type is %T\n", confrenceName, conferenceTickets, remainingTickets)
	fmt.Printf("Welcone to our %v booking app \n", confrenceName)
	fmt.Println("Total tickets: ", conferenceTickets)
	fmt.Println("Remaining tickets: ", remainingTickets)
	fmt.Println("Get your text here to attend")
}

func getFirstNames() []string {
	var firstNames []string
	for _, item := range booking {
		var names = strings.Fields(item.firstName)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}

func userInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)
	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)
	fmt.Println("Enter your email: ")
	fmt.Scan(&email)
	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)
	return firstName, lastName, email, userTickets
}

func bookingTicketa(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	booking = append(booking, userData)
	fmt.Printf("The list of bookings is: %v \n", booking)
	fmt.Printf("Remaining tickets: %v\n", remainingTickets)
	fmt.Printf("Thank you %v %v for booking %v tickets.\n You will receive a confirmation email on %v .\n", firstName, lastName, userTickets, email)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets send to %v %v", userTickets, firstName, lastName)
	fmt.Println("#############################")
	fmt.Printf("Sending ticket card\n %v \nto  %v \n\n", ticket, email)
	fmt.Println("#############################")
	wg.Done()
}
