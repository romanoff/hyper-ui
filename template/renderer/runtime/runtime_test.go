package runtime

import (
	"encoding/json"
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
	content, err := renderer.Render(nil)
	if err != nil {
		t.Errorf("Expected not to get error while rendering text node, but got %v", err)
	}
	expected := "Hello world"
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
	content, err := renderer.Render("Hello")
	if err != nil {
		t.Errorf("Expected not to get error while rendering variable node, but got %v", err)
	}
	expected := "Hello"
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
	var context map[string]interface{}
	if err := json.Unmarshal([]byte(`{"a": {"b": "Hello from json"}}`), &context); err != nil {
		t.Errorf("Undefined json: %v", err)
	}
	content, err := renderer.Render(context)
	if err != nil {
		t.Errorf("Expected not to get error while rendering variable node, but got %v", err)
	}
	expected := "Hello from json"
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
	content, err := renderer.Render(nil)
	if err != nil {
		t.Errorf("Expected not to get error while rendering tag node, but got %v", err)
	}
	expected := "<div class=\"a b\" id=\"identifier\">Hello world</div>"
	if expected != string(content) {
		t.Errorf("Expected to get:\n%v\n,but got:\n%v\n", expected, string(content))
	}
}
