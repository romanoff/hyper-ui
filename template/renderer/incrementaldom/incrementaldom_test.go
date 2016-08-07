package incrementaldom

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
	expected := "text('Hello world')\n"
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
	expected := "text(ctx)\n"
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
	expected := "text(ctx.a.b)\n"
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
	expected := `elementOpen('div', '', ['classes', 'a b', 'id', 'identifier'])
text('Hello world')
elementClose('div')
`
	if expected != string(content) {
		t.Errorf("Expected to get:\n%v\n,but got:\n%v\n", expected, string(content))
	}
}
