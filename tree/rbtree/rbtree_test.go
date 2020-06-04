// Copyright 2020 wongoo@apache.org. All rights reserved.

package rbtree

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func RandNumRbTree(t *testing.T, count int) *Node {
	return NumRbTree(t, rand.Perm(count))
}

func NumRbTree(t *testing.T, arr []int) *Node {
	var root *Node

	t.Log("rbtree rand build seq:", arr)

	for _, n := range arr {
		root = AddNode(root, Int(n))
	}

	return root
}

func TestRbTreeGraph(t *testing.T) {
	root := NumRbTree(t, []int{0, 4, 2, 3, 6, 5, 1, 7, 9, 8, 11, 30, 10, 40, 20, 200, 500, 300, 400})
	t.Log("root:", root.Item)
	generateTreeSvg(t, root)
}

func TestRbTreeAddFindDelete(t *testing.T) {
	tree := New()

	tree.Add(Int(4))
	tree.Add(Int(5))
	tree.Add(Int(6))
	tree.Add(Int(1))
	tree.Add(Int(2))
	tree.Add(Int(3))

	tree.Add(Int(7))
	assert.Equal(t, Int(7), tree.Find(Int(7)))

	tree.Add(Int(8))
	assert.Equal(t, Int(8), tree.Find(Int(8)))

	assert.Nil(t, tree.Delete(Int(10)))

	assert.Equal(t, Int(1), tree.Delete(Int(1)))
	assert.Equal(t, Int(2), tree.Delete(Int(2)))
	assert.Equal(t, Int(3), tree.Delete(Int(3)))
	assert.Equal(t, Int(4), tree.Delete(Int(4)))
	assert.Equal(t, Int(5), tree.Delete(Int(5)))
	assert.Equal(t, Int(6), tree.Delete(Int(6)))
	assert.Equal(t, Int(7), tree.Delete(Int(7)))
	assert.Equal(t, Int(8), tree.Delete(Int(8)))

	assert.Nil(t, tree.Delete(Int(8)))
}

func TestFindDelete(t *testing.T) {
	root := RandNumRbTree(t, 8)
	val := Find(root, Int(7))
	assert.Equal(t, Int(7), val)
	root, ret := Delete(root, Int(7))
	assert.Equal(t, Int(7), ret)
	root, ret = Delete(root, Int(6))
	assert.Equal(t, Int(6), ret)
	root, ret = Delete(root, Int(5))
	assert.Equal(t, Int(5), ret)
	root, ret = Delete(root, Int(4))
	assert.Equal(t, Int(4), ret)
	root, ret = Delete(root, Int(3))
	assert.Equal(t, Int(3), ret)
	root, ret = Delete(root, Int(2))
	assert.Equal(t, Int(2), ret)
	root, ret = Delete(root, Int(1))
	assert.Equal(t, Int(1), ret)
	root, ret = Delete(root, Int(0))
	assert.Equal(t, Int(0), ret)
	assert.Nil(t, root)
}

func TestRbTreeFindDelete2(t *testing.T) {
	root := RandNumRbTree(t, 8)

	val := Find(root, Int(7))
	assert.Equal(t, Int(7), val)

	root, ret := Delete(root, Int(7))
	assert.Equal(t, Int(7), ret)
	root, ret = Delete(root, Int(0))
	assert.Equal(t, Int(0), ret)
	root, ret = Delete(root, Int(6))
	assert.Equal(t, Int(6), ret)
	root, ret = Delete(root, Int(1))
	assert.Equal(t, Int(1), ret)
	root, ret = Delete(root, Int(5))
	assert.Equal(t, Int(5), ret)
	root, ret = Delete(root, Int(2))
	assert.Equal(t, Int(2), ret)
	root, ret = Delete(root, Int(4))
	assert.Equal(t, Int(4), ret)
	root, ret = Delete(root, Int(3))
	assert.Equal(t, Int(3), ret)
	assert.Nil(t, root)
}

func generateTreeSvg(t *testing.T, root *Node) {
	buf := bytes.NewBuffer(nil)
	buf.WriteString(`digraph G {
node [shape=circle,style=solid]
edge [arrowhead=none]
`)
	if root != nil {
		fillTreeDot(buf, root)
	}

	buf.WriteString("}")

	dir := os.TempDir()
	t.Log(dir)

	dotpath := filepath.Join(dir, "goalg_rbtree.dot")
	svgpath := filepath.Join(dir, "goalg_rbtree.svg")
	err := ioutil.WriteFile(dotpath, buf.Bytes(), 0660)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	command := fmt.Sprintf("dot %s -Tsvg -o%s", dotpath, svgpath)
	t.Log(command)
	result, err := exec.Command("/bin/sh", "-c", command).CombinedOutput()
	if err != nil {
		t.Logf("%s, %v", result, err)
		return
	}

	_ = exec.Command("open", svgpath).Run()
}

func fillTreeDot(buf *bytes.Buffer, node *Node) {
	buf.WriteString(fmt.Sprintf("%d[color=%s];\n", node.Item, node.Color))
	if node.Left != nil {
		buf.WriteString(fmt.Sprintf("%d -> %d[color=%s];\n", node.Item, node.Left.Item, node.Left.Color))
		fillTreeDot(buf, node.Left)
	}
	if node.Right != nil {
		buf.WriteString(fmt.Sprintf("%d -> %d[color=%s];\n", node.Item, node.Right.Item, node.Right.Color))
		fillTreeDot(buf, node.Right)
	}
}

var (
	benchmarkTestArr = rand.Perm(128)
)

// BenchmarkAdd the performance of addTreeNode is better than AddNode
func BenchmarkAdd(b *testing.B) {
	var root *Node
	stack := newStack(root)
	for i := 0; i < b.N; i++ {
		for _, n := range benchmarkTestArr {
			root = addTreeNode(stack, root, Int(n))
		}
	}
}

func BenchmarkAddOne(b *testing.B) {
	var root *Node
	for i := 0; i < b.N; i++ {
		for _, n := range benchmarkTestArr {
			root = AddNode(root, Int(n))
		}
	}
}
