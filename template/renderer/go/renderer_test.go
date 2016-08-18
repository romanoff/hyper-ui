package gorenderer

import (
	"github.com/romanoff/hyper-ui/template/ast"
	"testing"
)

func TestRenderText(t *testing.T) {
	tr := &ast.Tree{
		Root: &ast.ListNode{Nodes: []ast.Node{
			&ast.TextNode{
				NodeType: ast.NodeText,
				Content:  []byte("Hello world"),
			},
		}},
	}
	renderer := &Renderer{Tree: tr}
	renderer.Init()
	content, err := renderer.Render()
	if err != nil {
		t.Errorf("Expected not to get error while rendering text node, but got %v", err)
	}
	expected := `buffer := new(bytes.Buffer)
buffer.Write([]byte{'H','e','l','l','o',' ','w','o','r','l','d'})
return buffer.Bytes(), nil
`
	if expected != string(content) {
		t.Errorf("Expected to get:\n%v\n,but got:\n%v\n", expected, string(content))
	}
}


func TestRenderVariable(t *testing.T) {
	tr := &ast.Tree{
		Root: &ast.ListNode{Nodes: []ast.Node{
			&ast.VariableNode{
				NodeType: ast.NodeVariable,
				Ident:    []string{"."},
			},
		}},
	}
	renderer := &Renderer{Tree: tr}
	renderer.Init()
	content, err := renderer.Render()
	if err != nil {
		t.Errorf("Expected not to get error while rendering variable node, but got %v", err)
	}
	expected := `buffer := new(bytes.Buffer)
buffer.Write([]byte(fmt.Sprintf("%v", ctx)))
return buffer.Bytes(), nil
`
	if expected != string(content) {
		t.Errorf("Expected to get:\n%v\n,but got:\n%v\n", expected, string(content))
	}
}

func TestRenderVariableNested(t *testing.T) {
	tr := &ast.Tree{
		Root: &ast.ListNode{Nodes: []ast.Node{
			&ast.VariableNode{
				NodeType: ast.NodeVariable,
				Ident:    []string{".", "a", "b"},
			},
		}},
	}
	renderer := &Renderer{Tree: tr}
	renderer.Init()
	content, err := renderer.Render()
	if err != nil {
		t.Errorf("Expected not to get error while rendering variable node, but got %v", err)
	}
	expected := `buffer := new(bytes.Buffer)
buffer.Write([]byte(fmt.Sprintf("%v", hui.Get(ctx, "a", "b"))))
return buffer.Bytes(), nil
`
	if expected != string(content) {
		t.Errorf("Expected to get:\n%v\n,but got:\n%v\n", expected, string(content))
	}
}

func TestRenderTag(t *testing.T) {
	tr := &ast.Tree{
		Root: &ast.ListNode{Nodes: []ast.Node{
			&ast.TagNode{
				Name:       "div",
				Attributes: map[string]string{"id": "identifier"},
				Classes:    []string{"a", "b"},
				ListNode: &ast.ListNode{
					Nodes: []ast.Node{
						&ast.TextNode{
							NodeType: ast.NodeText,
							Content:  []byte("Hello world"),
						},
					},
				},
			},
		}},
	}
	renderer := &Renderer{Tree: tr}
	renderer.Init()
	content, err := renderer.Render()
	if err != nil {
		t.Errorf("Expected not to get error while rendering tag node, but got %v", err)
	}
	expected := `buffer := new(bytes.Buffer)
buffer.Write([]byte{'<','d','i','v'})
buffer.Write([]byte{' ', 'c', 'l', 'a', 's', 's', '=', '\''})
buffer.Write([]byte{'a', ' ', 'b'})
buffer.Write([]byte{'\''})
buffer.Write([]byte{' ','i','d', '=', '\''})
buffer.Write([]byte{'i','d','e','n','t','i','f','i','e','r', '\''})
buffer.Write([]byte{'>'})
buffer.Write([]byte{'H','e','l','l','o',' ','w','o','r','l','d'})
buffer.Write([]byte{'<', '/','d','i','v', '>'})
return buffer.Bytes(), nil
`
	if expected != string(content) {
		t.Errorf("Expected to get:\n%v\n,but got:\n%v\n", expected, string(content))
	}
}
