// Copyright 2020 wongoo@apache.org. All rights reserved.

// a red-black tree implement
// the node only contains pointers to left/right child, not for the parent, for saving storage space for large tree.
package rbtree

import "sync"

// Color node color
type Color bool

func (c Color) String() string {
	if c {
		return "red"
	}
	return "black"
}

// Position tree path position, left or right.
type Position bool

const (
	Red   = Color(true)
	Black = Color(false)
	Left  = Position(true)
	Right = Position(false)
)

// Node the node of red-Black tree.
type Node struct {
	Key         int
	Color       Color
	Left, Right *Node
	Value       interface{}
}

// Black a node is black if nil or its color is black.
func (n *Node) Black() bool {
	return n == nil || n.Color == Black
}

// Red a node is red if nil or its color is red.
func (n *Node) Red() bool {
	return n == nil || n.Color == Red
}

// LeftBlack the left child of a node is black if nil or its color is black.
func (n *Node) LeftBlack() bool {
	return n.Left == nil || n.Left.Color == Black
}

// LeftRed the left child of a node is black if not nil and its color is black.
func (n *Node) LeftRed() bool {
	return n.Left != nil && n.Left.Color == Red
}

// RightBlack the right child of a node is black if nil or its color is black.
func (n *Node) RightBlack() bool {
	return n.Right == nil || n.Right.Color == Black
}

// RightRed the right child of a node is black if not nil and its color is black.
func (n *Node) RightRed() bool {
	return n.Right != nil && n.Right.Color == Red
}

// RBTree red-black tree
type RBTree struct {
	Node  *Node
	lock  sync.Mutex
	stack *stack
}

// New create a new red-black tree
func New() *RBTree {
	return &RBTree{
		lock:  sync.Mutex{},
		Node:  nil,
		stack: newStack(nil),
	}
}

// LeftRotate left rotate a node.
func LeftRotate(n *Node) *Node {
	r := n.Right
	if r == nil {
		return n
	}

	n.Right = r.Left
	r.Left = n

	return r
}

// RightRotate right rotate a node.
func RightRotate(n *Node) *Node {
	l := n.Left
	if l == nil {
		return n
	}

	n.Left = l.Right
	l.Right = n

	return l
}
func (t *RBTree) Add(key int, value interface{}) {
	t.lock.Lock()
	defer t.lock.Unlock()

	t.Node = addTreeNode(t.stack, t.Node, key, value)
}

func (t *RBTree) Find(key int) interface{} {
	t.lock.Lock()
	defer t.lock.Unlock()

	return Find(t.Node, key)
}

func (t *RBTree) Delete(key int) (ret interface{}) {
	t.lock.Lock()
	defer t.lock.Unlock()

	t.Node, ret = Delete(t.Node, key)

	return ret
}

func addTreeNode(stack *stack, node *Node, key int, value interface{}) *Node {
	stack.init(node)
	defer stack.reset()

	if node == nil {
		// case 1: new root
		return &Node{
			Key:   key,
			Color: Black,
			Value: value,
		}
	}

	for node != nil {
		if node.Key == key {
			node.Value = value
			return stack.root()
		}

		if key < node.Key {
			stack.push(node, Left)
			node = node.Left
		} else {
			stack.push(node, Right)
			node = node.Right
		}
	}

	stack.bindChild(&Node{
		Key:   key,
		Color: Red,
		Value: value,
	})

	addTreeNodeBalance(stack)

	root := stack.root()
	root.Color = Black

	return root
}

// addTreeNodeBalance balance the tree after adding a node
// the pre condition is the child of current stack is red
func addTreeNodeBalance(stack *stack) {
	for stack.index > 0 {
		p := stack.node()

		// case 2: P is black, balance finish
		if p.Color == Black {
			return
		}

		// P is red

		pp := stack.parent()
		// case 1: reach the root
		if pp == nil {
			return
		}

		s := stack.sibling()

		// case 3: P is red, S is red, PP is black
		// execute: set P,S to black, PP to red
		// result: black count through PP is not change, continue balance on parent of PP
		if s != nil && s.Color == Red {
			p.Color = Black
			s.Color = Black
			pp.Color = Red

			stack.pop().pop()

			continue
		}

		// case 4: P is red, S is black, PP is black, the position of N and P are diff.
		// execute: rotate up the red child
		// result: let match the case 5.
		pos, ppos := stack.position(), stack.parentPosition()
		if pos != ppos {
			if pos == Left {
				p = RightRotate(p)
				pp.Right = p
			} else {
				p = LeftRotate(p)
				pp.Left = p
			}
		}

		// case 5: P is red, S is black, PP is black, the position of N and P are the same.
		// execute: set P to black, PP to red, and rotate P up
		// result: black count through P will not change, balance finish.
		p.Color = Black
		pp.Color = Red
		var ppn *Node
		if ppos == Left {
			ppn = RightRotate(pp)
		} else {
			ppn = LeftRotate(pp)
		}

		stack.pop().pop().bindChild(ppn)

		return
	}
}

