package main

import "fmt"

// func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
// 	returnListNode := ListNode{}
// 	l1Str := lnToString(l1)
// 	l2Str := lnToString(l2)
// 	len1 := len(l1Str)
// 	len2 := len(l2Str)
// 	if len1 != len2 {
// 		if len1 > len2 {
// 			remainZeros := strings.Repeat("0", len1-len2)
// 			l2Str += remainZeros

// 		}
// 		if len2 > len1 {
// 			remainZeros := strings.Repeat("0", len2-len1)
// 			l1Str += remainZeros

// 		}
// 	}
// 	l1Int := strToInt(l1Str)
// 	l2Int := strToInt(l2Str)

// 	resultInt := l1Int + l2Int
// 	resultStr := strconv.Itoa(resultInt)
// 	setReturnNodesValue(resultStr, &returnListNode, len(resultStr)-1)
// 	return &returnListNode

// }

// // tyrna use some bruteforce
// // get each l1 andc l2 and turn it to an integer
// // then turn the result to a reversed linked list
// // func lnToInt(l *ListNode) int {
// // 	str := lnToString(l)
// // 	i, err := strconv.Atoi(str)
// // 	if err != nil {
// // 		return 0
// // 	}
// // 	return i
// // }

// func strToInt(str string) int {

// 	i, err := strconv.Atoi(str)
// 	if err != nil {
// 		return 0
// 	}
// 	return i
// }
// func lnToString(l *ListNode) string {
// 	if l == nil {
// 		return ""
// 	}
// 	return fmt.Sprintf("%d", l.Val) + lnToString(l.Next)
// }

// func setReturnNodesValue(str string, listNode *ListNode, index int) {
// 	if listNode == nil || index < 0 {
// 		return
// 	}

// 	// get the value and save it
// 	valInt, err := strconv.Atoi(string(str[index]))
// 	if err != nil {
// 		return
// 	}
// 	// fmt.Println(valInt)
// 	listNode.Val = valInt
// 	if index > 0 {
// 		nextNode := ListNode{}
// 		listNode.Next = &nextNode
// 	}
// 	setReturnNodesValue(str, listNode.Next, index-1)
// }

// TODO all that wont work my man, worked for your idea but not leetcode's

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	resultNode := ListNode{}
	helper(l1, l2, &resultNode, 0)
	return &resultNode
}

func helper(l1, l2, resultNode *ListNode, remainder int) {
	if l1 == nil || l2 == nil {
		if l1 == nil && l2 == nil {
			// if this is the end and there is a remainder
			if remainder == 1 {
				resultNode.Val = 1
			}
			return
		}
		if l1 == nil && l2 != nil {
			l1 = &ListNode{}
		}
		if l2 == nil && l1 != nil {
			l2 = &ListNode{}
		}
	}
	sum, remain := adder(l1.Val, l2.Val, remainder)
	resultNode.Val = sum
	// resultNode.Next = &ListNode{}
	// helper(l1.Next, l2.Next, resultNode.Next, remain)
	// Only create the next node if there's more data to process
	if l1.Next != nil || l2.Next != nil || remain != 0 {
		resultNode.Next = &ListNode{}
		helper(l1.Next, l2.Next, resultNode.Next, remain)
	}

}

func adder(num1, num2, remainder int) (sum int, remain int) {
	// fmt.Println("Input.......")
	// fmt.Println(num1, num2, remainder)

	s := num1 + num2 + remainder

	if s >= 10 {
		sum = s - 10
		remain = 1

	} else {
		remain = 0
		sum = s
	}
	fmt.Println("Output.......")

	fmt.Println(sum, remain)
	return sum, remain
}
