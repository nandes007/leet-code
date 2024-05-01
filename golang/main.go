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
	s1 := "abcabcbb"
	s2 := "bbbbb"
	s3 := "pwwkew"

	fmt.Println("Length of longest substring without repeating characters in '"+s1+":", exercise.LengthOfLongestSubstring(s1))
	fmt.Println("Length of longest substring without repeating characters in '"+s2+":", exercise.LengthOfLongestSubstring(s2))
	fmt.Println("Length of longest substring without repeating characters in '"+s3+":", exercise.LengthOfLongestSubstring(s3))
}
