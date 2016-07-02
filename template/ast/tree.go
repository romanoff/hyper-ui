package ast

type Tree struct {
	Name string
	Root *ListNode
	text string
	treeSet   map[string]*Tree
}
