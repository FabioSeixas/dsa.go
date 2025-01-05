package btree

import (
	utils "fabioseixas/dsa/utils"
)

type Node struct {
	value int
	left  *Node
	right *Node
}

type BinaryTree struct {
	value *Node
}

func (t *BinaryTree) search(v int) (*Node, int) {

	queue := []*Node{t.value}

	level := -1

	for queue[0] != nil {

		level++

		current := queue[0]

		if current.value == v {
			return queue[0], level
		}

		if current.value > v {
			queue = append(queue, queue[0].left)
		}

		if current.value < v {
			queue = append(queue, queue[0].right)
		}

		queue = queue[1:]
	}

	return nil, 0
}

func (t *BinaryTree) Find(v int) *Node {
	node, _ := t.search(v)
	return node
}

func (t *BinaryTree) FindWithLevel(v int) (*Node, int) {
	return t.search(v)
}

// TODO
// As a public method
// This method must be able to rebalance the tree if needed
func (t *BinaryTree) Add(v int) {
	existingNode := t.Find(v)
	if existingNode != nil {
		return
	}

}

// BSF - level order traversal
func (t *BinaryTree) GetList() []int {

	result := []int{}
	queue := []*Node{t.value}

	for len(queue) > 0 {

		current := queue[0]
		result = append(result, current.value)

		// delete the first item
		queue = queue[1:]

		if current.left != nil {
			queue = append(queue, current.left)
		}

		if current.right != nil {
			queue = append(queue, current.right)
		}
	}

	return result
}

func BuildFromBSF(input []utils.NilInt) *BinaryTree {

	// todo: accept any list of integers
	// for now, only works with lists that perfect represents binary tree (BSF)

	tree := &BinaryTree{&Node{}}
	tree.value.value = input[0].Value

	inputCurrentIndex := 1

	queue := []*Node{tree.value}

	for len(queue) > 0 {
		current := queue[0]

		if len(input) == inputCurrentIndex {
			queue = queue[1:]
			continue
		}

		if input[inputCurrentIndex].Null == false {
			current.left = &Node{}
			current.left.value = input[inputCurrentIndex].Value
			queue = append(queue, current.left)
		}

		inputCurrentIndex++

		if input[inputCurrentIndex].Null == false {
			current.right = &Node{}
			current.right.value = input[inputCurrentIndex].Value
			queue = append(queue, current.right)
		}

		inputCurrentIndex++

		// delete the first item
		queue = queue[1:]
	}

	return tree
}
