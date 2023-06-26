package main

import "fmt"

func main() {
	var inputStr string
	var todoList []string
	fmt.Println("GO TODO APP")
	fmt.Scanln(&inputStr)
	todoList = append(todoList, inputStr)
	fmt.Println(todoList)
}
