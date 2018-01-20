package main

import (
	"fmt"
)

type ListNode struct {
	Val int
	Next *ListNode
}

func addSingleDigits(d1, d2 , carry int)(int, int) {
	sum := d1 + d2 + carry
	return sum / 10, sum % 10
}

func printList(root *ListNode) {
	node := root
	for node != nil {
		fmt.Println(node.Val)
		node = node.Next
	}
}

func createList(num int) *ListNode {
	var root, node *ListNode
	for num > 0 {
		digit := num % 10
		num = num / 10
		newNode := ListNode{digit, nil}
		if node == nil {
			node = &newNode
			root = &newNode
		} else {
			node.Next = &newNode
		}
		node = &newNode
	}
	return root
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var carry, sum int
	var node, root *ListNode

	for l1 != nil && l2 != nil {
		carry, sum = addSingleDigits(l1.Val, l2.Val, carry)
		newNode := ListNode{sum, nil}
		if node == nil {
			root = &newNode
		} else {
			node.Next = &newNode
		}
		l1 = l1.Next
		l2 = l2.Next
		node = &newNode
	}

	var remaining *ListNode
	if l1 != nil {
		remaining = l1
	} else if l2 != nil {
		remaining = l2
	}
	for remaining != nil  {
		carry, sum = addSingleDigits(remaining.Val, 0, carry)
		newNode := ListNode{sum, nil}
		node.Next = &newNode
		remaining = remaining.Next
		node = node.Next
	}
	if carry > 0 {
		newNode := ListNode{carry, nil}
		node.Next = &newNode
	}
	return root
}


func main() {
	num1 := createList(1)
	num2 := createList(99)
	result := addTwoNumbers(num1, num2)
	fmt.Println("answer:")
	printList(result)
}