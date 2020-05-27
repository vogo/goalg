// Copyright 2020 wongoo@apache.org. All rights reserved.

package rbtree_test

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
	"github.com/wongoo/goalg/tree/rbtree"
)

func TestRbTreeGraph(t *testing.T) {
	root := NumRbTree(t, []int{0, 4, 2, 3, 5, 1, 7, 9, 8})
	t.Log("root:", root.Key)
	generateTreeSvg(t, root)
}

func TestRbTreeAdd(t *testing.T) {
	root := RandNumRbTree(t, 7)

	root = rbtree.Add(root, 7, "7")
	assert.Equal(t, "7", rbtree.Find(root, 7))

	root = rbtree.Add(root, 7, "77")
	assert.Equal(t, "77", rbtree.Find(root, 7))

	root = rbtree.Add(root, 8, "8")
	assert.Equal(t, "8", rbtree.Find(root, 8))
	root = rbtree.Add(root, 8, "88")
	assert.Equal(t, "88", rbtree.Find(root, 8))
}

func TestRbTreeFindDelete(t *testing.T) {
	root := RandNumRbTree(t, 8)
	val := rbtree.Find(root, 7)
	assert.Equal(t, "7", val)
	root, ret := rbtree.Delete(root, 7)
	assert.Equal(t, "7", ret)
	root, ret = rbtree.Delete(root, 6)
	assert.Equal(t, "6", ret)
	root, ret = rbtree.Delete(root, 5)
	assert.Equal(t, "5", ret)
	root, ret = rbtree.Delete(root, 4)
	assert.Equal(t, "4", ret)
	root, ret = rbtree.Delete(root, 3)
	assert.Equal(t, "3", ret)
	root, ret = rbtree.Delete(root, 2)
	assert.Equal(t, "2", ret)
	root, ret = rbtree.Delete(root, 1)
	assert.Equal(t, "1", ret)
	root, ret = rbtree.Delete(root, 0)
	assert.Equal(t, "0", ret)
	assert.Nil(t, root)
}

func TestRbTreeFindDelete2(t *testing.T) {
	root := RandNumRbTree(t, 8)

	val := rbtree.Find(root, 7)
	assert.Equal(t, "7", val)

	root, ret := rbtree.Delete(root, 7)
	assert.Equal(t, "7", ret)
	root, ret = rbtree.Delete(root, 0)
	assert.Equal(t, "0", ret)
	root, ret = rbtree.Delete(root, 6)
	assert.Equal(t, "6", ret)
	root, ret = rbtree.Delete(root, 1)
	assert.Equal(t, "1", ret)
	root, ret = rbtree.Delete(root, 5)
	assert.Equal(t, "5", ret)
	root, ret = rbtree.Delete(root, 2)
	assert.Equal(t, "2", ret)
	root, ret = rbtree.Delete(root, 4)
	assert.Equal(t, "4", ret)
	root, ret = rbtree.Delete(root, 3)
	assert.Equal(t, "3", ret)
	assert.Nil(t, root)
}

func RandNumRbTree(t *testing.T, count int) *rbtree.Node {
	return NumRbTree(t, rand.Perm(count))
}

func NumRbTree(t *testing.T, arr []int) *rbtree.Node {
	var root *rbtree.Node

	t.Log("rbtree rand build seq:", arr)

	for _, n := range arr {
		root = rbtree.Add(root, n, strconv.Itoa(n))
	}

	return root
}

func generateTreeSvg(t *testing.T, root *rbtree.Node) {
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
		t.FailNow()
	}

	_ = exec.Command("open", svgpath).Run()
}

func fillTreeDot(buf *bytes.Buffer, node *rbtree.Node) {
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
	benchmarkTestArr = rand.Perm(64)
)

// BenchmarkAdd the performance of rbtree.Add is not better than rbtree.AddOne, because creating stack objects.
// will try to use pool to optimize it.
func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var root *rbtree.Node
		for _, n := range benchmarkTestArr {
			root = rbtree.Add(root, n, strconv.Itoa(n))
		}
	}
}

func BenchmarkAddOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var root *rbtree.Node
		for _, n := range benchmarkTestArr {
			root = rbtree.AddOne(root, n, strconv.Itoa(n))
		}
	}
}
