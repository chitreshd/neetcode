package main

import (
	"fmt"
	"strconv"
	"strings"
)

/**
 * how can I serialize? BFS or DFS, looking at the sample output it feels like BFS. So lets just do that.
 * gotchas
 ** serialize: couldn't understand how to properly terminate the loop.
 */

func main() {
	fmt.Println("hello world")
	root := ConstructTree()
	serializedTree := Serialize(root)
	fmt.Println("BFS: " + serializedTree)
	fmt.Println("DFS through recursion: " + SerializeRecursive(root))
	deserializedNode := Deserialize(serializedTree)
	againSerialized := Serialize(deserializedNode)
	fmt.Println("deserialized: " + againSerialized)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Serialize(node *TreeNode) string {
	if node == nil {
		return "null"
	}
	var res []string

	queue := []*TreeNode{node}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node == nil {
			res = append(res, "null")
		} else {
			res = append(res, fmt.Sprintf("%d", node.Val))
			queue = append(queue, node.Left, node.Right)
		}
	}

	return strings.Join(res, ",")
}

func SerializeRecursive(node *TreeNode) string {
	if node == nil {
		return "null"
	}
	return fmt.Sprintf("%d,%s,%s", node.Val, SerializeRecursive(node.Left), SerializeRecursive(node.Right))
}

func Deserialize(data string) *TreeNode {
	if data == "null" {
		return nil
	}
	vals := strings.Split(data, ",")
	index := 0
	val, _ := strconv.Atoi(vals[index])
	index++
	root := TreeNode{
		Val: val,
	}
	queue := []*TreeNode{&root}

	for index < len(vals) {
		currentNode := queue[0]
		queue = queue[1:]

		// left processing
		if vals[index] != "null" {
			val, _ := strconv.Atoi(vals[index])
			leftNode := &TreeNode{
				Val: val,
			}
			currentNode.Left = leftNode
			queue = append(queue, leftNode)
		}
		index++

		if vals[index] != "null" {
			val, _ = strconv.Atoi(vals[index])
			rightNode := &TreeNode{
				Val: val,
			}
			currentNode.Right = rightNode
			queue = append(queue, rightNode)
		}
		index++

	}
	return &root
}

func ConstructTree() *TreeNode {
	root := TreeNode{
		Val: 1,
	}
	root.Left = &TreeNode{
		Val: 2,
	}
	root.Right = &TreeNode{
		Val: 3,
	}
	root.Right.Left = &TreeNode{
		Val: 4,
	}
	root.Right.Right = &TreeNode{
		Val: 5,
	}

	return &root
}
