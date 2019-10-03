package main

import "fmt"

func reverse( a int)( b int){
	b = 0
	if a < 0 {
		return
	}

	for ; a > 0; a = a/10 {
		b *= 10
		b += a%10
	}
	return
}

func main() {
	fmt.Println("Hello this is program to reverse the given number")
	var num int
	fmt.Println("Enter any valid number : ")
	fmt.Scan(&num)
	fmt.Println("You have entered number is ", num)

	var rNum = reverse(num)

	fmt.Println("Your reverse number is ", rNum)
}