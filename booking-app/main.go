package main // Every code in go is written in packages and we also start it from main package here

import (
	"fmt" //Its a built-in package for GO and we import it like this
	"strings"
)

func main() {

	confName := "Go Conference" // We can also define variables with values like this but its not possible with const
	const confTickets = 50
	var remainingTickets uint = 50 // unit is unsigned integer so that value cant be zero
	bookings := []string{}         // This is how we define slice, slices are subset of array but much more powerful

	fmt.Printf("Data type of confName is %T\n", confName)

	greetUser(confName, confTickets, remainingTickets)

	// In Go we have only for loop, but with different types
	for { // This is the example of for loop for infinite loop

		fName, lName, email, userTickets := getUserInput()

		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(fName, lName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, fName+" "+lName) // This is how we append values to slice

			fmt.Printf("Now remaining tickets are %v\n", remainingTickets)

			fmt.Printf("User %v %v has booked %v tickets and it will be sent to %v mail\n", fName, lName, userTickets, email)

			firstNames := getFirstName(bookings)
			fmt.Printf("These first names of bookings are: %v\n", firstNames)

			if remainingTickets == 0 {
				//end program
				fmt.Println("Our conference tickets are booked, please try next year")
				break
			}
		} else {

			if !isValidName {
				fmt.Println("First Name or Last Name is too short, please provide a name with minimum 2 characters")
			}
			if !isValidEmail {
				fmt.Println("Email you enetered is incorrect!")
			}
			if !isValidTicketNumber {
				fmt.Println("No of Tickets you enetered is invalid")
			}
		}
	}

}

func greetUser(cName string, cTickets int, rTickets uint) {
	fmt.Printf("Welcome to %v booking application\n", cName) //Here %v is a refernce variable
	fmt.Printf("We have total %v number of tickets out of which %v are remaining.\n", cTickets, rTickets)
	fmt.Println("Get your tickets here to attend...")
	//Printf is used for easy formating but it excludes bydefault next line.
}

func getFirstName(bookings []string) []string {
	firstNames := []string{} // New slice for getting only first name and storing there.

	// _ is used for unsed variables
	for _, booking := range bookings { // This for loop iterates over bookings slice and range function will provide 2 values one is index and the value of element
		var names = strings.Fields(booking) // Here we are using built-in function of GO to seprate fileds with white spaces using string function
		var firstName = names[0]
		firstNames = append(firstNames, firstName)
	}
	return firstNames
}

func validateUserInput(fName string, lName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(fName) >= 2 && len(lName) >= 2
	isValidEmail := strings.Contains(email, "@") // In strings module we have function contains which we used here for validation
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidEmail, isValidTicketNumber //this is how we return values
}

func getUserInput() (string, string, string, uint) {
	var fName string     // When we are just declaring varaible we must provide the data type of variable
	var userTickets uint // But when we are defining with value data type is optional.
	var lName string
	var email string

	fmt.Println("Enter your First name:")
	fmt.Scan(&fName)
	fmt.Println("Enter your Last name:")
	fmt.Scan(&lName)
	fmt.Println("No. Of Tickets need to book:")
	fmt.Scan(&userTickets)
	fmt.Println("Email id where tickets will be send:")
	fmt.Scan(&email)

	return fName, lName, email, userTickets
}
