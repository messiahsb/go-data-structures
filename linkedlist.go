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

func removeNode(data int) bool {
	cur := head
	prev := head
	for cur != nil {

		if cur.data == data {
			if cur == head {
				head = cur.next
			} else {
				prev.next = cur.next
			}
			return true
		}
		prev = cur
		cur = cur.next

	}
	return false

}

func insertNode() {

}

func printList() {

	cur := head

	for cur != nil {

		fmt.Printf("%d->", cur.data)
		cur = cur.next

	}
	fmt.Printf("\n")

}

func printMenu() {
	fmt.Println("What operation would you like to do?")
	fmt.Println("1. Add")
	fmt.Println("2. Remove")
	fmt.Println("3. Insert")
	fmt.Println("4. Print")
	fmt.Println("5. Quit")
}

func main() {

	var option int
	var new int
	for option != 5 {
		printMenu()
		fmt.Scan(&option)
		switch option {
		case 1:

			fmt.Println("What value should be added?")

			fmt.Scan(&new)
			addNode(new)
			fmt.Println("Node Added!")
			break
		case 2:
			fmt.Println("What value should be removed?")
			fmt.Scan(&new)
			x := removeNode(new)
			if x {
				fmt.Println("Node Removed!")
			} else {
				fmt.Println("Node not found")
			}
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
