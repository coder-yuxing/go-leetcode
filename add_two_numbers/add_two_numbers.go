package add_two_numbers

type ListNode struct {
	Val int
	Next *ListNode
}

// AddTwoNumbers you are given two non-empty linked lists representing two non-negative integers.
// The digits are stored in reverse order, and each of their nodes contains a single digit.
// Add the two numbers and return the sum as a linked list.
//
// You may assume the two numbers do not contain any leading zero, except the number 0 itself.
// Example:
//
// Input: l1 = [2,4,3], l2 = [5,6,4]
//
// Output: [7,0,8]
//
// Explanation: 342 + 465 = 807
func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	maxLen, minLen := getMaxAndMinArr(link2Arr(l1), link2Arr(l2))

	var temp []int
	for i := 0; i < len(maxLen); i++ {
		if i < len(minLen) {
			total := maxLen[i] + minLen[i]
			isEqual := len(temp) - 1 == i
			if isEqual && temp != nil {
				total += temp[i]
			}
			temp = appendTemp(temp, i, total, isEqual)
			continue
		}
		isEqual := len(temp) - 1 == i
		val := maxLen[i]
		if isEqual && temp != nil {
			val += temp[i]
		}

		temp = appendTemp(temp, i, val, isEqual)
	}
	node := convert2ListNode(temp)
	return node
}

func appendTemp(temp []int, index, total int, isEqual bool) []int {
	if total < 10 {
		temp = setVal(temp, index, total, isEqual)
	} else {
		temp = setVal(temp, index, total%10, isEqual)
		temp = append(temp, total/10%10)
	}
	return temp
}

func setVal(temp []int, index, val int, isEqual bool) []int {
	if isEqual && temp != nil {
		temp[index] = val
	} else {
		temp = append(temp, val)
	}
	return temp
}

func convert2ListNode(arr []int) *ListNode {
	var result *ListNode
	var lastNode *ListNode
	for _, a := range arr {
		if result == nil {
			result = &ListNode{
				Val: a,
			}
			lastNode = result
			continue
		}
		preNode := lastNode
		lastNode = &ListNode{
			Val: a,
		}
		preNode.Next = lastNode
	}

	return result
}

func getMaxAndMinArr(a, b []int) (maxArr, minArr []int) {
	if len(a) >= len(b) {
		return a, b
	} else {
		return b, a
	}
}

func link2Arr(l *ListNode) []int {
	var result []int
	for ; l != nil ; {
		result = append(result, l.Val)
		l = l.Next
	}
	return result
}
