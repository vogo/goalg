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
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRbTreeGraph(t *testing.T) {
	root := NumRbTree(t, []int{0, 4, 2, 3, 5, 1, 7, 9, 8, 30, 10, 20, 200, 300, 400})
	t.Log("root:", root.Key)
	generateTreeSvg(t, root)
}

func TestRbTreeAddFindDelete(t *testing.T) {
	tree := New()

	tree.Add(4, "4")
	tree.Add(5, "5")
	tree.Add(6, "6")
	tree.Add(1, "1")
	tree.Add(2, "2")
	tree.Add(3, "3")

	tree.Add(7, "7")
	assert.Equal(t, "7", tree.Find(7))

	tree.Add(7, "77")
	assert.Equal(t, "77", tree.Find(7))

	tree.Add(8, "8")
	assert.Equal(t, "8", tree.Find(8))

	assert.Nil(t, tree.Delete(10))

	assert.Equal(t, "1", tree.Delete(1))
	assert.Equal(t, "2", tree.Delete(2))
	assert.Equal(t, "3", tree.Delete(3))
	assert.Equal(t, "4", tree.Delete(4))
	assert.Equal(t, "5", tree.Delete(5))
	assert.Equal(t, "6", tree.Delete(6))
	assert.Equal(t, "77", tree.Delete(7))
	assert.Equal(t, "8", tree.Delete(8))

	assert.Nil(t, tree.Delete(8))
}

func TestFindDelete(t *testing.T) {
	root := RandNumRbTree(t, 8)
	val := Find(root, 7)
	assert.Equal(t, "7", val)
	root, ret := Delete(root, 7)
	assert.Equal(t, "7", ret)
	root, ret = Delete(root, 6)
	assert.Equal(t, "6", ret)
	root, ret = Delete(root, 5)
	assert.Equal(t, "5", ret)
	root, ret = Delete(root, 4)
	assert.Equal(t, "4", ret)
	root, ret = Delete(root, 3)
	assert.Equal(t, "3", ret)
	root, ret = Delete(root, 2)
	assert.Equal(t, "2", ret)
	root, ret = Delete(root, 1)
	assert.Equal(t, "1", ret)
	root, ret = Delete(root, 0)
	assert.Equal(t, "0", ret)
	assert.Nil(t, root)
}

func TestRbTreeFindDelete2(t *testing.T) {
	root := RandNumRbTree(t, 8)

	val := Find(root, 7)
	assert.Equal(t, "7", val)

	root, ret := Delete(root, 7)
	assert.Equal(t, "7", ret)
	root, ret = Delete(root, 0)
	assert.Equal(t, "0", ret)
	root, ret = Delete(root, 6)
	assert.Equal(t, "6", ret)
	root, ret = Delete(root, 1)
	assert.Equal(t, "1", ret)
	root, ret = Delete(root, 5)
	assert.Equal(t, "5", ret)
	root, ret = Delete(root, 2)
	assert.Equal(t, "2", ret)
	root, ret = Delete(root, 4)
	assert.Equal(t, "4", ret)
	root, ret = Delete(root, 3)
	assert.Equal(t, "3", ret)
	assert.Nil(t, root)
}

func RandNumRbTree(t *testing.T, count int) *Node {
	return NumRbTree(t, rand.Perm(count))
}

func NumRbTree(t *testing.T, arr []int) *Node {
	var root *Node

	t.Log("rbtree rand build seq:", arr)

	for _, n := range arr {
		root = AddNode(root, n, strconv.Itoa(n))
	}

	return root
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
	err := ioutil.WriteFile(dotpath, buf.Bytes(), 0666)
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
	buf.WriteString(fmt.Sprintf("%d[color=%s];\n", node.Key, node.Color))
	if node.Left != nil {
		buf.WriteString(fmt.Sprintf("%d -> %d[color=%s];\n", node.Key, node.Left.Key, node.Left.Color))
		fillTreeDot(buf, node.Left)
	}
	if node.Right != nil {
		buf.WriteString(fmt.Sprintf("%d -> %d[color=%s];\n", node.Key, node.Right.Key, node.Right.Color))
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
	value := "1"
	for i := 0; i < b.N; i++ {
		for _, n := range benchmarkTestArr {
			root = addTreeNode(stack, root, n, value)
		}
	}
}

func BenchmarkAddOne(b *testing.B) {
	var root *Node
	value := "1"
	for i := 0; i < b.N; i++ {
		for _, n := range benchmarkTestArr {
			root = AddNode(root, n, value)
		}
	}
}
