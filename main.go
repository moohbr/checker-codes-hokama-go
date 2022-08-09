package main

import (
	"fmt"
)

// Definition of a node in a stack pile.
type Node struct {
	key  string
	line int
	next *Node
}

func createStack() *Node {
	return nil
}

func createNode(key string, line int) *Node {
	return &Node{key, line, nil}
}

func push(head *Node, key string, line int) (*Node, error) {
	newNode := createNode(key, line)
	if newNode == nil {
		return nil, fmt.Errorf("could not create node")
	}
	newNode.key = key
	newNode.line = line
	newNode.next = head
	return newNode, nil
}

func pop(head *Node) (*Node, error) {
	if head == nil {
		return nil, fmt.Errorf("stack is empty")
	}
	return head.next, nil
}

func showSize(head *Node) int {
	size := 0
	for head != nil {
		size++
		head = head.next
	}
	return size
}

func showStack(head *Node) {
	for head != nil {
		fmt.Println(head.key)
		head = head.next
	}
}

func printError(head *Node) {
	for head != nil {
		fmt.Printf("Error: %s is not closed in line %d \n", head.key, head.line)
		head = head.next
	}
}

func main() {
	var typed_string string
	var count_enter int = 0

	fmt.Scanln(&typed_string)

	head_key := createStack()
	head_paranthesis := createStack()
	head_brackets := createStack()

	for _, key := range typed_string {
		if string(key) == "(" {
			head_paranthesis, _ = push(head_paranthesis, string(key), count_enter)
		}
		if string(key) == "[" {
			head_brackets, _ = push(head_brackets, string(key), count_enter)
		}
		if string(key) == "{" {
			head_key, _ = push(head_key, string(key), count_enter)
		}
		if string(key) == ")" {
			head_paranthesis, _ = pop(head_paranthesis)
		}
		if string(key) == "]" {
			head_brackets, _ = pop(head_brackets)
		}
		if string(key) == "}" {
			head_key, _ = pop(head_key)
		}
		if string(key) == "\n" {
			count_enter++
		}
	}

	if showSize(head_paranthesis) != 0 {
		printError(head_paranthesis)
	} else if showSize(head_brackets) != 0 {
		printError(head_brackets)
	} else if showSize(head_key) != 0 {
		printError(head_key)
	} else {
		fmt.Println("No errors")
	}
}
