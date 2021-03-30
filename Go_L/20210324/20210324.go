package main

import "math"

/*
 * type ListNode struct{
 *   Val int
 *   Next *ListNode
 * }
 */

/**
 *
 * @param pHead ListNode类
 * @return ListNode类
 */
/*输入一个链表，反转链表后，输出新链表的表头。
示例1
输入
复制
{1,2,3}
返回值
复制
{3,2,1}*/
func ReverseList(pHead *ListNode) *ListNode {
	// write code here
	c := pHead
	var pre *ListNode //nil
	for c != nil {
		next := c.Next
		c.Next = pre
		pre = c
		c = next
	}
	return pre //pre
}

/*给定一个 n * m 的矩阵 a，从左上角开始每次只能向右或者向下走，最后到达右下角的位置，路径上所有的数字累加起来就是路径和，输出所有的路径中最小的路径和。
示例1
输入
复制
[[1,3,5,9],[8,1,3,4],[5,0,6,1],[8,8,4,0]]
返回值
复制
12*/
/*题解 动态规划
第一行 只能从左往右
第一个元素 的值为 原数组的第一个元素 dp[0][0] = a[0][0]
dp[0][j] = a[0][j] + dp[0][j-1];
第一列元素 只能从上往下
dp[i][0] = dp[i-1][0] + a[i][0]
第二行第二列元素的可能从 当前节点的左节点 和上节点过来
那么该节点的最小值应为 当前节点的值 加上 min （ 上节点 左节点）
dp[i][j] = a[i][j] + Math.min(dp[i][j-1],dp[i-1][j]);*/
func minPathSum(matrix [][]int) int {
	m := len(matrix)
	n := len(matrix[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	dp[0][0] = matrix[0][0]
	for i := 1; i < n; i++ {
		dp[0][i] = dp[0][i-1] + matrix[0][i]
	}
	for i := 1; i < m; i++ {
		dp[i][0] = dp[i-1][0] + matrix[i][0]
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = matrix[i][j] + int(math.Min(float64(dp[i-1][j]), float64(dp[i][j-1])))
		}
	}
	return dp[m-1][n-1]
}

/*将一个链表\ m m 位置到\ n n 位置之间的区间反转，要求时间复杂度 O(n)O(n)，空间复杂度 O(1)O(1)。
例如：
给出的链表为 1\to 2 \to 3 \to 4 \to 5 \to NULL1→2→3→4→5→NULL, m=2,n=4m=2,n=4,
返回 1\to 4\to 3\to 2\to 5\to NULL1→4→3→2→5→NULL.
注意：
给出的 mm,nn 满足以下条件：
1 \leq m \leq n \leq 链表长度1≤m≤n≤链表长度
示例1
输入
复制
{1,2,3,4,5},2,4
返回值
复制
{1,4,3,2,5}*/
/*public ListNode reverseBetween(ListNode head, int m, int n) {
ListNode dummy = new ListNode(0);
dummy.next = head;
ListNode preStart = dummy;
ListNode start = head;
for (int i = 1; i < m; i ++ ) {
preStart = start;
start = start.next;
}
// reverse
for (int i = 0; i < n - m; i ++ ) {
ListNode temp = start.next;
start.next = temp.next;
temp.next = preStart.next;
preStart.next = temp;
}
return dummy.next;
//????????????????????????????
}*/
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	//var pre *ListNode
	//c := head
	//for i := 0; i < m-1; i++ {
	//	pre = c
	//	c = c.Next
	//}
	//mCur := c
	//mBefore := pre
	//
	//c = head
	//for i := 0; i < n-1; i++ {
	//	c = c.Next
	//}
	//nAfter := c.Next
	//
	//c = mCur
	//for i := 0; i < n-m; i++ {
	//	next := c.Next
	//	c.Next = pre
	//	pre = c
	//	c = next
	//}
	//mCur.Next = nAfter
	//mBefore.Next = c
	//return head
}

/*给定一个二叉树和一个值\ sum sum，请找出所有的根节点到叶子节点的节点值之和等于\ sum sum 的路径，
例如：
给出如下的二叉树，\ sum=22 sum=22，
返回
[
[5,4,11,2],
[5,8,9]
]
*/
/*
 * type TreeNode struct {
 *   Val int
 *   Left *TreeNode
 *   Right *TreeNode
 * }
 */

/**
 *
 * @param root TreeNode类
 * @param sum int整型
 * @return int整型二维数组
 */
var way []int
var result [][]int

func pathSum(root *TreeNode, sum int) [][]int {
	dfs(root, 0, sum)
	return result
}

func dfs(root *TreeNode, total int, targetSum int) {
	if root == nil {
		return
	}
	total += root.Val
	way = append(way, root.Val)
	if root.Left == nil && root.Right == nil { //叶子节点
		if total == targetSum {
			result = append(result, way)
		}
	} else {
		dfs(root.Left, total, targetSum)
		dfs(root.Right, total, targetSum)
	}
	way = way[:len(way)-1]
}

/*用例输入???????????????????
{1},0
预期输出
[]
实际输出 [[1]]*/
