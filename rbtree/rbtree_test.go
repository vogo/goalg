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

	"github.com/wongoo/goalg/rbtree"
)

func TestNewNumRbTree(t *testing.T) {
	root := RandNumRbTree(32)
	generateTreeSvg(t, root)
}

func RandNumRbTree(count int) *rbtree.Node {
	var root *rbtree.Node

	arr := make([]int, count)
	for i := 0; i < count; i++ {
		arr[i] = i
	}

	for len(arr) > 0 {
		i := rand.Intn(len(arr))
		root = rbtree.AddKeyValue(root, arr[i], strconv.Itoa(arr[i]))
		arr = append(arr[:i], arr[i+1:]...)
	}

	return root
}

func generateTreeSvg(t *testing.T, root *rbtree.Node) {
	buf := bytes.NewBuffer(nil)
	buf.WriteString(`digraph G {
node [shape=circle,style=solid]
edge [arrowhead=none]
`)
	fillTreeDot(buf, root)

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
	if node.Left != nil {
		buf.WriteString(fmt.Sprintf("%d -> %d[color=%s];\n", node.Key, node.Left.Key, node.Left.Color))
		if node.Left.Color == rbtree.Red {
			buf.WriteString(fmt.Sprintf("%d[color=%s];\n", node.Left.Key, node.Left.Color))
		}
		fillTreeDot(buf, node.Left)
	}
	if node.Right != nil {
		buf.WriteString(fmt.Sprintf("%d -> %d[color=%s];\n", node.Key, node.Right.Key, node.Right.Color))
		if node.Right.Color == rbtree.Red {
			buf.WriteString(fmt.Sprintf("%d[color=%s];\n", node.Right.Key, node.Right.Color))
		}
		fillTreeDot(buf, node.Right)
	}
}
