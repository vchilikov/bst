package main

import (
	"fmt"
	"strconv"
	"strings"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type BSTNode struct {
	Val   int
	Left  *BSTNode
	Right *BSTNode
}

func (t *BSTNode) insert(val int) *BSTNode {

	return t
}

func (t *BSTNode) delete(val int) *BSTNode {

	return t
}

func (t *BSTNode) search(val int) *BSTNode {

	return t
}

func (t *BSTNode) height() int {
	if t == nil {
		return 0
	}
	return 1 + max(t.Left.height(), t.Right.height())
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

func pad(s string, length int) string {
	if len(s) >= length {
		return s
	}
	pLen := length - len(s)
	return strings.Repeat(" ", (pLen+1)/2) + s + strings.Repeat(" ", pLen-(pLen+1)/2)
}

func (t *BSTNode) print() {
	h := t.height()
	maxLen := max(len("nil"), len(strconv.Itoa(t.max())))
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
				fmt.Print(pad(strconv.Itoa(el.Val), maxLen))
				newLine = append(newLine, el.Left, el.Right)
			} else {
				fmt.Print(pad("nil", maxLen))
				newLine = append(newLine, nil, nil)
			}
			fmt.Print(emptySpace)
		}
		line = newLine
		fmt.Println()
	}
}

func main() {
	t := &BSTNode{
		Val: 10,
		Left: &BSTNode{
			Val: 5,
			Left: &BSTNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
			Right: &BSTNode{
				Val:  6,
				Left: nil,
				Right: &BSTNode{
					Val: 9,
				},
			},
		},
		Right: &BSTNode{
			Val:   152,
			Left:  nil,
			Right: nil,
		},
	}
	t.print()
}
