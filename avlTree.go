package main

import "fmt"

type AVLNode struct {
	Key    int
	Height int
	Left   *AVLNode
	Right  *AVLNode
}

type AVLTree struct {
	Root *AVLNode
}

func getHeight(n *AVLNode) int {
	if n == nil { return 0 }
	return n.Height
}

func getBalance(n *AVLNode) int {
	if n == nil { return 0 }
	return getHeight(n.Left) - getHeight(n.Right)
}

func updateHeight(n *AVLNode) {
	if n == nil { return }
	maxH := getHeight(n.Left)
	if getHeight(n.Right) > maxH {
		maxH = getHeight(n.Right)
	}
	n.Height = 1 + maxH
}

func rotateRight(y *AVLNode) *AVLNode {
	x := y.Left
	T2 := x.Right
	x.Right = y
	y.Left = T2
	updateHeight(y)
	updateHeight(x)
	return x
}

func rotateLeft(x *AVLNode) *AVLNode {
	y := x.Right
	T2 := y.Left
	y.Left = x
	x.Right = T2
	updateHeight(x)
	updateHeight(y)
	return y
}

func rebalance(n *AVLNode) *AVLNode {
	updateHeight(n)
	bal := getBalance(n)
	if bal > 1 {
		if getBalance(n.Left) < 0 {
			n.Left = rotateLeft(n.Left)
		}
		return rotateRight(n)
	}
	if bal < -1 {
		if getBalance(n.Right) > 0 {
			n.Right = rotateRight(n.Right)
		}
		return rotateLeft(n)
	}
	return n
}

func insertNode(n *AVLNode, key int) *AVLNode {
	if n == nil {
		return &AVLNode{Key: key, Height: 1}
	}
	if key < n.Key {
		n.Left = insertNode(n.Left, key)
	} else if key > n.Key {
		n.Right = insertNode(n.Right, key)
	} else {
		return n
	}
	return rebalance(n)
}

func (t *AVLTree) Insert(key int) {
	t.Root = insertNode(t.Root, key)
}

func searchNode(n *AVLNode, key int) bool {
	if n == nil { return false }
	if n.Key == key { return true }
	if key < n.Key { return searchNode(n.Left, key) }
	return searchNode(n.Right, key)
}

func (t *AVLTree) Search(key int) bool {
	return searchNode(t.Root, key)
}

func (t *AVLTree) PrintInOrder() {
	fmt.Print("InOrder: ")
	inorder_rec(t.Root)
	fmt.Println()
}

func inorder_rec(n *AVLNode) {
	if n == nil { return }
	inorder_rec(n.Left)
	fmt.Printf("%d ", n.Key)
	inorder_rec(n.Right)
}

func (t *AVLTree) PrintPreOrder() {
	fmt.Print("PreOrder: ")
	preorder_rec(t.Root)
	fmt.Println()
}

func preorder_rec(n *AVLNode) {
	if n == nil { return }
	fmt.Printf("%d ", n.Key)
	preorder_rec(n.Left)
	preorder_rec(n.Right)
}

func (t *AVLTree) PrintPostOrder() {
	fmt.Print("PostOrder: ")
	postorder_rec(t.Root)
	fmt.Println()
}

func postorder_rec(n *AVLNode) {
	if n == nil { return }
	postorder_rec(n.Left)
	postorder_rec(n.Right)
	fmt.Printf("%d ", n.Key)
}

func (t *AVLTree) PrintLevelOrder() {
	fmt.Print("LevelOrder: ")
	if t.Root == nil {
		fmt.Println("(empty)")
		return
	}
	
	// Используем срез как очередь для указателей на узлы
	queue := []*AVLNode{t.Root}
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:] // pop front
		
		fmt.Printf("%d ", curr.Key)
		
		if curr.Left != nil { queue = append(queue, curr.Left) }
		if curr.Right != nil { queue = append(queue, curr.Right) }
	}
	fmt.Println()
}