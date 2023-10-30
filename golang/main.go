package main

import (
	"fmt"

	"github.com/nandes007/leetcode/golang/exercise"
)

func printLinkedList(node *exercise.ListNode) {
	for node != nil {
		fmt.Printf("%d -> ", node.Val)
		node = node.Next
	}
	fmt.Println("nil")
}

func main() {
	numbers := []int{-1, -2, -3, -4, -5}
	fmt.Print(exercise.TwoSum(numbers, -8))

	l1 := &exercise.ListNode{Val: 2, Next: &exercise.ListNode{Val: 4, Next: &exercise.ListNode{Val: 3, Next: nil}}}
	l2 := &exercise.ListNode{Val: 5, Next: &exercise.ListNode{Val: 6, Next: &exercise.ListNode{Val: 4, Next: nil}}}
	result := exercise.AddTwoNumbers(l1, l2)
	printLinkedList(result)
}
