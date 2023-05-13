// display the pattern 1111
                    //2222
					//3333
					//4444

package main 

import "fmt"

func main() {
     for i:=1;i<5;i++ {
         for j:=1;j<5;j++ {
            fmt.Print(i," ")
		 }
		 fmt.Println()
	 }
}