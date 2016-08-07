package incrementaldom

import (
	"bytes"
	"github.com/romanoff/hyper-ui/template/ast"
	"strings"
)

type Renderer struct {
	Tree   *ast.Tree
	buffer *bytes.Buffer
}

func (self *Renderer) Init() {
	self.buffer = new(bytes.Buffer)
}

func (self *Renderer) write(parts ...[]byte) *Renderer {
	for _, v := range parts {
		self.buffer.Write(v)
	}
	return self
}

func (self *Renderer) writeLine(parts ...[]byte) *Renderer {
	self.write(parts...)
	self.buffer.WriteByte('\n')
	return self
}

func (self *Renderer) writeNode(node ast.Node) error {
	switch n := node.(type) {
	case *ast.TextNode:
		self.write([]byte("text('"))
		//TODO: escape content
		self.write(n.Content)
		self.writeLine([]byte("')"))
		return nil
	case *ast.VariableNode:
		value := ""
		if n.Ident[0] == "." {
			value = "ctx"
		} else {
			value = "variables"
		}
		for _, ident := range n.Ident[1:] {
			value += "." + ident
		}
		self.writeLine([]byte("text(" + value + ")"))
		return nil
	case *ast.TagNode:
		if len(n.Classes) > 0 || len(n.Attributes) > 0 {
			self.write([]byte("elementOpen('" + n.Name + "', '', ["))
			if len(n.Classes) > 0 {
				self.write([]byte("'classes', '" + strings.Join(n.Classes, " ") + "'"))
			}
			index := 0
			for attrName, value := range n.Attributes {
				index += 1
				if index != 1 || len(n.Classes) != 0 {
					self.write([]byte(", "))
				}
				self.write([]byte("'" + attrName + "', '" + value + "'"))
			}
			self.writeLine([]byte("])"))
		} else {
			self.writeLine([]byte("elementOpen('" + n.Name + "')"))
		}
		self.writeNode(n.ListNode)
		self.writeLine([]byte("elementClose('" + n.Name + "')"))
		return nil
	case *ast.ListNode:
		for _, node := range n.Nodes {
			err := self.writeNode(node)
			if err != nil {
				return err
			}
		}
		return nil
	}

	panic("unreachable")
}

func (self *Renderer) Render() ([]byte, error) {
	err := self.writeNode(self.Tree.Root)
	if err != nil {
		return nil, err
	}
	return self.buffer.Bytes(), nil
}
