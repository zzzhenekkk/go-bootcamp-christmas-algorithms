package main

import (
	"fmt"
)

// TreeNode represents a node in the binary tree.
type TreeNode struct {
	HasToy bool
	Left   *TreeNode
	Right  *TreeNode
}

// countToys counts the number of toys in the tree.
func countToys(root *TreeNode) int {
	if root == nil {
		return 0
	}
	count := 0
	if root.HasToy {
		count++
	}
	count += countToys(root.Left)
	count += countToys(root.Right)
	return count
}

// areToysBalanced checks if the number of toys in left and right subtrees are equal.
func areToysBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	leftCount := countToys(root.Left)
	rightCount := countToys(root.Right)
	return leftCount == rightCount
}

// unrollGarland returns a slice of bools representing the toys in a zig-zag level order.
func unrollGarland(root *TreeNode) []bool {
	if root == nil {
		return nil
	}

	var result []bool
	result = append(result, root.HasToy)

	var currentLevel []*TreeNode
	currentLevel = append(currentLevel, root)
	height := 0

	for len(currentLevel) > 0 {
		var levelValues []bool
		nextLevel := []*TreeNode{}

		for _, pointer := range currentLevel {
			levelValues = append(levelValues, pointer.HasToy)

			if pointer.Left != nil {
				nextLevel = append(nextLevel, pointer.Left)
			}
			if pointer.Right != nil {
				nextLevel = append(nextLevel, pointer.Right)
			}
		}

		if height%2 == 1 {
			for i, j := 0, len(levelValues)-1; i < j; i, j = i+1, j-1 {
				levelValues[i], levelValues[j] = levelValues[j], levelValues[i]
			}
		}

		result = append(result, levelValues...)
		currentLevel = nextLevel
		height++
	}

	return result
}

func main() {
	// Creating a sample tree:
	//      1
	//     / \
	//    1   0
	//   / \ / \
	//  1  0 1  1
	root := &TreeNode{HasToy: true}
	root.Left = &TreeNode{HasToy: true}
	root.Right = &TreeNode{HasToy: false}
	root.Left.Left = &TreeNode{HasToy: true}
	root.Left.Right = &TreeNode{HasToy: false}
	root.Right.Left = &TreeNode{HasToy: true}
	root.Right.Right = &TreeNode{HasToy: true}

	// Testing areToysBalanced function
	fmt.Println("Are toys balanced?", areToysBalanced(root)) // Output: true

	// Testing unrollGarland function
	garland := unrollGarland(root)
	fmt.Println("Unrolled garland:", garland) // Output: [true true false true true false true]
}
