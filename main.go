package main

import (
	"fmt"
)

func main() {
	// reader := bufio.NewReader(os.Stdin)
	// char, _, _ := reader.ReadRune()

	var todoList []string

	fmt.Println("GO TODO APP")
	for {
		var inputStr string
		fmt.Scanln(&inputStr)

		todoList = append(todoList, inputStr)
		fmt.Println(todoList)

		if inputStr == "x" {
			fmt.Println("Exiting...")
			break
		}

	}

}
