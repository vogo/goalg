// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This binary compares memory usage between btree and rbtree.
// copy from https://github.com/google/btree/blob/master/btree_mem.go
package main

import (
	"flag"
	"fmt"
	"math/rand"
	"runtime"
	"time"

	"github.com/google/btree"
	"github.com/vogo/goalg/compare"
	"github.com/vogo/goalg/tree/rbtree"
)

var (
	size     = flag.Int("size", 1000000, "size of the tree to build")
	degree   = flag.Int("degree", 8, "degree of btree")
	gorbtree = flag.Bool("rbtree", false, "use rbtree instead of btree")
)

func main() {
	flag.Parse()
	vals := rand.Perm(*size)
	var t, v interface{}
	v = vals
	var stats runtime.MemStats
	for i := 0; i < 10; i++ {
		runtime.GC()
	}
	fmt.Println("-------- BEFORE ----------")
	runtime.ReadMemStats(&stats)
	fmt.Printf("%+v\n", stats)
	start := time.Now()
	if *gorbtree {
		tr := rbtree.New()
		for _, v := range vals {
			tr.Add(compare.Int(v))
		}
		t = tr // keep it around
	} else {
		tr := btree.New(*degree)
		for _, v := range vals {
			tr.ReplaceOrInsert(btree.Int(v))
		}
		t = tr // keep it around
	}
	fmt.Printf("%v inserts in %v\n", *size, time.Since(start))
	fmt.Println("-------- AFTER ----------")
	runtime.ReadMemStats(&stats)
	fmt.Printf("%+v\n", stats)
	for i := 0; i < 10; i++ {
		runtime.GC()
	}
	fmt.Println("-------- AFTER GC ----------")
	runtime.ReadMemStats(&stats)
	fmt.Printf("%+v\n", stats)
	if t == v {
		fmt.Println("to make sure vals and tree aren't GC'd")
	}
}
