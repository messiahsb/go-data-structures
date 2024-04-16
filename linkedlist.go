package main

import "fmt"

type Node struct {
	data int
	next *Node
}

func addNode() {

}

func removeNode() {

}

func insertNode() {

}

func printList() {

}

func main() {

	fmt.Println("What operation would you like to do?")
	fmt.Println("1. Add")
	fmt.Println("2. Remove")
	fmt.Println("3. Insert")
	fmt.Println("4. Print")
	fmt.Println("5. Quit")

	var option int

	for option != 5 {
		switch option {
		case 1:
			addNode()
			fmt.Println("Node Added!")
		case 2:
			removeNode()
			fmt.Println("Node Removed!")
		case 3:
			insertNode()
			fmt.Println("Node Inserted!")
		case 4:
			printList()
		default:
			fmt.Printf("Please input a valid number:")
		}
	}
}
