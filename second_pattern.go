/*******************************
program for bellow pattern
1
2 3
4 5 6
7 8 9 10
11 12 13 14
********************************/
package main

import "fmt"

func main(){
	var num int
	fmt.Println("Enter a number:")
	fmt.Scan(&num)
	fmt.Println("You have entered number : ", num)

	var i, j int
	var k = 1
	for i = 1; i <= num; i++{
		for j = 0; j < i; j++ {
			fmt.Print(k," ")
			k++
		}
		fmt.Println()
	}
}
