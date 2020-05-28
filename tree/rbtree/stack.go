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

func (s *stack) insertBeforeCurrent(n *Node, pos Position) {
	s.nodes[s.index+1], s.positions[s.index+1] = s.nodes[s.index], s.positions[s.index]
	s.nodes[s.index], s.positions[s.index] = n, pos
	s.index++
}

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
func (s *stack) childSibling() *Node {
	if s.position() == Left {
		return s.node().Right
	} else {
		return s.node().Left
	}
}

func (s *stack) bindChild(n *Node) {
	if s.position() == Left {
		s.node().Left = n
	} else {
		s.node().Right = n
	}
}
