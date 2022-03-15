package main

import (
	"fmt"
)

// Node - Struct element for tree.
// Tree nodes have value, left child and
// right child fields.
type Node struct{
	Value		int
	Left		*Node
	Right		*Node
}

// Insert - Adds a new tree node which have value v.
// Function checks value is greater or equal than
// corresponding node's value.
func (n *Node) Insert(v int){
	if n.Value < v{
		// move right
		if n.Right == nil{
			n.Right = &Node{Value: v}
		}else{
			n.Right.Insert(v)
		}
	}else if n.Value > v{
		// move left
		if n.Left == nil{
			n.Left = &Node{Value: v}
		}else{
			n.Left.Insert(v)
		}
	}
}

// Search - Checks whether the tree has a node
// whose value is v or not and return boolean.
func (n *Node) Search(v int)bool{
	if n == nil {
		return false
	}
	
	if n.Value < v {
		// move right
		return n.Right.Search(v)
	}else if n.Value > v {
		// move left
		return n.Left.Search(v)
	}
	return true
}

func main(){
	tree := &Node{Value: 100}
	fmt.Println(*tree)
	tree.Insert(200)
	fmt.Println(*tree)
	tree.Insert(22)
	tree.Insert(55)
	tree.Insert(11)
	tree.Insert(33)
	tree.Insert(66)
	tree.Insert(112)
	tree.Insert(222)
	fmt.Println("66", tree.Search(66))
	fmt.Println("112", tree.Search(112))
	fmt.Println("99", tree.Search(99))
	fmt.Println("300", tree.Search(300))
}
