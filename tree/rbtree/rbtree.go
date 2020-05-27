// Copyright 2020 wongoo@apache.org. All rights reserved.

// a red-black tree implement
// the node only contains pointers to left/right child, not for the parent, for saving storage space for large tree.
package rbtree

// Color node color
type Color bool

// String node color desc.
func (c Color) String() string {
	if c {
		return "red"
	}
	return "Black"
}

// Position tree path position, left or right.
type Position bool

// String position desc.
func (p Position) String() string {
	if p {
		return "left"
	}
	return "right"
}

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

// Add add new key/value, return the new root node.
func Add(root *Node, key int, value interface{}) *Node {
	return addNew(root, &Node{
		Key:   key,
		Color: Red,
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

// nodePath the double linked list for the route path in a red-black tree.
type nodePath struct {
	*Node
	pos            Position
	previous, next *nodePath
}

// append a new node path.
func (np *nodePath) append(n *Node, pos Position) *nodePath {
	next := &nodePath{
		Node:     n,
		pos:      pos,
		previous: np,
		next:     nil,
	}

	np.next = next

	return next
}

// bindChild set the child at the position of the path.
func (np *nodePath) bindChild(n *Node) {
	if np.pos == Left {
		np.Left = n
	} else {
		np.Right = n
	}
}

// siblingChild get the sibling of the child of the node.
func (np *nodePath) siblingChild() *Node {
	if np.pos == Left {
		return np.Right
	} else {
		return np.Left
	}
}

// rotateUpSiblingChild rotate the sibling child up as the parent of current node.
func (np *nodePath) rotateUpSiblingChild() *nodePath {
	pp := &nodePath{
		previous: np.previous,
		next:     np,
	}
	np.previous = pp

	if np.pos == Left {
		pp.pos = Left
		pp.Node = LeftRotate(np.Node)
	} else {
		pp.pos = Right
		pp.Node = RightRotate(np.Node)
	}

	return pp
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

	var stack *nodePath
	var ret interface{}

	// add a EMPTY
	stack = &nodePath{
		Node: &Node{
			Left: node,
		},
		pos: Left,
	}
	rootStack := stack

	// find the node
	for node != nil {
		if node.Key == key {
			ret = node.Value
			break
		}

		if key < node.Key {
			stack = stack.append(node, Left)
			node = node.Left
		} else {
			stack = stack.append(node, Right)
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
		stack = stack.append(node, Right)

		inorderSuccessor = node.Right

		for inorderSuccessor.Left != nil {
			stack = stack.append(inorderSuccessor, Left)

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
		// get the node's previous from stack
		p := stack

		// delete N
		p.bindChild(nil)

		if node.Color == Red {
			return root, ret
		}

		deleteBalance(stack)
		if rootStack.Left != nil {
			rootStack.Left.Color = Black
		}
		return rootStack.Left, ret
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
	deleteBalance(stack)
	rootStack.Left.Color = Black
	return rootStack.Left, ret
}

// deleteBalance balance the tree after deleting.
// code comment use the following terms:
// - N as the balance node
// - P as the father of N
// - PP as the grand father of N
// - S as the sibling of N
// - SL as the left child of S
// - SR as the right child of S
func deleteBalance(stack *nodePath) {
	var (
		p, pp *nodePath
		s     *Node
	)

	// case 1: reach the root.
	// execute: nothing.
	// result: balance finish.
	for stack.previous != nil {
		p, pp = stack, stack.previous
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
				stack = stack.previous
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
