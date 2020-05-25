// 红黑树 red-black tree
package rbtree

type Color bool

func (c Color) String() string {
	if c {
		return "red"
	}
	return "black"
}

type Position bool

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

// Node the node of red-black tree
type Node struct {
	Key         int
	Color       Color
	Left, Right *Node
	Value       interface{}
}

// leftRotate left rotate
func leftRotate(n *Node) *Node {
	r := n.Right

	n.Right = r.Left
	r.Left = n

	return r
}

// rightRotate right rotate
func rightRotate(n *Node) *Node {
	l := n.Left

	n.Left = l.Right
	l.Right = n

	return l
}

// AddKeyValue add new key/value, return the new root node
func AddKeyValue(root *Node, key int, value interface{}) *Node {
	return AddNew(root, &Node{
		Key:   key,
		Color: Red,
		Value: value,
	})
}

// AddNew add new node, return the new root node
func AddNew(root *Node, new *Node) *Node {
	new.Color = Red
	root = addNode(root, Left, new)
	root.Color = Black
	return root
}

// addNode recursively down to leaf, and add the new node to the leaf, then rebuild the tree from the leaf to root
func addNode(node *Node, pos Position, new *Node) *Node {
	if node == nil {
		return new
	}

	if new.Key < node.Key {
		node.Left = addNode(node.Left, Left, new)
		if node.Left.Color == Black {
			return node
		}

		if node.Color == Red {
			if pos == Right {
				node = rightRotate(node)
			}
			return node
		}

		if node.Left.Left != nil && node.Left.Left.Color == Red {
			node = rightRotate(node)
			node.Left.Color = Black
		}

		return node
	}

	if new.Key > node.Key {
		node.Right = addNode(node.Right, Right, new)

		if node.Right.Color == Black {
			return node
		}

		if node.Color == Red {
			if pos == Left {
				node = leftRotate(node)
			}
			return node
		}

		if node.Right.Right != nil && node.Right.Right.Color == Red {
			node = leftRotate(node)
			node.Right.Color = Black
		}

		return node
	}

	new.Left = node.Left
	new.Right = node.Right
	new.Color = node.Color

	return new
}
