package main

import (
	"fmt"
)

func main() {
	s := []int{4, 5, 1, 6, 2, 7, 3, 8,}
	fmt.Println(GetLeastNumbers_Solution(s, 8))

	k := 1
	for i := 2*k + 1; i < 10; i = 2*i + 1 { //k的值确定好了
		fmt.Println(i)
		k = 10 //不影响
	}
}

/*判断给定的链表中是否有环。如果有环则返回true，否则返回false。
你能给出空间复杂度的解法么？

判断链表是否有环应该是老生常谈的一个话题了，最简单的一种方式就是快慢指针，慢指针针每次走一步，快指针每次走两步，如果相遇就说明有环，如果有一个为空说明没有环。代码比较简单
*/
//func hasCycle(head *ListNode) bool {
//	if head == nil {
//		return false
//	}
//	slow := head
//	quick := head
//	for quick.Next != nil && quick.Next.Next != nil {
//		slow = slow.Next
//		quick = quick.Next.Next
//		if slow == quick {
//			return true
//		}
//	}
//	return false
//}

/*
给定一个数组，找出其中最小的K个数。例如数组元素是4,5,1,6,2,7,3,8这8个数字，则最小的4个数字是1,2,3,4。如果K>数组的长度，那么返回一个空的数组
示例1
输入
复制
[4,5,1,6,2,7,3,8],4
返回值
复制
[1,2,3,4]*/
// topk
//https://www.cnblogs.com/chengxiao/p/6129630.html heap sort
//https://www.runoob.com/w3cnote/heap-sort.html all sort
func GetLeastNumbers_Solution(input []int, k int) []int {
	if input == nil || len(input) == 0 || k == 0 || k > len(input) { //pre check
		return nil
	}
	if k > len(input) {
		return input
	}
	heap := make([]int, k)
	for i := 0; i < k; i++ {
		heap[i] = input[i]
	}
	// init k size heap
	for i := k/2 - 1; i >= 0; i-- {
		adjustHeap(heap, i, k)
	}
	// compare top heap --k
	for i := k; i < len(input); i++ {
		if input[i] < heap[0] {
			heap[0] = input[i]
			adjustHeap(heap, 0, k)
		}
	}
	//sort
	for len := k; len > 0; len-- {
		heap[len-1], heap[0] = heap[0], heap[len-1]
		adjustHeap(heap, 0, len-1) //去掉1个再安排堆
	}
	return heap
}

//递归
func adjustHeap(heap []int, k int, len int) {
	left := 2*k + 1
	right := 2*k + 2
	largest := k
	if left < len && heap[left] > heap[largest] {
		largest = left
	}
	if right < len && heap[right] > heap[largest] {
		largest = right
	}
	if largest != k {
		heap[k], heap[largest] = heap[largest], heap[k]
		adjustHeap(heap, largest, len) //交换了再递归
	}
}

//非递归
//https://www.cnblogs.com/chengxiao/p/6129630.html
func adjustHeap(arr []int, i int, length int) {
	temp := arr[i]
	for k := i*2 + 1; k < length; k = k*2 + 1 { //从i结点的左子结点开始，也就是2i+1处开始
		if k+1 < length && arr[k] < arr[k+1] { //如果左子结点小于右子结点，k指向右子结点
			k++
		}
		if arr[k] > temp { //如果子节点大于父节点，将子节点值赋给父节点（不用进行交换），和最初的比较，判断是否还需要往下子节点遍历
			arr[i] = arr[k]
			i = k //子节点的索引，作为后续用
		} else {
			break
		}
	}
	arr[i] = temp //将temp值放到最终的位置
}

//给出一个仅包含字符'(',')','{','}','['和']',的字符串，判断给出的字符串是否是合法的括号序列
//括号必须以正确的顺序关闭，"()"和"()[]{}"都是合法的括号序列，但"(]"和"([)]"不合法。
func isValid(s string) bool {
	stack := make([]byte, 0)
	for i := range s {
		c := s[i]
		if c == '(' {
			stack = append(stack, ')')
		} else if c == '[' {
			stack = append(stack, ']')
		} else if c == '{' {
			stack = append(stack, '}')
		} else {
			if len(stack) == 0 {
				return false
			}
			e := stack[len(stack)-1]
			if e != c {
				return false
			}
			stack = append(stack[:len(stack)-1], nil...)
		}
	}
	return len(stack) == 0
}
