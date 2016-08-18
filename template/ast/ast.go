package ast

// NodeType identifies the type of a parse tree node.
type NodeType int

func (t NodeType) Type() NodeType {
	return t
}

const (
	NodeText NodeType = iota // Plain text.
	NodeList                 // A list of Nodes.
	NodeVariable
	NodeTag
)

type Node interface {
	Type() NodeType
	// String() string
	Position() Pos
	tree() *Tree
}

type Pos int

func (p Pos) Position() Pos {
	return p
}

type ListNode struct {
	NodeType
	Pos
	tr    *Tree
	Nodes []Node
}

func (self *ListNode) tree() *Tree {
	return self.tr
}

type TextNode struct {
	NodeType
	Pos
	tr      *Tree
	Content []byte
}

func (self *TextNode) tree() *Tree {
	return self.tr
}

type VariableNode struct {
	NodeType
	Pos
	tr    *Tree
	Ident []string // Variable name and fields in lexical order.
}

func (self *VariableNode) tree() *Tree {
	return self.tr
}

type TagNode struct {
	NodeType
	Name string
	Pos
	tr    *Tree
	ListNode *ListNode
	Attributes map[string]string
	Classes []string
	Void bool
}

func (self *TagNode) tree() *Tree {
	return self.tr
}
