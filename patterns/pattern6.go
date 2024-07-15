package main
import "fmt"

func main() {
	for i:=1;i<=5;i++ {
		for j:=1;j<=6-i;j++ {
			fmt.Print(j," ")
		}
	fmt.Println()
	}
}