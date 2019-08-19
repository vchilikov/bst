package main

import (
	"bst/utils"
	"fmt"
	"strconv"
	"strings"
)

type BST struct {
	root *BSTNode
}

func BSTNew() *BST {
	return new(BST)
}

func (t *BST) Insert(val int) {
	t.root = t.root.insert(val)
}

func (t *BST) Delete(val int) {

}

func (t *BST) Exists(val int) bool {
	return false
}

func (t *BST) Print() {
	t.root.print()
}

type BSTNode struct {
	Val   int
	Left  *BSTNode
	Right *BSTNode
}

func (t *BSTNode) insert(val int) *BSTNode {
	if t == nil {
		return &BSTNode{
			Val: val,
		}
	}
	if val < t.Val {
		t.Left = t.Left.insert(val)
	} else if val > t.Val {
		t.Right = t.Right.insert(val)
	}
	return t
}

func (t *BSTNode) height() int {
	if t == nil {
		return 0
	}
	return 1 + utils.Max(t.Left.height(), t.Right.height())
}

func (t *BSTNode) max() int {
	if t == nil {
		return 0
	}
	if t.Right == nil {
		return t.Val
	}
	return t.Right.max()
}

func (t *BSTNode) print() {
	h := t.height()
	maxLen := utils.Max(len("nil"), len(strconv.Itoa(t.max())))
	emptyString := strings.Repeat(" ", maxLen)
	maxLineLen := 1<<uint(h) - 1
	line := []*BSTNode{t}
	for i := 0; i < h; i++ {
		var newLine []*BSTNode
		linePad := 1<<uint(h-i-1) - 1
		fmt.Print(strings.Repeat(emptyString, linePad))
		cntEl := 1 << uint(i)
		emptySpace := ""
		if i > 0 {
			emptySpace = strings.Repeat(emptyString, (maxLineLen-linePad*2-cntEl)/(cntEl-1))
		}
		for _, el := range line {
			if el != nil {
				fmt.Print(utils.Pad(strconv.Itoa(el.Val), maxLen))
				newLine = append(newLine, el.Left, el.Right)
			} else {
				fmt.Print(utils.Pad("nil", maxLen))
				newLine = append(newLine, nil, nil)
			}
			fmt.Print(emptySpace)
		}
		line = newLine
		fmt.Println()
	}
}

func main() {
	t := BSTNew()
	t.Insert(10)
	t.Insert(5)
	t.Insert(11)
	t.Insert(12)
	t.Insert(6)
	t.Insert(7)
	t.Insert(1)
	t.Insert(8)
	t.Print()
}
