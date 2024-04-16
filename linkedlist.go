package main

import "fmt"

type Node struct {
	data int
	next *Node
}

var head *Node

func addNode(data int) {
	new := Node{data, nil}

	if head == nil {
		head = &new
	} else {
		new.next = head
		head = &new
	}

}

func removeNode() {

}

func insertNode() {

}

func printList() {

	cur := head

	for cur != nil {

		fmt.Printf("%d->", cur.data)
		cur = cur.next

	}

}

func printMenu() {
	fmt.Println("\nWhat operation would you like to do?")
	fmt.Println("1. Add")
	fmt.Println("2. Remove")
	fmt.Println("3. Insert")
	fmt.Println("4. Print")
	fmt.Println("5. Quit")
}

func main() {

	var option int
	for option != 5 {
		printMenu()
		fmt.Scan(&option)
		switch option {
		case 1:

			fmt.Println("What value should be added?")
			var new int
			fmt.Scan(&new)
			addNode(new)
			fmt.Println("Node Added!")
		case 2:
			removeNode()
			fmt.Println("Node Removed!")
		case 3:
			insertNode()
			fmt.Println("Node Inserted!")
		case 4:
			printList()
		case 5:
			fmt.Println("Happy Coding!")
			break
		default:
			fmt.Println("Please input a valid number:")
		}
	}
}
