package runtime

import (
	"bytes"
	"errors"
	"fmt"
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
	case *ast.TagNode:
		//TODO: escape tag name, attribute names and values
		self.write([]byte("<" + n.Name))
		if len(n.Classes) > 0 {
			self.write([]byte(" class=" + "\"" + strings.Join(n.Classes, " ") + "\""))
		}
		for attrName, value := range n.Attributes {
			self.write([]byte(" " + attrName + "=" + "\"" + value + "\""))
		}
		self.write([]byte(">"))
		self.writeNode(n.ListNode, context, variables)
		self.write([]byte("</" + n.Name + ">"))
		return nil
	case *ast.ListNode:
		for _, node := range n.Nodes {
			err := self.writeNode(node, context, variables)
			if err != nil {
				return err
			}
		}
		return nil
	}

	panic("unreachable")
}

func (self *Renderer) Render(context interface{}) ([]byte, error) {
	err := self.writeNode(self.Tree.Root, context, nil)
	if err != nil {
		return nil, err
	}
	return self.buffer.Bytes(), nil
}
