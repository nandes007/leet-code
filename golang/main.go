package main

import (
	"fmt"

	"github.com/nandes007/leetcode/golang/exercise"
	"github.com/nandes007/leetcode/golang/model"
)

func printLinkedList(node *model.ListNode) {
	for node != nil {
		fmt.Printf("%d -> ", node.Val)
		node = node.Next
	}
	fmt.Println("nil")
}

func main() {
	numbers := []int{-1, -2, -3, -4, -5}
	fmt.Print(exercise.TwoSum(numbers, -8))

	l1 := &model.ListNode{Val: 2, Next: &model.ListNode{Val: 4, Next: &model.ListNode{Val: 3, Next: nil}}}
	l2 := &model.ListNode{Val: 5, Next: &model.ListNode{Val: 6, Next: &model.ListNode{Val: 4, Next: nil}}}
	// result := exercise.AddTwoNumbers(l1, l2)
	result := exercise.MergeTwoLists(l1, l2)
	sorted := exercise.SortList(result)
	printLinkedList(result)
	printLinkedList(sorted)
}
