// to display the volume of a sphere.

package main 

import "fmt"

func main() {
	var radius float32
	fmt.Print("Enter the radius of the sphere: ");
	fmt.Scan(&radius)
	volume := (4.0/3.0)*3.14*radius*radius*radius
    fmt.Println("The volume of the sphere is: ",volume)
}