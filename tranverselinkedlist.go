package main

import "fmt"

func tranverlinkedlist(l *ListNode) {
	if l == nil {
		return
	}
	fmt.Println(l.Val)
	tranverlinkedlist(l.Next)
}
