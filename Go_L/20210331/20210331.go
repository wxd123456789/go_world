package main

import (
	"fmt"
	"math"
)

/*给定一棵二叉树以及这棵树上的两个节点 o1 和 o2，请找到 o1 和 o2 的最近公共祖先节点。
示例1
输入
复制
[3,5,1,6,2,0,8,#,#,7,4],5,1
返回值
复制
3*/
//LCA dfs深度优先搜索
type TreeNode struct {
	Val   int
	Right *TreeNode
	Left  *TreeNode
}

func lowestCommonAncestor(root *TreeNode, o1 int, o2 int) int {
	if root == nil {
		return -1
	}
	r := dfs(root, o1, o2)
	return r.Val
}

func dfs(root *TreeNode, o1 int, o2 int) *TreeNode {
	if root == nil || root.Val == o1 || root.Val == o2 {
		return root
	}
	left := dfs(root.Left, o1, o2)
	right := dfs(root.Right, o1, o2)
	if left == nil {
		return right
	} else if right == nil {
		return left
	} else {
		return root
	}
}

/*以字符串的形式读入两个数字，编写一个函数计算它们的和，以字符串形式返回。
（字符串长度不大于100000，保证字符串仅由'0'~'9'这10种字符组成）
示例1
输入
复制
"1","99"
返回值
复制
"100"
说明
1+99=100 */
func solve(s string, t string) string {
	i, j := len(s), len(t)
	max := int(math.Max(float64(i), float64(j)))
	result := make([]byte, max+1)

	indexS := i - 1
	indexT := j - 1
	n := 0
	resultIndex := max
	for indexS >= 0 || indexT >= 0 {
		sum := add(getE(indexS, s), getE(indexT, t), n)
		result[resultIndex] = byte(sum%10 + '0') //将数字1 转化为'1'  1+'0'
		n = sum / 10
		indexS--
		indexT--
		resultIndex--
	}
	result[0] = byte(n + '0') //将数字1 转化为'1'  1+'0' 不然后面转换为诶string失败
	if n == 0 {
		return string(result[1:])
	} else {
		return string(result)
	}
}

func getE(index int, s string) byte {
	if index >= 0 {
		return s[index]
	} else {
		return '0'
	}
}

func add(a byte, b byte, n int) int {
	sum := (a - '0') + (b - '0')
	return int(sum) + n
}

/*给定一个二叉树，返回该二叉树的之字形层序遍历，（第一层从左向右，下一层从右向左，一直这样交替）
例如：
给定的二叉树是{3,9,20,#,#,15,7},

该二叉树之字形层序遍历的结果是
[
[3],
[20,9],
[15,7]
]
示例1
输入
复制
{1,#,2}
返回值
复制
[[1],[2]]
*/
//stack bfs广度优先遍历
//https://mp.weixin.qq.com/s?__biz=MzU0ODMyNDk0Mw==&mid=2247487028&idx=1&sn=e06a0cd5760e62890e60e43a279a472b&chksm=fb419d14cc36140257eb220aaeac182287b10c3cab5c803ebd54013ee3fc120d693067c2e960&token=2095441666&lang=zh_CN#rd
func main() {
}

func testDeleteSliceE() {
	seq := []string{"a", "b", "c", "d", "e"}
	//删除某个元素
	index := 2
	seq = append(seq[:index], seq[index+1:]...)
	fmt.Println(seq) //[a b d e]
	//删除前4个元素
	index = 3
	seq = append(seq[:0], seq[index+1:]...)
	fmt.Println(seq) //[e]
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	result := [][]int{}
	queue := []*TreeNode{root}
	leftTOR := true
	for len(queue) > 0 {
		e := []int{}
		num := len(queue)
		//控制每一层
		for i := 0; i < num; i++ {
			q := queue[0]
			queue = append(queue[:0], queue[1:]...)
			e = append(e, q.Val)
			if q.Left != nil {
				queue = append(queue, q.Left)
			}
			if q.Right != nil {
				queue = append(queue, q.Right)
			}
		}
		//手动反转
		if !leftTOR {
			e = reverseSlice(e)
		}
		leftTOR = !leftTOR
		result = append(result, e)
	}
	return result
}

func reverseSlice(s []int) []int {
	r := make([]int, len(s))
	for i := range s {
		r[i] = s[len(s)-i-1]
	}
	return r
}
