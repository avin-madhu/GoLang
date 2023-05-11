package main 

import (
	"fmt"
	"time"
)

func main() {
    fmt.Print("Enter your birthdate: ")
	var input string
	fmt.Scanln(&input)

	dob, err := time.Parse("02-01-2006",input) // return a time.Time format string
	
	// check if error happened

	if err!= nil {
			fmt.Println("Invalid date format : ", err)
			return
		}

		age := time.Since(dob)
		years:= int(age.Hours()/24/365)
		days := int(age.Hours()/24)
		hours := int(age.Hours())
		minutes := int(age.Minutes())
		
		fmt.Printf("you are %d years old \n", years)
		fmt.Printf("you are %d days old \n", days)
		fmt.Printf("you are %d hours old \n", hours)
		fmt.Printf("you are %d minutes old \n", minutes)


}