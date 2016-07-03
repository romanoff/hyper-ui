package runtime

import (
	"bytes"
	"errors"
	"fmt"
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

func (self *Renderer) writeNode(node ast.Node,
	context interface{},
	variables map[string]interface{},
) error {
	switch n := node.(type) {
	case *ast.TextNode:
		self.write(n.Content)
		return nil
	case *ast.VariableNode:
		var value interface{}
		if n.Ident[0] == "." {
			value = context
		} else {
			value = variables[n.Ident[0]]
		}
		for _, ident := range n.Ident[1:] {
			mapValue, ok := value.(map[string]interface{})
			if !ok {
				return errors.New(fmt.Sprintf("runtime error: can't read variable %v", ident))
			}
			value = mapValue[ident]
		}
		//TODO: likely can be improved
		self.write([]byte(fmt.Sprintf("%v", value)))
		return nil
	}
	panic("unreachable")
}

func (self *Renderer) Render(context interface{}) ([]byte, error) {
	var err error
	for _, node := range self.Tree.Root.Nodes {
		err = self.writeNode(node, context, nil)
		if err != nil {
			return nil, err
		}
	}
	return self.buffer.Bytes(), nil
}
