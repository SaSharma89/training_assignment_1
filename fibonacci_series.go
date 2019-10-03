package main

import "fmt"

func fib( a int)([]int){

	var i int
	var list []int
	list = append(list, 0)
	list = append(list, 1)

	for i = 1; i < a; i++ {
		list = append(list, list[i] + list[i-1])
	}
	return list

}

func main(){
	var num int
	fmt.Println("Hello, This is fibonacci series program")
	fmt.Println("Enter any valid number : ")
	fmt.Scan(&num)
	fmt.Println("You have entered number is ", num)
	var intSlice = fib(num)
	fmt.Println(intSlice)
}