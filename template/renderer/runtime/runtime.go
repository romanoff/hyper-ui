package runtime

import (
	"bytes"
	"github.com/romanoff/hyper-ui/template/ast"
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

func (self *Renderer) writeNode(node ast.Node, context interface{}) error {
	switch n := node.(type) {
	case *ast.TextNode:
		self.write(n.Content)
		return nil
	case *ast.VariableNode:
		content, err := n.Value(context)
		if err != nil {
			return err
		}
		self.write(content)
		return nil
	}
	panic("unreachable")
}

func (self *Renderer) Render(context interface{}) ([]byte, error) {
	var err error
	for _, node := range self.Tree.Root.Nodes {
		err = self.writeNode(node, context)
		if err != nil {
			return nil, err
		}
	}
	return self.buffer.Bytes(), nil
}