// Add add new key/value, return the new root node.
// this method add node and balance the tree recursively, not using loop logic.
func AddOne(root *Node, key int, value interface{}) *Node {
	return addNew(root, &Node{
		Key:   key,
		Value: value,
	})
}

// addNew add new node, return the new root node.
func addNew(root *Node, new *Node) *Node {
	// set the new node to red
	new.Color = Red

	root = addNode(root, Left, new)

	// reset root color
	root.Color = Black

	return root
}

// addNode recursively down to leaf, and add the new node to the leaf,
// then rebuild the tree from the leaf to root.
// the main purpose is reduce two linked red nodes and keep the black count balance.
//
// code comment use the following terms:
// - N as the balance node
// - L as the left child of N
// - R as the right child of N
// - P as the parent of N
// - LL as the left child of left child of N
// - RR as the right child of right child of N
func addNode(node *Node, pos Position, new *Node) *Node {
	// case 1: first node
	if node == nil {
		return new
	}

	if new.Key < node.Key {
		node.Left = addNode(node.Left, Left, new)

		// case 2: L is black means it's already balance.
		if node.Left.Color == Black {
			return node
		}

		if node.Color == Red {
			// case 3: L is red, N is red, N is right child of P
			// execute: right rotate up the L
			// result: the black count through L,N will not change, but let it match the case 4
			if pos == Right {
				node = RightRotate(node)
			}

			// case 4: L is red, N is red, N is left child of P
			// execute: nothing
			// result: it's the case 5 of PP
			return node
		}

		if node.Left.Left != nil && node.Left.Left.Color == Red {
			// case 5: N is black, L is red, LL is red
			// execute: right rotate N, and make LL to black
			// result: black count through N is not change, while that through LL increase 1, tree is now balance.
			node = RightRotate(node)
			node.Left.Color = Black
		}

		return node
	}

	if new.Key > node.Key {
		node.Right = addNode(node.Right, Right, new)

		// case 2: R is black means it's already balance
		if node.Right.Color == Black {
			return node
		}

		if node.Color == Red {
			if pos == Left {
				// case 3: R is red, N is red, N is left child of P
				// execute: left rotate up the R
				// result: the black count through R,N will not change, but let it match the case 4
				node = LeftRotate(node)
			}

			// case 4: R is red, N is red, N is right child of P
			// execute: nothing
			// result: it's the case 5 of PP
			return node
		}

		// case 5: N is black, R is red, RR is red
		// execute: left rotate N, and make RR to black
		// result: black count through N is not change, while that through RR increase 1, tree is now balance.
		if node.Right.Right != nil && node.Right.Right.Color == Red {
			node = LeftRotate(node)
			node.Right.Color = Black
		}

		return node
	}

	// case 6: find the exists node, just replace the old value with the new
	node.Value = new.Value

	return node
}

// Find find the value of a key.
func Find(node *Node, key int) interface{} {
	for node != nil {
		if node.Key == key {
			return node.Value
		}
		if key < node.Key {
			node = node.Left
		} else {
			node = node.Right
		}
	}
	return nil
}

