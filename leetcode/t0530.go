package leetcode

import (
	"sort"
)

/*
解题思路：
	1.遍历二叉树（此处使用前序遍历）
	2.排序
	3.计算相邻元素差值，取最小
*/

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getMinimumDifference(root *TreeNode) int {
	arr := preorderTraversal(root)
	sort.Ints(arr)
	ret := arr[1] - arr[0] // 采用数组前两位数，计算差值，作为初始值（初始值的来源，不要简单地置为0）
	for i := 0; i < len(arr)-1; i++ {
		//使用arr[i+1] - arr[i]，减少处理步骤
		difference := arr[i+1] - arr[i]
		if ret > difference {
			ret = difference
		}
	}

	return ret
}

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var nodes = make([]int, 0)
	nodes = append(nodes, root.Val)
	if root.Left != nil {
		nodes = append(nodes, preorderTraversal(root.Left)...)
	}
	if root.Right != nil {
		nodes = append(nodes, preorderTraversal(root.Right)...)
	}
	return nodes
}
