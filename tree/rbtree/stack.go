// Copyright 2020 wongoo@apache.org. All rights reserved.

package rbtree

const (
	maxStackDeep = 32
)

type stack struct {
	nodes     []*Node
	positions []Position
	index     int
}

func newStack(root *Node) *stack {
	s := &stack{
		nodes:     make([]*Node, maxStackDeep),
		positions: make([]Position, maxStackDeep),
		index:     -1,
	}
	s.push(&Node{}, Left)
	s.nodes[0].Left = root
	return s
}

func (s *stack) init(root *Node) {
	s.nodes[0].Left = root
}

func (s *stack) reset() {
	for s.index > 0 {
		s.nodes[s.index] = nil
		s.index--
	}

	s.nodes[0].Left = nil
}

func (s *stack) push(n *Node, pos Position) {
	s.index++
	s.nodes[s.index], s.positions[s.index] = n, pos
}

//func (s *stack) insertBeforeCurrent(n *Node, pos Position) {
//	s.nodes[s.index+1], s.positions[s.index+1] = s.nodes[s.index], s.positions[s.index]
//	s.nodes[s.index], s.positions[s.index] = n, pos
//	s.index++
//}

func (s *stack) pop() *stack {
	if s.index == 0 {
		return s
	}

	s.nodes[s.index] = nil
	s.index--
	return s
}

func (s *stack) root() *Node {
	return s.nodes[0].Left
}

func (s *stack) node() *Node {
	return s.nodes[s.index]
}

func (s *stack) position() Position {
	return s.positions[s.index]
}

func (s *stack) parentPosition() Position {
	return s.positions[s.index-1]
}

func (s *stack) parent() *Node {
	if s.index > 0 {
		return s.nodes[s.index-1]
	}
	return nil
}

func (s *stack) sibling() *Node {
	if s.index > 0 {
		i := s.index - 1
		if s.positions[i] == Left {
			return s.nodes[i].Right
		} else {
			return s.nodes[i].Left
		}
	}
	return nil
}

func (s *stack) bindChild(n *Node) {
	if s.position() == Left {
		s.node().Left = n
	} else {
		s.node().Right = n
	}
}

// nodePath the double linked list for the route path in a red-black tree.
// while there is not a link to parent for a Node, the nodePath is used to save the route path
// of the tree, it will be used in loop search/balance of the red-black tree.
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

//// rotateUpChild rotate the child up as the parent of current node.
//func (np *nodePath) rotateUpChild() *nodePath {
//	npp := &nodePath{
//		previous: np.previous,
//		next:     np,
//		pos:      np.previous.pos,
//	}
//
//	if np.pos == Left {
//		npp.Node = RightRotate(np.Node)
//	} else {
//		npp.Node = LeftRotate(np.Node)
//	}
//
//	np.previous.next = npp
//	np.previous = npp
//
//	return npp
//}

// rotateUpSiblingChild rotate the sibling child up as the parent of current node.
func (np *nodePath) rotateUpSiblingChild() *nodePath {
	npp := &nodePath{
		previous: np.previous,
		next:     np,
	}

	if np.pos == Left {
		npp.pos = Left
		npp.Node = LeftRotate(np.Node)
	} else {
		npp.pos = Right
		npp.Node = RightRotate(np.Node)
	}

	np.previous.next = npp
	np.previous = npp

	return npp
}
