# Age Calculator

## Project Code

```go
package main 

import (
	"fmt"
	"time"
)

func main() {

    fmt.Print("Enter your birthdate: ")
	  var input string
	  fmt.Scanln(&input)

	dob, err := time.Parse("02-01-2006",input) 

	if err!= nil {
			fmt.Println("Invalid date format : ", err)
			return
		}

		age := time.Since(dob)
		years:= int(age.Hours()/24/365)
		days := int(age.Hours()/24)
		hours := int(age.Hours())
		minutes := int(age.Minutes())
		seconds := int(age.Seconds())
		
		fmt.Printf("you are %d years old \n", years)
		fmt.Printf("you are %d days old \n", days)
		fmt.Printf("you are %d hours old \n", hours)
		fmt.Printf("you are %d minutes old \n", minutes)
		fmt.Printf("you are %d seconds old \n", seconds)

}
```

## Documentation

This is a simple Age calculator in which the program will ask the user their Birthdate in a specific format ( DD-MM-YYYY) and give their age in years, days, hour, minutes and seconds.

So for this project we import two packages “fmt” and “time” 

- The “fmt” package is used for us to access the `Printf()` which we use to display the age after calculations.
- The “time” package is used for us to access certain built in methods like `Parse()` , `Since()` etc. which I’m gonna explain now.
- So in Go the first function that is gonna start executing is always the `main()`
- `fmt.Print("Enter your birthdate: (DD-MM-YYYY)")`

first we ask the user for their Birthdate in the mentioned format.

- `var input string`

we declare a string type variable named input that is for storing the input of the user (their birthdate)

- `fmt.Scanln(&input)`

we take the input that the user typed in and store it in the input method

- `dob, err := time.Parse("02-01-2006",input)`

 So now that we got the birthdate of the user as string, how do we tell our program that it is a Birth date and not any ordinary string? We can achieve this by the `time.Parse()`. What this method does is that it simply takes two arguments and return in `time.Time` format

       

      : - the format of the string that we are providing (a.k.a the format of the date: 02-01-2006) 

      :- the string itself that we want to convert into the date format. (input)

Here, we have provided the format as "02-01-2006" and the user input as input. This will convert the input string into a date format that we can use for further calculations which in our case calculating the age of the user.

- `if err!= nil {
 fmt.Println("Invalid date format : ", err)
return
	}` :

We also gave a err parameter didn’t we? this is to catch any err that may occur during the input read for example: user type in invalid date format compared to the specified date format

- `age := time.Since(dob)`

Now using `time.Since()` we calculate the difference between the current time and the input date.

- `years:= int(age.Hours()/24/365)
days := int(age.Hours()/24)
hours := int(age.Hours())
minutes := int(age.Minutes())
seconds := int(age.Seconds())`

Now we turn the age that is given into each of the shown time quantities using simple calculations and methods and store them into each variables with meaningful names, pretty self-explanatory.

- `fmt.Printf("you are %d years old \n", years)
 fmt.Printf("you are %d days old \n", days)
 fmt.Printf("you are %d hours old \n", hours)
 fmt.Printf("you are %d minutes old \n", minutes)
 fmt.Printf("you are %d seconds old \n", seconds)`

Now we simply Print the results we calculated using `Printf()` and since we display the time in integers we use the type specifier %d to let the program know it’s a decimal value.