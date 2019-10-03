/*******************************
program for bellow pattern
   *
  **
 ***
****
********************************/

package main

import "fmt"

func main(){
	var num int
	fmt.Println("Enter a number:")
	fmt.Scan(&num)
	fmt.Println("You have entered number : ", num)

	var i, j, k int
	for i = 1; i <= num; i++{
		for k = i; k < num; k++{
			fmt.Print(" ")
		}
		for j =0; j < i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
