package main

import "fmt"

type Node struct {
	Key   int
	left  *Node
	right *Node
}

// Insert will add a node to the tree
// the key to add should not be already in the tree
func (n *Node) Insert(k int) {
	if n.Key > k {

		// move to left node
		if n.left == nil {
			n.left = &Node{Key: k}
		} else {
			n.left.Insert(k)
		}
	} else if n.Key < k {

		// move to right node
		if n.right == nil {
			n.right = &Node{Key: k}
		} else {
			n.right.Insert(k)
		}
	}
}

// Search will take in a key value
// and RETURN true if there is a node with that value
func (n *Node) Search(k int) bool {
	if n == nil {
		return false
	}

	if n.Key > k {
		// Check right child
		return n.left.Search(k)
	} else if n.Key < k {
		// Check right child
		return n.left.Search(k)
	}

	return true
}

func main() {
	tree := &Node{Key: 100}
	tree.Insert(200)
	tree.Insert(300)

	fmt.Println(tree.Search(301))
}
