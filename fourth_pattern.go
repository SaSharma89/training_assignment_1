/*******************************
program for bellow pattern
   *
  ***
 *****
*******
 *****
  ***
   *
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
		for k = i; k <= (num - 1); k++{
			fmt.Print(" ")
		}
		for j = num - i; j < num; j++ {
			if j == (num - i) {
				fmt.Print("*")
			}else {
				fmt.Print("**")
			}
		}

		fmt.Println()
	}

	for i = 1; i <= num-1; i++{
		for k = 1; k <= i; k++{
			fmt.Print(" ")
		}

		for j = 1; j <= num-i; j++ {
			if j == 1 {
				fmt.Print("*")
			} else {
				fmt.Print("**")
			}
		}

		fmt.Println()
	}
}