// Delete delete a node.
// return the new root node, and the value of the deleted node.
// the new root node will be nil if no node exists in the tree after deleted.
// the deleted node value will be nil if not found.
func Delete(node *Node, key int) (*Node, interface{}) {
	if node == nil {
		return nil, nil
	}

	root := node

	var router *nodePath
	var ret interface{}

	// add a EMPTY
	router = &nodePath{
		Node: &Node{
			Left: node,
		},
		pos: Left,
	}
	rootRouter := router

	// find the node
	for node != nil {
		if node.Key == key {
			ret = node.Value
			break
		}

		if key < node.Key {
			router = router.append(node, Left)
			node = node.Left
		} else {
			router = router.append(node, Right)
			node = node.Right
		}

	}

	// not find
	if node == nil {
		return root, nil
	}

	var inorderSuccessor *Node

	// find the inorder successor
	if node.Right != nil {
		router = router.append(node, Right)

		inorderSuccessor = node.Right

		for inorderSuccessor.Left != nil {
			router = router.append(inorderSuccessor, Left)

			inorderSuccessor = inorderSuccessor.Left
		}

		node.Key = inorderSuccessor.Key
		node.Value = inorderSuccessor.Value

		node = inorderSuccessor
	}

	// get the child of node
	c := node.Left
	if c == nil {
		c = node.Right
	}

	// N has no child
	if c == nil {
		// get the node's previous from router
		p := router

		// delete N
		p.bindChild(nil)

		if node.Color == Red {
			return root, ret
		}

		deleteBalance(router)
		if rootRouter.Left != nil {
			rootRouter.Left.Color = Black
		}
		return rootRouter.Left, ret
	}

	// N has one next
	// then copy key/value from next to N
	node.Key = c.Key
	node.Value = c.Value

	// delete the next
	node.Left = nil
	node.Right = nil

	// N has diff color with next
	if node.Color != c.Color {
		// set color of N to black
		node.Color = Black

		return root, ret
	}

	// the color of N and next are both Black
	deleteBalance(router)
	rootRouter.Left.Color = Black
	return rootRouter.Left, ret
}

// deleteBalance balance the tree after deleting.
// code comment use the following terms:
// - N as the balance node
// - P as the father of N
// - PP as the grand father of N
// - S as the sibling of N
// - SL as the left child of S
// - SR as the right child of S
func deleteBalance(router *nodePath) {
	var (
		p, pp *nodePath
		s     *Node
	)

	// case 1: reach the root.
	// execute: nothing.
	// result: balance finish.
	for router.previous != nil {
		p, pp = router, router.previous
		s = p.siblingChild()

		// case 2: S is red.
		// execute: rotate S up as the PP of N, and exchange the color of P and S.
		// result: the black number not change, but N has a black sibling now.
		if s.Color == Red {
			p.Color, s.Color = s.Color, p.Color

			pp.bindChild(p.rotateUpSiblingChild().Node)

			// reset PP (original S)
			pp = p.previous

			// reset S (a black node, original SL/SR)
			s = p.siblingChild()
		}

		// now S is black.

		if s.LeftBlack() && s.RightBlack() {
			// case 3: color of P, S, SL, SR are all Black.
			// execute: set S to red.
			// result: the path through S will reduce one black, and the left and right of P now balance,
			//         set N to p, and continue execute balance.
			if p.Black() {
				s.Color = Red
				router = router.previous
				continue
			}

			// case4: S, SL, SR are black, P is red.
			// execute: exchange the color of S and P.
			// result: add one black on the path through N, while that is not change for path through S, balance finish.
			p.Color, s.Color = s.Color, p.Color
			return
		}

		//  now SL and SR has diff color

		if p.pos == Left {
			// case 5: N is left child of P, S is black, SL is red, SR is black.
			// execute: right rotate on S, then exchange color of SL(parent of S now) and S.
			// result: N has a new black sibling S(original SL), and S has a red right child SR(original S),
			//         while the black count through S will not change.
			if s.LeftRed() {
				s = RightRotate(s)
				s.Color, s.Right.Color = s.Right.Color, s.Color
				p.Right = s
			}

			// case6: N is left child of P, S is black, SL is black, SR is red.
			// execute: set SR to black, left rotate P, the exchange the color of P and S.
			// result: S is now the parent of P, the black count through N increase 1,
			//         the black count through S keep the same,
			//         balance finish.
			s.Right.Color = Black
			pp.bindChild(LeftRotate(p.Node))
			p.Color, s.Color = s.Color, p.Color
			return
		}

		// case 5: N is right child of P, S is black, SL is black, SR is red.
		// execute: left rotate on S, then exchange color of SR(parent of S now) and S.
		// result: N has a new black sibling S(original SR), and S has a red left child SL(original S),
		//         while the black count through S will not change.
		if s.RightRed() {
			s = LeftRotate(s)
			s.Color, s.Left.Color = s.Left.Color, s.Color
			p.Left = s
		}

		// case6: N is right child of P, S is black, SL is red, SR is black.
		// execute: set SL to black, right rotate P, the exchange the color of P and S.
		// result: S is now the parent of P, the black count through N increase 1,
		//         the black count through S keep the same,
		//         balance finish.
		s.Left.Color = Black
		pp.bindChild(RightRotate(p.Node))
		p.Color, s.Color = s.Color, p.Color
		return
	}
}